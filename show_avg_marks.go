package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	studentids := flag.String("studentids", "", "input student ids to get average for")
	flag.Parse()

	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string([]byte(*studentids)))
	conn.Write([]byte(*studentids))

	var buf bytes.Buffer
	io.Copy(&buf, conn)
	fmt.Printf(buf.String())
	conn.Close()

}
