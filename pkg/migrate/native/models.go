package nativemigrate

import "strings"

type Script struct {
	Name    string
	HasUp   bool
	HasDown bool
}

func (s Script) IsValid() bool {
	return len(strings.TrimSpace(s.Name)) > 0 && s.HasUp && s.HasDown
}
