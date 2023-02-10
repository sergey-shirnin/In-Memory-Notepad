package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
  "strings"
)


const (
	spaceChar = " "
    entryMsg = "Enter a command and data: "
    noData = "<data undefined>"
    exitComm = "exit"
	exitMsg = "[Info] Bye!"
)

func main() {
    var command string
    scanner := bufio.NewScanner(os.Stdin)
  
    for {
      var content string = "<no data>"
      fmt.Printf("%s", entryMsg)
      scanner.Scan()
    
      command, content = func() (string, string) {
        entry := regexp.MustCompile(`\s+`).Split(scanner.Text(), 2)
        entry = append(entry, content)
        return entry[0], entry[1]
      }()
      
      if command == exitComm { break }
      fmt.Printf("%s\n", strings.Join([]string{command, content}, spaceChar))
    }
  fmt.Printf("%s\n", exitMsg)
}
