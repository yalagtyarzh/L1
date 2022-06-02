package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
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

	*num = changeBit(*num, *idx, *bit)

	fmt.Printf("Number: %d. Bit Representation: %b.\n", *num, *num)
}

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

func setBit(num int64, idx uint) int64 {
	num |= 1 << idx
	return num
}

func clearBit(num int64, idx uint) int64 {
	var mask int64 = ^(1 << idx)
	num &= mask
	return num
}
