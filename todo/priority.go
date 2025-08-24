package todo

import "fmt"

type Priority int

const (
	None = iota
	Low
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case None:
		return ""
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	default:
		return "No Status"
	}
}

func ParsePriority(s string) (Priority, error) {
	switch s {
	case "none", "":
		return None, nil
	case "low", "Low":
		return Low, nil
	case "medium", "Medium":
		return Medium, nil
	case "high", "High":
		return High, nil
	default:
		return None, fmt.Errorf("invalid property: %s", s)
	}
}
