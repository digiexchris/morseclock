package timing_test

import (
	"github.com/digiexchris/morseclock/pkg/timing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timing", func() {
	Context("When WPM is 10", func() {
		It("Should return 120ms dit", func() {
			time := timing.Dit(10)

			Expect(time).To(Equal(120))
		})

		It("Should return 120ms segment space", func() {
			time := timing.CharacterSegmentSpace(10)

			Expect(time).To(Equal(120))
		})

		It("Should return 360ms dah", func() {
			time := timing.Dah(10)

			Expect(time).To(Equal(360))
		})

		It("Should return 360ms character space", func() {
			time := timing.CharacterSpace(10)

			Expect(time).To(Equal(360))
		})

		It("Should return 840ms word space", func() {
			time := timing.WordSpace(10)

			Expect(time).To(Equal(840))
		})
	})

	Context("When WPM is 20", func() {
		It("Should return 60ms dit", func() {
			time := timing.Dit(20)

			Expect(time).To(Equal(60))
		})

		It("Should return 60ms segment space", func() {
			time := timing.CharacterSegmentSpace(20)

			Expect(time).To(Equal(60))
		})

		It("Should return 180ms dah", func() {
			time := timing.Dah(20)

			Expect(time).To(Equal(180))
		})

		It("Should return 180ms character space", func() {
			time := timing.CharacterSpace(20)

			Expect(time).To(Equal(180))
		})

		It("Should return 420ms word space", func() {
			time := timing.WordSpace(20)

			Expect(time).To(Equal(420))
		})
	})

})
