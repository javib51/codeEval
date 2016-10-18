
package main

import "unsafe"
import "fmt"

func main() {
	var x int = 0x012345678
	var p unsafe.Pointer = unsafe.Pointer(&x)

	var endian string = ""

	if *(*byte)(p) == 0x01 {
		endian = "BigEndian"
	} else {
		endian = "LittleEndian"
	} 
	fmt.Println(endian)
}
//LittleEndian or BigEndian.
