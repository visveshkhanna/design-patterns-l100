package main

import "strings"

// ProfanityFilter rejects messages containing banned words.
type ProfanityFilter struct{ BaseHandler }

// SpamDetector rejects messages that look like spam.
type SpamDetector struct{ BaseHandler }

// LengthLimiter trims overly long messages but does not reject them.
type LengthLimiter struct {
	BaseHandler
	Limit int
}

// AuditLogger appends an audit flag and always passes.
type AuditLogger struct{ BaseHandler }

func (h *ProfanityFilter) Handle(msg *Message) {
	lower := strings.ToLower(msg.Text)
	banned := []string{"darn", "heck", "foobar"}
	for _, w := range banned {
		if strings.Contains(lower, w) {
			msg.Rejected = true
			msg.Reason = "profanity: " + w
			return
		}
	}
	h.callNext(msg)
}

func (h *SpamDetector) Handle(msg *Message) {
	lower := strings.ToLower(msg.Text)
	if strings.Contains(lower, "buy now") || strings.Count(lower, "!!!") > 2 {
		msg.Rejected = true
		msg.Reason = "spam detected"
		return
	}
	h.callNext(msg)
}

func (h *LengthLimiter) Handle(msg *Message) {
	if h.Limit > 0 && len(msg.Text) > h.Limit {
		msg.Text = msg.Text[:h.Limit]
		msg.Flags = append(msg.Flags, "trimmed")
	}
	h.callNext(msg)
}

func (h *AuditLogger) Handle(msg *Message) {
	msg.Flags = append(msg.Flags, "audited")
	h.callNext(msg)
}
