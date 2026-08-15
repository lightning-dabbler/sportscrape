package wnba

type BoxScoreType int
type FeedType int

const (
	UndefinedBoxScoreType BoxScoreType = iota
	Traditional
	Advanced
	Misc
	Scoring
	Usage
	FourFactors
)

const (
	UndefinedFeedType FeedType = iota
	BoxScore
	PlayByPlay
)

func (bst BoxScoreType) Type() string {
	switch bst {
	case Traditional:
		return "traditional"
	case Advanced:
		return "advanced"
	case Misc:
		return "misc"
	case Scoring:
		return "scoring"
	case Usage:
		return "usage"
	case FourFactors:
		return "fourfactors"
	}
	return "undefined"
}

func (bst BoxScoreType) Undefined() bool {
	switch bst {
	case Traditional, Advanced, Misc, Scoring, Usage, FourFactors:
		return false
	}
	return true
}

func (ft FeedType) Type() string {
	switch ft {
	case BoxScore:
		return "box-score"
	case PlayByPlay:
		return "play-by-play"
	}
	return "undefined"
}

func (ft FeedType) Undefined() bool {
	switch ft {
	case BoxScore, PlayByPlay:
		return false
	}
	return true
}
