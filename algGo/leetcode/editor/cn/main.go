package main

import (
	"flag"
)

var problemNum = flag.Int("num", 0, "配置")

var SolutionMap = map[int]func(){}

func Register(problemNum int, solution func()) {
	SolutionMap[problemNum] = solution
}

//func main() {
//	flag.Parse()
//
//	if *problemNum == 0 {
//		fmt.Println("Please input problem num!")
//		return
//	}
//
//	if fn, ok := SolutionMap[*problemNum]; !ok {
//		fmt.Printf("Not found problem by num: %d", *problemNum)
//	} else {
//		fn()
//	}
//
//}
