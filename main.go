package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const strSize int = 4

	fmt.Println("=======================================================")

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
	fmt.Println("Done!")
	fmt.Println("=======================================================")

	fmt.Println("CONVERTING BASE 16 to BASE 64")
	start := time.Now()
	b64 := base16to64(str)
	runTime := time.Since(start)
	fmt.Println("time: ", runTime)

	if len(b64) <= 100 {
		fmt.Printf("b16: %s\n", str)
		fmt.Printf("b64: %s\n", b64)
		fmt.Printf("b64: %v\n", b64)
	}

	fmt.Printf("SIZE: %d => %d\n", len(str), len(b64))

	fmt.Println("=======================================================")

	fmt.Println("CONVERTING BASE 64 BACK TO BASE 16")
	start = time.Now()
	b16 := base64to16(b64)
	runTime = time.Since(start)
	fmt.Println("time: ", runTime)

	if len(b64) <= 100 {
		fmt.Printf("b64: %s\n", b64)
		fmt.Printf("b16: %s\n", b16)
		fmt.Printf("b16: %v\n", b16)
	}

	fmt.Printf("SIZE: %d => %d\n", len(b64), len(b16))

	fmt.Println("=======================================================")

	fmt.Println("TEST FOR EQUALITY")
	var eq = true
	if len(b16) != len(str) {
		eq = false
	} else {
		for i := 0; i < len(str); i++ {
			if b16[i] != str[i] {
				eq = false
			}
		}
	}
	if eq == true {
		fmt.Println("SUCCESS! End result matches initial string!")
	} else {
		fmt.Println("FAIL! The result string does NOT match the initial string!")
	}
	
}
