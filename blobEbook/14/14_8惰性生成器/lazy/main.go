package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any, Any) (Any, Any)

func main() {
	evenFunc := func(state Any, state2 Any) (Any, Any) {
		// 前
		os := state.(int)
		// 后
		os2 := state2.(int)

		ns := os + os2
		// 后变前 插入通道  前+后 后
		return os2, ns
	}
	even := BuildLazyIntEvalluator(evenFunc, -1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even:%v\n", i, even())
	}
}
func BuildLazyEvalluator(evenFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		//前
		var actState Any = initState
		//后
		var retVal Any = 1
		for {
			actState, retVal = evenFunc(actState, retVal)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}
func BuildLazyIntEvalluator(evenFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvalluator(evenFunc, initState)
	return func() int {
		return ef().(int)
	}
}

/* Output:

Fib nr 0: 0
Fib nr 1: 1
Fib nr 2: 1
Fib nr 3: 2
Fib nr 4: 3
Fib nr 5: 5
Fib nr 6: 8
Fib nr 7: 13
Fib nr 8: 21
Fib nr 9: 34
*/
