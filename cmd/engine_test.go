package main

import (

  // this is a test.
  "testing"

  // prints.
  "fmt"

  // other imports.
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  inputString       string
  expectedString    string
}

type TestSuite struct {
  functionUnderTest func(string) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test DateAsHeader.
  */
  {
    DateAsHeader,
    []TestList{
      {"Date 1 digit day, leading space",    "2025-02-01", " 1. February, 2025"},
      {"Date 2 digit day, no leading space", "2025-12-12", "12. December, 2025"},
    },
  },
  /*
  * Test DateAsHeader.
  */
  {
    DateAsHeader,
    []TestList{
      {"Date 1 digit day, leading space",    "2025-02-01", " 1. February, 2025"},
      {"Date 2 digit day, no leading space", "2025-12-12", "12. December, 2025"},
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      exp := test.expectedString
      got := suite.functionUnderTest(test.inputString)
      if exp != got {
        t.Errorf("In '%s': Exptected '%#v', got actually '%#v'.\n", name, exp, got)
      }
    }
  }
}

// import ("time")
// parsedDate, err := time.Parse(layout, dateStr)
//     if err != nil {
//         fmt.Println("Error parsing date:", err)
//         return
//     }

//   actualArrays := DateAsArrays()
//   expectArrays := [][]string {
//     { "Mo", "Tu", "We", "Th", "Fr", "Sa", "Su" },
//     { "  ", "  ", " 1", " 2", " 3", " 4", " 5" },
//     { " 6", " 7", " 8", " 9", "10", "11", "12" },
//     { "13", "14", "15", "16", "17", "18", "19" },
//     { "20", "21", "22", "23", "24", "25", "26" },
//     { "27", "28", "  ", "  ", "  ", "  ", "  " },
//   }
//   actualBlob = DateAsBlob()
//   expectedBlob = "Mo Tu We Th Fr Sa Su" + NL +
//            "       1  2  3  4  5" + NL +
//            " 6  7  8  9 10 11 12" + NL +
//            "13 14 15 16 17 18 19" + NL +
//            "20 21 22 23 24 25 26" + NL +
//            "27 28               " + NL
// }

