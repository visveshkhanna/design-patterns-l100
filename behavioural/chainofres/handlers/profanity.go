package handlers

import (
	"strings"
	"design-patterns/behavioural/chainofres/models"
)

type ProfanityFilter struct{ BaseHandler }

func (h *ProfanityFilter) Handle(msg *models.Message) {
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
