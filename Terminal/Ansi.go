package Terminal

import "fmt"

// --- Text Styles ---
const (
	Reset         = "\033[0m"
	Bold          = "\033[1m"
	Dim           = "\033[2m"
	Italic        = "\033[3m"
	Underline     = "\033[4m"
	Blink         = "\033[5m"
	Reverse       = "\033[7m" // Invert foreground and background
	Hidden        = "\033[8m"
	Strikethrough = "\033[9m"
)

// --- Standard Foreground Colors (3-bit/4-bit) ---
const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Bright variants
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// --- Standard Background Colors (3-bit/4-bit) ---
const (
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	// Bright background variants
	BgBrightBlack   = "\033[100m"
	BgBrightRed     = "\033[101m"
	BgBrightGreen   = "\033[102m"
	BgBrightYellow  = "\033[103m"
	BgBrightBlue    = "\033[104m"
	BgBrightMagenta = "\033[105m"
	BgBrightCyan    = "\033[106m"
	BgBrightWhite   = "\033[107m"
)

// --- Screen Control & Erasing ---
const (
	ClearScreen      = "\033[2J" // Clear entire screen
	ClearLine        = "\033[2K" // Clear entire line
	ClearLineToRight = "\033[0K" // Clear from cursor to end of line
	ClearLineToLeft  = "\033[1K" // Clear from cursor to start of line
	ClearScreenDown  = "\033[0J" // Clear from cursor down to end of screen
	ClearScreenUp    = "\033[1J" // Clear from cursor up to start of screen
)

// --- Alternate Screen Buffer ---
const (
	EnterAltScreen = "\033[?1049h" // Switch to alternate full-screen buffer
	ExitAltScreen  = "\033[?1049l" // Switch back to primary screen buffer
)

// --- Cursor Visibility ---
const (
	HideCursor = "\033[?25l"
	ShowCursor = "\033[?25h"
)

// --- Dynamic Cursor Functions ---

// MoveUp moves the cursor up by n rows.
func MoveUp(n int) string { return fmt.Sprintf("\033[%dA", n) }

// MoveDown moves the cursor down by n rows.
func MoveDown(n int) string { return fmt.Sprintf("\033[%dB", n) }

// MoveRight moves the cursor right by n columns.
func MoveRight(n int) string { return fmt.Sprintf("\033[%dC", n) }

// MoveLeft moves the cursor left by n columns.
func MoveLeft(n int) string { return fmt.Sprintf("\033[%dD", n) }

// MoveToColumn moves the cursor to a specific column index in the current row.
func MoveToColumn(col int) string { return fmt.Sprintf("\033[%dG", col) }

// MoveTo position the cursor at an exact row and column coordinate (1-indexed).
func MoveTo(row, col int) string { return fmt.Sprintf("\033[%d;%dH", row, col) }
