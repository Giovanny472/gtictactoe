package model

type TypeUI int

const (
	TuiMenu  TypeUI = 0
	TuiScene TypeUI = 1
)

type UI interface {
	Load()
	Show()
}

type Menu interface {
	UI
}

type Scene interface {
	UI
}
