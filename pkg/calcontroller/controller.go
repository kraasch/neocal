
package calcontroller

import (
  "time"
  "fmt"
)

type Controller struct {
  cal      time.Time
  startCal time.Time
}

func NewCalNow() (c Controller, ok bool) {
  ok = c.SetDate(time.Now().Format(time.DateOnly))
  return
}

func (c *Controller) SetDate(in string) (ok bool) {
  t, err := time.Parse(time.DateOnly, in)
  c.cal = t
  c.startCal = t
  if err == nil {
    ok = true
  }
  return
}

func (c *Controller) Control(action string, unit string) (ok bool) {
  ok = true // default to true.

  // interpret special actions and units first.
  if action == "go" && unit == "start" {
    c.cal = c.startCal
    return
  }

  // evaluate action.
  dir := 0
  if action == "next" {
    dir = 1
  } else if action == "prev" {
    dir = -1
  } else {
    ok = false
    return
  }

  // evaluate unit.
  days   := 0
  months := 0
  years  := 0
  if unit == "day" {
    days = 1
  } else if unit == "week" {
    days = 7
  } else if unit == "month" {
    months = 1
  } else if unit == "year" {
    years = 1
  } else {
    ok = false
    return
  }

  // update date according to aciton.
  c.cal = c.cal.AddDate(years * dir, months * dir, days * dir)

  return
}

func (c *Controller) ReadDate() (out string) {
  out = c.cal.Format(time.DateOnly)
  return
}

func (c *Controller) ReadDateYM() (out string) {
  out = c.ReadDate()[:7]
  return
}

func (c *Controller) ReadDateD() (out string) {
  str := c.ReadDate()
  out = str[len(str)-2:]
  return
}

func (c *Controller) ReadDateHuman() (out string) {
  month := fmt.Sprint(c.cal.Month())
  year  := fmt.Sprint(c.cal.Year())
  day   := fmt.Sprint(c.cal.Day())
  out    = fmt.Sprintf("%s. %s, %s", day, month, year)
  return
}


