package main

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
)

const iter = 4

func iteration(i string) string {
	return strings.Replace(i, "F", "F+F--F+F", -1)
}

func nextPoint(a, l, x, y float64) (newX, newY float64) {

	deltaX := l * math.Cos(a)
	deltaY := l * math.Sin(a)

	newX = deltaX + x
	newY = deltaY + y

	return
}

func drawSnowflake(sys string, gc *draw2dimg.GraphicContext) {

	angle := float64(0)
	lineLength := float64(10)
	x, y := float64(1000), float64(1000)

	gc.MoveTo(x, y)

	for _, i := range sys {

		switch string(i) {

		case "F":
			x, y = nextPoint(angle, lineLength, x, y)
			gc.LineTo(x, y)
		case "+":
			angle += math.Pi / 3
		case "-":
			angle -= math.Pi / 3
		}
	}
}

func snowflake() {

	// Create the L System
	sys := "F--F--F"
	for i := 0; i < iter; i++ {
		sys = iteration(sys)
	}

	// Create graphic context for drawing
	dest := image.NewRGBA(image.Rect(0, 0, 2000, 2000))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	gc.SetLineWidth(5)

	//Draw the image
	drawSnowflake(sys, gc)
	//gc.Close()
	gc.Stroke()

	//Save image
	draw2dimg.SaveToPngFile("test.png", dest)
}

func main() {

	snowflake()
}
