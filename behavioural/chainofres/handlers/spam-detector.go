package handlers

import (
	"strings"
	"design-patterns/behavioural/chainofres/models"
)

type SpamDetector struct{ BaseHandler }

func (h *SpamDetector) Handle(msg *models.Message) {
	lower := strings.ToLower(msg.Text)
	if strings.Contains(lower, "buy now") || strings.Count(lower, "!!!") > 2 {
		msg.Rejected = true
		msg.Reason = "spam detected"
		return
	}
	h.callNext(msg)
}
