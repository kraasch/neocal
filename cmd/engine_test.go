package main

import (

  // this is a test.
  "testing"

  // prints.
  "fmt"

  // other imports.
  //"strings"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  inputArr          []string
  expectedValue     string
}

type TestSuite struct {
  functionUnderTest func(...string) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the function TEMPLATE().
  */
  {
    nil,
    []TestList{
    },
  },
  /*
  * Test for the function DateAsHeader().
  */
  {
    func(in ...string) (out string) {
      targetDate := in[0]
      out = DateAsHeader(targetDate)
      return
    },
    []TestList{
      {"date-1-digit-day+leading-space",    []string{"2025-02-01"}, " 1. February, 2025"},
      {"date-2-digit-day+no-leading-space", []string{"2025-12-12"}, "12. December, 2025"},
    },
  },
  /*
  * Test for the function MonthAsCalendar().
  */
  {
    func(in ...string) (out string) {
      targetDate    := in[0]
      formatCulture := in[1]
      out = MonthAsCalendar(targetDate, formatCulture)
      return
    },
    []TestList{
      {
        testName:       "february-with-28-days",
        inputArr:       []string{"2025-02", "eu"},
        expectedValue:
        "Mo Tu We Th Fr Sa Su" + NL +
        "                1  2" + NL +
        " 3  4  5  6  7  8  9" + NL +
        "10 11 12 13 14 15 16" + NL +
        "17 18 19 20 21 22 23" + NL +
        "24 25 26 27 28      " + NL,
      },
      {
        testName:       "february-with-28-days",
        inputArr:       []string{"2025-02", "us"},
        expectedValue:
        "Su Mo Tu We Th Fr Sa" + NL +
        "                   1" + NL +
        " 2  3  4  5  6  7  8" + NL +
        " 9 10 11 12 13 14 15" + NL +
        "16 17 18 19 20 21 22" + NL +
        "23 24 25 26 27 28   " + NL,
      },
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      exp := test.expectedValue
      got := suite.functionUnderTest(test.inputArr...)
      if exp != got {
        t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
      }
    }
  }
}


