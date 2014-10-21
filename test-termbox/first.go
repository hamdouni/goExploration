package main

import "github.com/nsf/termbox-go"

func mainMenu(menu []string, sel int) {
	for i := 1; i <= len(menu); i++ {
		DrawRichLine(1, i, menu[i-1], (sel == i))
	}
	termbox.Flush()
}

func DrawRichLine(x, y int, text string, highlight bool) {
	fg := termbox.ColorWhite
	bg := termbox.ColorBlack

	if highlight {
		tmp := fg
		fg = bg
		bg = tmp
	}
	for index, ch := range text {
		termbox.SetCell(x+index, y, ch, fg, bg)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	menu := []string{"Scanner vers reseau", "Scanner vers email", "ESC pour quitter"}
	selMenu := 1
	maxMenu := len(menu)
	mainMenu(menu, selMenu)

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowDown && selMenu < maxMenu {
				selMenu++
				mainMenu(menu, selMenu)
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowUp && selMenu > 1 {
				selMenu--
				mainMenu(menu, selMenu)
			}
		}
	}
}
