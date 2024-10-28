```go
package main

import (
	"fmt"
	
	"github.com/shuaibingn/go-extension/ordered_map"
)

func main() {
	om := ordered_map.NewOrderedMap[string, string]()
	om.Set("key1", "value1")
	om.Set("key2", "value2")
	
	for item := range om.Iterator() {
		fmt.Println(item.Key, item.Value)
    }
}
```