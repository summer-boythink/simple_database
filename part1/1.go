package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// store connection info
type Connection struct{
	file *os.File
}

// store input buffer
type InputBuffer struct {
	buffer string
	buffer_len int
	input_len int
}

func (B *InputBuffer)readInput(){
	n,err := fmt.Scanln(&B.buffer)
	if err != nil {
		log.Fatal(err)
	}
	B.input_len = n
	B.buffer_len = n
}

func NewInputBuffer() *InputBuffer{
	return &InputBuffer{
		buffer: "",
		buffer_len: 0,
		input_len: 0,
	}
}

func GetDbFilename() string {
	if(len(os.Args) < 2){
		panic("Must supply a filename for the database")
	}
	return os.Args[1]
}

func OpenConnection(dbFilename string) *Connection{
	f,err := os.OpenFile(dbFilename,os.O_RDWR | os.O_CREATE,0755)
	if err != nil {
		log.Fatal(err)
	}
	return &Connection{
		f,
	}
}

func prompt(){
	fmt.Printf("db >")
}

func main(){
	dbFilename := GetDbFilename()
	OpenConnection(dbFilename)
	inputBuffer := NewInputBuffer()
	for {
		prompt()
		inputBuffer.readInput()
		if strings.Compare(".exit",inputBuffer.buffer) == 0{
			fmt.Println("say goodbye !")
			os.Exit(-1)
		}else{
			fmt.Printf("no support %s",inputBuffer.buffer)
		}
	}
}