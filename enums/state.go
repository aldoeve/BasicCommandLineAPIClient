package enums

type State uint

// The states the application can be in.
const (
	INTRO = iota
	MAIN_VIEW
	PASTE
	QUICK_BUILD
	RECENTS
	MANUAL
	FIRE_N_SHOW_HTTP
	ENDING
)
