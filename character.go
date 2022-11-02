//go:generate fyne bundle -o bundled.go characters

package main

import "fyne.io/fyne/v2"

func resourceForCharacter(isUser bool) fyne.Resource {
	if isUser {
		return resourcePacmanIconSvg
	}
	return nil
}
