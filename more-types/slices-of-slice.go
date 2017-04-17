package main

import (
  "fmt"
  "strings"
)

func main() {
  borad := [][]string{
    []string{"_","_","_"},
    []string{"_","_","_"},
    []string{"_","_","_"},
  }

  borad[0][0] = "X"
  borad[2][2] = "O"
  borad[1][2] = "X"
  borad[1][0] = "O"
  borad[0][2] = "X"

  for i := 0; i < len(borad); i++ {
    fmt.Printf("%s\n",strings.Join(borad[i], "　"))
  }
}