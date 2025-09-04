package models

type SignaturePen struct {
	Color string
}

func (s *SignaturePen) Sign(artwork string) string {
	return "Signed '" + artwork + "' with " + s.Color + " pen"
}
