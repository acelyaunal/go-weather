package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

func supportArgs() (city string, days int, wantJSON bool) {
	days = 3
	wantJSON = false
	jsonFlag := flag.Bool("json", false, "json")
	daysFlag := flag.Int("days", 3, "days")
	flag.Parse()
	argsAfter := flag.Args()
	if len(argsAfter) > 0 {
		city = strings.Join(argsAfter, " ")
		days = *daysFlag
		wantJSON = *jsonFlag
	}
	raw := os.Args[1:]
	var cityParts []string
	skip := 0
	for i := 0; i < len(raw); i++ {
		if skip > 0 {
			skip--
			continue
		}
		a := raw[i]
		if a == "-days" && i+1 < len(raw) {
			if n, err := strconv.Atoi(raw[i+1]); err == nil {
				days = n
				skip = 1
				continue
			}
		}
		if a == "-json" {
			wantJSON = true
			continue
		}
		if !strings.HasPrefix(a, "-") {
			cityParts = append(cityParts, a)
		}
	}
	if len(cityParts) > 0 {
		city = strings.Join(cityParts, " ")
	}
	return
}
