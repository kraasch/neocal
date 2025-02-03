
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
  firstDayDate := parseYearAndMonth(targetDate)
  firstDay := firstDayDate.Day()

  // Get the total number of days in the month.
  lastDayDate := firstDayDate.AddDate(0, 1, -1)
  lastDay := lastDayDate.Day()

  // Get the weekday of the first day in the month.
  // US format, week starts with Sunday:
  // 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat.
  weekday := int(firstDayDate.Weekday())
  if culture == "eu" {
    // turn US format: 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat 0=sun
    // into EU format: 6=sun 0=mon 1=tue 2=wed 3=thu 4=fri 5=sat 6=sun
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

  // Print leading spaces for the
  for i := 0; i < weekday; i++ {
    s += "   "
  }
  daysInFirstWeek := 7 - weekday

  // TODO: implement: shift SUN to MON.

  // Print the days of the month.
  for day := 1; day <= lastDay; day++ {
    // Print the day, formatted to fit in 2 characters.
    s += fmt.Sprintf("%2d ", day)
    // Move to the next line after 7 days.
    i := day - 1
    if (firstDay + i) % 7 == daysInFirstWeek {
      s += fmt.Sprintln()
    }
  }

  return
}

