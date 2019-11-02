/*
	Parallel Merge Sort
*/

package main

import(
	"fmt"
	"time"
	"math/rand"
	"testing"
	"sync"
)

func Merge(arr []int, s_idx int, m_idx int, e_idx int){
	/*
	This function takes two arrays and merge them into one array sorted array.
	Two arrays which are passed contigoues in memory.
	First array lies between s_idx and m_idx indices and next array lies between m_idx+1 and
	e_idx indices of the parent array.
	Here instead of arrays slices are used.
	*/

	idx_1 := 0
	idx_2 := 0
	idx_main := s_idx

	temp1 := make([]int, m_idx+1-s_idx)
	temp2 := make([]int, e_idx-m_idx)

	copy(temp1, arr[s_idx:m_idx+1])
	copy(temp2, arr[m_idx+1:e_idx+1])

	for idx_1 < len(temp1) && idx_2 < len(temp2){
		if(temp1[idx_1] < temp2[idx_2]){
			arr[idx_main] = temp1[idx_1]
			idx_1 = idx_1+1
		} else{
			arr[idx_main] = temp2[idx_2]
			idx_2 = idx_2+1
		}
		idx_main = idx_main+1
	}

	if(idx_1 == len(temp1)){
		copy(arr[idx_main:e_idx+1], temp2[idx_2:])
		idx_2 = len(temp2)
	} else{
		copy(arr[idx_main:e_idx+1], temp1[idx_1:])
		idx_1 = len(temp1)
	}
	idx_main = e_idx+1
}

func SeqMergeSort(arr []int, s_idx int, e_idx int){
	/*
	This function takes the data slice or parent slice and recursively splits into
	slice of half size upto a minimum size ie. 2 and then merge those slices until
	we merge all the children slice.
	*/

	m_idx := int((s_idx + e_idx)/2)
	if(s_idx < m_idx || e_idx > m_idx){
		SeqMergeSort(arr, s_idx, m_idx)
		SeqMergeSort(arr, m_idx+1, e_idx)
		Merge(arr, s_idx, m_idx, e_idx)
	}
}

func ParallelMergeSort(arr []int, s_idx int, e_idx int, process int, wg *sync.WaitGroup){
	/*
	This function sorts the passed array in prallel using mergesort in background.
	If we directly try to make merge sort parallel, the no of processes in the system
	will due to the code will explode. This function limits te number of process which
	will be spawned during te execution of the code.
	Note: Pass number of process in power of 2.
	*/
	
	var inner_wg sync.WaitGroup
	m_idx := int((s_idx + e_idx)/2)
	if process != 1{
		inner_wg.Add(2)
		ParallelMergeSort(arr, s_idx, m_idx, int(process/2), &inner_wg)
		ParallelMergeSort(arr, m_idx+1, e_idx, int(process/2), &inner_wg)
		inner_wg.Wait()
		Merge(arr, s_idx, m_idx, e_idx)
	} else{
		SeqMergeSort(arr, s_idx, e_idx)
		wg.Done()
	}
}

func BenchmarkFunction(b *testing.B){
	/*
	This function is used to mesure the performance of function after each improvement.
	*/

	size := 1e4
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, uint(size))
	for i:=0 ; i<len(arr) ; i++{
		arr[i] = rand.Intn(100)
	}
	SeqMergeSort(arr, 0, len(arr)-1)
}

func main(){
	/*
	This is a driver functions. It will allow us to encapsulte different function.
	*/

	size := 16
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, uint(size))
	for i:=0 ; i<len(arr) ; i++{
		arr[i] = rand.Intn(100)
	}
	
	var wg sync.WaitGroup
	fmt.Printf("\nOriginal Array: %v\n", arr)
	start := time.Now()
	
	wg.Add(1)
	ParallelMergeSort(arr, 0, len(arr)-1, 4, &wg)
	wg.Wait()
	
	fmt.Println("Time Elpased", time.Since(start))
	fmt.Printf("Sorted Array: %v\n", arr)
}