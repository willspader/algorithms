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

func (hashMap HashMap[K, V]) get(key K) (V, bool) {
	var hash int = hashMap.getHash(key)

	var linkedList LinkedList[K, V] = hashMap.arr[hash]

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

func main() {

}