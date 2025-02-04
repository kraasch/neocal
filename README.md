
# A command-line calendar

Prints the current month with the current day highlighted.
Similar to `cal`, `ncal`, etc.

<p align="center">
  <img src="./resources/example.png" width="300"/>
</p>

## to-dos

 - [ ] make main return a date string (format: time.DateOnly, ie '2006-01-02').
 - [ ] test for negative years, eg year -1 (do not let user navigate into pre-historic past).
 - [ ] test for 5-digit years, eg year 10000 (do not let user navigate in post-historic future).
 - [ ] auto-update GUI when next day came.
 - [ ] extract the calendar array into the model.

Later/maybe:

 - [ ] have interactive events.
   - [ ] let user associated input forms with certain days of the calendar.
   - [ ] show interactive events on days (eg blue font highlight).
   - [ ] activate interactive events with special key (eg enter/space).
 - [ ] have script events.
   - [ ] let user associated script with certain days of the calendar.
   - [ ] show script events on days (eg blue font highlight).
   - [ ] activate script events with special key (eg enter/space).
   - [ ] run the associated script.

## done

 - [X] make the model navigable via vim home row keys (hjkl).
   - [X] go to next month when leaving left or right.

## misc info

For formatting [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) were used.
