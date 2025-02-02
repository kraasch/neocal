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
  inputString       string
  expectedValue     interface{}
}

type TestSuite struct {
  functionUnderTest func(string) string
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
    DateAsHeader,
    []TestList{
      {"Date 1 digit day, leading space",    "2025-02-01", " 1. February, 2025"},
      {"Date 2 digit day, no leading space", "2025-12-12", "12. December, 2025"},
    },
  },
  /*
  * Test for the function MonthAsCalendar().
  */
  {
    MonthAsCalendar,
    []TestList{
      {
        testName:       "February with 28 days",
        inputString:    "2025-02",
        expectedValue:
        "Mo Tu We Th Fr Sa Su" + NL +
        "       1  2  3  4  5" + NL +
        " 6  7  8  9 10 11 12" + NL +
        "13 14 15 16 17 18 19" + NL +
        "20 21 22 23 24 25 26" + NL +
        "27 28               " + NL,
      },
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      exp := test.expectedValue
      got := suite.functionUnderTest(test.inputString)
      if exp != got {
        t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
      }
    }
  }
}


