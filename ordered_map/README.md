```go
package main

import (
	"fmt"
	
	"github.com/shuaibingn/go-extension/ordered_map"
)

func main() {
	om := ordered_map.NewOrderedMap[string, string]() // 初始化有序map
	om.Set("key1", "value1") // 设置key, value
	om.Set("key2", "value2")
	om.Set("key3", "value3")
	
	value, ok := om.Get("key1") // 获取key1的值
    fmt.Println(value, ok)
	
	om.Remove("key2") // 删除key2
	
	keys := om.Keys() // 获取所有的key
	fmt.Println(keys)
	
	values := om.Values() // 获取所有的value
	fmt.Println(values)
	
	// 有序map遍历
	for item := range om.Iterator() {
		fmt.Println(item.Key, item.Value)
    }
	
	om.Clear() // 清空有序map
	fmt.Println(om.Len()) // 获取有序map的长度
}
```