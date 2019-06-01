package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type MarksStudent struct {
	UID     int64  `json:"uid"`
	Name    string `json:"name"`
	Maths   int    `json:"maths"`
	Physics int    `json:"physics"`
	Chem    int    `json:"chemistry"`
}

func loadStudentMarksJson(filepath string) *[]MarksStudent {
	arr := new([]MarksStudent)
	dataJson, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(dataJson, &arr)
	if err != nil {
		log.Fatal("Invalid Json! Please give a valid json input")
	}

	return arr
}

// Insert into sqlite db
func Insert(tid int, database *sql.DB, entries <-chan MarksStudent, results chan<- MarksStudent) {

	statement, _ := database.Prepare("INSERT INTO student_marks (uid, name, maths, physics, chemistry) VALUES (?, ?, ?, ?, ?)")
	for entry := range entries {
		//fmt.Println(tid, entry)
		statement.Exec(entry.UID, entry.Name, entry.Maths,
			entry.Physics, entry.Chem)
		results <- entry
	}

}

func main() {

	inputfile := flag.String("f", "", "input file path")
	workerThreads := flag.Int("w", 1, "number of worker threads to load data into sqlite")
	flag.Parse()
	studentMarksArr := loadStudentMarksJson(*inputfile)
	database, _ := sql.Open("sqlite3", "./students_marks.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS student_marks (uid INTEGER PRIMARY KEY, name TEXT, maths INTEGER, physics INTEGER, chemistry INTEGER)")
	statement.Exec()

	fmt.Println("loading json file into sqlite...")
	var entries = make(chan MarksStudent, len(*studentMarksArr))
	var results = make(chan MarksStudent, len(*studentMarksArr))

	for w1 := 0; w1 < *workerThreads; w1++ {
		go Insert(w1, database, entries, results)
	}

	for i := 0; i < len(*studentMarksArr); i++ {
		entries <- (*studentMarksArr)[i]
	}
	close(entries)

	for a := 0; a < len(*studentMarksArr); a++ {
		<-results
	}

	fmt.Printf("loading completed!")
	database.Close()

}
