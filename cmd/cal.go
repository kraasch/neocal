
package main

import (
  // for making a nice centred box.
  tea "github.com/charmbracelet/bubbletea"
  lip "github.com/charmbracelet/lipgloss"

  // basics.
  "fmt"
  "os"
  "flag"

  // local packages.
  engine "github.com/kraasch/neocal/pkg/calengine"
  ctrl "github.com/kraasch/neocal/pkg/calcontroller"
)

var (
  // return value.
  output = ""
  // flags.
  verbose = false
  suppress = false
  // styles.
  styleBox = lip.NewStyle().
    BorderStyle(lip.NormalBorder()).
    BorderForeground(lip.Color("56"))
)

type model struct {
  width     int
  height    int
  c         ctrl.Controller
  startDate string
}

func (m model) Init() tea.Cmd {
  return func() tea.Msg { return nil }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
  case tea.KeyMsg:
    switch msg.String() {
    case ".", "t":
      m.c.Control("go", "start")
    case "h", "left":
      m.c.Control("prev", "day")
    case "l", "right":
      m.c.Control("next", "day")
    case "k", "up":
      m.c.Control("prev", "week")
    case "j", "down":
      m.c.Control("next", "week")
    case "K", "space":
      m.c.Control("prev", "month")
    case "J", "backpace":
      m.c.Control("next", "month")
    case "H", "pgup":
      m.c.Control("prev", "year")
    case "L", "pgdown":
      m.c.Control("next", "year")
    case "Q":
      suppress = true
      return m, tea.Quit
    case "q":
      return m, tea.Quit
    }
  }
  output = m.c.ReadDate() // update the return value.
  return m, cmd
}

func (m model) View() string {
  if m.width == 0 {
    return ""
  }
  r   := m.c.ReadDateHuman() + "\n"
  bgDay := []string{m.startDate}
  str := engine.HMonthAsCalendar(m.c.ReadDateYM(), "eu", m.c.ReadDate(), bgDay)
  r   += styleBox.Render(str)
  if verbose {
    bottomMsg := "\nQuit (q), move (hjklt)."
    r += bottomMsg + "\n"
  }
  return lip.Place(m.width, m.height, lip.Center, lip.Center, r)
}

func main() {

  // parse flags.
  flag.BoolVar(&verbose,  "verbose",  false, "Show info")
  flag.BoolVar(&suppress, "suppress", false, "Silence output")
  flag.Parse()

  // init model.
  cal, ok := ctrl.NewCalNow()
  if !ok {
    return
  }
  m := model{0, 0, cal, cal.ReadDate()}
  output = m.c.ReadDate() // initialize output value for CLI.

  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }

  // print the last highlighted value in calendar to stdout.
  if !suppress {
    fmt.Println(output)
  }

} // fin.

