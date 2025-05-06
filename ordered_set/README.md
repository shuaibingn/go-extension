```go
package main

import (
	"fmt"
	
	"github.com/shuaibingn/go-extension/ordered_set"
)


func main() {
	orderedSet := pkg.NewOrderedSet(1, 2, 3, 4, 5, 6)
	fmt.Println(orderedSet.Last())

	orderedSet.Remove(4, 3, 1)
	fmt.Println(orderedSet)
	for item := range orderedSet.Iterator() {
		fmt.Printf("index: %v, value: %v\n", orderedSet.Index(item), item)
	}

	orderedSet.Add(1, 2, 3, 4, 5, 6)

	// for item := range orderedSet.Iterator() {
	// 	fmt.Printf("index: %v, value: %v\n", orderedSet.Index(item), item)
	// }

	orderedSet.Remove(2, 5, 6, 10)
	fmt.Println(orderedSet)
	for item := range orderedSet.Iterator() {
		fmt.Printf("index: %v, value: %v\n", orderedSet.Index(item), item)
	}

	orderedSet.Add(11)

	orderedSet.Remove(12)
	fmt.Println(orderedSet)
	for item := range orderedSet.Iterator() {
		fmt.Printf("index: %v, value: %v\n", orderedSet.Index(item), item)
	}
}
```