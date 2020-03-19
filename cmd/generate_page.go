package main

import (
	"fmt"
	"github.com/marceloneil/colour/pkg/colours"
	"github.com/marceloneil/colour/pkg/webpage"
	"io/ioutil"
	"sort"
)

func main() {
	colourList, err := colours.GetColourList()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of colours downloaded:", len(colourList))

	sort.Sort(colourList)

	page, err := webpage.GetColourPage(colourList)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("index.html", []byte(page), 0644)
	if err != nil {
		panic(err)
	}
}
