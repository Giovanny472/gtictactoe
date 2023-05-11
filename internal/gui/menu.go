package gui

import (
	"log"

	"github.com/Giovanny472/gtictactoe/model"
	tcell "github.com/gdamore/tcell/v2"
)

type menu struct {
	scr tcell.Screen
	err error
}

const countMenuOpts int = 3

var menuOptions = [countMenuOpts]string{"играть", "играть с другом", "выйти"}

var me *menu

func NewMenu() model.Menu {
	if me == nil {
		me = &menu{}
	}
	return me
}

func (m *menu) Load() {

	//  загрузка и инициализация экрана
	m.scr, m.err = tcell.NewScreen()
	if m.err != nil {
		log.Fatalf("%+v", m.err)
	}
	m.err = m.scr.Init()
	if m.err != nil {
		log.Fatalf("%+v", m.err)
	}

	// загрузка background
	m.scr.Clear()
	m.loadBackground()
	m.loadMenu()

}

func (m *menu) loadBackground() {

	styleBackgroud := tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorYellow)
	m.scr.SetStyle(styleBackgroud)
}

func (m *menu) loadMenu() {

	m.drawBox()

	m.drawText()
}

func (m *menu) drawBox() {

	w, h := m.scr.Size()
	w, h = w/2-25, h/2-10
	we := w + 50
	he := h + 15

	// стиль
	styleBorder := tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorGreen)

	// рисуем borders
	for col := w; col <= we; col++ {
		m.scr.SetContent(col, h, tcell.RuneHLine, nil, styleBorder)
		m.scr.SetContent(col, he, tcell.RuneHLine, nil, styleBorder)
	}
	for row := h; row < he; row++ {
		m.scr.SetContent(w, row, tcell.RuneVLine, nil, styleBorder)
		m.scr.SetContent(we, row, tcell.RuneVLine, nil, styleBorder)
	}

	// Only draw corners if necessary
	if h != he && w != we {
		m.scr.SetContent(w, h, tcell.RuneULCorner, nil, styleBorder)
		m.scr.SetContent(we, h, tcell.RuneURCorner, nil, styleBorder)
		m.scr.SetContent(w, he, tcell.RuneLLCorner, nil, styleBorder)
		m.scr.SetContent(we, he, tcell.RuneLRCorner, nil, styleBorder)
	}

}

func (m *menu) drawText() {

	w, h := m.scr.Size()
	w, h = w/2-10, h/2-4
	we := w + 50
	he := h + 15

	// стиль
	styleWords := tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorDarkBlue)
	styleWords = styleWords.Bold(true)

	for idx, itemStr := range menuOptions {

		row := h + idx
		col := w

		for _, letter := range itemStr {

			m.scr.SetContent(col, row, letter, nil, styleWords)
			col++
			if col >= we {
				row++
				col = w
			}
			if row > he {
				break
			}
		}
	}

}

func (m *menu) Show() {

	m.scr.Show()

}
