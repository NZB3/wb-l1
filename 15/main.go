package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = 'A'
	}
	return string(b)
}

func SomeFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = v[:100]
}

func HugeString(justString *string) {
	*justString = createHugeString(1 << 10)
}

func SmallString(justString *string) {
	*justString = createHugeString(100)
}
func main() {
	SomeFunc()

	fmt.Println(justString)
}
