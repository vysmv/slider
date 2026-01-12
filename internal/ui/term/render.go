package term

import (
	"fmt"
	"strings"

	"golang.org/x/term"
)

type RenderModel struct {
	Width     int
	Height    int
	Indicator string // "X/Y"
	Content   string // slide text OR error text
	Prompt    string
}

func (u *UI) Render(m RenderModel) {
	clearScreen(u)

	w := m.Width
	h := m.Height
	if w <= 0 || h <= 0 {
		w, h = 80, 24
	}

	// minimal sane size
	if w < 10 {
		w = 10
	}
	if h < 6 {
		h = 6
	}

	top := "┌" + strings.Repeat("─", w-2) + "┐"
	bot := "└" + strings.Repeat("─", w-2) + "┘"

	u.Print(top + "\n")

	// inner height excludes top & bottom borders
	innerH := h - 2
	innerW := w - 2

	// Lines inside the frame:
	// line 0: indicator at top-right
	// line 1-2: blank
	// line 3.. : content
	contentStartLine := 3

	contentLines := splitAndClampLines(m.Content, innerW)
	contentIdx := 0

	for i := 0; i < innerH; i++ {
		var line string

		switch {
		case i == 0:
			// indicator at top-right inside the frame
			ind := clampToWidth(m.Indicator, innerW)
			spaces := innerW - runeLen(ind)
			if spaces < 0 {
				spaces = 0
			}
			line = strings.Repeat(" ", spaces) + ind
			line = padRight(line, innerW)

		case i < contentStartLine:
			line = strings.Repeat(" ", innerW)

		default:
			if contentIdx < len(contentLines) {
				if contentLines[contentIdx] != "" {
					line = padRight("  " + contentLines[contentIdx], innerW)
				} else {
					line = padRight(contentLines[contentIdx], innerW)
				}
				
				contentIdx++
			} else {
				line = strings.Repeat(" ", innerW)
			}
		}

		u.Print("│" + line + "│\n")
	}

	u.Print(bot + "\n")
	u.Print(m.Prompt)
}

func (u *UI) TermSize() (int, int) {
	w, h, err := term.GetSize(0) // stdin fd
	h = h - 3
	if err != nil || w <= 0 || h <= 0 {
		return 80, 24
	}
	return w, h
}

func clearScreen(u *UI) {
	// ANSI clear + home
	u.Print("\x1b[2J\x1b[H")
}

func splitAndClampLines(s string, width int) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	raw := strings.Split(s, "\n")

	out := make([]string, 0, len(raw))
	for _, ln := range raw {
		// do not wrap; clamp
		out = append(out, clampToWidth(ln, width))
	}
	return out
}

func clampToWidth(s string, width int) string {
	if width <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= width {
		return s
	}
	return string(r[:width])
}

func padRight(s string, width int) string {
	n := width - runeLen(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(" ", n)
}

func runeLen(s string) int {
	return len([]rune(s))
}

func indicator(cur, total int) string {
	return fmt.Sprintf("%d/%d", cur, total)
}
