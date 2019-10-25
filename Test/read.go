package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

    type name struct {
        first string
        last string
    }

    var names []name = make( []name, 0, 10 )
	var filename string

	scanner := bufio.NewScanner(os.Stdin)

	// Get the filename
	fmt.Print( "Enter name of the file to read: " )
	scanner.Scan()
	filename = scanner.Text()

	// Open the file
	fmt.Printf( "\nOpening file %s\n", filename )
	f, err := os.Open( filename )
	if nil != err { 
		fmt.Printf( "Error opening file %s: %v\n", filename, err.Error() )
		return
	}
	defer f.Close()

	// Read the file
	filescanner := bufio.NewScanner( f )
	line := 1
	for filescanner.Scan() {
		namearray := strings.Split( filescanner.Text(), " " )
		if len( namearray ) != 2 {
			fmt.Printf( "Bad input on line %d, %s\n", line, filescanner.Text() )
		} else {
			names = append( names, name{first: namearray[0], last: namearray[1]} )
		}

		line++
	}

	fmt.Printf( "\nPrinting %d names:\n", len(names) )
	for _, r := range names {
		fmt.Printf( "\t%s %s\n", r.first, r.last )
	}
}
