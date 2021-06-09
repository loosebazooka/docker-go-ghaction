package main

import (
  "fmt"
  "os"

  "github.com/google/uuid"
)

func main() {
  args := os.Args[1:]
  fmt.Printf("Peace be the journey. %v", args[0])

  fmt.Printf("\n::set-output name=time::%v", uuid.New())
}
