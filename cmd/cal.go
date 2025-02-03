
package main

import (

  // local packages.
  "github.com/kraasch/neocal/pkg/calengine"

  // basics.
  "fmt"
  "os"
  "flag"

  // calculations.
  "time"
  // "math/rand"

  // for making a nice centred box.
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

var (
  // strings.
  top_msg = ""
  bot_msg = "\nQuit (q), move (hjkl)."
  // styles.
  styleBox = lipgloss.NewStyle().
  BorderStyle(lipgloss.NormalBorder()).
  BorderForeground(lipgloss.Color("56"))
  // flags.
  verbose = false
)

///// Add cursor formatting to current day formatting (ie fg and bg color).
// styleToday = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
// styleCursor = lipgloss.NewStyle().Background(lipgloss.Color("#404040"))
// str := styleCursor.Render(styleToday.Render("24")) // TODO: use.

type model struct {
  width    int
  height   int
  cursor_x int
  cursor_y int
  day      string
  month    string
  year     string
  content  string
}

func (m model) Init() tea.Cmd { return func() tea.Msg { return nil } }

// TODO: navigate the fields of the calendar array (within the model).
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
  case tea.KeyMsg:
    switch msg.String() {
      case "h", "left": // move left.
      if m.cursor_x > 0 {
        m.cursor_x--
      }
      case "l", "right": // move right.
      if m.cursor_x < 4 {
        m.cursor_x++
      }
      case "k", "up": // move up.
      if m.cursor_y > 0 {
        m.cursor_y--
      }
      case "j", "down": // move down.
      if m.cursor_y < 4 {
        m.cursor_y++
      }
    case "q":
      return m, tea.Quit
    }
  }
  return m, cmd
}

func (m model) View() string {
  if m.width == 0 {
    return ""
  }
  r := top_msg + "\n"
  r += styleBox.Render(m.content)
  if verbose {
    r += bot_msg + "\n"
  }
  return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, r)
}

func main() {

  // parse flags.
  flag.BoolVar(&verbose, "verbose", false, "Show info")
  flag.Parse()

  // some stuff.
  currentTime := time.Now() // Get the current time
  monthNum := int(currentTime.Month())
  month    := fmt.Sprint(currentTime.Month())
  year     := fmt.Sprint(currentTime.Year())
  day      := fmt.Sprint(currentTime.Day())

  // init model.
  currentMonth := fmt.Sprintf("%s-%02d", year, monthNum)
  culture      := "eu"
  str          := calengine.MonthAsCalendar(currentMonth, culture)
  m            := model{0, 0, 0, 0, day, month, year, str}

  // init variables.
  top_msg = fmt.Sprintf("%s. %s, %s", day, month, year)

  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }

} // fin.
