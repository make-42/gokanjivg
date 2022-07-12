# GoKanjiVG
This go module lets you interact with the [KanjiVG](https://github.com/KanjiVG/kanjivg) dataset in Go.

# Documentation
### Methods
#### kanjivg.GetSVGForCharacter
```
kanjivg.GetSVGForCharacter(character string) (string, error)
```
Returns a string containing the SVG data for the character passed as input.
#### colorize.Colorize
```
colorize.Colorize(character, mode string, saturation, value float64) (string, error)
```
Returns a string containing colorized SVG data for the character passed as input.

`mode` can be `"spectrum"`, `"double-spectrum"` or `"pastel"`.

`saturation` and `value` must be contained within 0 and 1. They let you control the saturation and value of the color of the strokes when using `"spectrum"` or `"double-spectrum"` mode.

## More documentation
Example usage can be found at [gokanjivg-demo](https://github.com/make-42/gokanjivg-demo).

## Demo
![demo](https://user-images.githubusercontent.com/17462236/178521880-a3bd679b-9197-4b63-a221-173a15493a97.png)

## Inspiration
The Colorize function is inspired by [Kanji Colorise](https://github.com/cayennes/kanji-colorize).
