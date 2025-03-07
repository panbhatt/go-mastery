package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Instructions:
// Complete each task below without looking at references.
// Try to recall the syntax from memory and write valid Go code.

func main() {

	// basic string operations
	// 1. Declare a string variable with the value "Hello, Go!" and print its length.
	var hello string = "Hello, Go!"
	fmt.Println("Length of hello =", len(hello))

	// 2. Concatenate the strings "Hello" and "World" with a space in between and print the result.
	var helloWorld = "hello" + " " + "world"
	fmt.Println(helloWorld)

	// 3. Extract and print the first 5 characters of the string "Golang is fun".
	var golangisFun string = "Golang is fun"
	fmt.Println(golangisFun[:5])

	// formatting
	// === Basic Print Functions ===
	// 1. Print "Hello, World!" using `fmt.Print`.
	fmt.Print("Hello, World!")
	// 2. Print "Hello, World!" with a newline using `fmt.Println`.
	fmt.Println("Hello, World!")
	// 3. Print "Hello, World!" using `fmt.Printf`.
	fmt.Printf("%s", "Hello, World!")

	// === String Formatting with `fmt.Sprintf` ===
	// 4. Format and store "My name is Alice and I am 25 years old." in a variable.
	var aliceString string = fmt.Sprintf("My name is %s and I am %d years old.", "Alice", 25)
	// 5. Format a float `3.14159` to 2 decimal places.
	var twoFloat string = fmt.Sprintf("%.2f", 3.14159)
	// 6. Format an integer `42` in hexadecimal representation.
	var fourtyTwoInHex string = fmt.Sprintf("%X", 42)
	fmt.Println(aliceString, " - ", twoFloat, " - ", fourtyTwoInHex)

	// === Formatting Numbers ===
	// 7. Print a number `1234567` with thousand separators.
	p := message.NewPrinter(language.English)
	p.Printf("%d\n", 1234567)
	// 8. Print a floating number `1234.5678` with scientific notation.
	fmt.Printf("%e\n", 1234.5678)
	// 9. Print a floating number `1234.5678` using default width and precision.
	fmt.Printf("%f\n", 1234.5678)

	// === Formatting Booleans ===
	// 10. Print the boolean value `true` using `fmt.Printf`.
	fmt.Printf("%t\n", true)

	// === Formatting Complex Numbers ===
	// 11. Print the complex number `3 + 4i` in default format.
	// 12. Print the real and imaginary parts separately.

	// === Formatting Width and Alignment ===
	// 13. Print "GoLang" right-aligned in a 10-character width space.
	fmt.Printf("%10s\n", "GoLang")
	// 14. Print "GoLang" left-aligned in a 10-character width space.
	fmt.Printf("%-10s\n", "GoLang")
	// 15. Print an integer `42` with leading zeros, padded to 6 digits.
	fmt.Printf("%06d", 42)

	// === Formatting Structs ===
	// 16. Print a struct `{Name: "Alice", Age: 25}` using `%v`.
	person := struct {
		Name string
		Age  int8
	}{Name: "Alice", Age: 25}
	fmt.Printf("\n%v", person)
	// 17. Print the same struct with field names using `%+v`.
	fmt.Printf("\n%+v", person)
	// 18. Print the struct in Go syntax using `%#v`.
	fmt.Printf("\n%#v", person)

	// === Formatting Slices and Maps ===
	// 19. Print a slice `[1, 2, 3, 4]` using `%v`.
	var intSlice []int = []int{1, 2, 3, 4}
	fmt.Printf("\n%v", intSlice)

	// 20. Print a map `{"name": "Alice", "age": 25}` using `%v`.
	var aliceMap map[string]interface{} = map[string]interface{}{
		"name": "Alice", "age": 25,
	}
	fmt.Printf("\n%v", aliceMap)

	// === Formatting Pointers ===
	// 21. Print the memory address of an integer variable.
	var intVar int = 10
	var ptrIntVar *int = &intVar
	fmt.Println("\n", ptrIntVar)

	// === Formatting Type Information ===
	// 22. Print the type of a variable using `%T`.
	fmt.Printf("%T", intVar)

	// === Error Formatting ===
	// 23. Create a custom error message `"File not found"` and print it.
	var err error = errors.New("File Not Found")
	fmt.Println("\n", err)

	// === Advanced Formatting ===
	// 24. Use `fmt.Fprint` to write "Hello, File!" to a string buffer.
	var sb strings.Builder
	sb.WriteString("Hello, File!")
	fmt.Println(sb.String())
	sb.Reset()
	// 25. Use `fmt.Fprintln` to write "Buffered Output" to a string buffer.
	bSBuffer := bytes.NewBufferString("")
	fmt.Fprintln(bSBuffer, "Buffered Output")
	// 26. Use `fmt.Fscanf` to read formatted input from a string.
	var nm1 string
	var age1 int8

	fmt.Fscanf(strings.NewReader("Shiva 38"), "%s %d", &nm1, &age1)
	fmt.Println(nm1, age1)

	// using the go std library package for strings
	// === Basic String Operations ===
	// 1. Clone the string "GoLang" and print it.
	var golang string = "Golang"
	fmt.Println("Clone -> ", strings.Clone(golang))
	// 2. Compare two strings "hello" and "world" and print the result.
	fmt.Println(strings.Compare("hello", "world"))

	// === String Searching ===
	// 3. Check if "Go" is contained in "I love Golang" and print the result.
	var iLGolang = "I love Golang"
	fmt.Println("Contain -> ", strings.Contains(iLGolang, "Go"))
	// 4. Check if any character from "xyz" is found in "abcdef" and print the result.
	fmt.Println("Contains Any ->", strings.ContainsAny("abcdef", "xyz"))
	// 5. Check if the string contains a rune 'o' and print the result.
	fmt.Println("Contains Rune ->", strings.ContainsRune("omygod", 'o'))
	// 6. Count how many times "is" appears in "This is a test. This is only a test."
	fmt.Println("Count -> ", strings.Count("This is a test. This is only a test", "is"))
	// 7. Find the index of "fun" in "Golang is fun" and print it.
	fmt.Println("INDEX -> ", strings.Index("Golang is fun", "fun"))
	// 8. Find the last occurrence of "is" in "This is a test. This is only a test."
	fmt.Println("Last occured -> ", strings.LastIndex("This is a test. This is only a test.", "is"))
	// 9. Find the first index of any character in "xyz" inside "abcdef" and print it.
	fmt.Println("First Index", strings.IndexAny("abcdef", "xyz"))
	// 10. Find the last index of any character in "xyz" inside "abcdef" and print it.
	fmt.Println("First Index", strings.LastIndexAny("abcdef", "xyz"))

	// === String Modification ===
	// 11. Cut the string "hello:world" at ":" and print both parts.
	fmt.Println(strings.Split("hello:world", ":"))
	// 12. Cut the prefix "Hello " from "Hello World" and print the result.
	fmt.Println(strings.CutPrefix("Hello World", "Hello "))
	// 13. Cut the suffix " World" from "Hello World" and print the result.
	fmt.Println(strings.CutSuffix("Hello World", " World"))
	// 14. Replace "bad" with "good" in "This is a bad idea" and print it.
	fmt.Println(strings.Replace("this is a bad idea", "bad", "good", 1))
	// 15. Replace all occurrences of "a" with "o" in "banana" and print the result.
	fmt.Println(strings.ReplaceAll("banana", "a", "o"))
	// 16. Repeat the string "Go" 5 times and print the result.
	fmt.Println(strings.Repeat("Go", 5))
	// 17. Use `strings.Replacer` to replace "Java" with "Go" and "C++" with "Rust" in "I love Java and C++".
	rp := strings.NewReplacer("Java", "Go", "C++", "Rust")
	fmt.Println("Replacer ->", rp.Replace("I love Java and C++"))

	// === String Splitting & Joining ===
	// 18. Split "Go is simple and powerful" into words and print the slice.
	allWords := strings.FieldsSeq("Go is simple and powerful")
	for word := range allWords {
		fmt.Printf(": %s", word)
	}
	// 19. Use `FieldsFunc` to split "apple,banana,,orange" by commas but ignore empty parts.
	commaSep := strings.FieldsFunc("apple,banana,orange", func(c rune) bool { return c == ',' })
	fmt.Println()
	for _, word := range commaSep {
		fmt.Printf(": %s", word)
	}
	// 20. Use `SplitAfter` to split "one,two,three" after each comma.
	fmt.Println("\n", strings.SplitAfter("apple,banana,orange", ","))
	// 21. Use `SplitN` to split "apple-banana-orange" at "-" but limit the splits to 2 parts.
	fmt.Println("\n", strings.SplitAfterN("apple,banana,orange", ",", 2))
	// 22. Join a slice of strings `["Go", "is", "great"]` into a single string separated by spaces.
	fmt.Println(strings.Join([]string{"go", "is", "Great"}, " "))

	// === String Case Modifications ===
	// 23. Convert "golang is fun" to title case and print it.
	fmt.Println(strings.Title("golang is fun"))
	fmt.Println(cases.Title(language.English, cases.Compact).String("golang is fun"))
	// 24. Convert "Go Is Fun" to lowercase and print it.
	fmt.Println(strings.ToLower("golang is fun"))
	// 25. Convert "Go Is Fun" to uppercase and print it.
	fmt.Println(strings.ToUpper("golang is fun"))

	// === String Trimming ===
	// 26. Trim all spaces from "  Go  " and print the result.
	fmt.Println(strings.Trim("   Go    ", " "))
	// 27. Trim only leading spaces from "    Hello World!" and print the result.
	fmt.Println(strings.TrimLeft("   Hello World!", " "))
	// 28. Trim only trailing spaces from "Hello World!    " and print the result.
	fmt.Println(strings.TrimRight("Hello World!", " "))
	// 29. Trim the prefix "Hello" from "Hello, World!" and print the result.
	fmt.Println(strings.TrimPrefix("Hello, World!", "Hello"))
	// 30. Trim the suffix " World!" from "Hello, World!" and print the result.
	fmt.Println(strings.TrimSuffix("Hello, World!", "World!"))
	// 31. Trim spaces using `TrimSpace` from "  Too much space  " and print the result.
	fmt.Println(strings.TrimSpace("  Too much space   "))

	// === UTF-8 and Rune Operations ===
	// 32. Use `Map` to convert all vowels in "golang" to uppercase.
	fmt.Println(strings.Map(func(r rune) rune {
		if r >= 97 && r <= 122 {
			return rune(r - 32)
		} else {
			return r
		}
	}, "golang"))
	// 33. Convert invalid UTF-8 sequences to valid ones using `ToValidUTF8`.
	fmt.Println(strings.ToValidUTF8("SHIVARATRI \xc5 is today", " "))
	// 34. Use `IndexFunc` to find the first vowel in "golang".
	fmt.Println(strings.IndexFunc("golang", func(r rune) bool { return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' }))
	// 35. Use `TrimFunc` to remove all leading numbers from "12345hello".
	fmt.Println(strings.TrimFunc("12345hello", func(r rune) bool {
		if r >= 48 && r <= 57 {
			return true
		}
		return false
	}))

	// === Advanced String Builders ===
	// 36. Use `strings.Builder` to build the string "GoLang is awesome!" efficiently and print it.
	var sb3 strings.Builder
	sb3.WriteString("Golang")
	sb3.WriteRune(' ')
	sb3.WriteString("is")
	sb3.WriteRune(' ')
	sb3.WriteString("Awesome")
	fmt.Println("Final String -> ", sb3.String())
	// 37. Use `strings.Builder` to append multiple lines efficiently.
	// 38. Use `strings.Builder.Grow()` to preallocate memory for a string.
	var sb4 strings.Builder
	sb4.WriteString("SHIVA RATRI")
	sb4.Grow(10)
	sb4.WriteString(" PARVATI")
	fmt.Println("SB 4 = ", sb4.String(), "\n")

	// === String Reading ===
	// 39. Use `strings.Reader` to read "Hello, Go!" byte by byte and print each byte.
	sR := strings.NewReader("Hello, Go!")
	for {
		bt, ok := sR.ReadByte()
		if ok == nil {
			fmt.Print(bt, " ")
		} else {
			break
		}
	}

	// 40. Use `strings.Reader` to read "Hello, Go!" rune by rune and print each rune.
	fmt.Println("\n")
	sR1 := strings.NewReader("Hello, Go!")
	for {
		bt, _, ok := sR1.ReadRune()
		if ok == nil {
			fmt.Print(string(rune(bt)), " ")
		} else {
			break
		}
	}
	// 41. Use `strings.Reader.Seek()` to move the read position in "Hello, Go!".
	fmt.Println("\n")
	sR2 := strings.NewReader("Hello, Go!")
	sR2.Seek(2, io.SeekCurrent)
	for {
		bt, _, ok := sR2.ReadRune()
		if ok == nil {
			fmt.Print(string(rune(bt)), " ")
		} else {
			break
		}
	}
	// 42. Use `strings.Reader.Size()` to get the total size of "Hello, Go!".
	sR3 := strings.NewReader("Hello, Go!")
	fmt.Println("\n ", sR3.Size())

	// === Miscellaneous Functions ===
	// 43. Use `HasPrefix` to check if "Gopher" starts with "Go".
	fmt.Println("Gopher has go ->", strings.HasPrefix("Gopher", "Go"))
	// 44. Use `HasSuffix` to check if "Gopher" ends with "her".
	fmt.Println("Gopher has her at last -> ", strings.HasSuffix("Gopher", "her"))
	// 45. Use `CutPrefix` to remove "Go" from "Gopher".
	gopherRemovePrefix, _ := strings.CutPrefix("Gopher", "Go")
	fmt.Println("Gopher Remove prefix go ->", gopherRemovePrefix)
	// 46. Use `CutSuffix` to remove "er" from "Gopher"
	gopherRemoveSufix, _ := strings.CutSuffix("Gopher", "her")
	fmt.Println("Gopher Remove Suffix her ->", gopherRemoveSufix)

	// === Deprecated or Rare Functions ===
	// 47. Use `Title()` to convert "go is fun" to title case (Deprecated, but still try it).
	fmt.Println(strings.Title("go is fun"))

	// using the go std library package for string conversion
	// === Integer and Float Conversions ===
	// 1. Convert the string "42" to an integer and print the result.
	fortyTwoNumber, err := strconv.Atoi("42")
	if err == nil {
		fmt.Println(fortyTwoNumber)
	}
	// 2. Convert an integer 123 to a string and print it.
	fmt.Println(strconv.Itoa(123))
	// 3. Convert the string "3.14" to a float and print the result.
	piInFloat, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("PI in Float ", piInFloat)
	// 4. Convert a float 9.81 to a string with 2 decimal places.
	fmt.Println(fmt.Sprintf("%0.2f", 9.82))
	// 5. Convert an integer to a hexadecimal string.
	fmt.Printf("%0x\n", 100)
	// 6. Convert a hexadecimal string "1a" to an integer.
	a1InHex, _ := strconv.ParseInt("1a", 16, 32)
	fmt.Println("1a -> HEX ", a1InHex)

	// === Boolean Conversions ===
	// 7. Convert the string "true" to a boolean and print the result.
	// 8. Convert a boolean `false` to a string and print it.

	// === Appending to Byte Slices ===
	// 9. Append an integer to a byte slice and print the result.
	// 10. Append a float to a byte slice with a precision of 3 decimal places.
	// 11. Append a boolean to a byte slice.

	// === Quoting and Escaping ===
	// 12. Quote the string `"Hello, Go!"` and print it.
	// 13. Quote the rune `'♥'` and print it.
	// 14. Convert `"Hello, 世界"` to an ASCII-safe quoted string.
	// 15. Convert a rune `'世'` to an ASCII-safe quoted string.
	// 16. Check if a string can be safely backquoted (raw string).

	// === Parsing and Unquoting ===
	// 17. Unquote the string `"\"GoLang\""` and print it.
	// 18. Extract the first quoted value from `"'Hello' 'World'"` and print it.
	// 19. Parse a quoted string and return its rune and remaining text.

	// === Handling Numeric Errors ===
	// 20. Try to parse an invalid number `"abc123"` and handle the error.
	// 21. Use `NumError` to unwrap an error from an invalid conversion.

	// === ASCII & Graphic Character Checks ===
	// 22. Check if the rune `'a'` is printable.
	// 23. Check if the rune `'@'` is a graphic character.

	// === Basic Rune Operations ===
	// 1. Declare a rune variable with the character 'A' and print it.
	// 2. Print the Unicode code point of the rune 'A' using `%U` format.
	// 3. Convert the rune 'A' to an integer and print it.
	// 4. Convert an integer 65 to a rune and print it.

	// === Rune Iteration ===
	// 5. Iterate over the string "GoLang" and print each rune.
	// 6. Iterate over the string "你好, 世界" and print each rune with its Unicode value.

	// === Rune Conversion ===
	// 7. Convert the rune 'a' to uppercase and print it.
	// 8. Convert the rune 'A' to lowercase and print it.

	// === Rune Classification ===
	// 9. Check if the rune '9' is a digit.
	// 10. Check if the rune 'A' is a letter.
	// 11. Check if the rune ' ' (space) is a whitespace character.
	// 12. Check if the rune '#' is a graphic character.
	// 13. Check if the rune '♥' is printable.

	// === Working with UTF-8 ===
	// 14. Convert the string "Go" to a slice of runes and print it.
	// 15. Convert a slice of runes `[71, 111]` back to a string and print it.
	// 16. Find the number of runes in "你好, 世界".

	// === Modifying Strings with Runes ===
	// 17. Replace all lowercase vowels in "golang" with uppercase ones using `strings.Map`.
	// 18. Remove all non-letter characters from "Go 123!".

	// === Advanced Rune Manipulation ===
	// 19. Use `utf8.DecodeRuneInString` to decode the first rune of "GoLang".
	// 20. Use `utf8.RuneCountInString` to count the runes in "GoLang".
	// 21. Use `utf8.ValidRune` to check if a rune is valid.
	// 22. Create a rune slice `[]rune{'你', '好'}` and append another rune to it.
	// 23. Reverse the string "GoLang" using runes.

	// === Unicode Basic Properties ===
	// 1. Print the Unicode code point of the rune '你' using `%U`.
	// 2. Check if the rune 'A' is a letter using `unicode.IsLetter`.
	// 3. Check if the rune '9' is a digit using `unicode.IsDigit`.
	// 4. Check if the rune ' ' (space) is a whitespace character using `unicode.IsSpace`.
	// 5. Check if the rune '#' is a punctuation mark using `unicode.IsPunct`.

	// === Unicode Category Classification ===
	// 6. Check if the rune '©' is a symbol.
	// 7. Check if the rune 'Ω' belongs to the Greek script using `unicode.In`.

	// === Unicode Normalization (NFC, NFD, NFKC, NFKD) ===
	// 8. Normalize a string "é" to NFC and print the result.
	// 9. Normalize a string "é" to NFD and print the result.
	// 10. Check if two visually identical strings are Unicode-normalized.

	// === UTF-8 Encoding & Decoding ===
	// 11. Convert a string "Hello, 世界" to a slice of UTF-8 bytes and print it.
	// 12. Convert a UTF-8 byte slice back to a string.
	// 13. Decode the first rune in a UTF-8 encoded string.
	// 14. Count the number of runes in "你好, 世界".
	// 15. Check if a given byte slice is valid UTF-8.

	// === Handling Invalid UTF-8 ===
	// 16. Fix an invalid UTF-8 string by replacing errors with '?'.
	// 17. Remove all non-UTF-8 characters from a corrupted string.

	// === Advanced Unicode Handling ===
	// 18. Iterate over a string containing mixed language characters and print each rune with its Unicode category.
	// 19. Convert a string to its ASCII equivalent (removing accents) using `runenames` or `unicode` mapping.
	// 20. Strip diacritics (accents) from a string (e.g., "é" → "e").
	// 21. Reverse a string while keeping Unicode characters intact.

}
