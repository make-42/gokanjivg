package colorize

import (
	"fmt"

	"github.com/make-42/gokanjivg/kanjivg"

	"github.com/PerformLine/go-stockutil/colorutil"
)

// strokeNumber is a 0 based index
func GetSpectrumColor(strokeNumber, strokeCount int, saturation, value float64) string {
	hue := float64(strokeNumber) / float64(strokeCount) * 360.0
	r, g, b := colorutil.HsvToRgb(hue, saturation, value)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func GetDoubleSpectrumColor(strokeNumber, strokeCount int, saturation, value float64) string {
	hue := float64(strokeNumber) / float64(strokeCount) * 360.0 * 2.0
	r, g, b := colorutil.HsvToRgb(hue, saturation, value)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// Mode can be "spectrum"
func Colorize(character, mode string, saturation, value float64) (string, error) {
	svgData, err := kanjivg.GetSVGForCharacter(character)
	strokeCount := kanjivg.GetStrokeCount(svgData)
	hexColor := "#000000"
	for i := 0; i < strokeCount; i++ {
		switch mode {
		case "spectrum":
			hexColor = GetSpectrumColor(i, strokeCount, saturation, value)
			break
		case "double-spectrum":
			hexColor = GetDoubleSpectrumColor(i, strokeCount, saturation, value)
			break
		default:
			hexColor = "#000000"
		}
		svgData = kanjivg.SetColorOfStroke(svgData, hexColor, i)
	}
	return svgData, err
}
