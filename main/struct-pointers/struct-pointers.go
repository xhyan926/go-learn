package main

import (
	"fmt"
)

type Vertext struct {
	X int
	Y int
}

func main() {
	v := Vertext{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

}
