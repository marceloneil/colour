package colours

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const colortRegex = `{{Colort/ColorShort\|[^\n\r]*\|r=([0-9 ]*)\|g=([0-9 ]*)\|b=([0-9 ]*)\|h=([0-9 \-—]*)\|s=([0-9 ]*)\|v=([0-9 ]*)\|name=(\[\[([^\n\r\|]*|([^\n\r]*)\|([^\n\r]*))\]\]|[^\n\r]*)}}`

func parseColoursFromArticle(article string) (colourList ColourList, err error) {
	colortRegexp, err := regexp.Compile(colortRegex)
	if err != nil {
		return nil, err
	}
	matches := colortRegexp.FindAllStringSubmatch(article, -1)
	expectedCount := strings.Count(article, "Colort")
	actualCount := len(matches)
	if expectedCount != actualCount {
		return nil, fmt.Errorf("error parsing article, found %d colours, expected %d",
			actualCount, expectedCount)
	}

	colourList = make([]*Colour, actualCount)
	for index, match := range matches {
		red, err := parseCode(match[1])
		if err != nil {
			return nil, err
		}
		green, err := parseCode(match[2])
		if err != nil {
			return nil, err
		}
		blue, err := parseCode(match[3])
		if err != nil {
			return nil, err
		}

		hue, err := parseCode(match[4])
		if err != nil {
			return nil, err
		}
		saturation, err := parseCode(match[5])
		if err != nil {
			return nil, err
		}
		value, err := parseCode(match[6])
		if err != nil {
			return nil, err
		}

		var name string
		var article string
		if match[10] != "" {
			name = match[10]
			article = match[9]
		} else if match[8] != "" {
			name = match[8]
			article = match[8]
		} else if match[7] != "" {
			name = match[7]
			article = match[7]
		} else {
			return nil, fmt.Errorf("error parsing name from %s", match[0])
		}

		colourList[index] = &Colour{
			Name: name,
			Article: article,
			RGB: RGB{
				R: uint8(red),
				G: uint8(green),
				B: uint8(blue),
			},
			HSV: HSV{
				H: float64(hue) / 360,
				S: float64(saturation) / 100,
				V: float64(value) / 100,
			},
		}
	}

	return colourList, nil
}

func parseCode(data string) (code uint64, err error) {
	trimmedData := strings.TrimSpace(data)
	if trimmedData == "-" || trimmedData == "—" {
		return 0, nil
	}
	return strconv.ParseUint(strings.TrimSpace(data), 10, 64)
}
