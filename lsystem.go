package main

import (
    "strings"
)

type Point struct {
    x, y float64
}

func (p *Point) add(other Point) {
    p.x += other.x
    p.y += other.y
}

type LSystem struct {
    rules map[byte][]byte
    contents string
}

func NewLSystem(rules map[byte]string, start byte) LSystem {
    result := LSystem {
        rules: make(map[byte][]byte),
        contents: string([]byte{start}),
    }
    for key, val := range rules {
        result.rules[key] = []byte(val)
    }
    return result
}

func (system *LSystem) Step() {
    var buffer strings.Builder
    for i := 0; i < len(system.contents); i++ {
        lhs := system.contents[i]
        if rhs, ok := system.rules[lhs]; ok {
            buffer.Write(rhs)
        } else {
            buffer.WriteByte(lhs)
        }
    }
    system.contents = buffer.String()
}

func (system *LSystem) Iterate(n uint) {
    for i := uint(0); i < n; i++ {
        system.Step()
    }
}