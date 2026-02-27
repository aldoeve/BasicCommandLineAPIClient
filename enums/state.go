package enums

type State uint

// The states the application can be in.
const (
	INTRO = iota
	MAIN_VIEW
	ENDING
)
