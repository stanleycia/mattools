package github.com/stanleycia/mattools

import (
    "gonum.org/v1/gonum/mat"
)

func TopN(m *mat.Dense, l int) {
    le := l
    r, c := m.Dims()
    if le > r {
        le = r
    }
    if le <= 0 {
        panic("2nd parameter must be positive")
    }
    for i := 0; i < le; i++ {
        for j := 0; j < c; j++ {
            print(m.At(i, j), "\t")
        }
        println()
    }
}
