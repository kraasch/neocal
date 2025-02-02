
package main

import (

  // other calculations.
  // "math/rand"

  // calendar's logic.
  "time"

  // command-line arguments.
  "flag"

  // print and exit.
  "fmt"
  "os"

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
  styleToday = lipgloss.NewStyle().
  Bold(true).
  Foreground(lipgloss.Color("#FF0000"))
  styleCursor = lipgloss.NewStyle().
  Background(lipgloss.Color("#404040"))
  // flags.
  verbose = false
)

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

// TODO: make this into a calendar array (within the model).
func makeMonthGrid(year int, month time.Month, today int) string {
  s := ""
  firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC) // Get the first day of the month.
  lastDay := firstDay.AddDate(0, 1, -1) // Get the total number of days in the month.
  s += fmt.Sprintln("Mo Tu We Th Fr Sa Su") // Add the header (day names).
  // Print leading spaces for the first day.
  for i := 0; i < int(firstDay.Weekday()); i++ {
    s += "  "
  }
  // Print the days of the month
  for day := 1; day <= lastDay.Day(); day++ {
    if day == today {
      // TODO: add cursor formatting to current day formatting (ie fg and bg color).
      a := fmt.Sprint(day)
      b := styleCursor.Render(styleToday.Render(a))
      s += fmt.Sprintf("%2s ", b)
    } else {
      s += fmt.Sprintf("%2d ", day) // Print the day, formatted to fit in 2 characters.
    }
    if (firstDay.Day()+day)%7 == 6 { // Move to the next line after 7 days.
      s += fmt.Sprintln()
    }
  }
  // Return.
  return s
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
  month := fmt.Sprint(currentTime.Month())
  year  := fmt.Sprint(currentTime.Year())
  day   := fmt.Sprint(currentTime.Day())

  // init model.
  str := makeMonthGrid(int(currentTime.Year()), currentTime.Month(), int(currentTime.Day()))
  m := model{0, 0, 0, 0, day, month, year, str}

  // init variables.
  top_msg = fmt.Sprintf("%s. %s, %s", day, month, year)

  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }

} // fin.
