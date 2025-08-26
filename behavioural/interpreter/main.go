package main

import "fmt"

func main() {
	// Access Rule: (role==admin OR (beta AND region==EU)) AND requests<100
	rule := &And{
		left: &Or{
			left:  &IsRole{role: "admin"},
			right: &And{left: &IsBeta{}, right: &IsRegion{region: "EU"}},
		},
		right: &RequestsBelow{max: 100},
	}

	ctxs := []Ctx{
		{Role: "admin", Region: "US", Beta: false, Requests: 12},
		{Role: "user", Region: "EU", Beta: true, Requests: 2},
		{Role: "user", Region: "EU", Beta: true, Requests: 120},
		{Role: "user", Region: "US", Beta: false, Requests: 5},
	}

	for _, c := range ctxs {
		fmt.Println(rule.Eval(&c))
	}
}
