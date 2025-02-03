
package main

import (
  "fmt"
  "time"
  "strconv"
)

func parseYearAndMonth(in string) (out time.Time) {
  const layout = "2006-01"
  parsedDate, err := time.Parse(layout, in)
  if err != nil {
    fmt.Println("Error parsing date:", err)
  }
  out = parsedDate
  return
}

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

  // Get the first day of the target month.
  firstDay := parseYearAndMonth(targetDate)

  // Get the total number of days in the month.
  lastDay := firstDay.AddDate(0, 1, -1)

  // Get the weekday of the first day in the month.
  // US format, week starts with Sunday:
  // 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat.
  weekday := int(firstDay.Weekday())
  if culture == "eu" {
    // turn this: 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat 0=sun
    // into this: 6=sun 0=mon 1=tue 2=wed 3=thu 4=fri 5=sat 6=sun
    weekday = (weekday - 1) % 7
  }

  fmt.Println("firstDay:", firstDay)
  fmt.Println("lastDay:", lastDay)
  fmt.Println("weekday:", weekday)

  // Add the header (day names).
  if culture == "us" {
    s += fmt.Sprintln("Su Mo Tu We Th Fr Sa")
  } else {
    s += fmt.Sprintln("Mo Tu We Th Fr Sa Su")
  }

  // TODO: implement: shift SUN to MON.

  // Print leading spaces for the first day.
  for i := 0; i < weekday; i++ {
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

