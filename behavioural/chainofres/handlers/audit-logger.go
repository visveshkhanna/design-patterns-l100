package handlers

import "design-patterns/behavioural/chainofres/models"

type AuditLogger struct{ BaseHandler }

func (h *AuditLogger) Handle(msg *models.Message) {
	msg.Flags = append(msg.Flags, "audited")
	h.callNext(msg)
}
