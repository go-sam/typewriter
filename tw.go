package tw

import (
	"fmt"
	"time"
)

type Typewriter struct {
	Text  string
	Speed int
}

func (t *Typewriter) SetText(text string) {
	t.Text = text
}

func (t *Typewriter) SetSpeed(speed int) {
	t.Speed = speed
}

func (tw *Typewriter) Type() {
	visibleCount := 0
	inANSISequence := false

	for _, char := range tw.Text {
		fmt.Print(string(char))

		if char == '\x1b' {
			inANSISequence = true
		} else if inANSISequence && char == 'm' {
			inANSISequence = false
			continue
		}

		if !inANSISequence && char != '\n' {
			visibleCount++
			time.Sleep(time.Duration(tw.Speed) * time.Millisecond)
		}
	}
}
