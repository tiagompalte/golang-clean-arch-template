package nativemigrate

import "strings"

type Migrate struct {
	Name    string
	HasUp   bool
	HasDown bool
}

func (m Migrate) IsValid() bool {
	return len(strings.TrimSpace(m.Name)) > 0 && m.HasUp && m.HasDown
}
