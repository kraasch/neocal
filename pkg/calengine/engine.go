
package calengine

import (
  "fmt"
  "time"
  "strconv"
)

const (
  F1 = "\x1b[1;38;2;255;0;0m" // ANSI foreground color (= red).
  B1 = "\x1b[38;5;56m"        // ANSI background color (= purple).
  N0 = "\x1b[0m"              // ANSI clear formatting.
  // TL = "┌" // top left corner.
  // TR = "┐" // top right corner.
  // BL = "└" // bottom left corner.
  // BR = "┘" // bottom right corner.
  // HO = "─" // horizontal line.
  // VE = "│" // vertical line.
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

func MonthAsCalendar(targetDate string, culture string) string {
  return CMonthAsCalendar(targetDate, culture, "")
}

// color day in month.
func CMonthAsCalendar(targetDate string, culture string, dayToHighlight string) (s string) {
  days := []string{dayToHighlight}
  hls  := []string{F1}
  return hlMonthAsCalendar(targetDate, culture, days, hls)
}

// highlight days in month (without explicity highlights).
func HMonthAsCalendar(targetDate string, culture string, dayToFg string, daysToBg []string) (s string) {
  daysToHl := []string{dayToFg}
  hls      := []string{F1}
  for _, day := range daysToBg {
    daysToHl = append(daysToHl, day)
    hls      = append(hls,      B1)
  }
  return hlMonthAsCalendar(targetDate, culture, daysToHl, hls)
}

func mergeHighlights(targetYear int, targetMonth int, days []string, highlights []string) map[int][]string {
  out := make(map[int][]string)
  if len(days) == len(highlights) {
    for i := range days {
      y, _ := strconv.Atoi(days[i][:4])
      m, _ := strconv.Atoi(days[i][5:7])
      d, _ := strconv.Atoi(days[i][8:10])
      if y == targetYear && m == targetMonth {
        out[d] = append(out[d], highlights[i])
      }
    }
  }
  return out
}

func format(day int, dayToHighlight string, highlight string) (s string) {
  theDay := 0
  if dayToHighlight != "" {
    converted, err := strconv.Atoi(dayToHighlight)
    if err != nil {
      panic(err)
    }
    theDay = converted
  }

  if day == theDay {
    s += highlight
  }
  s += fmt.Sprintf(" %2d", day)
  if day == theDay {
    s += N0
  }
  return
}

// highlight days in month (with explicity highlights).
func hlMonthAsCalendar(targetDate string, culture string, daysToHl []string, highlights []string) (s string) {

  // Get the first day of the target month.
  firstDayDate := parseYearAndMonth(targetDate)
  firstDay := firstDayDate.Day()

  // Get the total number of days in the month.
  lastDayDate := firstDayDate.AddDate(0, 1, -1)
  lastDay := lastDayDate.Day()
  lastWeekday := int(lastDayDate.Weekday())
  if culture == "us" {
    // turn EU format: 6=sun 0=mon 1=tue 2=wed 3=thu 4=fri 5=sat 6=sun
    // into US format: 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat 0=sun
    lastWeekday = (lastWeekday + 1) % 7
  }

  // Get the weekday of the first day in the month.
  // US format, week starts with Sunday:
  // 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat.
  weekday := int(firstDayDate.Weekday())
  if culture == "eu" {
    // turn US format: 0=sun 1=mon 2=tue 3=wed 4=thu 5=fri 6=sat 0=sun
    // into EU format: 6=sun 0=mon 1=tue 2=wed 3=thu 4=fri 5=sat 6=sun
    weekday = (weekday + 6) % 7
  }

  // Add the header (day names).
  if culture == "us" {
    s += fmt.Sprintln(" Su Mo Tu We Th Fr Sa ")
  } else {
    s += fmt.Sprintln(" Mo Tu We Th Fr Sa Su ")
  }

  // Print leading spaces for the
  for i := 0; i < weekday; i++ {
    s += "   "
  }
  daysInFirstWeek := (7 - weekday) % 7

  // Print the days of the month.
  for day := 1; day <= lastDay; day++ {
    s += format(day, daysToHl[0], F1)
    i := day - 1
    if (firstDay + i) % 7 == daysInFirstWeek {
      // Move to the next line after 7 days.
      s += " " // Add right side padding.
      s += fmt.Sprintln()
    }
  }

  // Print trailing spaces.
  for i := lastWeekday; i < 7; i++ {
    s += "   "
  }
  s += " " // Add right side padding.

  return
}

