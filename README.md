# Essential Go Commands

| Command                  | Description                                                        |
|--------------------------|--------------------------------------------------------------------|
| `go run main.go`         | Runs the Go file directly (compiles and executes).                 |
| `go build main.go`       | Builds the Go file into an executable.                             |
| `go mod init <module>`   | Initializes a new Go module (creates go.mod).                      |
| `go mod tidy`            | Cleans up dependencies in go.mod and go.sum.                       |
| `go get <package>`       | Downloads and adds a dependency to your project.                   |
| `go test`                | Runs all tests in the current directory.                           |
| `go fmt`                 | Formats your Go source files according to Go standards.            |
| `go clean`               | Removes object files and cached files.                             |
| `go doc <pkg or symbol>` | Shows documentation for a package or symbol.                       |
| `go list`                | Lists packages in your module or workspace.                        |
| `go env`                 | Displays Go environment variables.                                 |
| `go version`             | Shows the installed Go version.                                    |


# Annotations
```go
// Variables types
// Primitives
// int
var a = 10       // Infered type int
var b int = 20   // Explicit type int
var c int64 = 30 // Explicit type int64( size = 2.n^(64-1)) -> -32767 to 32767

// uint -> positive integers only
var aa = 10      // Infered type uint
var bb uint = 20 // Explicit type uint

// string
var d = "Hello"        // Infered type string
var e string = "World" // Explicit type string

// float
var floatVar = 3.14         // Infered type float64
var floatVar float32 = 2.71 // Explicit type float32(worst precision but less memory)

// alias
var alias1 rune = 'A' // rune is an alias for int32
var alias2 byte = 'B' // byte is an alias for uint8

// boolean
var isTrue = true // Infered type bool

// Another way to declare variables
variable := "This is a variable" // Short variable declaration with infered type string


```
# Running a file
### First build to an .exe and then run
```go
go build filename.go
./filename
```
### or just execute
```bash
go run filename.go
```
### Example code of a function call with loop/nil error treatment
```go
package main

import (
	"fmt"
)

func main() {

	// Doing integer division
	numerator := 10
	denominator := 0
	result, remainder, error := intDivision(numerator, denominator)
	// One way to check conditions is with if-else
	if error != nil {
		fmt.Println("Error:", error)
	} else if remainder == 0 {
		fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	} else {
		fmt.Printf("The result of %d divided by %d is %d and the remainder is %d\n", numerator, denominator, result, remainder)
	}
	// another way to do this is using switch
	switch {
	case error != nil:
		fmt.Println("Error:", error)
	case remainder == 0:
		fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	default:
		fmt.Printf("The result of %d divided by %d is %d and the remainder is %d\n", numerator, denominator, result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var error error
	if denominator == 0 {
		return 0, 0, fmt.Errorf("denominator cannot be zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator

	// nil is used to indicate that there is no error during the execution of the program
	// if error return nil, it means that there is no error
	// return the result of the division, remainder an nil
	return result, remainder, error
}
```
### Arrays, Slices, Maps and Speed Test between traditional and make constructor
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Array()

	// Slice()

	// Map()

	speedTest() // Speed test for slice preallocation
}

func Array() {
	// The type specify how much memory is allocated for the variable
	var integerArray [10]int16 = [10]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Array of integers initialized with a fixed size of 10 [2 bytes each]
	intArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}                    // Short way to initialize array of integers initialized with a fixed size of 10 [4 bytes each]
	int64Array := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}                  // Short way to initialize array of integers with inferred size [8 bytes each]

	fmt.Printf("Integer Array: %v, Type: %T\n", integerArray, integerArray[0])
	for i := range integerArray { // Works exactly as -> for i := 0; i < len(integerArray); i++
		fmt.Printf("Position: %d, Value: %d, Memory Position: %p \n", i, integerArray[i], &integerArray[i]) // %p -> pointer
	}

	fmt.Printf("Integer Array: %v, Type: %T\n", intArray, intArray[0])
	for j := range intArray {
		fmt.Printf("Position: %d, Value: %d, Memory Position: %p \n", j, intArray[j], &intArray[j]) // %p -> pointer
	}

	fmt.Printf("Integer Array: %v, Type: %T\n", int64Array, int64Array[0])
	for k := range int64Array {
		fmt.Printf("Position: %d, Value: %d, Memory Position: %p \n", k, int64Array[k], &int64Array[k]) // %p -> pointer
	}

}

func Slice() {
	// Slices are a more flexible way to work with arrays in Go
	slice := []int64{1, 2} // Slice of integers initialized with a variable size
	fmt.Print("--------------- SLICE ---------------\n")
	fmt.Printf("Initial slice %v length %d and capacity %d\n", slice, len(slice), cap(slice)) // len() -> length, cap() -> capacity
	slice = append(slice, 12, 7, 45)
	fmt.Printf("Slice after appending values %v and their length %d and capacity %d\n", slice, len(slice), cap(slice)) // Appending values to the slice

	slice2 := []int64{201, 38, 21} // Creating another slice of integers initialized with a variable size to append to the first one

	slice = append(slice, slice2...)                                                                                          // Appending another slice to the first one with spread operator
	fmt.Printf("Slice after appending another slice %v and their length %d and capacity %d\n", slice, len(slice), cap(slice)) // Printing the final slice

	fmt.Print("--------------- SLICE WITH MAKE ---------------\n")
	makeSlice := make([]int64, 5, 10) // Creating a slice with make() function, length of 5 and capacity of 10
	fmt.Printf("Slice created with make %v and their length %d and capacity %d\n", makeSlice, len(makeSlice), cap(makeSlice))
	makeSlice = append(makeSlice, 1, 2, 3, 4, 5)                                                                                            // Appending values to the slice created with make()
	fmt.Printf("Slice after appending values %v and their length %d and capacity %d\n", makeSlice, len(makeSlice), cap(makeSlice))          // Printing the final slice created with make()
	makeSlice = append(makeSlice, slice...)                                                                                                 // Appending the first slice to the slice created with make()
	fmt.Printf("Slice after appending the first slice %v and their length %d and capacity %d\n", makeSlice, len(makeSlice), cap(makeSlice)) // Printing the final slice created with make() after appending the first slice
}

