package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	flag.Usage = usage

	var (
		date time.Time
	)

	flag.Parse()
	if flag.NArg() == 1 {
		t := flag.Arg(0)

		d, err := time.Parse(time.DateOnly, t)
		if err != nil {
			fmt.Printf("invalid date string: %s\n", t)

			flag.Usage()
			return
		}
		date = d
	} else {
		date = time.Now()
	}

	fmt.Print(formatDate(date))
}

func formatDate(t time.Time) string {
	_, weekNumber := t.ISOWeek()

	var s strings.Builder
	fmt.Fprintln(&s, t.Format(time.DateOnly))
	fmt.Fprintf(&s, "week: %d\n", weekNumber)

	return s.String()
}

func usage() {
	var s strings.Builder

	fmt.Fprintf(&s, "Usage:\n")
	fmt.Fprintf(&s, "  %s", "ukenr\n")
	fmt.Fprintf(&s, "  %s", "ukenr [yyyy-mm-dd]\n")

	fmt.Print(s.String())
}
