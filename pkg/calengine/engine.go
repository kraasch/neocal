
package calengine

import (
  "fmt"
  "time"
  "strconv"
)

const (
  F1 = "\x1b[1;38;2;255;0;0m" // ANSI foreground color (= red).
  B1 = "\x1b[48;5;56m"        // ANSI background color (= purple).
  N0 = "\x1b[0m"              // ANSI clear formatting.
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

// print month.
func MonthAsCalendar(targetDate string, culture string) (s string) {
  days := []string{}
  hls  := []string{}
  return hlMonthAsCalendar(targetDate, culture, days, hls)
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
      // TODO: check for input to be of valid length of valid format.... maybe. Or use time.Parse() with layout.
      y, erry := strconv.Atoi(days[i][:4])
      m, errm := strconv.Atoi(days[i][5:7])
      d, errd := strconv.Atoi(days[i][8:10])
      if erry != nil || errm != nil || errd != nil {
        _ = fmt.Errorf("Bad date conversion for days '%s': %#v, %#v, %#v", days, erry, errm, errd)
        return out
      }
      if y == targetYear && m == targetMonth {
        out[d] = append(out[d], highlights[i])
      }
    }
  }
  return out
}

func format(day int, highlights []string) (s string) {
  // If there is highlights, add them around the day string.
  for _, hl := range highlights {
    s += hl
  }
  s += fmt.Sprintf(" %2d", day)
  for range highlights {
    s += N0
  }
  return
}

// highlight days in month (with explicit highlights).
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
  targetY := firstDayDate.Year()
  targetM := int(firstDayDate.Month())
  hlsForEachDay := mergeHighlights(targetY, targetM, daysToHl, highlights)
  for day := 1; day <= lastDay; day++ {
    hlsForDay := hlsForEachDay[day]
    if len(hlsForDay) > 0 {
      s += format(day, hlsForDay)
    } else {
      s += fmt.Sprintf(" %2d", day)
    }
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

