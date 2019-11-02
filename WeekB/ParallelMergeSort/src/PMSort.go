/*
	Parallel Merge Sort
*/

package main

import(
	"fmt"
	"math/rand"
	"testing"
	"sync"
	"time"
)

var N_PROC int = 4
/*
	Number of thread to be spawn. It should be equlat to number of logical processer.
	Logical Processor = Hyperthreading_Factor * Processors
*/

var VERBOSE bool = true
/*
	Flag variable to suppress outputs during Testing and Benchmarking.
*/

var POPULATION int = 40
/*
	Number of Eleemnts in Array for Result Verification.
*/


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

func ParallelMergeSort(arr []int, s_idx int, e_idx int, process int, par_wg *sync.WaitGroup){
	/*
	This function sorts the passed array in prallel using mergesort in background.
	If we directly try to make merge sort parallel, the no of processes in the system
	will due to the code will explode. This function limits te number of process which
	will be spawned during te execution of the code.
	Note: Pass number of process in power of 2.
	*/
	if process != 1{
		m_idx := int((s_idx + e_idx)/2)
		var wg sync.WaitGroup
		
		wg.Add(2)
			go ParallelMergeSort(arr, s_idx, m_idx, int(process/2), &wg)
			go ParallelMergeSort(arr, m_idx+1, e_idx, int(process/2), &wg)
		wg.Wait()
		go par_wg.Done()
		Merge(arr, s_idx, m_idx, e_idx)
	} else{
		SeqMergeSort(arr, s_idx, e_idx)
		if VERBOSE{
			fmt.Printf("Sorted Sub Array: %v\n", arr[s_idx:e_idx+1])
		}
		par_wg.Done()
	}
}

func getRandomArray(size int) []int{
	/*
	Function to populate the array with random values.
	*/
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, uint(size))
	for i:=0 ; i<len(arr) ; i++{
		arr[i] = rand.Intn(size*100)
	}
	return arr
}

func getUsersArray() []int{
	/*
	Function to populate an array with values passed through command line.
	*/
	var n,val int
	fmt.Printf("\nEnter the number of Elements to be Sorted(>8): ")
	fmt.Scanf("%d\n", &n)
	
	if n<8{
		n = 8
	}
	arr := make([]int, n)
	
	fmt.Println("\nEnter the values: <Seprate two values by Pressing Enter>")
	for i:=0 ; i<n ; i++{
		fmt.Printf("%d > ", i+1)
		fmt.Scanf("%d\n", &val)
		arr[i] = val
	}
	return arr
}

func BenchmarkFunSequential(b *testing.B){
	/*
	This function is used to mesure the performance of function after each improvement.
	*/

	size := 1e6
	arr := getRandomArray(int(size))
	SeqMergeSort(arr, 0, len(arr)-1)
}

func BenchmarkFunParallel(b *testing.B){
	/*
	This function is used to mesure the performance of function after each improvement.
	*/

	size := 1e6
	
	arr := getRandomArray(int(size))
	var wg sync.WaitGroup
	
	wg.Add(1)
		ParallelMergeSort(arr, 0, len(arr)-1, N_PROC, &wg)
	wg.Wait()
}

func main(){
	var CHOICE int
	
	for {
		fmt.Println("\n############################################################")
		fmt.Println("\nChoice 1: Function Testing.")
		fmt.Println("Choice 2: Peform sorting on Randomly populated Array.")
		fmt.Println("Choice 3: Sorts the user defined array.)")
		fmt.Printf("\nEnter the Choice in Integer: ")
		fmt.Scanf("%d\n", &CHOICE)
		
		switch CHOICE{
		
		case 1:
				VERBOSE = false
				fmt.Printf("\nParallel Execution with %d Threads\n", N_PROC)
				fmt.Println(testing.Benchmark(BenchmarkFunParallel))
				fmt.Println("Serial Execution with 1 Main Thread")
				fmt.Println(testing.Benchmark(BenchmarkFunSequential))
				VERBOSE = true
				
		case 2: 
				arr := getRandomArray(POPULATION)
				fmt.Printf("\nOriginal Array:\n%v\n\n", arr)
				var wg sync.WaitGroup

				wg.Add(1)
					start := time.Now()
					ParallelMergeSort(arr, 0, len(arr)-1, N_PROC, &wg)
				wg.Wait()

				fmt.Println("\nTime Elpased", time.Since(start))
				fmt.Printf("\nSorted Array: %v\n", arr)
				
		case 3:
				arr := getUsersArray()
				fmt.Printf("\n\nOriginal Array:\n%v\n\n", arr)
				var wg sync.WaitGroup

				wg.Add(1)
					start := time.Now()
					ParallelMergeSort(arr, 0, len(arr)-1, N_PROC, &wg)
				wg.Wait()

				fmt.Println("\nTime Elpased", time.Since(start))
				fmt.Printf("\nSorted Array: %v\n", arr)
				
		default:
				fmt.Println("End")
				return	
		}
	}
}
