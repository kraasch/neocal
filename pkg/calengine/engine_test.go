package calengine

import (

	// this is a test.
	"regexp"
	"testing"

	// printing and formatting.
	"fmt"
	"strings"

	// other imports.
	"github.com/kraasch/godiff/godiff"
)

type TestList struct {
	testName      string
	isMulti       bool
	inputArr      []string
	inputArr2     []string
	expectedValue string
}

type TestSuite struct {
	testingFunction func(in TestList) string
	tests           []TestList
}

func removeAnsii(input string) string {
	// Regular expression to match ANSI escape sequences
	re := regexp.MustCompile(`\x1b\[[0-9;?]*[A-Za-z]`)
	return re.ReplaceAllString(input, "")
}

var suites = []TestSuite{
	/*
	 * Test for the function DateAsHeader().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			out = DateAsHeader(targetDate)
			return out
		},
		tests: []TestList{
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
	 * Test for the function MonthAsThree().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			hlDay := targetDate + "-01"
			formatCulture := in.inputArr[1]
			formatStyle := "week"
			out = ThreeMonthAsCalendar(targetDate, formatCulture, hlDay, []string{}, formatStyle)
			out = removeAnsii(out)
			return out
		},
		tests: []TestList{
			{
				testName: "three_calendar_eu_week-starts-sun_00", // start month with last day of week.
				isMulti:  true,
				inputArr: []string{"2024-12", "eu"},
				expectedValue: "  | Mo Tu We Th Fr Sa Su " + NL +
					"44| 28 29 30 31  1  2  3 " + NL +
					"45|  4  5  6  7  8  9 10 " + NL +
					"46| 11 12 13 14 15 16 17 " + NL +
					"47| 18 19 20 21 22 23 24 " + NL +
					"48| 25 26 27 28 29 30  1 " + NL +
					"49|  2  3  4  5  6  7  8 " + NL +
					"50|  9 10 11 12 13 14 15 " + NL +
					"51| 16 17 18 19 20 21 22 " + NL +
					"52| 23 24 25 26 27 28 29 " + NL +
					"53| 30 31  1  2  3  4  5 " + NL +
					" 2|  6  7  8  9 10 11 12 " + NL +
					" 3| 13 14 15 16 17 18 19 " + NL +
					" 4| 20 21 22 23 24 25 26 " + NL +
					" 5| 27 28 29 30 31  1  2 ",
			},
		},
	},

	/*
	 * Test for the function MonthAsCalendar().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			formatCulture := in.inputArr[1]
			fillStyle := in.inputArr[2]
			formatStyle := "week"
			out = MonthAsCalendar(targetDate, formatCulture, fillStyle, formatStyle)
			return out
		},
		tests: []TestList{
			{
				testName: "weeknums_calendar_eu_week-starts-sun_00", // start month with last day of week.
				isMulti:  true,
				inputArr: []string{"2024-12", "eu", "none"},
				expectedValue: "  | Mo Tu We Th Fr Sa Su " + NL +
					"48|                    1 " + NL +
					"49|  2  3  4  5  6  7  8 " + NL +
					"50|  9 10 11 12 13 14 15 " + NL +
					"51| 16 17 18 19 20 21 22 " + NL +
					"52| 23 24 25 26 27 28 29 " + NL +
					"53| 30 31                ",
			},
			{
				testName: "weeknums_calendar_us_week-starts-mid_00", // start month with middle of week.
				isMulti:  true,
				inputArr: []string{"2030-01", "us", "none"},
				expectedValue: "  | Su Mo Tu We Th Fr Sa " + NL +
					" 1|        1  2  3  4  5 " + NL +
					" 2|  6  7  8  9 10 11 12 " + NL +
					" 3| 13 14 15 16 17 18 19 " + NL +
					" 4| 20 21 22 23 24 25 26 " + NL +
					" 5| 27 28 29 30 31       ",
			},
			{
				testName: "weeknums_calendar_us_week-starts-sun_00", // start month with first day of week.
				isMulti:  true,
				inputArr: []string{"2024-12", "us", "none"},
				expectedValue: "  | Su Mo Tu We Th Fr Sa " + NL +
					"49|  1  2  3  4  5  6  7 " + NL +
					"50|  8  9 10 11 12 13 14 " + NL +
					"51| 15 16 17 18 19 20 21 " + NL +
					"52| 22 23 24 25 26 27 28 " + NL +
					"53| 29 30 31             ",
			},
			{
				testName: "weeknums_calendar_us_week-starts-sat_00", // start month with last day of week.
				isMulti:  true,
				inputArr: []string{"2025-02", "us", "none"},
				expectedValue: "  | Su Mo Tu We Th Fr Sa " + NL +
					" 5|                    1 " + NL +
					" 6|  2  3  4  5  6  7  8 " + NL +
					" 7|  9 10 11 12 13 14 15 " + NL +
					" 8| 16 17 18 19 20 21 22 " + NL +
					" 9| 23 24 25 26 27 28    ",
			},
			{
				testName: "weeknums_calendar_eu_month-ends-with-sun_00",
				isMulti:  true,
				inputArr: []string{"2026-05", "eu", "line"},
				expectedValue: "  | Mo Tu We Th Fr Sa Su " + NL +
					"18| 27 28 29 30  1  2  3 " + NL +
					"19|  4  5  6  7  8  9 10 " + NL +
					"20| 11 12 13 14 15 16 17 " + NL +
					"21| 18 19 20 21 22 23 24 " + NL +
					"22| 25 26 27 28 29 30 31 " + NL +
					"  |  1  2  3  4  5  6  7 ",
			},
			{
				testName: "weeknums_calendar_eu_month-ends-not-with-sun_00",
				isMulti:  true,
				inputArr: []string{"2025-02", "eu", "line"},
				expectedValue: "  | Mo Tu We Th Fr Sa Su " + NL +
					" 5| 27 28 29 30 31  1  2 " + NL +
					" 6|  3  4  5  6  7  8  9 " + NL +
					" 7| 10 11 12 13 14 15 16 " + NL +
					" 8| 17 18 19 20 21 22 23 " + NL +
					" 9| 24 25 26 27 28  1  2 ",
			},
		},
	},

	/*
	 * Test for the function MonthAsCalendar().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			formatCulture := in.inputArr[1]
			fillStyle := "none"
			formatStyle := "none"
			out = MonthAsCalendar(targetDate, formatCulture, fillStyle, formatStyle)
			return out
		},
		tests: []TestList{
			{
				testName: "calendar_eu_week-starts-sun_00", // start month with last day of week.
				isMulti:  true,
				inputArr: []string{"2024-12", "eu"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					"                    1 " + NL +
					"  2  3  4  5  6  7  8 " + NL +
					"  9 10 11 12 13 14 15 " + NL +
					" 16 17 18 19 20 21 22 " + NL +
					" 23 24 25 26 27 28 29 " + NL +
					" 30 31                ",
			},
			{
				testName: "calendar_us_week-starts-mid_00", // start month with middle of week.
				isMulti:  true,
				inputArr: []string{"2030-01", "us"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					"        1  2  3  4  5 " + NL +
					"  6  7  8  9 10 11 12 " + NL +
					" 13 14 15 16 17 18 19 " + NL +
					" 20 21 22 23 24 25 26 " + NL +
					" 27 28 29 30 31       ",
			},
			{
				testName: "calendar_us_week-starts-sun_00", // start month with first day of week.
				isMulti:  true,
				inputArr: []string{"2024-12", "us"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					"  1  2  3  4  5  6  7 " + NL +
					"  8  9 10 11 12 13 14 " + NL +
					" 15 16 17 18 19 20 21 " + NL +
					" 22 23 24 25 26 27 28 " + NL +
					" 29 30 31             ",
			},
			{
				testName: "calendar_us_week-starts-sat_00", // start month with last day of week.
				isMulti:  true,
				inputArr: []string{"2025-02", "us"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					"                    1 " + NL +
					"  2  3  4  5  6  7  8 " + NL +
					"  9 10 11 12 13 14 15 " + NL +
					" 16 17 18 19 20 21 22 " + NL +
					" 23 24 25 26 27 28    ",
			},
			{
				testName: "design_eu_month-ends-with-31-days_00",
				isMulti:  true,
				inputArr: []string{"2026-05", "eu"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					"              1  2  3 " + NL +
					"  4  5  6  7  8  9 10 " + NL +
					" 11 12 13 14 15 16 17 " + NL +
					" 18 19 20 21 22 23 24 " + NL +
					" 25 26 27 28 29 30 31 " + NL +
					// " 25 26 27 28 29 30 31 ", // NOTE: see below.
					"                      ",
				// NOTE: this could also end without a new line.
				// ie. " 25 26 27 28 29 30 31 " + NL,
				// But even traditional CLI calendars print a line full of spaces here.
			},
		},
	},

	/*
	 * Test for the function CMonthAsCalendar().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			formatCulture := in.inputArr[1]
			dayToHighlight := in.inputArr[2]
			out = CMonthAsCalendar(targetDate, formatCulture, dayToHighlight, "none", "none")
			return out
		},
		tests: []TestList{
			{
				testName: "color_february-with-28-days_eu_00",
				isMulti:  true,
				inputArr: []string{"2025-02", "eu", "2025-02-15"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					"                 1  2 " + NL +
					"  3  4  5  6  7  8  9 " + NL +
					" 10 11 12 13 14 " +
					F1 + "15" + N0 +
					" 16 " + NL +
					" 17 18 19 20 21 22 23 " + NL +
					" 24 25 26 27 28       ",
			},
			{
				testName: "color_february-with-28-days_us_00",
				isMulti:  true,
				inputArr: []string{"2025-02", "us", "2025-02-15"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					"                    1 " + NL +
					"  2  3  4  5  6  7  8 " + NL +
					"  9 10 11 12 13 14 " +
					F1 + "15" + N0 + " " + NL +
					" 16 17 18 19 20 21 22 " + NL +
					" 23 24 25 26 27 28    ",
			},
		},
	},

	/*
	 * Test for the function mergeHighlights().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			dd := in.inputArr
			hh := in.inputArr2
			targetMonth := 2
			targetYear := 2025
			highlightMap := mergeHighlights(targetYear, targetMonth, dd, hh)

			// TODO: test whole map: sort and convert to string.
			out += " 7 => " + strings.Join(highlightMap[7], ", ") + "." + NL
			out += fmt.Sprintf("11 => %s\n", highlightMap[11])
			out += "17 => " + strings.Join(highlightMap[17], ", ") + "." + NL
			out += fmt.Sprintf("22 => %s\n", highlightMap[22])

			return out
		},
		tests: []TestList{
			{
				testName:  "merge-highlights_00",
				inputArr:  []string{"2025-02-07", "2025-02-07", "2025-02-17", "2025-01-11", "2024-02-22"},
				inputArr2: []string{F1, B1, B1, F1, B1},
				expectedValue: " 7 => " + F1 + ", " + B1 + "." + NL +
					"11 => []" + NL +
					"17 => " + B1 + "." + NL +
					"22 => []" + NL,
			},
		},
	},

	/*
	 * Test for the function HMonthAsCalendar().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			formatCulture := in.inputArr[1]
			dayToFg := in.inputArr[2]
			daysToBg := in.inputArr2
			out = HMonthAsCalendar(targetDate, formatCulture, dayToFg, daysToBg, "none", "none", false)
			return out
		},
		tests: []TestList{
			{
				testName:  "highlight+color_eu_00",
				isMulti:   true,
				inputArr:  []string{"2025-02", "eu", "2025-02-15"},
				inputArr2: []string{"2025-02-02", "2025-03-03"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					"                 1 " + B1 + " 2" + N0 + " " + NL +
					"  3  4  5  6  7  8  9 " + NL +
					" 10 11 12 13 14 " + F1 + "15" + N0 + " 16 " + NL +
					" 17 18 19 20 21 22 23 " + NL +
					" 24 25 26 27 28       ",
			},
			{
				testName:  "highlight+color_eu_01",
				isMulti:   true,
				inputArr:  []string{"2025-02", "us", "2025-02-14"},
				inputArr2: []string{"2025-02-27", "2025-02-28", "2025-03-01"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					"                    1 " + NL +
					"  2  3  4  5  6  7  8 " + NL +
					"  9 10 11 12 13 " + F1 + "14" + N0 + " 15 " + NL +
					" 16 17 18 19 20 21 22 " + NL +
					" 23 24 25 26 " + B1 + "27" + N0 + " " + B1 + "28" + N0 + "    ",
			},
			{
				testName:  "highlight+color_eu_02",
				isMulti:   true,
				inputArr:  []string{"2025-02", "eu", "2025-02-15"},
				inputArr2: []string{"2025-02-15"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					"                 1  2 " + NL +
					"  3  4  5  6  7  8  9 " + NL +
					" 10 11 12 13 14 " + F1 + B1 + "15" + N0 + N0 + " 16 " + NL +
					" 17 18 19 20 21 22 23 " + NL +
					" 24 25 26 27 28       ",
			},
		},
	},

	/*
	 * Test for the function MonthAsCalendar() and fill each week.
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			formatCulture := in.inputArr[1]
			fillStyle := in.inputArr[2]
			formatStyle := "none"
			out = MonthAsCalendar(targetDate, formatCulture, fillStyle, formatStyle)
			return out
		},
		tests: []TestList{
			{
				testName: "fill-calendar_eu_week-starts-mid_00",
				isMulti:  true,
				inputArr: []string{"2025-02", "eu", "line"}, // start month with middle of week.
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					" 27 28 29 30 31  1  2 " + NL +
					"  3  4  5  6  7  8  9 " + NL +
					" 10 11 12 13 14 15 16 " + NL +
					" 17 18 19 20 21 22 23 " + NL +
					" 24 25 26 27 28  1  2 ",
			},
			{
				testName: "fill-calendar_eu_week-starts-mon_00", // start month with first day of week.
				isMulti:  true,
				inputArr: []string{"2003-09", "eu", "line"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					"  1  2  3  4  5  6  7 " + NL +
					"  8  9 10 11 12 13 14 " + NL +
					" 15 16 17 18 19 20 21 " + NL +
					" 22 23 24 25 26 27 28 " + NL +
					" 29 30  1  2  3  4  5 ",
			},
			{
				testName: "fill-calendar_us_week-starts-sun_00", // start month with first day of week.
				isMulti:  true,
				inputArr: []string{"2024-12", "us", "line"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					"  1  2  3  4  5  6  7 " + NL +
					"  8  9 10 11 12 13 14 " + NL +
					" 15 16 17 18 19 20 21 " + NL +
					" 22 23 24 25 26 27 28 " + NL +
					" 29 30 31  1  2  3  4 ",
			},
			{
				testName: "fill-calendar_us_week-starts-sat_00", // start month with last day of week.
				isMulti:  true,
				inputArr: []string{"2025-02", "us", "line"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					" 26 27 28 29 30 31  1 " + NL +
					"  2  3  4  5  6  7  8 " + NL +
					"  9 10 11 12 13 14 15 " + NL +
					" 16 17 18 19 20 21 22 " + NL +
					" 23 24 25 26 27 28  1 ",
			},
		},
	},

	/*
	 * Test for the function CMonthAsCalendar() and fill each week.
	 */
	{
		testingFunction: func(in TestList) (out string) {
			targetDate := in.inputArr[0]
			formatCulture := in.inputArr[1]
			dayToHighlight := in.inputArr[2]
			fillStyle := in.inputArr[3]
			out = CMonthAsCalendar(targetDate, formatCulture, dayToHighlight, fillStyle, "none")
			return out
		},
		tests: []TestList{
			{
				testName: "color_february-with-28-days_eu_00",
				isMulti:  true,
				inputArr: []string{"2025-02", "eu", "2025-02-15", "line"},
				expectedValue: " Mo Tu We Th Fr Sa Su " + NL +
					B2 + " 27 28 29 30 31" + N0 + "  1  2 " + NL +
					"  3  4  5  6  7  8  9 " + NL +
					" 10 11 12 13 14 " +
					F1 + "15" + N0 +
					" 16 " + NL +
					" 17 18 19 20 21 22 23 " + NL +
					" 24 25 26 27 28" + B2 + "  1  2" + N0 + " ",
			},
			{
				testName: "color_february-with-28-days_us_00",
				isMulti:  true,
				inputArr: []string{"2025-02", "us", "2025-02-15", "line"},
				expectedValue: " Su Mo Tu We Th Fr Sa " + NL +
					B2 + " 26 27 28 29 30 31" + N0 + "  1 " + NL +
					"  2  3  4  5  6  7  8 " + NL +
					"  9 10 11 12 13 14 " +
					F1 + "15" + N0 + " " + NL +
					" 16 17 18 19 20 21 22 " + NL +
					" 23 24 25 26 27 28" + B2 + "  1" + N0 + " ",
			},
		},
	},
} // Fin of test suite.

func TestAll(t *testing.T) {
	for _, suite := range suites {
		for _, test := range suite.tests {
			name := test.testName
			t.Run(name, func(t *testing.T) {
				exp := test.expectedValue
				got := suite.testingFunction(test)
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
