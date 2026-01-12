package app

import (
	"fmt"
	"strconv"

	"github.com/vysmv/slider/internal/domain"
	"github.com/vysmv/slider/internal/store"
	"github.com/vysmv/slider/internal/ui/term"
)

type App struct {
	store store.SlideStore
	ui    *term.UI
	state domain.State
	msg   string // ephemeral message (e.g., "slide k not found.")
}

func New(st store.SlideStore, ui *term.UI) *App {
	s := domain.NewState(st.Total())
	return &App{
		store: st,
		ui:    ui,
		state: s,
	}
}

func (a *App) Run() error {
	for {
		content, err := a.store.Content(a.state.Current)
		if err != nil {
			// store error is real fatal here
			return err
		}

		viewText := content
		if a.msg != "" {
			viewText = a.msg
		}

		w, h := a.ui.TermSize()

		a.ui.Render(term.RenderModel{
			Width:     w,
			Height:    h,
			Indicator: fmt.Sprintf("%d/%d", a.state.Current, a.state.Total),
			Content:   viewText,
			Prompt:    "\n\x1b[33m[1...n]=open a slide by its number / [n]=next / [p]=prev [q]=quit > \x1b[0m ",
		})

		a.msg = "" // message shown once

		cmd, err := a.ui.ReadLine()
		if err != nil {
			return err
		}

		switch cmd {
		case "q":
			return nil
		case "n":
			a.state.Next()
		case "p":
			a.state.Prev()
		default:
			if cmd == "" {
				continue
			}
			k, convErr := strconv.Atoi(cmd)
			if convErr != nil {
				continue
			}
			if !a.state.Open(k) {
				a.msg = fmt.Sprintf("slide %d not found.", k)
			}
		}
	}
}
