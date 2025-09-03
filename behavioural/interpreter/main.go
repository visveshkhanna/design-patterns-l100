package main

import (
	"design-patterns/behavioural/interpreter/expressions"
	"design-patterns/behavioural/interpreter/parser"
	"fmt"
	"log"
)

func main() {
	accessRule := "(role==admin OR (beta AND region==EU)) AND requests<100"

	rule, err := parser.Parse(accessRule)
	if err != nil {
		log.Fatalf("Failed to parse access rule: %v", err)
	}

	fmt.Printf("Parsed access rule: %s\n\n", accessRule)

	ctxs := []expressions.Ctx{
		{Role: "admin", Region: "US", Beta: false, Requests: 12},
		{Role: "user", Region: "EU", Beta: true, Requests: 2},
		{Role: "user", Region: "EU", Beta: true, Requests: 120},
		{Role: "user", Region: "US", Beta: false, Requests: 5},
	}

	for i, c := range ctxs {
		result := rule.Eval(&c)
		fmt.Printf("Context %d: %+v -> Access: %t\n", i+1, c, result)
	}
}
