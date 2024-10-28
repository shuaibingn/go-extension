```go
package main

import (
	"fmt"
	
	"github.com/shuaibingn/go-extension/set"
)

func main() {
	newSet := set.NewSet[int](1, 2, 3, 4)
	newSet.Add(1)
	fmt.Println(newSet)

	fmt.Println(newSet.Contains(1))
	
	newSet.Remove(1, 2)
	fmt.Println(newSet)
	
	fmt.Println(newSet.Contains(1))
	
	fmt.Println(newSet.Slice())
}
```