package main

import (
  "fmt"
  "os"

  "github.com/google/uuid"
)

func main() {
  args := os.Args[1:]
  fmt.Printf("Peace be the journey. %v", args[0])
  fmt.Println()
  fmt.Printf("::set-output name=some-output::%v", uuid.New())
  fmt.Println()
}
