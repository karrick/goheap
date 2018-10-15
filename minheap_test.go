package goheap_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/karrick/goheap"
)

func TestMinHeap(t *testing.T) {
	t.Run("get from empty", func(t *testing.T) {
		var mh goheap.MinHeap
		_, ok := mh.Get()
		if got, want := ok, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("get from heap with single item", func(t *testing.T) {
		var mh goheap.MinHeap
		mh.Put(13, "13")

		t.Run("returns item", func(t *testing.T) {
			v, ok := mh.Get()
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := v, "13"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("leaves heap empty", func(t *testing.T) {
			_, ok := mh.Get()
			if got, want := ok, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("items inserted out of order", func(t *testing.T) {
		var mh goheap.MinHeap
		mh.Put(42, "42")
		mh.Put(13, "13")
		mh.Put(8, "8")

		t.Run("first out is the smallest", func(t *testing.T) {
			v, ok := mh.Get()
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := v, "8"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("second out is the middle", func(t *testing.T) {
			v, ok := mh.Get()
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := v, "13"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("third out is the largest", func(t *testing.T) {
			v, ok := mh.Get()
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := v, "42"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("update key's value", func(t *testing.T) {
		t.Run("empty heap", func(t *testing.T) {
			var mh goheap.MinHeap

			mh.Update(23, func(previous interface{}, ok bool) interface{} {
				if ok {
					t.Fatalf("GOT: %v; WANT: %v", ok, false)
				}
				return "23"
			})

			for {
				v, ok := mh.Get()
				if !ok {
					t.Fatalf("Could not find key: %v", 23)
				}
				value, ok := v.(string)
				if !ok {
					t.Fatalf("GOT: %v; WANTED: %v", ok, true)
				}
				if value == "23" {
					return
				}
			}
		})

		t.Run("new key", func(t *testing.T) {
			var mh goheap.MinHeap
			mh.Put(42, "42")
			mh.Put(13, "13")
			mh.Put(8, "8")

			mh.Update(23, func(previous interface{}, ok bool) interface{} {
				if ok {
					t.Fatalf("GOT: %v; WANT: %v", ok, false)
				}
				return "23"
			})

			for {
				v, ok := mh.Get()
				if !ok {
					t.Fatalf("Could not find key: %v", 23)
				}
				value, ok := v.(string)
				if !ok {
					t.Fatalf("GOT: %v; WANTED: %v", ok, true)
				}
				if value == "23" {
					return
				}
			}
		})

		t.Run("key exists", func(t *testing.T) {
			t.Run("first", func(t *testing.T) {
				var mh goheap.MinHeap
				mh.Put(42, "42")
				mh.Put(13, "13")
				mh.Put(8, "8")
				mh.Put(23, "23")

				mh.Update(8, func(previous interface{}, ok bool) interface{} {
					if !ok {
						t.Fatalf("GOT: %v; WANT: %v", ok, true)
					}
					return "*8*"
				})

				for {
					v, ok := mh.Get()
					if !ok {
						t.Fatalf("Could not find key: %v", 23)
					}
					value, ok := v.(string)
					if !ok {
						t.Fatalf("GOT: %v; WANTED: %v", ok, true)
					}
					if value == "*8*" {
						return
					}
				}
			})

			t.Run("middle", func(t *testing.T) {
				var mh goheap.MinHeap
				mh.Put(42, "42")
				mh.Put(13, "13")
				mh.Put(8, "8")
				mh.Put(23, "23")

				mh.Update(23, func(previous interface{}, ok bool) interface{} {
					if !ok {
						t.Fatalf("GOT: %v; WANT: %v", ok, true)
					}
					return "*23*"
				})

				for {
					v, ok := mh.Get()
					if !ok {
						t.Fatalf("Could not find key: %v", 23)
					}
					value, ok := v.(string)
					if !ok {
						t.Fatalf("GOT: %v; WANTED: %v", ok, true)
					}
					if value == "*23*" {
						return
					}
				}
			})

			t.Run("last", func(t *testing.T) {
				var mh goheap.MinHeap
				mh.Put(42, "42")
				mh.Put(13, "13")
				mh.Put(8, "8")
				mh.Put(23, "23")

				mh.Update(42, func(previous interface{}, ok bool) interface{} {
					if !ok {
						t.Fatalf("GOT: %v; WANT: %v", ok, true)
					}
					return "*42*"
				})

				for {
					v, ok := mh.Get()
					if !ok {
						t.Fatalf("Could not find key: %v", 23)
					}
					value, ok := v.(string)
					if !ok {
						t.Fatalf("GOT: %v; WANTED: %v", ok, true)
					}
					if value == "*42*" {
						return
					}
				}
			})
		})
	})
}

func BenchmarkMinHeapmarkBuildHeap(b *testing.B) {
	values := rand.Perm(b.N)
	mh := new(goheap.MinHeap)
	b.ReportAllocs()
	b.ResetTimer()

	for _, v := range values {
		mh.Put(int64(v), strconv.Itoa(v))
	}
}

func BenchmarkMinHeapmarkBuildHeapWithStartingSize(b *testing.B) {
	values := rand.Perm(b.N)
	mh := goheap.NewMinHeap(len(values))

	b.ReportAllocs()
	b.ResetTimer()

	for _, v := range values {
		mh.Put(int64(v), strconv.Itoa(v))
	}
}
