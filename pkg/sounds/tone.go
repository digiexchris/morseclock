package sounds

import "math/rand"

type Tone struct{}

func (n Tone) Err() error {
	return nil
}

func (n Tone) Stream(samples [][2]float64) (m int, ok bool) {
	for i := range samples {
		samples[i][0] = rand.Float64()*2 - 1
		samples[i][1] = rand.Float64()*2 - 1
	}
	return len(samples), true
}