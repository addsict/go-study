//https://blog.golang.org/strings
package main

import "fmt"

func test1() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// sample := []byte{'\xbd', '\xb2', '\x3d', '\xbc', '\x20', '\xe2', '\x8c', '\x98'}
	fmt.Println(sample)

	fmt.Println("By Loop")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")
	fmt.Println("By % x")
	fmt.Printf("% x\n", sample)

	fmt.Println("By %+q")
	fmt.Printf("%+q\n", sample)

	fmt.Println("By Loop with %q format")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q", sample[i])
	}
}

func test2() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: %s\n", placeOfInterest)

	fmt.Printf("quoted string: %+q\n", placeOfInterest)

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func test3() {
	const nihongo = "日本語"
	for idx, runeValue := range nihongo {
		// fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
		// fmt.Printf("%q starts at byte position %d\n", runeValue, idx)
		// fmt.Printf("%s starts at byte position %d\n", string(runeValue), idx)
		fmt.Printf("%s starts at byte position %d\n", runeValue, idx)
	}
}

func test4() {
	const invalid = "\xaa\xab\xac\xad\xae\xaf"
	for idx, runeValue := range invalid {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}
}

func main() {
	// test1()
	// test2()
	test3()
	// test4()
	// fmt.Printf("length of 日本語 is %d", len("日本語"))
}
