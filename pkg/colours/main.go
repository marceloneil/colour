package colours

func GetColourList() (colourList ColourList, err error) {
	article, err := downloadColourArticle()
	if err != nil {
		return nil, err
	}

	return parseColoursFromArticle(article)
}
