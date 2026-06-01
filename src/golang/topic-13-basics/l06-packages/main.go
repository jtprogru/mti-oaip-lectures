package main

import (
	"fmt"

	"github.com/jtprogru/mti-oaip-lectures/src/golang/topic-13-basics/l06-packages/mathx"
)

func main() {
	fmt.Println("Sum(2, 3) =", mathx.Sum(2, 3))
	fmt.Println("Max(7, 4) =", mathx.Max(7, 4))
}
