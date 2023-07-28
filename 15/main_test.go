package main

import (
	"testing"
)

func BenchmarkSomeFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomeFunc()
	}
}

func BenchmarkHugeString(b *testing.B) {
	var anotherString string
	for i := 0; i < b.N; i++ {
		HugeString(&anotherString)
	}
}

func BenchmarkSmallString(b *testing.B) {
	var justString string
	for i := 0; i < b.N; i++ {
		SmallString(&justString)
	}
}
