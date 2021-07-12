package keyer

import "time"

type Keyer interface {
	Dit(ms time.Duration)
	Dah(ms time.Duration)
	CharacterSpace(ms time.Duration)
	WordSpace(ms time.Duration)
}
