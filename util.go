package main

import (
    "github.com/fogleman/gg"
)

func DrawSierpinskiTriangle(context *gg.Context, numIterations uint) {
    context.SetRGB(1, 1, 1)
    context.Clear()
    sierpinskiTriangle := NewSierpinskiTriangle()
    sierpinskiTriangle.Iterate(numIterations)
    sierpinskiTriangle.Draw(context)
    context.Stroke()
    context.SavePNG("outputs/sierpinski_triangle.png")
}

func DrawHexagonalGosperCurve(context *gg.Context, numIterations uint) {
    context.SetRGB(1, 1, 1)
    context.Clear()
    hexagonalGosperCurve := NewHexagonalGosperCurve()
    hexagonalGosperCurve.Iterate(numIterations)
    hexagonalGosperCurve.Draw(context)
    context.Stroke()
    context.SavePNG("outputs/hexagonal_gosper_curve.png")
}