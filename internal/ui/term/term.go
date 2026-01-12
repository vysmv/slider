package term

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type UI struct {
	in  *bufio.Reader
	out io.Writer
}

func NewUI() *UI {
	return &UI{
		in:  bufio.NewReader(os.Stdin),
		out: os.Stdout,
	}
}

func (u *UI) ReadLine() (string, error) {
	line, err := u.in.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func (u *UI) Print(s string) {
	fmt.Fprint(u.out, s)
}
