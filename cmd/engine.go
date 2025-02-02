
package main

import (
  "fmt"
  "time"
  "strconv"
)

const (
  layout = "9999-99-99"
)

func DateAsHeader(targetDate string) (header string) {
  parsedDate, err := time.Parse(layout, targetDate)
  if err != nil {
      fmt.Println("Error parsing date:", err)
  }
  y := strconv.Itoa(parsedDate.Year())
  m := parsedDate.Month().String()
  d := strconv.Itoa(parsedDate.Day())
  header = d + ". " + m + " " + y
  return
}

