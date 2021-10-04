package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [32]byte{1, 2, 3, 4}
	zero(&a)
	zeroNew(&a)
	/*	c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
		//2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
		//4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
		//false
		//[32]uint8
		c := fmt.Sprintf("%T", c1)
		fmt.Printf("%T\n", c)*/
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zeroNew(ptr *[32]byte) {
	*ptr = [32]byte{}
	fmt.Println(reflect.TypeOf(ptr))
	fmt.Println(reflect.TypeOf(*ptr))
}
