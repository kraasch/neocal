
package calcontroller

import (
  // test.
  "testing"
  "github.com/kraasch/godiff/godiff"

  // other.
  //"time"
  //"fmt"
)

type TestList struct {
  testName          string
  isMulti           bool
  initialValue      string
  inputArr          []string
  expectedValue     string
}

type TestSuite struct {
  functionUnderTest func(string, ...string) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the function ReadDateHuman().
  */
  {
    functionUnderTest:
    func(setDate string, in ...string) string {
      c := Controller{}
      c.SetDate(setDate)
      return c.ReadDateHuman()
    },
    tests:
    []TestList{
      {
        testName:      "controller_next_00",
        isMulti:       false,
        initialValue:  "2025-02-04",
        inputArr:      nil,
        expectedValue: "4. February, 2025",
      },
      {
        testName:      "controller_next_00",
        isMulti:       false,
        initialValue:  "2000-01-31",
        inputArr:      nil,
        expectedValue: "31. January, 2000",
      },
      {
        testName:      "controller_next_00",
        isMulti:       false,
        initialValue:  "0290-01-31",
        inputArr:      nil,
        expectedValue: "31. January, 290",
      },
    },
  },
  /*
  * Test for the function ReadDateD().
  */
  {
    functionUnderTest:
    func(setDate string, in ...string) string {
      c := Controller{}
      ok := c.SetDate(setDate)
      if ok {
        action := ""
        for i, val := range in {
          unit := val // 2nd loop: treat val as unit.
          if i % 2 == 1 { // only trigger every 2nd loop.
            c.Control(action, unit)
          }
          action = val // 2nd loop + 1: treat val as new action for next loop.
        }
      }
      return c.ReadDateD()
    },
    tests:
    []TestList{
      {
        testName:      "controller_next_00",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "next", "day",
          "next", "year",
          "prev", "month",
        },
        expectedValue: "05",
      },
    },
  },
  /*
  * Test for the function ReadDateYM().
  */
  {
    functionUnderTest:
    func(setDate string, in ...string) string {
      c := Controller{}
      ok := c.SetDate(setDate)
      if ok {
        action := ""
        for i, val := range in {
          unit := val // 2nd loop: treat val as unit.
          if i % 2 == 1 { // only trigger every 2nd loop.
            c.Control(action, unit)
          }
          action = val // 2nd loop + 1: treat val as new action for next loop.
        }
      }
      return c.ReadDateYM()
    },
    tests:
    []TestList{
      {
        testName:      "controller_next_00",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "next", "day",
          "next", "year",
          "prev", "month",
        },
        expectedValue: "2025-12",
      },
    },
  },
  /*
  * Test for creating and manipulating dates with Controller.
  */
  {
    functionUnderTest:
    func(setDate string, in ...string) string {
      c := Controller{}
      ok := c.SetDate(setDate)
      if ok {
        action := ""
        for i, val := range in {
          unit := val // 2nd loop: treat val as unit.
          if i % 2 == 1 { // only trigger every 2nd loop.
            c.Control(action, unit)
          }
          action = val // 2nd loop + 1: treat val as new action for next loop.
        }
      }
      return c.ReadDate()
    },
    tests:
    []TestList{
      {
        testName:      "controller_next_00",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "next", "day",
        },
        expectedValue: "2025-01-05",
      },
      {
        testName:      "controller_next_01",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "next", "day",
          "next", "day",
        },
        expectedValue: "2025-01-06",
      },
      {
        testName:      "controller_next_02",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "next", "day",
          "next", "week",
        },
        expectedValue: "2025-01-12",
      },
      {
        testName:      "controller_next_03",
        isMulti:       false,
        initialValue:  "2000-01-01",
        inputArr:      []string{
          "next", "year",
          "next", "month",
          "next", "day",
        },
        expectedValue: "2001-02-02",
      },
      {
        testName:      "controller_prev_00",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "prev", "day",
        },
        expectedValue: "2025-01-03",
      },
      {
        testName:      "controller_prev_01",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "prev", "day",
          "prev", "day",
        },
        expectedValue: "2025-01-02",
      },
      {
        testName:      "controller_prev_02",
        isMulti:       false,
        initialValue:  "2025-01-04",
        inputArr:      []string{
          "prev", "day",
          "prev", "week",
        },
        expectedValue: "2024-12-27",
      },
      {
        testName:      "controller_prev_03",
        isMulti:       false,
        initialValue:  "2000-01-01",
        inputArr:      []string{
          "prev", "year",
          "prev", "month",
          "prev", "day",
        },
        expectedValue: "1998-11-30",
      },
      {
        testName:      "controller_04",
        isMulti:       false,
        initialValue:  "2000-01-01",
        inputArr:      []string{
          "next", "year",
          "next", "year",
          "prev", "month",
          "next", "week",
          "next", "month",
          "prev", "day", // all, but this one cancel each other.
          "prev", "year",
          "prev", "week",
          "prev", "year",
        },
        expectedValue: "1999-12-31",
      },
      {
        testName:      "leap-day_no_100-years_00",
        isMulti:       false,
        initialValue:  "1900-02-28",
        inputArr:      []string{
          "next", "day",
        },
        expectedValue: "1900-03-01",
      },
      {
        testName:      "leap-day_no_100-years_01",
        isMulti:       false,
        initialValue:  "2100-02-28",
        inputArr:      []string{
          "next", "day",
        },
        expectedValue: "2100-03-01",
      },
      {
        testName:      "leap-day_yes_all-4-years_00",
        isMulti:       false,
        initialValue:  "1904-02-28",
        inputArr:      []string{
          "next", "day",
        },
        expectedValue: "1904-02-29",
      },
      {
        testName:      "leap-day_yes_all-4-years_01",
        isMulti:       false,
        initialValue:  "1996-02-28",
        inputArr:      []string{
          "next", "day",
        },
        expectedValue: "1996-02-29",
      },
      {
        testName:      "leap-day_yes_all-400-years_00",
        initialValue:  "2000-02-28",
        inputArr:      []string{
          "next", "day",
        },
        expectedValue: "2000-02-29",
      },
      {
        testName:      "controller_order-matters_00",
        isMulti:       false,
        initialValue:  "2000-01-01",
        inputArr:      []string{
          "prev", "day",
          "prev", "month",
        },
        expectedValue: "1999-12-01",
      },
      {
        testName:      "controller_order-matters_01",
        isMulti:       false,
        initialValue:  "2000-01-01",
        inputArr:      []string{
          "prev", "month",
          "prev", "day",
        },
        expectedValue: "1999-11-30",
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
        got := suite.functionUnderTest(test.initialValue, test.inputArr...)
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

