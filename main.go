package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const strSize int = 1000

	// generate a random string of hex chars
	fmt.Println("GENERATING RANDOM NUMBERS")
	str := make([]byte, strSize) // 0000 0001 0010 -> AS
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(str); i++ {
		val := rand.Intn(16)
		if val < 10 {
			str[i] = byte(val + 48) // 48 is 0
		} else {
			str[i] = byte(val - 10 + 97) // 97 is a
		}
	}

	fmt.Println("CONVERTING BASE 16 to BASE 64")
	start := time.Now()
	b64 := base16to64(str)
	runTime := time.Since(start)
	fmt.Println("time: ", runTime)

	if len(b64) <= 100 {
		fmt.Println("=======================================================")
		fmt.Printf("%s\n", str)
		fmt.Println("=======================================================")
		fmt.Printf("%s\n", b64)
		fmt.Println(b64)
		fmt.Println("=======================================================")
	}

	fmt.Println("*******************************************************")
	fmt.Println("=======================================================")
	fmt.Println("*******************************************************")

	fmt.Printf("%d => %d\n", len(str), len(b64))

	fmt.Println("CONVERTING BASE 64 BACK TO BASE 16")
	start = time.Now()
	b16 := base64to16(b64)
	runTime = time.Since(start)
	fmt.Println("time: ", runTime)

	if len(b64) <= 100 {
		fmt.Println("=======================================================")
		fmt.Printf("%s\n", b64)
		fmt.Println("=======================================================")
		fmt.Printf("%s\n", b16)
		fmt.Println(b16)
		fmt.Println("=======================================================")
	}

	fmt.Printf("%d => %d\n", len(b64), len(b16))

	// b16 := []byte(`0000`)
	// b64 := base16to64(b16)
	// fmt.Printf("\"%s\" => \"%s\"\n\n", b16, b64)

	// newb16 := base64to16(b64)
	// fmt.Printf("\n\"%s\" => \"%s\"\n", b64, newb16)

}
