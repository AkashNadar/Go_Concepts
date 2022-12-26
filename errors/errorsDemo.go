package main

import (
	"errors"
	"fmt"
)

type Stuffs struct {
	ele []int
}

func (s *Stuffs) getEle(index int) (int, error) {
	if index > len(s.ele) {
		return 0, errors.New(fmt.Sprintf("index out of range: %v", index))
	} else {
		return s.ele[index], nil
	}
}

func main() {
	stuff := Stuffs{
		ele: []int{45, 64, 23, 46},
	}
	val, err := stuff.getEle(10)
	fmt.Printf("the value is :%v, and the err is :%v", val, err)
}
