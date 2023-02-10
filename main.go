package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
  inpMsg = "Enter a command and data: "
  addMsg = "[OK] The note was successfully created"
  delMsg = "[OK] All notes were successfully deleted"
  fullMsg, listFmt = "[Error] Notepad is full", "[Info] %d: %s"
  errCmd, errVal = "!command unknown", "!content empty"
  exitMsg, exitCmd = "[Info] Bye!", "exit"
  arrLimit = 5
)

func newSlice() []string {
  return make([]string, 0, arrLimit)
}

func add(notes []string, content string) ([]string, string) {
  if content == "" {return notes, errVal}
  if len(notes) == arrLimit { return notes, fullMsg }
  notes = append(notes, content); return notes, addMsg
}

func list(notes []string) string {
  var listed = newSlice()
  for i, note := range notes { 
    listed = append(listed, fmt.Sprintf(listFmt, i + 1, note))
  }; return strings.Join(listed, "\n")
}

func clear(notes []string) ([]string, string) {
  notes = nil; return notes, delMsg
}

func main() {
    var command string
    var notes = newSlice()
    scanner := bufio.NewScanner(os.Stdin)

    mainLoop:
      for {
        var content, response string
        fmt.Printf("%s", inpMsg)
        scanner.Scan()
      
        command, content = func() (string, string) {
          entry := regexp.MustCompile(`\s+`).Split(scanner.Text(), 2)
          entry = append(entry, content)
          return entry[0], entry[1]
        }()
        
        switch command {
          case "create":
            notes, response = add(notes, content)
          case "list":
            response = list(notes)
          case "clear":
            notes, response = clear(notes)
          case exitCmd:
            break mainLoop
          default:
            fmt.Print(errCmd)
        }
        fmt.Println(response)
      }
  fmt.Printf(exitMsg)
}
