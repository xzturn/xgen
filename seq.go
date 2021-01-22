package xgen

import(
    "math/rand"
    "time"
)

type XGenerator struct {
}

func NewXGenerator() *XGenerator {
    rand.Seed(time.Now().UnixNano())
    return &XGenerator{}
}

func genKnuth(seq *[]int, n, minv, maxv int) {
    if n <= 0 {
        return
    }
    k := maxv - minv + 1
    if rand.Intn(k) < n {
        (*seq)[n-1] = maxv
        genKnuth(seq, n-1, minv, maxv-1)
    } else {
        genKnuth(seq, n, minv, maxv-1)
    }
}

func (g XGenerator) GenKnuth(n, minv, maxv int) []int {
    seq := make([]int, n)
    genKnuth(&seq, n, minv, maxv)
    return seq
}

func (g XGenerator) GenShuffle(n, minv, maxv int) []int {
    seq := g.GenKnuth(n, minv, maxv)
    rand.Shuffle(n, func(i, j int) {
        seq[i], seq[j] = seq[j], seq[i]
    })
    return seq
}
