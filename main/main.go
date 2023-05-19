package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var year *int
var file *string

func init() {
	nowYear := time.Now().Year()
	year = flag.Int("year", nowYear, "Year")
	file = flag.String("file", fmt.Sprintf("%d.json", nowYear), "file path")
	flag.Parse()
}

func main() {
	if *year < 2019 {
		os.Exit(1)
	}
	RequestHolidayJson(*year, *file)
}
