package main

import (
    "github.com/fogleman/gg"
)

const (
    canvasWidth = 1000
    canvasHeight = 1000
)

func main() {
    context := gg.NewContext(canvasWidth, canvasHeight)

    DrawSierpinskiTriangle(context, 8);
    DrawHexagonalGosperCurve(context, 5);

}