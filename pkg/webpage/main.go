package webpage

import (
	"bytes"
	"github.com/marceloneil/colour/pkg/colours"
	"html/template"
	"strings"
)

func GetColourPage(colourList colours.ColourList) (string, error) {
	indexTmpl := template.New("index.tmpl").Funcs(template.FuncMap{
		"stringToLower": strings.ToLower,
		"rgbToBase64Square": RGBToBase64Square,
		"formatArticle": func(article string) string { return strings.ReplaceAll(article, " ", "_") },
	})
	indexTmpl, err := indexTmpl.ParseFiles("pkg/webpage/index.tmpl")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = indexTmpl.Execute(&buf, colourList)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
