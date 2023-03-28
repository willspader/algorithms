package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"math"
	"math/rand"
)

type HashTable[K any, V any] struct {
	arr   []V
	size  int
	a     int
	b     int
	prime int
}

func (hashTable HashTable[K, V]) Get(key K) V {
	var hash int = hashTable.getHash(key)

	return hashTable.arr[hash]
}

func (hashTable HashTable[K, V]) Put(key K, value V) {
	var hash int = hashTable.getHash(key)

	hashTable.arr[hash] = value
}

// universal hashing formula
// hab(k) = (((ak + b) mod p) mod m)
func (hashTable HashTable[K, V]) getHash(key K) int {
	var buf *bytes.Buffer = new(bytes.Buffer)

	var enc *gob.Encoder = gob.NewEncoder(buf)

	if err := enc.Encode(key); err != nil {
		println(err)
	}

	var intFromBytes uint32 = binary.LittleEndian.Uint32(buf.Bytes())

	return (((hashTable.a * int(intFromBytes)) + hashTable.b) % hashTable.prime) % hashTable.size
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

func printTable(arr []int) {
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}

func New(size int) HashTable[string, int] {
	var a int = getPositiveRandom()
	var b int = getPositiveRandom()

	var prime int = getPrimeBiggerThan(size)

	hashTable := HashTable[string, int]{arr: make([]int, size), size: size, a: a, b: b, prime: prime}

	return hashTable
}

func main() {
	t := New(5)

	println("random a = ", t.a)
	println("random b = ", t.b)
	println("prime number = ", t.prime)

	t.Put("first-key", 43)
	printTable(t.arr)

	println()
	t.Put("seco9q09nd-key", 29)
	printTable(t.arr)

	println()
	t.Put("ajodoiwhduasd-kdfey", 8855)
	printTable(t.arr)

}
