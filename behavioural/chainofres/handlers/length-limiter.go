package handlers

import "design-patterns/behavioural/chainofres/models"

type LengthLimiter struct {
	BaseHandler
	Limit int
}

func (h *LengthLimiter) Handle(msg *models.Message) {
	if h.Limit > 0 && len(msg.Text) > h.Limit {
		msg.Text = msg.Text[:h.Limit]
		msg.Flags = append(msg.Flags, "trimmed")
	}
	h.callNext(msg)
}
