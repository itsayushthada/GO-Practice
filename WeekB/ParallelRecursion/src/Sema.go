package main

import (
	"fmt"
	"sync"
	"time"
)

func serial_exec(name string, a int, parent *sync.WaitGroup, self *sync.WaitGroup){
	/*
	This function is created for testing purpose only.
	*/
	
	if(a==1){
		a = a-1
		name = fmt.Sprintf("%s.L", name)
		self.Done()
		fmt.Printf("Leaf: %s | Parent: %v | Own: %v \n", name, *parent, *self)
	}else{
		var child sync.WaitGroup
		
		a = int(a/2)
		name1  := fmt.Sprintf("%s.1", name)
		name2  := fmt.Sprintf("%s.2", name)
		name3  := fmt.Sprintf("%s.M", name)
		child.Add(2)
			fmt.Printf("Split: %s ### Parent: %v ### Own: %v \n", name1, *self, child)
			serial_exec(name1, a, self, &child)
			fmt.Printf("Split: %s ### Parent: %v ### Own: %v \n", name2, *self, child)
			serial_exec(name2, a, self, &child)
			self.Done()
		child.Wait()
		fmt.Printf("Merge: %s ### Parent: %v ### Own: %v \n", name3, *self, child)
	}
}

func parallel_exec(name string, a int, parent *sync.WaitGroup, self *sync.WaitGroup){
	if(a==1){
		a = a-1
		name = fmt.Sprintf("Leaf: %s.L\n", name)
		self.Done()
		fmt.Printf(name)
	}else{
		var child sync.WaitGroup
		
		a = int(a/2)
		name1  := fmt.Sprintf("%s.1", name)
		name2  := fmt.Sprintf("%s.2", name)
		name3  := fmt.Sprintf("%s.M", name)
		child.Add(2)
			fmt.Printf("Split: %s\n", name1)
			go parallel_exec(name1, a, self, &child)
			fmt.Printf("Split: %s\n", name2)
			go parallel_exec(name2, a, self, &child)
			self.Done()
		child.Wait()
		fmt.Printf("Merge: %s\n", name3)
	}
}

func main(){
	var parent sync.WaitGroup
	var self sync.WaitGroup
	a := 4
	name := "O"
	start := time.Now()
	
	self.Add(1)
		fmt.Printf("Origin:	%s\n", name)
		parallel_exec(name, a, &parent, &self)
	self.Wait()
	
	fmt.Println("End: O\n\nTime Elapsed", time.Since(start))

}
