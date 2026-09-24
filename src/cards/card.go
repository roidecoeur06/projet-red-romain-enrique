package cards

import "fmt"

type Card struct {
	Suit  string
	Value string
}

func (c Card) GetASCII() []string {

	vTop := c.Value
	vBot := c.Value
	if len(vTop) == 1 {
		vTop = vTop + " "
		vBot = " " + vBot
	}

	return []string{
		"┌─────────┐",
		fmt.Sprintf("│ %s      │", vTop),
		"│         │",
		fmt.Sprintf("│    %s    │", c.Suit),
		"│         │",
		fmt.Sprintf("│      %s │", vBot),
		"└─────────┘",
	}
}

func (c Card) GetBackASCII() []string {
	return []string{
		"┌─────────┐",
		"│░░░░░░░░░│",
		"│░┌─────┐░│",
		"│░│░░░░░│░│",
		"│░│░░░░░│░│",
		"│░└─────┘░│",
		"└─────────┘",
	}
}
