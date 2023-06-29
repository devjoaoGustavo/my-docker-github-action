package main

import (
  "fmt"
  "os"
  "time"
)

func main() {
  var output string = "Hello, World!"

  if len(os.Args) > 1 {
    output = fmt.Sprintln("Hello, " + os.Args[1] + "!")
  }

  var current_time = time.Now()

  fmt.Println(output)

  os.Setenv("GITHUB_OUTPUT", current_time.Format(time.RFC3339))
}
