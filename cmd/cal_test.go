package main

import (

  // this is a test.
  "testing"

  // prints.
  "fmt"

  // other imports.
)

func TestPass(t *testing.T) {
  data0 := [][]string {
    { "Mo", "Tu", "We", "Th", "Fr", "Sa", "Su" },
    { "  ", "  ", " 1", " 2", " 3", " 4", " 5" },
    { " 6", " 7", " 8", " 9", "10", "11", "12" },
    { "13", "14", "15", "16", "17", "18", "19" },
    { "20", "21", "22", "23", "24", "25", "26" },
    { "27", "28", "  ", "  ", "  ", "  ", "  " },
  }
  NL := fmt.Sprintln()
  data1 := "Mo Tu We Th Fr Sa Su" + NL +
           "       1  2  3  4  5" + NL +
           " 6  7  8  9 10 11 12" + NL +
           "13 14 15 16 17 18 19" + NL +
           "20 21 22 23 24 25 26" + NL +
           "27 28               " + NL
  fmt.Printf("%#v\n", data0)
  fmt.Printf("%s\n",  data1)
}


