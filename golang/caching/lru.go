package main

// https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"math"
	"math/rand"
)

type HashMap[K comparable, V any] struct {
	arr   []LinkedList[K, V]
	size  int
	a     int
	b     int
	prime int
}

type LinkedList[K comparable, V any] struct {
	head *Node[K, V]
	size int
}

type Node[K comparable, V any] struct {
	key         K
	data        V
	accessCount int
	next        *Node[K, V]
}

func (hashMap HashMap[K, V]) Get(key K) (V, bool) {
	var idx int = hashMap.getHash(key)

	var linkedList LinkedList[K, V] = hashMap.arr[idx]

	if linkedList.size == 0 {
		return Node[K, V]{}.data, false
	}

	var head *Node[K, V] = linkedList.head

	for head != nil {
		if head.key == key {
			head.accessCount++
			return head.data, true
		}
		head = head.next
	}

	return Node[K, V]{}.data, false
}

func (hashMap HashMap[K, V]) Put(key K, value V) {
	var idx int = hashMap.getHash(key)

	var linkedList *LinkedList[K, V] = &hashMap.arr[idx]

	if linkedList.size == 0 {
		linkedList.head = &Node[K, V]{key: key, data: value, accessCount: 0, next: nil}
		linkedList.size++
		return
	}

	if linkedList.size < 5 {
		linkedList.insert(linkedList, &Node[K, V]{key: key, data: value, accessCount: 0, next: nil})
		return
	}

	var prevLru *Node[K, V] = linkedList.findPrevLru(linkedList)

	if prevLru.next == nil {
		prevLru.key = key
		prevLru.accessCount = 0
		prevLru.data = value
		prevLru.next = nil
	} else {
		prevLru.key = key
		prevLru.accessCount = 0
		prevLru.data = value
		prevLru.next = prevLru.next.next
	}
}

func (LinkedList[K, V]) findPrevLru(linkedList *LinkedList[K, V]) *Node[K, V] {
	var lru *Node[K, V] = linkedList.head
	var compare *Node[K, V] = lru
	for compare.next != nil {
		if compare.accessCount < lru.accessCount {
			lru = compare
		}
		compare = compare.next
	}

	return compare
}

func (LinkedList[K, V]) insert(linkedList *LinkedList[K, V], node *Node[K, V]) {
	if linkedList.size == 0 {
		linkedList.head = node
		linkedList.size++
		return
	}

	var head *Node[K, V] = linkedList.head
	for head.next != nil {
		head = head.next
	}

	linkedList.size++
	head.next = node
}

// universal hashing formula
// hab(k) = (((ak + b) mod p) mod m)
func (hashMap HashMap[K, V]) getHash(key K) int {
	var buf *bytes.Buffer = new(bytes.Buffer)

	var enc *gob.Encoder = gob.NewEncoder(buf)

	if err := enc.Encode(key); err != nil {
		println(err)
	}

	var intFromBytes uint32 = binary.LittleEndian.Uint32(buf.Bytes())

	return (((hashMap.a * int(intFromBytes)) + hashMap.b) % hashMap.prime) % hashMap.size
}

func getPositiveRandom() int {
	var random int

	random = rand.Int()
	for random <= 0 {
		random = rand.Int()
	}

	return random
}

func getPrimeBiggerThan(min int) int {
	var found bool = false
	var test int = min

	for !found {
		test = test + 1
		var sqrt int = int(math.Sqrt(float64(test)))
		var zeroRemainder bool = false
		for i := 2; i <= sqrt; i++ {
			if test%i == 0 {
				zeroRemainder = true
				break
			}
		}
		if !zeroRemainder {
			found = true
		}
	}
	return test
}

func printLRU(hashMap HashMap[string, int]) {
	println("printing each linked list size")
	for i := 0; i < len(hashMap.arr); i++ {
		println(hashMap.arr[i].size)
	}

	println()

	println("printing lru")
	for i := 0; i < len(hashMap.arr); i++ {
		println("array position = ", i)
		var head *Node[string, int] = hashMap.arr[i].head
		for head != nil {
			println(head.data)
			head = head.next
		}
	}
}

func New(size int) HashMap[string, int] {
	var a int = getPositiveRandom()
	var b int = getPositiveRandom()

	var prime int = getPrimeBiggerThan(size)

	hashMap := HashMap[string, int]{arr: make([]LinkedList[string, int], size), size: size, a: a, b: b, prime: prime}

	return hashMap
}

func main() {
	t := New(5)

	println("random a = ", t.a)
	println("random b = ", t.b)
	println("prime number = ", t.prime)

	t.Put("first-key", 10)

	t.Put("second-key", 20)

	t.Put("third-key", 30)
	t.Put("fourth-key", 40)
	t.Put("fifth-key", 50)
	t.Put("sixth-key", 70)
	t.Put("sdfsdf-key", 80)
	//t.Put("sixakoaskth-key", 90) // -> hash -1 index out of range
	t.Put("vsdv-key", 100)
	t.Put("eoqeqe-key", 120)
	t.Put("swe-key", 130)
	t.Put("adwdw-key", 140)
	t.Put("adwdww-key", 150) // -> eoqeqe-key substitute

	printLRU(t)
}
