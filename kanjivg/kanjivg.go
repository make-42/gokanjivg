package kanjivg

import (
	"embed"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/make-42/gokanjivg/utils"
)

//go:embed deps/kanjivg/kanji/*
var f embed.FS

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

	dat, err := f.ReadFile(filepath.Join("./deps/kanjivg/kanji/", fmt.Sprintf("%05x.svg", rune)))
	return string(dat), err
}
