package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         //指针p指向i
	fmt.Println(*p) //通过p读取i的值

	*p = 21 //通过指针p设置i的值
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
