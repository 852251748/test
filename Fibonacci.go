package main

import "fmt"

//
//type Any interface{}
//type EvalFunc func(Any, Any) (Any, Any, Any)
//
//func main() {
//	logic := func(f1 Any, f2 Any) (Any, Any, Any) {
//		os := f1.(int) + f2.(int)
//		f1, f2 = f2, f1
//		f1 = os
//		return f1, f2, os
//	}
//	even := BuildLazyInitEvaluator(logic, 0, 1)
//	for i := 0; i < 20; i++ {
//		fmt.Print(even(), ",")
//	}
//}
//
//func BuildLazyEvaluator(evalFunc EvalFunc, initState1 Any, initState2 Any) func() Any {
//	yiled := make(chan Any)
//	f1, f2 := initState1, initState2
//	var result Any
//	loop := func() {
//		for {
//			f1, f2, result = evalFunc(f1, f2)
//			yiled <- result
//		}
//	}
//	ret := func() Any {
//		return <-yiled
//	}
//	go loop()
//	return ret
//}
//func BuildLazyInitEvaluator(evalFunc EvalFunc, initState1 Any, initState2 Any) func() int {
//	fc := BuildLazyEvaluator(evalFunc, initState1, initState2)
//	return func() int {
//		return fc().(int)
//	}
//}

//----------------------------------------------------------------------
//var resume chan int
//
//func interger() chan int {
//	yield := make(chan int)
//
//	go func() {
//		os := 0
//		f1, f2 := 0, 1
//		//yield <- 1
//		for {
//			os = f1 + f2
//			yield <- os
//			f1, f2 = f2, f1
//			f1 = os
//		}
//	}()
//	return yield
//}
//
//func getinter() int {
//	return <-resume
//}
//func main() {
//	resume = interger()
//	start := time.Now()
//	for i := 1; i < 21; i++ {
//		fmt.Print(getinter(), ",")
//	}
//	end := time.Now()
//	delta := end.Sub(start)
//	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
//}

//--------------------------------------
//使用闭包实现斐波那契数
//func main() {
//	f1, f2 := 0, 1
//	fibonacciFunc := func() {
//		os := f1 + f2
//		f1, f2 = f2, f1
//		f1 = os
//		fmt.Print(os, " ")
//	}
//	start := time.Now()
//	for i := 0; i < 20; i++ {
//		fibonacciFunc()
//	}
//	end := time.Now()
//	delta := end.Sub(start)
//	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
//}

//------------------------------
//使用数组打印斐波那契数
func main() {
	fibonacci := [20]int64{}

	for i := 0; i < len(fibonacci); i++ {
		if i < 2 {
			fibonacci[i] = 1
		} else {
			fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		}
	}
	fmt.Println(fibonacci)
}
