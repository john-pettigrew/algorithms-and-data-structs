package main

import (
	"fmt"
	"math"
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

type line struct {
	X1, Y1, X2, Y2 float64
}


const (
	height = 1000
	width = 1000
)

var iterations = 9

func getMiddle(l line) (line, float64) {

	// Take the total length of the line and divide by 3 to find the length of the middle.
	length := (math.Sqrt(math.Pow(l.X2-l.X1, 2)+math.Pow(l.Y2-l.Y1, 2)) / 3)

	angle := math.Atan2(l.Y2 - l.Y1, l.X2 - l.X1)

	deltaX := length * math.Cos(angle)
	deltaY := length * math.Sin(angle)


	mid := line{deltaX + l.X1, deltaY + l.Y1, l.X1 + (2 * deltaX), l.Y1 + (2 * deltaY)}

	return mid, length
}

func drawLine(l line, gc *draw2dimg.GraphicContext) {
	gc.SetLineWidth(5)
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	gc.MoveTo(l.X1, l.Y1)
	gc.LineTo(l.X2, l.Y2)
	gc.FillStroke()

}

func eraseLine(l line, gc *draw2dimg.GraphicContext) {

	gc.SetLineWidth(5)
	gc.SetFillColor(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	gc.SetStrokeColor(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})


	for i := 0; i < 10; i++ {
		gc.MoveTo(l.X1, l.Y1)
		gc.LineTo(l.X2, l.Y2)
		gc.FillStroke()
	}
}

func branch(l line, i int, gc *draw2dimg.GraphicContext) {

	if i == iterations {

		return
	}

	drawLine(l, gc)

	middle, length := getMiddle(l)

	eraseLine(middle, gc)

	fmt.Println(length)

	midPointX := middle.X1 + length
	midPointY := middle.Y1 + length

	left := line{middle.X1, middle.Y1, midPointX, midPointY}
	right := line{middle.X2, middle.Y2, midPointX, midPointY}

	branch(left, i+1, gc)
	branch(right, i+1, gc)

}

func drawSnowflake(gc *draw2dimg.GraphicContext) {
	l1 := line{500, 250, 250, 750}
	branch(l1, 1, gc)

	l2 := line{250, 750, 750, 750}
	branch(l2, 1, gc)

	l3 := line{750, 750, 500, 250}
	branch(l3, 1, gc)
}

func main() {
	fmt.Println("drawing")


	dest := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFillColor(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	gc.ClearRect(0,0,width,height)

	drawSnowflake(gc)

	gc.Close()

	draw2dimg.SaveToPngFile("test.png", dest)

}
