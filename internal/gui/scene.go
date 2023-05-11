package gui

import (
	"github.com/Giovanny472/gtictactoe/model"
)

type scene struct {
}

var sc *scene

func NewScene() model.Scene {
	if sc == nil {
		sc = &scene{}
	}
	return sc
}

func (s *scene) Load() {

}

func (s *scene) loadBackground() {

}

func (s *scene) Show() {

}
