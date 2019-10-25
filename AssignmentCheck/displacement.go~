// by Ardavan Izadiyar @ardawanizadi
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Displacement struct {
	Acceleration float64
	Velocity     float64
	Displacement float64
	Time         float64
}

var input_data Displacement
var input_map = make(map[string]float64)

func main() {
	inputs := reflect.ValueOf(input_data)

	for i := 0; i < inputs.NumField(); i++ {
		scanner := bufio.NewScanner(os.Stdin)
		field := inputs.Type().Field(i).Name
		fmt.Print(field, ": ")
		scanner.Scan()

		if v, err := strconv.Atoi(scanner.Text()); err == nil {
			input_map[field] = float64(v)
		} else if v, err := strconv.ParseFloat(scanner.Text(), 64); err == nil {
			input_map[field] = float64(v)
		} else if strings.ToLower(scanner.Text()) == "x" {
			return
		} else {
			fmt.Println(">> please enter inreger or float numbers. press X to exit")
			i--
		}
	}

	input_data.Acceleration = input_map["Acceleration"]
	input_data.Velocity = input_map["Velocity"]
	input_data.Displacement = input_map["Displacement"]
	input_data.Time = input_map["Time"]

	fn := GenDisplaceFn(input_data.Acceleration, input_data.Velocity, input_data.Displacement)

	fmt.Println(fn(input_data.Time))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*(a*math.Pow(t, 2)) + (v * t) + s
	}
}
