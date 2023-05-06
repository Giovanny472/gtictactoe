package gui

import (
	"github.com/Giovanny472/gtictactoe/model"
)

func MakeUI(tui model.TypeUI) model.UI {

	switch tui {
	case model.TuiMenu:
		return NewMenu()
	case model.TuiScene:
		return NewScene()
	default:
		return nil
	}
}
