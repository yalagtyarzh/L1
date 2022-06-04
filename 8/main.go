package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Из cli получаем число, индекс изменяемого бита и на какой бит меняем
	num := flag.Int64("num", 1<<10, "")
	idx := flag.Uint("idx", 8, "")
	bit := flag.Uint("bit", 1, "")

	if *bit != 1 && *bit != 0 {
		log.Fatal("invalid bit value")
	}

	if *idx < 0 || *idx > 63 {
		log.Fatal("invalid index")
	}

	fmt.Printf("Number: %d. Bit Representation: %b.\n", *num, *num)

	// Меняем бит
	*num = changeBit(*num, *idx, *bit)

	fmt.Printf("Number: %d. Bit Representation: %b.\n", *num, *num)
}

// changeBit возвращает значение, представляющее из себе num, в котором поменяли бит на bit
// в позиции idx
func changeBit(num int64, idx, bit uint) int64 {
	switch bit {
	case 0:
		return clearBit(num, idx)
	case 1:
		return setBit(num, idx)
	default:
		return num
	}
}

// setBit ставит бит в num на позицию idx
func setBit(num int64, idx uint) int64 {
	num |= 1 << idx
	return num
}

// clearBit очищает бит в num на позицию idx
func clearBit(num int64, idx uint) int64 {
	var mask int64 = ^(1 << idx)
	num &= mask
	return num
}
