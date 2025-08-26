package main

import "strings"

func main() {
	profanity := &ProfanityFilter{}
	spam := &SpamDetector{}
	limiter := &LengthLimiter{Limit: 24}
	audit := &AuditLogger{}

	profanity.SetNext(spam)
	spam.SetNext(limiter)
	limiter.SetNext(audit)

	msgs := []*Message{
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
