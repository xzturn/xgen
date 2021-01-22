package xgen

import (
    "fmt"
    "testing"
)

const (
    n    = 100
    minv = 1
    maxv = 10000
)

var g *XGenerator = NewXGenerator()

func TestGenKunth(t *testing.T) {
    fmt.Printf("%v\n", g.GenKnuth(n, minv, maxv))
}

func TestGenShuffle(t *testing.T) {
    fmt.Printf("%v\n", g.GenShuffle(n, minv, maxv))
}

func BenchmarkGenKunth(b *testing.B) {
    for i := 0; i < b.N; i++ {
        g.GenKnuth(n, minv, maxv)
    }
}

func BenchmarkGenShuffle(b *testing.B) {
    for i := 0; i < b.N; i++ {
        g.GenShuffle(n, minv, maxv)
    }
}
