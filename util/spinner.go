package util

import (
	"BCLAC/styles"

	"github.com/charmbracelet/bubbles/spinner"
)

func SetSpinnerDefaults(s *spinner.Model) {
	s.Spinner = spinner.Line
	s.Style = styles.SPINNER_STYLE
}
