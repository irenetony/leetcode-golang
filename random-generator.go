package main

import (
	"math"
	"math/rand"
	"time"
	"fmt"
)

func random(N int) int {
	i := int(math.Ceil(math.Log2(float64(N))))
	for {
		tmp := 0
		for j := 0; j < i; j++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			tmp = 2*tmp + r.Intn(2)
		}
		if tmp <= N-1 {
			return tmp
		}
	}
}

func main() {
	N := 10
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\n", random(N))
	}
}
