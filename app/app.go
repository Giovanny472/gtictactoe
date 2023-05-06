package app

import "github.com/Giovanny472/gtictactoe/model"

type app struct {
}

var ap *app

func NewApp() model.App {

	if ap == nil {
		ap = &app{}
	}
	return ap
}

func (a *app) Exec() {

}
