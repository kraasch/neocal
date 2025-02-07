
package calengine

import (

  // this is a test.
  "testing"

  // prints.
  "fmt"

  // other imports.
  // "strings"
  "github.com/kraasch/godiff/godiff"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  isMulti           bool
  inputArr          []string
  expectedValue     string
}

type TestSuite struct {
  functionUnderTest func(...string) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the function DateAsHeader().
  */
  {
    functionUnderTest:
    func(in ...string) (out string) {
      targetDate := in[0]
      out = DateAsHeader(targetDate)
      return
    },
    tests: 
    []TestList{
      {
        testName:      "date-1-digit-day_leading-space_00",
        isMulti:       false,
        inputArr:      []string{"2025-02-01"},
        expectedValue: " 1. February, 2025",
      },
      {
        testName:      "date-2-digit-day_no-leading-space_00",
        isMulti:       false,
        inputArr:      []string{"2025-12-12"},
        expectedValue: "12. December, 2025",
      },
    },
  },
  /*
  * Test for the function MonthAsCalendar().
  */
  {
    functionUnderTest:
    func(in ...string) (out string) {
      targetDate    := in[0]
      formatCulture := in[1]
      out = MonthAsCalendar(targetDate, formatCulture)
      return
    },
    tests: 
    []TestList{
      {
        testName:       "calendar_eu_week-starts-mid_00",
        isMulti:        true,
        inputArr:       []string{"2025-02", "eu"}, // start month with middle of week.
        expectedValue:
        " Mo Tu We Th Fr Sa Su " + NL +
        "                 1  2 " + NL +
        "  3  4  5  6  7  8  9 " + NL +
        " 10 11 12 13 14 15 16 " + NL +
        " 17 18 19 20 21 22 23 " + NL +
        " 24 25 26 27 28       ",
      },
      {
        testName:       "calendar_eu_week-starts-mon_00", // start month with first day of week.
        isMulti:        true,
        inputArr:       []string{"2003-09", "eu"},
        expectedValue:
        " Mo Tu We Th Fr Sa Su " + NL +
        "  1  2  3  4  5  6  7 " + NL +
        "  8  9 10 11 12 13 14 " + NL +
        " 15 16 17 18 19 20 21 " + NL +
        " 22 23 24 25 26 27 28 " + NL +
        " 29 30                ",
      },
      {
        testName:       "calendar_eu_week-starts-sun_00", // start month with last day of week.
        isMulti:        true,
        inputArr:       []string{"2024-12", "eu"},
        expectedValue:
        " Mo Tu We Th Fr Sa Su " + NL +
        "                    1 " + NL +
        "  2  3  4  5  6  7  8 " + NL +
        "  9 10 11 12 13 14 15 " + NL +
        " 16 17 18 19 20 21 22 " + NL +
        " 23 24 25 26 27 28 29 " + NL +
        " 30 31                ",
      },
      {
        testName:       "calendar_us_week-starts-mid_00", // start month with middle of week.
        isMulti:        true,
        inputArr:       []string{"2030-01", "us"},
        expectedValue:
        " Su Mo Tu We Th Fr Sa " + NL +
        "        1  2  3  4  5 " + NL +
        "  6  7  8  9 10 11 12 " + NL +
        " 13 14 15 16 17 18 19 " + NL +
        " 20 21 22 23 24 25 26 " + NL +
        " 27 28 29 30 31       ",
      },
      {
        testName:       "calendar_us_week-starts-sun_00", // start month with first day of week.
        isMulti:        true,
        inputArr:       []string{"2024-12", "us"},
        expectedValue:
        " Su Mo Tu We Th Fr Sa " + NL +
        "  1  2  3  4  5  6  7 " + NL +
        "  8  9 10 11 12 13 14 " + NL +
        " 15 16 17 18 19 20 21 " + NL +
        " 22 23 24 25 26 27 28 " + NL +
        " 29 30 31             ",
      },
      {
        testName:       "calendar_us_week-starts-sat_00", // start month with last day of week.
        isMulti:        true,
        inputArr:       []string{"2025-02", "us"},
        expectedValue:
        " Su Mo Tu We Th Fr Sa " + NL +
        "                    1 " + NL +
        "  2  3  4  5  6  7  8 " + NL +
        "  9 10 11 12 13 14 15 " + NL +
        " 16 17 18 19 20 21 22 " + NL +
        " 23 24 25 26 27 28    ",
      },
    },
  },
  /*
  * Test for the function MonthAsCalendar().
  */
  {
    functionUnderTest:
    func(in ...string) (out string) {
      targetDate    := in[0]
      formatCulture := in[1]
      dayInMonth    := in[2]
      out = CMonthAsCalendar(targetDate, formatCulture, dayInMonth)
      return
    },
    tests: 
    []TestList{
      {
        testName:       "color_february-with-28-days_eu_00",
        isMulti:        true,
        inputArr:       []string{"2025-02", "eu", "15"},
        expectedValue:
        " Mo Tu We Th Fr Sa Su " + NL +
        "                 1  2 " + NL +
        "  3  4  5  6  7  8  9 " + NL +
        " 10 11 12 13 14" +
                   F1 + " 15" + N0 +
                          " 16 " + NL +
        " 17 18 19 20 21 22 23 " + NL +
        " 24 25 26 27 28       ",
      },
      {
        testName:       "color_february-with-28-days_us_00",
        isMulti:        true,
        inputArr:       []string{"2025-02", "us", "15"},
        expectedValue:
        " Su Mo Tu We Th Fr Sa " + NL +
        "                    1 " + NL +
        "  2  3  4  5  6  7  8 " + NL +
        "  9 10 11 12 13 14" +
                      F1 + " 15" + N0 + " " + NL +
        " 16 17 18 19 20 21 22 " + NL +
        " 23 24 25 26 27 28    ",
      },
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      t.Run(name, func(t *testing.T) {
        exp := test.expectedValue
        got := suite.functionUnderTest(test.inputArr...)
        if exp != got {
          if test.isMulti {
            t.Errorf("In '%s':\n", name)
            diff := godiff.CDiff(exp, got)
            t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
            t.Errorf("exp/got:\n%s\n", diff)
          } else {
            t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
          }
        }
      })
    }
  }
}

