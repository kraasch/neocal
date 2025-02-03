
package main

import (
  "fmt"
  "time"
  "strconv"
)

func parseTime(in string) (out time.Time) {
  const layout = "2006-01-02"
  parsedDate, err := time.Parse(layout, in)
  if err != nil {
    fmt.Println("Error parsing date:", err)
  }
  out = parsedDate
  return
}

func DateAsHeader(targetDate string) (layouted string) {
  parsedDate := parseTime(targetDate)
  y := strconv.Itoa(parsedDate.Year())
  m := parsedDate.Month().String()
  d := strconv.Itoa(parsedDate.Day())
  if len(d) == 1 {
    d = " " + d
  }
  layouted = d + ". " + m + ", " + y
  return
}

func MonthAsCalendar(targetDate string, culture string) (s string) {
  currentTime := time.Now() // Get the current time.
  year := int(currentTime.Year())
  // today  := int(currentTime.Day())
  month := currentTime.Month()
  firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC) // Get the first day of the month.
  lastDay := firstDay.AddDate(0, 1, -1) // Get the total number of days in the month.
  // Add the header (day names).
  if culture == "us" {
    s += fmt.Sprintln("Su Mo Tu We Th Fr Sa")
  } else {
    s += fmt.Sprintln("Mo Tu We Th Fr Sa Su")
  }
  // Print leading spaces for the first day.
  for i := 0; i < int(firstDay.Weekday()); i++ {
    s += "  "
  }
  // Print the days of the month.
  for day := 1; day <= lastDay.Day(); day++ {
    s += fmt.Sprintf("%2d ", day) // Print the day, formatted to fit in 2 characters.
    if (firstDay.Day()+day)%7 == 6 { // Move to the next line after 7 days.
      s += fmt.Sprintln()
    }
  }
  return
}

