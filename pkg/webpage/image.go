package webpage

import (
	"bytes"
	"encoding/base64"
	"github.com/marceloneil/colour/pkg/colours"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)

const SquareSize = 512

func RGBToBase64Square(rgb colours.RGB) string {
	src := image.NewUniform(color.RGBA{
		R: rgb.R,
		G: rgb.G,
		B: rgb.B,
		A: 255,
	})
	rect := image.Rect(0, 0, SquareSize, SquareSize)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), src, image.Point{}, draw.Src)

	var jpegBuffer bytes.Buffer
	err := jpeg.Encode(&jpegBuffer, img, nil)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(jpegBuffer.Bytes())
}
