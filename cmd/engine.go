
package main

import (
  "fmt"
  "time"
  "strconv"
)

const (
  layout = "2006-01-02"
)

func DateAsHeader(targetDate string) (layouted string) {
  parsedDate, err := time.Parse(layout, targetDate)
  if err != nil {
      fmt.Println("Error parsing date:", err)
  }
  y := strconv.Itoa(parsedDate.Year())
  m := parsedDate.Month().String()
  d := strconv.Itoa(parsedDate.Day())
  if len(d) == 1 {
    d = " " + d
  }
  layouted = d + ". " + m + ", " + y
  return
}

func MonthAsCalendar(targetDate string, formatCulture string) (layouted string) {
  return
}

