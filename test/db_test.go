package test

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"testing"
)

func runScript(name string, cmds []string) {
	cmd := exec.Command("E:/go_wheel/db/part3/part3.exe", name)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	defer stdin.Close()
	cmd.Start()
	for _, v := range cmds {
		io.WriteString(stdin, v+"\r\n")
		reader := bufio.NewReader(stdout)
		bytes, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
		}
		output := string(bytes)

		log.Println(output)
	}
	cmd.Wait()
}

func TestRunScript(t *testing.T) {
	runScript("aaa", []string{"insert 1 user1 person1@example.com",
		"select",
		".exit"})
}
