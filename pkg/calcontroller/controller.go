
package calcontroller

import (
  "fmt"
  "strconv"
)

type Controller struct {
  state int
}

func (c *Controller) Control(in string) (ok bool) {
  val, err := strconv.Atoi(in)
  if err == nil {
    ok = true
  }
  c.state = val
  return
}

func (c *Controller) ReadState() (out string) {
  out = fmt.Sprintf("%d", c.state)
  return
}

