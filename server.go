package main

import (
	"avitest/commons"
	"avitest/lrucache"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3000"
	CONN_TYPE = "tcp"
)

var cache *lrucache.LRUCache

func getAvgMarks(querystmt *sql.Stmt, uid string) (float32, error) {
	if cache.Get(uid) == nil {
		// load data from db
		// put data in cache
		studentMarks := new(commons.MarksStudent)

		fmt.Println(uid)
		rows, err := querystmt.Query(uid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rows)
		for rows.Next() {
			rows.Scan(&studentMarks.UID, &studentMarks.Name, &studentMarks.Maths, &studentMarks.Physics, &studentMarks.Chem)
		}
		defer rows.Close()
		if studentMarks.UID == 0 {
			return 0.0, errors.New("Entry Not Found")
		}
		fmt.Println(studentMarks)
		cache.Put(uid, studentMarks)

	}
	marks := cache.Get(uid).(*commons.MarksStudent)
	return float32(marks.Maths+marks.Physics+marks.Chem) / float32(3), nil
}

func generateInputUidList(uidStr string) []string {
	uidStr = strings.Trim(uidStr, ",")
	uidStrList := strings.Split(uidStr, ",")
	return uidStrList
}

func getOutputBytes(resultMap map[string]string) []byte {
	var buffer bytes.Buffer
	for k, v := range resultMap {
		buffer.WriteString(fmt.Sprintf("Avg Marks of %s: %s", k, v))
		buffer.Write([]byte("\n"))
	}
	return buffer.Bytes()
}

func handleRequest(querystmt *sql.Stmt, conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		conn.Close()
		return
	}
	uidStr := string(buf[:n-1])
	if strings.Trim(uidStr, " ") == "" {
		fmt.Println("uid cannot be empty")
		conn.Close()
		return
	}
	uidStrList := generateInputUidList(uidStr)
	resultMap := make(map[string]string)
	for _, uid := range uidStrList {
		avgMarks, err := getAvgMarks(querystmt, uid)
		if err != nil {
			resultMap[uid] = fmt.Sprintf("Entry Not Found.")
			continue
		}
		resultMap[uid] = fmt.Sprintf("%.2f", avgMarks)
	}

	conn.Write(getOutputBytes(resultMap))
	conn.Close()
}

func main() {

	cache = new(lrucache.LRUCache).Init(100000)
	database, _ := sql.Open("sqlite3", "./students_marks.db")
	querystmt, err := database.Prepare("SELECT uid, name, maths, physics, chemistry FROM student_marks  WHERE uid in ($1)")
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(querystmt, conn)
	}
	querystmt.Close()
	database.Close()

}
