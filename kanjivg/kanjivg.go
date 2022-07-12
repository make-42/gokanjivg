package kanjivg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/make-42/gokanjivg/utils"
)

var KanjiVGDataPath = "./deps/kanjivg/kanji/"

func SetKanjiVGDataPath(NewKanjiVGDataPath string) {
	KanjiVGDataPath = NewKanjiVGDataPath
}

func GetStrokeCount(svgData string) int {
	return strings.Count(svgData, "<path")
}

func SetColorOfStroke(svgData, hexColor string, strokeNumber int) string {
	svgData = utils.ReplaceNth(svgData, "<path", "<path style=\"stroke:"+hexColor+";\"", strokeNumber)
	svgData = utils.ReplaceNth(svgData, "<text", "<text style=\"fill:"+hexColor+";\"", strokeNumber)
	return svgData
}

func GetSVGForCharacter(character string) (string, error) {
	rune, _ := utf8.DecodeRuneInString(character)
	dat, err := os.ReadFile(filepath.Join(KanjiVGDataPath, fmt.Sprintf("%05x.svg", rune)))
	return string(dat), err
}
