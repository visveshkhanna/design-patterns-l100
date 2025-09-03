package main

import (
	"design-patterns/behavioural/chainofres/handlers"
	"design-patterns/behavioural/chainofres/models"
	"strings"
)

func main() {
	profanity := &handlers.ProfanityFilter{}
	spam := &handlers.SpamDetector{}
	limiter := &handlers.LengthLimiter{Limit: 24}
	audit := &handlers.AuditLogger{}

	profanity.SetNext(spam)
	spam.SetNext(limiter)
	limiter.SetNext(audit)

	msgs := []*models.Message{
		{Text: "Hello team, standup at 10"},
		{Text: "BUY NOW!!! Amazing DEAL!!!"},
		{Text: "This is a darn long message that should be trimmed"},
	}

	for _, m := range msgs {
		profanity.Handle(m)
		if m.Rejected {
			println("rejected:", m.Reason)
		} else {
			println("accepted:", m.Text, "flags:", strings.Join(m.Flags, ","))
		}
	}
}
