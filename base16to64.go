package main

const b0011 = 3

var table = []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`)

// using this is SLOW
var b16map = map[byte]byte{
	byte('0'): 0,
	byte('1'): 1,
	byte('2'): 2,
	byte('3'): 3,
	byte('4'): 4,
	byte('5'): 5,
	byte('6'): 6,
	byte('7'): 7,
	byte('8'): 8,
	byte('9'): 9,
	byte('a'): 10,
	byte('b'): 11,
	byte('c'): 12,
	byte('d'): 13,
	byte('e'): 14,
	byte('f'): 15,
}

func base16to64(src16 []byte) []byte {

	if src16 == nil {
		return nil
	}

	if len(src16) == 0 {
		return make([]byte, 0, 0)
	}

	n := len(src16)
	dst := make([]byte, (n+1)*2/3+((3-n%3)%3)) // see SCRATCHWORK, and modulus of negative returns negative
	var srcIndex int
	var dstIndex int
	var stored byte
	var buffer byte
	for i := 0; i < len(src16); i++ {
		var b = src16[srcIndex] - 48 - ((src16[srcIndex] >> 6) * 39) // b16map[src16[srcIndex]] // about 18x slower :O
		srcIndex++
		switch stored {
		case 0: // i mod 3 == 0
			buffer = b << 2
			stored = 4
		case 4: // i mod 3 == 1
			dst[dstIndex] = table[buffer+(b>>2)]
			dstIndex++
			buffer = (b & b0011) << 4
			stored = 2
		case 2: // i mod 3 == 2
			dst[dstIndex] = table[buffer+b]
			dstIndex++
			stored = 0
		}
	}
	if stored > 0 {
		dst[dstIndex] = table[buffer]
		dstIndex++
		dst[dstIndex] = 61 // 61 is =
	}
	if stored > 2 {
		dstIndex++
		dst[dstIndex] = 61 // 61 is =
	}
	return dst
}

// CONVERT BASE 64 encoding to BASE 16 encoding:

var b64table = map[byte]byte{
	byte('A'): 0,
	byte('B'): 1,
	byte('C'): 2,
	byte('D'): 3,
	byte('E'): 4,
	byte('F'): 5,
	byte('G'): 6,
	byte('H'): 7,
	byte('I'): 8,
	byte('J'): 9,
	byte('K'): 10,
	byte('L'): 11,
	byte('M'): 12,
	byte('N'): 13,
	byte('O'): 14,
	byte('P'): 15,
	byte('Q'): 16,
	byte('R'): 17,
	byte('S'): 18,
	byte('T'): 19,
	byte('U'): 20,
	byte('V'): 21,
	byte('W'): 22,
	byte('X'): 23,
	byte('Y'): 24,
	byte('Z'): 25,
	byte('a'): 26,
	byte('b'): 27,
	byte('c'): 28,
	byte('d'): 29,
	byte('e'): 30,
	byte('f'): 31,
	byte('g'): 32,
	byte('h'): 33,
	byte('i'): 34,
	byte('j'): 35,
	byte('k'): 36,
	byte('l'): 37,
	byte('m'): 38,
	byte('n'): 39,
	byte('o'): 40,
	byte('p'): 41,
	byte('q'): 42,
	byte('r'): 43,
	byte('s'): 44,
	byte('t'): 45,
	byte('u'): 46,
	byte('v'): 47,
	byte('w'): 48,
	byte('x'): 49,
	byte('y'): 50,
	byte('z'): 51,
	byte('0'): 52,
	byte('1'): 53,
	byte('2'): 54,
	byte('3'): 55,
	byte('4'): 56,
	byte('5'): 57,
	byte('6'): 58,
	byte('7'): 59,
	byte('8'): 60,
	byte('9'): 61,
	byte('+'): 62,
	byte('/'): 63,
}

var b16table = []byte(`0123456789abcdef`)
var b1111 byte = 15

func base64to16(src64 []byte) []byte {

	if src64 == nil {
		return nil
	}

	if len(src64) == 0 {
		return make([]byte, 0, 0)
	}

	n := len(src64)
	var nEq int
	if n >= 2 && src64[n-2] == byte('=') {
		nEq = 2
	} else if n >= 1 && src64[n-1] == byte('=') {
		nEq = 1
	}
	// fmt.Println("nEq:", nEq)
	n -= nEq // don't bother reading the equal signs

	dst := make([]byte, 1000)
	var srcIndex int
	var dstIndex int
	var stored byte
	var buffer byte

	// fmt.Println("srcIndex:", srcIndex, "n:", n)
	// fmt.Println("------------------------------------")
	for srcIndex < n {
		// if stored != 0 {
		// 	fmt.Printf("Stored: %d Buffer: %04b\n", stored, buffer)
		// } else {
		// 	fmt.Printf("Stored: %d Buffer: XXXX\n", stored)
		// }
		switch stored {
		case 0:
			b := b64table[src64[srcIndex]]
			srcIndex++
			// fmt.Println("Storing:", b16table[b>>2])
			dst[dstIndex] = b16table[b>>2]
			buffer = (b & b0011) << 2
			stored = 2
			// fmt.Println("srcIndex:", srcIndex, "n:", n)
		case 2:
			b := b64table[src64[srcIndex]]
			srcIndex++
			// fmt.Println("Storing:", b16table[b>>2])
			dst[dstIndex] = b16table[buffer+(b>>4)]
			buffer = b & b1111
			stored = 4
			// fmt.Println("srcIndex:", srcIndex, "n:", n)
		case 4:
			// fmt.Println("Storing:", b16table[buffer])
			dst[dstIndex] = b16table[buffer]
			stored = 0
		}
		dstIndex++
		// fmt.Println("------------------------------------")
	}
	// fmt.Println("POST loop")
	// fmt.Printf("Stored: %d Buffer: %04b\n", stored, buffer)
	if stored == 4 && nEq != 1 {
		// fmt.Println("Storing:", b16table[buffer])
		dst[dstIndex] = b16table[buffer]
	}
	// fmt.Println("------------------------------------")
	return dst
}
