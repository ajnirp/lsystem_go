package main

type SierpinskiTriangle struct {
    LSystem
}

func NewSierpinskiTriangle() SierpinskiTriangle {
    rules := map[byte]string {
        'A': "B-A-B",
        'B': "A+B+A",
    }
    return SierpinskiTriangle { NewLSystem(rules, 'A') }
}

type FractalTree struct {
    LSystem
}

func NewFractalTree() FractalTree {
    rules := map[byte]string {
        '1': "11",
        '0': "1[0]0",
    }
    return FractalTree { NewLSystem(rules, '0') }
}

type HexagonalGosperCurve struct {
    LSystem
}

func NewHexagonalGosperCurve() HexagonalGosperCurve {
    rules := map[byte]string {
        'A': "A+B++B-A--AA-B+",
        'B': "-A+BB++B+A--A-B",
    }
    return HexagonalGosperCurve { NewLSystem(rules, 'A') }
}