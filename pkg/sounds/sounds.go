package sounds

import (
	"github.com/faiface/beep/speaker"
	"time"

	"github.com/digiexchris/morseclock/pkg/keyer"
	"github.com/faiface/beep"
)

type player struct {
	sampleRate beep.SampleRate
}

func NewPlayer() keyer.Keyer {
	p := player{
		sampleRate: beep.SampleRate(44100),
	}

	speaker.Init(p.sampleRate, p.sampleRate.N(time.Second/10))

	return p
}

func (p player) Dit(ms time.Duration) {

	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Take(p.sampleRate.N(ms), Tone{}), beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (p player) Dah(ms time.Duration) {
	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Take(p.sampleRate.N(ms), Tone{}), beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (p player) CharacterSpace(ms time.Duration) {
	//blocking wait
}

func (p player) WordSpace(ms time.Duration) {
	//blocking wait
}