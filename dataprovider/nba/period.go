package nba

type Period int

const (
	UndefinedPeriod Period = iota
	Full
	Q1
	Q2
	Q3
	Q4
	H1
	H2
	AllOT
)

func (p Period) Period() string {
	switch p {
	case Full:
		return "All"
	case Q1:
		return "Q1"
	case Q2:
		return "Q2"
	case Q3:
		return "Q3"
	case Q4:
		return "Q4"
	case H1:
		return "1stHalf"
	case H2:
		return "2ndHalf"
	case AllOT:
		return "AllOT"
	}
	return "Undefined"
}

func (p Period) Undefined() bool {
	switch p {
	case Full, Q1, Q2, Q3, Q4, H1, H2, AllOT:
		return false
	}
	return true
}
