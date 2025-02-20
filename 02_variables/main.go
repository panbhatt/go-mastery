package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Instructions:
// Complete each task below without looking at references.
// Try to recall the syntax from memory and write valid Go code.

// 1. Declare a package-level variable.
const PI = 3.14

func main() {

	fmt.Println("PI =", PI)
	// 2. Declare a variable `name` of type string and assign it your name.
	var name string = "Pankaj Bhatt"
	fmt.Println("String NAME = ", name)

	// 3. Declare an integer variable `age` and assign it a value.
	var age int = 38
	fmt.Println("Integer Age = ", age)

	// 4. Declare a float variable `height` and assign it a value.
	var height float32 = 5.6
	fmt.Println("FLOAT -> Height = ", height)

	// 5. Declare a boolean variable `isCodingFun` and assign it `true`.
	var isCodingFun bool = true
	fmt.Println("BOOLEAN isCodingFun = ", isCodingFun)

	// 6. Declare multiple variables in a single line `city`, `country`, `population`.
	var city, country, population string = "Dehradun", "india", "8989898"
	fmt.Println("City = ", city, " Country = ", country, " Population = ", population)

	// 7. Use the shorthand syntax to declare and initialize a variable for your favorite language.
	lang := "GoLang"
	fmt.Println("Language = ", lang)

	// 8. Declare a constant `pi` with a value of 3.14159.
	const pi = 3.1459

	// 9. Swap two variables without using a temporary variable.
	a, b := 10, 20
	fmt.Println("Before Swapping ", a, b)
	a, b = b, a
	fmt.Println("After Swapping ", a, b)

	// 10. Print the **type** of a variable using the golang formatting lib
	fmt.Printf("\nType = %T  Via Reflection = ", a, reflect.TypeOf(a))

	// 11. Declare a zero-valued variable without assigning anything. Print its default value.
	var salary int64
	fmt.Println("\nDEFAULT VALUE = ", salary)

	// 12. Declare a pointer variable that points to an integer and modify the value using the pointer.
	var salaryPtr *int64 = &salary
	*salaryPtr = 58
	fmt.Println(" SALARY Via Ptr = ", *salaryPtr)

	// 13. Create a function that returns multiple values and store them in variables.
	myName, myAge, mySal := getMultipleValues()
	fmt.Println("MY Name = ", myName, " My Age = ", myAge, " My Salary = ", mySal)

	// 14. Convert an integer to a float and vice versa.
	var salaryFloat = float32(age)
	fmt.Println(" SALARY as float = ", salaryFloat)

	// 15. Use `const` with an untyped constant and assign it to different typed variables.
	const MgrSalary = "34.5"
	var mgrSalary1, _ = strconv.ParseFloat(MgrSalary, 64)
	fmt.Println("MGR SALARY = ", mgrSalary1)

	// 16. Declare and use a shadowed variable inside a function.
	shadowfn := shadowVarUsage()
	fmt.Println("1 call ", shadowfn())
	fmt.Println("2 call ", shadowfn())

	// 17. Use the new keyword to allocate memory for an integer and modify its value.
	var number *int = new(int)
	*number = 120
	fmt.Println("NEW(int) ", *number)

	// 18. Use shorthand assignment inside an `if` statement.
	if i := 0; i > -1 {
		fmt.Println("I is greater then -1 ")
	}

	// 19. Use a function closure to modify an outer variable.
	callMe := func() {
		*number++
	}
	callMe()
	fmt.Println(" Value of Number = ", *number)

}

func getMultipleValues() (string, int, float64) {
	return "BHAGWAN", 38, 678.7
}

func shadowVarUsage() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}
