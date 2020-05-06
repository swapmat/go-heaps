package benchmark

import (
	"container/heap"
	"math/rand"
	"testing"

	"github.com/theodesp/go-heaps"
	"github.com/theodesp/go-heaps/binomial"
	"github.com/theodesp/go-heaps/fibonacci"
	"github.com/theodesp/go-heaps/leftist"
	"github.com/theodesp/go-heaps/pairing"
	rpheap "github.com/theodesp/go-heaps/rank_pairing"
)

type Stat struct {
	data     int
	priority uint32
}

func (s Stat) Compare(than go_heaps.Item) int {
	other := than.(*Stat)
	if s.priority < other.priority {
		return -1
	} else if s.priority == other.priority {
		return 0
	}
	return 1
}

var heap1 binomial.BinomialHeap

func BenchmarkBinomialHeapInsert(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap1.Insert(&Stat{data: n, priority: rand.Uint32()})
	}
}

func BenchmarkBinomialHeapFind(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap1.FindMin()
	}
}

func BenchmarkBinomialHeapDelete(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap1.DeleteMin()
	}
}

var heap2 fibonacci.FibonacciHeap

func BenchmarkFibonacciHeapInsert(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap2.Insert(&Stat{data: n, priority: rand.Uint32()})
	}
}

func BenchmarkFibonacciHeapFind(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap2.FindMin()
	}
}

func BenchmarkFibonacciHeapDelete(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap2.DeleteMin()
	}
}

var heap3 leftist.LeftistHeap

func BenchmarkLeftistHeapInsert(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap3.Insert(&Stat{data: n, priority: rand.Uint32()})
	}
}

func BenchmarkLeftistHeapFind(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap3.FindMin()
	}
}

func BenchmarkLeftistHeapDelete(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap3.DeleteMin()
	}
}

var heap4 *rpheap.RPHeap

func BenchmarkRPHeapInsert(b *testing.B) {
	b.ReportAllocs()
	heap4 = rpheap.New()
	for n := 0; n < b.N; n++ {
		heap4.Insert(&Stat{data: n, priority: rand.Uint32()})
	}
}

func BenchmarkRPHeapFind(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap4.FindMin()
	}
}

func BenchmarkRPHeapDelete(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap4.DeleteMin()
	}
}

var heap5 *pairing.PairHeap

func BenchmarkPairHeapInsert(b *testing.B) {
	heap5 = pairing.New()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap5.Insert(&Stat{data: n, priority: rand.Uint32()})
	}
}

func BenchmarkPairHeapFind(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap5.FindMin()
	}
}

func BenchmarkPairHeapDelete(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap5.DeleteMin()
	}
}

type MinHeap []*Stat

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].priority < h[j].priority
}

func (h *MinHeap) Push(val interface{}) {
	item := val.(*Stat)
	*h = append(*h, item)
}

func (h *MinHeap) FindMin() interface{} {
	old := *h
	n := len(*h)
	val := old[n-1]
	return val
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	val := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return val
}

var heap6 *MinHeap

func BenchmarkSimpleHeapInsert(b *testing.B) {
	b.ReportAllocs()
	heap6 = &MinHeap{}
	heap.Init(heap6)
	for n := 0; n < b.N; n++ {
		heap.Push(heap6, &Stat{data: n, priority: rand.Uint32()})
	}
}

func BenchmarkSimpleHeapFind(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap6.FindMin()
	}
}

func BenchmarkSimpleHeapDelete(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		heap.Pop(heap6)
	}
}
