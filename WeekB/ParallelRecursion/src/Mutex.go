package main

import (
	"fmt"
	"sync"
	"time"
)

func parallel_exec(a int){
	if(a==1){
		a = a-1
		fmt.Printf("Leaf: %d\n", a)
	}else{		
		var inner_wg sync.Mutex
		a = int(a/2)
		
		inner_wg.Lock()
			fmt.Printf("Split: %d | %+v \n", a, inner_wg)
			go parallel_exec(a)
			fmt.Printf("Split: %d | %+v \n", a, inner_wg)
			go parallel_exec(a)
		inner_wg.Unlock()
	}
}

func main(){
	var wg sync.Mutex
	a := 4
	start := time.Now()
	fmt.Printf("Origin: %d | %+v \n", a, wg)
	parallel_exec(a)
	fmt.Println("Done. %v", wg, time.Since(start))

}
