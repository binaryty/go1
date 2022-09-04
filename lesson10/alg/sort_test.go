package alg_test

import (
	"github.com/binaryty/go1/lesson10/alg"
	"testing"
)

var arr = []int{13, -5, 9, 12, 13, 0, 21, -8, 55, 32, 99}

func TestInsSort(t *testing.T) {
	a := arr
	a = alg.InsSort(a)
	if !alg.IsSort(a) {
		t.Errorf("sorted %v", arr)
		t.Errorf("   got %v", a)
	}
}

func TestBubbleSort(t *testing.T) {
	a := arr
	a = alg.BubbleSort(a)
	if !alg.IsSort(a) {
		t.Errorf("sorted %v", arr)
		t.Errorf("   got %v", a)
	}
}

func TestQuickSort(t *testing.T) {
	a := arr
	a = alg.QuickSort(a)
	if !alg.IsSort(a) {
		t.Errorf("sorted %v", arr)
		t.Errorf("   got %v", a)
	}
}

func BenchmarkInsSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		a := make([]int, 10000)
		alg.FillArr(a)
		b.StartTimer()
		alg.InsSort(a)
		b.StopTimer()
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		a := make([]int, 10000)
		alg.FillArr(a)
		b.StartTimer()
		alg.BubbleSort(a)
		b.StopTimer()
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		a := make([]int, 10000)
		alg.FillArr(a)
		b.StartTimer()
		alg.QuickSort(a)
		b.StopTimer()
	}
}
