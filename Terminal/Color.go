package Terminal

import "fmt"

type Color struct {
	R    uint8
	G    uint8
	B    uint8
	Font bool
}

func (c Color) ToString() string {

	if c.Font == true {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
	}

	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
}
