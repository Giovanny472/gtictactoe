package gui

import "github.com/Giovanny472/gtictactoe/model"

type menu struct {
}

var me *menu

func NewMenu() model.Menu {
	if me == nil {
		me = &menu{}
	}
	return me
}

func (m *menu) Show() {

}
