package todo

type Priority int 

const (
	Low = iota + 1
	Medium
	High
)

func (p Priority) String() string {
	switch p {
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

// todo: add parser, implement in addTask