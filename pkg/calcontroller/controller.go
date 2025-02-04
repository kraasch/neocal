
package calcontroller

import (
  "time"
)

type Controller struct {
  cal time.Time
}

func (c *Controller) SetDate(in string) (ok bool) {
  t, err := time.Parse(time.DateOnly, in)
  c.cal = t
  if err == nil {
    ok = true
  }
  return
}

func (c *Controller) Control(action string, unit string) (ok bool) {
  ok = true // default to true.

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

