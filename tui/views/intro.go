// Package views contains the other views of the application.
package views

// Returns the into title of the application.
func IntroSequence() string {
	const title = `
  ____   _____ _               _____ 
 |  _ \ / ____| |        /\   / ____|
 | |_) | |    | |       /  \ | |     
 |  _ <| |    | |      / /\ \| |     
 | |_) | |____| |____ / ____ \ |____ 
 |____/ \_____|______/_/    \_\_____|
        WELCOME
`
	return title
}
