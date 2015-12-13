package main

import (
	"fmt"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

type line struct {
	X1, Y1, X2, Y2 float64
}

var iterations = 1

func getMiddle(l line) (line, float64) {

	// Take the total length of the line and divide by 3 to find the length of the middle.
	length := (math.Sqrt(math.Pow(l.X2-l.X1, 2)+math.Pow(l.Y2-l.Y1, 2)) / 3)

	mid := line{l.X1 + length, l.Y1 + length, l.X1 + (2 * length), l.Y1 + (2 * length)}

	return mid, length
}

func drawLine(l line) {

}

func eraseLine(l line) {

}

func branch(l line, i int) {

	if i == iterations {

		return
	}

	drawLine(l)

	middle, length := getMiddle(l)

	eraseLine(middle)

	left := line{middle.X1, middle.Y1, middle.X2 + length, middle.Y2}
	right := line{middle.X1, middle.Y1, middle.X2, middle.Y2 + length}

	branch(left, i+1)
	branch(right, i+1)

}

func drawSnowflake(gc *draw2dimg.GraphicContext) {
	// l1 :=
	// l2 :=
	// l3 :=
}

func main() {
	//dest := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	//gc := draw2dimg.NewGraphicContext(dest)

	//drawSnowflake(gc)

	//draw2dimg.SaveToPngFile("test.png", dest)

	line, l := getMiddle(line{1, 0, 0, 3})
	fmt.Println("Line is ", line)
	fmt.Println("length is ", l)
}
