package colours

const repetitions = 8

type ColourList []*Colour

func (c ColourList) Len() int {
	return len(c)
}

func (c ColourList) Less(i, j int) bool {
	h1, l1, v1 := c[i].HLV(repetitions)
	h2, l2, v2 := c[j].HLV(repetitions)
	return h1 < h2 ||
		(h1 == h2 && l1 < l2) ||
		(h1 == h2 && l1 == l2 && v1 < v2)
}

func (c ColourList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
