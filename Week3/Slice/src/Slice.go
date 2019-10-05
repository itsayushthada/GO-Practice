package main 
  
import ( 
    "fmt"
    "sort"
	"strconv"
) 
  
func main() { 

    scl := make([]int , 0, 100)
    fmt.Println("Slice Instantiated: ", scl)

	n := ""
	for{
		_, _ = fmt.Scanln(&n)
		if !(n[0] == 'x' || n[0] == 'X'){
			val, _ := strconv.Atoi(n)
			scl = append(scl, val)
			if !sort.IntsAreSorted(scl){
				sort.Ints(scl) 
			}
			fmt.Println("Sorted Slice is: ", scl)
		} else{
			fmt.Println("Program Over")
			return
		}
	}
}