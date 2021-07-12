package characters

import (
	"strings"

	"github.com/alwindoss/morse"
)

func getCodeString(t string) (error,[]byte) {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(t))
	if err != nil {
		return err,nil
	}

	return err,morseCode
}