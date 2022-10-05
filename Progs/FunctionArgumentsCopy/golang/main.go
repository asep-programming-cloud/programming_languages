package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	numInt := 100
	fmt.Println("Inside Main: Number is: ", numInt)
	passInt(numInt)
	fmt.Println("Inside Main: Number is: ", numInt)
	fmt.Println()

	message := "Hi Hello"
	fmt.Println("Inside Main: Message is: ", message)
	passString(message)
	fmt.Println("Inside Main: Message is: ", message)
	fmt.Println()

	userDetails := User{Name: "Ethan Hunt", Age: 45}
	fmt.Println("Inside Main: User is: ", userDetails)
	passStruct(userDetails)
	fmt.Println("Inside Main: User is: ", userDetails)
	fmt.Println()

	sliceOfInts_1 := []int{50, 500, 5000, 50000}
	fmt.Println("Inside Main: Slice of Ints is: ", sliceOfInts_1)
	passSliceInt_ModifyIndex(sliceOfInts_1)
	fmt.Println("Inside Main: Slice of Ints is: ", sliceOfInts_1)
	fmt.Println()

	sliceOfInts_2 := []int{80, 800, 8000, 80000}
	fmt.Println("Inside Main: Slice of Ints is: ", sliceOfInts_2)
	passSliceInt_NewSlice(sliceOfInts_2)
	fmt.Println("Inside Main: Slice of Ints is: ", sliceOfInts_2)
	fmt.Println()

	// numInt := 100
	// passInt(numInt)
	// fmt.Println("Inside Main: Number is: ", numInt)

	// numInt := 100
	// passInt(numInt)
	// fmt.Println("Inside Main: Number is: ", numInt)

}

func passInt(num int) {
	fmt.Println("Inside Function passInt: Number is: ", num)
	num = 200
	fmt.Println("Inside Function passInt: Number is: ", num)
}

func passString(message string) {
	fmt.Println("Inside Function passString: Message is: ", message)
	message = "Welcome"
	fmt.Println("Inside Function passString: Message is: ", message)
}

func passStruct(userDetails User) {
	fmt.Println("Inside Function passStruct: User is: ", userDetails)
	userDetails = User{Name: "Miller Roy", Age: 40}
	fmt.Println("Inside Function passStruct: User is: ", userDetails)
}

func passSliceInt_ModifyIndex(sliceOfInts []int) {
	fmt.Println("Inside Function passSliceInt_ModifyIndex: Slice of Ints is: ",
		sliceOfInts)
	sliceOfInts[1] = 23
	sliceOfInts[2] = 45
	fmt.Println("Inside Function passSliceInt_ModifyIndex: Slice of Ints is: ",
		sliceOfInts)
}

func passSliceInt_NewSlice(sliceOfInts []int) {
	fmt.Println("Inside Function passSliceInt_NewSlice: Slice of Ints is: ",
		sliceOfInts)
	sliceOfInts = []int{1, 2, 3, 4, 5}
	fmt.Println("Inside Function passSliceInt_NewSlice: Slice of Ints is: ",
		sliceOfInts)
}

func passSliceString(sliceOfStrings []string) {
	fmt.Println("Inside Function: Slice of Strings is: ", sliceOfStrings)
}

func passSliceStruct(sliceOfStructs []User) {
	fmt.Println("Inside Function: Slice of Users is: ", sliceOfStructs)
}
