
package calcontroller

import (
  "testing"
  "fmt"
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
      inputData := in[0]
      var c Controller
      ok := c.Control(inputData)
      if ok {
        out = "Wrote value: "
      }
      out += c.ReadState()
      return
    },
    tests: 
    []TestList{
      {
        testName:      "controller_create_00",
        isMulti:       false,
        inputArr:      []string{"42"},
        expectedValue: "Wrote value: 42",
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

