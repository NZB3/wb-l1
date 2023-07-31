package main

import (
	"fmt"
)

func createString(size int) string {
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = 'A'
	}
	return string(b)
}

func ToHugeString(justString *string) {
	*justString = createString(1024)
}

func ToSmallString(justString *string) {
	*justString = createString(100)
}
func main() {
	var justString string

	ToHugeString(&justString)
	fmt.Println(justString)

	ToSmallString(&justString)
	fmt.Println(justString)

}

// var justString string // глобальная переменная
// func someFunc() { // не понятное название функции
// не понятно что она изменяет какую-то пременную
//		v := createHugeString(1 << 10) // трудочитаемый размер строки
//	 	justString = v[:100] // утечка памяти
// }
// func main() {
//		someFunc()
// }
