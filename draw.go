package main

import (
    "math"

    "github.com/fogleman/gg"
)

func (s *SierpinskiTriangle) Draw(context *gg.Context) {
    context.SetRGB(0, 0, 0)
    context.SetLineWidth(1)

    const deltaAngle = -math.Pi / 3
    const segmentLen = 3.25
    angle := 0.0
    p := Point { 100, 900 }
    for _, b := range s.contents {
        if b == 'A' || b == 'B' {
            dp := Point { segmentLen*math.Cos(angle), segmentLen*math.Sin(angle) }
            context.DrawLine(p.x, p.y, p.x + dp.x, p.y + dp.y)
            p.add(dp)
        } else if b == '+' {
            angle += deltaAngle
        } else if b == '-' {
            angle -= deltaAngle
        }
    }
}

func (s *HexagonalGosperCurve) Draw(context *gg.Context) {
    context.SetRGB(0, 0, 0)
    context.SetLineWidth(1)

    const deltaAngle = -math.Pi / 3
    const segmentLen = 6
    angle := 0.0
    p := Point { 750, 850 }
    for _, b := range s.contents {
        if b == 'A' || b == 'B' {
            dp := Point { segmentLen*math.Cos(angle), segmentLen*math.Sin(angle) }
            context.DrawLine(p.x, p.y, p.x + dp.x, p.y + dp.y)
            p.add(dp)
        } else if b == '+' {
            angle += deltaAngle
        } else if b == '-' {
            angle -= deltaAngle
        }
    }
}