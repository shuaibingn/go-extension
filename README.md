# go-extension

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

使用泛型实现的 Go 语言数据结构扩展库，提供 Set、OrderedSet、OrderedMap 等常用数据结构。

## 特性

- 🚀 基于 Go 1.22+ 泛型实现，类型安全
- 💾 优化的内存管理，防止内存泄漏
- ⚡ 高性能设计，关键操作 O(1) 复杂度
- 📦 零依赖，仅使用标准库

## 安装

```shell
go get -u github.com/shuaibingn/go-extension
```

## 快速开始

```go
import "github.com/shuaibingn/go-extension"
```

### Set - 无序集合

```go
s := extension.NewSet(1, 2, 3)
s.Add(4, 5)
s.Remove(1)

fmt.Println(s.Contains(2)) // true
fmt.Println(s.Len())       // 4

// 遍历
s.ForEach(func(item int) bool {
    fmt.Println(item)
    return true // 继续遍历
})
```

### OrderedSet - 有序集合

```go
os := extension.NewOrderedSet("a", "b", "c")
os.Add("d")

fmt.Println(os.Index("b"))  // 1 (O(1) 复杂度)
fmt.Println(os.First())     // "a", true
fmt.Println(os.Last())      // "d", true

// 按索引获取
val, ok := os.Get(0)        // "a", true

// 遍历（保持插入顺序）
os.ForEach(func(index int, item string) bool {
    fmt.Printf("[%d] %s\n", index, item)
    return true
})
```

### OrderedMap - 有序映射

```go
om := extension.NewOrderedMap[string, int]()
om.Set("one", 1)
om.Set("two", 2)
om.Set("three", 3)

val, ok := om.Get("two")    // 2, true
keys := om.Keys()           // ["one", "two", "three"]
values := om.Values()       // [1, 2, 3]

// 遍历（保持插入顺序）
om.ForEach(func(key string, value int) bool {
    fmt.Printf("%s: %d\n", key, value)
    return true
})
```

## API 参考

### Set[T comparable]

| 方法 | 描述 | 复杂度 |
|-----|------|-------|
| `NewSet[T](items ...T)` | 创建新集合 | O(n) |
| `Add(items ...T)` | 添加元素 | O(1) |
| `Remove(items ...T)` | 移除元素 | O(1) |
| `Contains(item T) bool` | 检查元素是否存在 | O(1) |
| `Len() int` | 返回元素数量 | O(1) |
| `Clear()` | 清空集合 | O(n) |
| `Slice() []T` | 返回元素切片 | O(n) |
| `Equal(other Set[T]) bool` | 比较两个集合是否相等 | O(n) |
| `Join(sep string) string` | 连接元素为字符串 | O(n) |
| `ForEach(fn func(T) bool)` | 遍历元素 | O(n) |

### OrderedSet[T comparable]

| 方法 | 描述 | 复杂度 |
|-----|------|-------|
| `NewOrderedSet[T](items ...T)` | 创建新有序集合 | O(n) |
| `Add(items ...T)` | 添加元素 | O(1) |
| `Remove(items ...T)` | 移除元素 | O(n) |
| `Contains(item T) bool` | 检查元素是否存在 | O(1) |
| `Index(item T) int` | 获取元素索引 | O(1) |
| `Get(index int) (T, bool)` | 按索引获取元素 | O(1) |
| `First() (T, bool)` | 获取第一个元素 | O(1) |
| `Last() (T, bool)` | 获取最后一个元素 | O(1) |
| `Len() int` | 返回元素数量 | O(1) |
| `Clear()` | 清空集合 | O(1) |
| `Slice() []T` | 返回元素切片（拷贝） | O(n) |
| `SliceRef() []T` | 返回内部切片引用 | O(1) |
| `Join(sep string) string` | 连接元素为字符串 | O(n) |
| `ForEach(fn func(int, T) bool)` | 遍历元素 | O(n) |

### OrderedMap[K comparable, V any]

| 方法 | 描述 | 复杂度 |
|-----|------|-------|
| `NewOrderedMap[K, V]()` | 创建新有序映射 | O(1) |
| `Set(key K, value V)` | 设置键值对 | O(1) |
| `Get(key K) (V, bool)` | 获取值 | O(1) |
| `Remove(key K)` | 移除键值对 | O(n) |
| `Keys() []K` | 返回所有键 | O(n) |
| `Values() []V` | 返回所有值 | O(n) |
| `Len() int` | 返回键值对数量 | O(1) |
| `Clear()` | 清空映射 | O(1) |
| `ForEach(fn func(K, V) bool)` | 遍历键值对 | O(n) |

## 项目结构

```
go-extension/
├── go.mod
├── LICENSE
├── README.md
├── set.go               # Set 实现
├── set_test.go          # Set 测试
├── ordered_set.go       # OrderedSet 实现
├── ordered_set_test.go  # OrderedSet 测试
├── ordered_map.go       # OrderedMap 实现
└── ordered_map_test.go  # OrderedMap 测试
```

## License

[MIT License](LICENSE)
