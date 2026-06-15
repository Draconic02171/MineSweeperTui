package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/Draconic02171/MineSweeper/Terminal"
	"golang.org/x/term"
)

const Height uint8 = 16
const Width uint8 = 32
const MaxBomb uint8 = 32
const EnumBomb uint8 = 9
const FullBlock rune = '█'

type Vec2 struct{ X, Y int8 }

type Block struct {
	ForgroundColor  Terminal.Color
	BackgroundColor Terminal.Color
	Selected        bool
	Checked         bool
	Value           uint8
}

func PrintBlock(b Block) {
	if b.Selected == true {
		fmt.Printf(
			"%s%s%d%s ",
			b.ForgroundColor.ToString(),
			b.BackgroundColor.ToString(),
			b.Value,
			Terminal.Reset,
		)
		return
	}
	fmt.Printf(
		"%s%s.%s ",
		Terminal.Color{R: 150, G: 150, B: 150, Font: true}.ToString(),
		b.BackgroundColor.ToString(),
		Terminal.Reset,
	)
}

func Render(Field *[Height][Width]Block) {
	for i := range int(Height) {
		for j := range int(Width) {
			PrintBlock(Field[i][j])
		}
		fmt.Printf("%s\n", Terminal.MoveToColumn(0))
	}
}

func CleanBlocks(Field *[Height][Width]Block) {
	for Y := range int(Height) {
		for X := range int(Width) {
			Field[Y][X].Checked = false
		}
	}
}

func GenerateMap() [Height][Width]Block {

	field := [Height][Width]Block{}
	bombPos := [MaxBomb]Vec2{}

	for i := 0; i < int(MaxBomb); i++ {

		Bomb := Vec2{
			X: int8(rand.IntN(int(Width))),
			Y: int8(rand.IntN(int(Height))),
		}

		if field[Bomb.Y][Bomb.X].Value == EnumBomb {
			i--
			continue
		}

		field[Bomb.Y][Bomb.X].Value = EnumBomb
		bombPos[i] = Bomb
	}

	for i := range int(MaxBomb) {

		Bomb := bombPos[i]
		Bomb.X -= 1
		Bomb.Y -= 1

		for Y := Bomb.Y; Y < (Bomb.Y + 3); Y++ {
			for X := Bomb.X; X < (Bomb.X + 3); X++ {

				if Y < 0 || Y >= int8(Height) {
					continue
				}
				if X < 0 || X >= int8(Width) {
					continue
				}
				if field[Y][X].Value == 9 {
					continue
				}

				field[Y][X].Value += 1
			}
		}
	}

	for Y := range int(Height) {
		for X := range int(Width) {

			var RedColor uint8 = ((255 / 9) * field[Y][X].Value)
			var GreenColor uint8 = 0
			var BlueColor uint8 = 0

			if field[Y][X].Value != 0 {
				GreenColor = 255 / (field[Y][X].Value * field[Y][X].Value * field[Y][X].Value)
			}

			if field[Y][X].Value != 0 {
				BlueColor = 255 / field[Y][X].Value
			}

			field[Y][X].ForgroundColor = Terminal.Color{
				R:    RedColor,
				G:    GreenColor,
				B:    BlueColor,
				Font: true}

			field[Y][X].BackgroundColor = Terminal.Color{
				R:    0,
				G:    0,
				B:    0,
				Font: false}
			field[Y][X].Selected = false
		}
	}

	return field
}

func RevealBlocks(Field *[Height][Width]Block, Position Vec2, Attempt uint32) {

	{ // check if its exceeds max possible attempts
		if Attempt > uint32(Height)*uint32(Width) {
			return
		}
		Attempt++
	}
	{ // check out of bound
		if Position.X < 0 || Position.X >= int8(Width) {
			return
		}
		if Position.Y < 0 || Position.Y >= int8(Height) {
			return
		}
	}

	{ // check if its already get checked or the value is non zero
		if Field[Position.Y][Position.X].Value != 0 {
			Field[Position.Y][Position.X].Selected = true
			Field[Position.Y][Position.X].Checked = true
			return
		}
		if Field[Position.Y][Position.X].Checked == true {
			return
		}
		if Field[Position.Y][Position.X].Selected == true {
			return
		}
	}

	Field[Position.Y][Position.X].Selected = true
	Field[Position.Y][Position.X].Checked = true

	RevealBlocks(Field, Vec2{Position.X - 1, Position.Y}, Attempt)
	RevealBlocks(Field, Vec2{Position.X + 1, Position.Y}, Attempt)
	RevealBlocks(Field, Vec2{Position.X, Position.Y - 1}, Attempt)
	RevealBlocks(Field, Vec2{Position.X, Position.Y + 1}, Attempt)

}

func main() {

	// Enable raw mode
	OldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), OldState)
	InputChannel := make(chan byte, 1)

	// Goroutine to read input
	go func() {
		Buffer := make([]byte, 1)
		for {
			os.Stdin.Read(Buffer)
			InputChannel <- Buffer[0]
		}
	}()

	fmt.Print(Terminal.HideCursor)
	fmt.Printf(Terminal.EnterAltScreen)
	//////////////////////////////////////////////////////////

	CursorPosition := Vec2{X: int8(Width / 2), Y: int8(Height / 2)}
	Field := GenerateMap()
	OldPosition := CursorPosition
	IsChanged := true
	IsRunning := true

	for {

		fmt.Printf("%s", Terminal.MoveTo(0, 0))
		if IsRunning == false {
			break
		}

		Field[OldPosition.Y][OldPosition.X].BackgroundColor = Terminal.Color{R: 0, G: 0, B: 0, Font: false}
		Field[CursorPosition.Y][CursorPosition.X].BackgroundColor = Terminal.Color{R: 255, G: 255, B: 255, Font: false}

		if IsChanged == true {
			CleanBlocks(&Field)
			Render(&Field)
			IsChanged = false
		}

		OldPosition = CursorPosition

		select {
		case key := <-InputChannel:
			IsChanged = true
			switch key {
			case 27:
				IsRunning = false
			case 'w':
				CursorPosition.Y--
			case 's':
				CursorPosition.Y++
			case 'a':
				CursorPosition.X--
			case 'd':
				CursorPosition.X++
			case 13:
				RevealBlocks(&Field, CursorPosition, 0)

				if Field[CursorPosition.Y][CursorPosition.X].Value == 9 {
					Field[CursorPosition.Y][CursorPosition.X].BackgroundColor = Terminal.Color{R: 255, G: 0, B: 0, Font: false}
					Render(&Field)
					time.Sleep(time.Second)
					IsRunning = false
				}

			default:
				break
			}

		default:
			fmt.Print("")
		}

		if CursorPosition.X < 0 {
			CursorPosition.X = 0
		}
		if CursorPosition.X >= int8(Width) {
			CursorPosition.X = int8(Width) - 1
		}
		if CursorPosition.Y < 0 {
			CursorPosition.Y = 0
		}
		if CursorPosition.Y >= int8(Height) {
			CursorPosition.Y = int8(Height) - 1
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Print(Terminal.ShowCursor)
	fmt.Printf(Terminal.ExitAltScreen)
}