func Map() {
	// Maps are a way to store key-value pairs in Go
	fmt.Print("--------------- MAP ---------------\n")
	//myMap := make(map[string]uint) // Creating a map with make() function

	map2 := map[string]uint{ // Creating another map with key-value pairs
		"Matheus": 25,
		"João":    30,
		"Maria":   28,
		"Pedro":   22,
		"Lucas":   27,
	}

	fmt.Printf("ACCESSING MAP VALUE WITH KEY %v\n", map2["Matheus"])
	fmt.Println(map2["Matheus"]) // Accessing a value in the map with the key -> Map always return value
	fmt.Print("CHECKING IF KEY EXISTS\n")
	age, ok := map2["Matheus"] // Accessing a value in the map with the key and checking if it exists
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Key not found")
	}
	// Deleting a key-value pair from the map
	// delete(map2, "Matheus") // Delete method receive (map, key)
	// Iterating over a map
	for key, value := range map2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value) // Printing the key and value of the map with random order
	}
}

func speedTest() {
	n := 1000000                    // Number of iterations for the speed test
	testSlice := []int{}            // Slice without preallocation
	testSlice2 := make([]int, 0, n) // Preallocated slice with capacity n

	fmt.Printf("Total time without preallocation: %v\n", timeTrack(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeTrack(testSlice2, n))
}

func timeTrack(slice []int, n int) time.Duration {
	t0 := time.Now()     // Start time tracking
	for len(slice) < n { // Loop until the slice reaches the desired length
		slice = append(slice, 1) // Appending values to the slice
	}
	return time.Since(t0) // Return the elapsed time
}
```
### Strings
```go
package main

import (
	"fmt"
	"strings"
)

func main() {

	String()

	concatenateString()
}

func String() {
	// Dealing with strings
	myString := "Matheus"
	index := myString[0]
	fmt.Printf("Hex: %x, Char: %c, Decimal: %d, Octal: %o, Type: %T\n", index, index, index, index, index) // Will return the value based on ASCII table	-> x = hex, c -> char, d -> decimal, %o -> octal
	for k, v := range myString {
		fmt.Println(k, v)
	}

	// Now with runes -> int32
	myRune := []rune("Matheus")
	indexRune := myRune[0]
	fmt.Printf("Hex: %x, Char: %c, Decimal: %d, Octal: %o, Type: %T\n", indexRune, indexRune, indexRune, indexRune, indexRune) // Will return the value based on Unicode table	-> x = hex, c -> char, d -> decimal, %o -> octal
	for k, v := range myRune {
		fmt.Println(k, v)
	}
}

func concatenateString() {
	// Concatenating strings
	newString := []string{"M", "a", "t", "h", "e", "u", "s"}
	concatString := ""
	for i := range newString {
		concatString += newString[i]
	}
	fmt.Println("Concatenated String:", concatString) // Worst case because it creates a new string every time you concatenate, so it is better to use strings.Builder or bytes.Buffer for large strings

	//  Using strings.Builder
	newStringBuilder := []string{"M", "a", "t", "h", "e", "u", "s"}
	stringBuilder := strings.Builder{}
	for i := range newStringBuilder {
		stringBuilder.WriteString(newStringBuilder[i])
	}
	stringBuilderString := stringBuilder.String()                         // Only now the string is created
	fmt.Println("Concatenated String with Builder:", stringBuilderString) // Better performance for large strings

}
```
### Pointers
```go
// Pointers are variables that store the memory address of another variable.
// & is used to get the address of a variable.
// * is used to declare a pointer or to dereference (access the value pointed by) a pointer.

a := 10
p := &a // p is a pointer to a

fmt.Println("Value of a:", a)      // 10
fmt.Println("Address of a:", &a)   // address of a
fmt.Println("Value of p:", p)      // same address as &a
fmt.Println("Value pointed by p:", *p) // 10
fmt.Println("Address of b:", &p)   // address of p

// Changing the value of a via pointer
*p = 20
fmt.Println("New value of a:", a) // 20

// Pointers with functions
func increment(x *int) {
    *x = *x + 1
}

value := 5
increment(&value)
fmt.Println("Incremented value:", value) // 6

// Pointers with structs
type Person struct {
    Name string
}

func changeName(p *Person, newName string) {
    p.Name = newName
}

person := Person{Name: "Maria"}
changeName(&person, "John")
fmt.Println("Changed name:", person.Name) // John

// The zero value of a pointer is nil
var ptr *int
if ptr == nil {
    fmt.Println("ptr is nil")
}




































































































































































































































































































































































































































































































































































































































































































```
