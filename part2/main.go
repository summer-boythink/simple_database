package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type MetaCommandResult int

const (
	META_COMMAND_SUCCESS MetaCommandResult = iota
	META_COMMAND_UNRECOGNIZED_COMMAND
)

// store connection info
type Connection struct {
	file *os.File
}

// store input buffer
type InputBuffer struct {
	buffer     string
	buffer_len int
	input_len  int
}

func (B *InputBuffer) readInput() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	B.buffer = text[:len(text)-2]
	if err != nil {
		log.Fatal(err)
	}
	B.input_len = len(text)
	B.buffer_len = len(B.buffer)
}

func (B *InputBuffer) doMetaCmd() MetaCommandResult {
	if strings.Compare(".exit", B.buffer) == 0 {
		fmt.Println("say goodbye !")
		os.Exit(-1)
		return META_COMMAND_SUCCESS
	} else {
		return META_COMMAND_UNRECOGNIZED_COMMAND
	}
}

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{
		buffer:     "",
		buffer_len: 0,
		input_len:  0,
	}
}

func GetDbFilename() string {
	if len(os.Args) < 2 {
		panic("Must supply a filename for the database")
	}
	return os.Args[1]
}

func OpenConnection(dbFilename string) *Connection {
	f, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return &Connection{
		f,
	}
}

func prompt() {
	fmt.Printf("db >")
}

func main() {
	dbFilename := GetDbFilename()
	OpenConnection(dbFilename)
	inputBuffer := NewInputBuffer()
	for {
		prompt()
		inputBuffer.readInput()
		if inputBuffer.buffer[0] == '.' {
			switch inputBuffer.doMetaCmd() {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGNIZED_COMMAND:
				fmt.Printf("no support %s\n", inputBuffer.buffer)
				continue
			}
		}

		// prepare cmd
		statement := NewStatement()
		switch statement.PrepareStatement(inputBuffer) {
		case PREPARE_SUCCESS:

		case PREPARE_UNRECOGNIZED_STATEMENT:
			fmt.Printf("no support keyword at start of %s\n", inputBuffer.buffer)
			continue
		}

		// exec cmd
		statement.ExecuteStatement()
		fmt.Println("exec ok!")
	}
}
