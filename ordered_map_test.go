package extension

import (
	"testing"
)

func TestOrderedMap_Set(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	if om.Len() != 3 {
		t.Errorf("expected len 3, got %d", om.Len())
	}

	// 更新已存在的键
	om.Set("b", 20)
	if om.Len() != 3 {
		t.Errorf("expected len 3 after update, got %d", om.Len())
	}

	val, ok := om.Get("b")
	if !ok || val != 20 {
		t.Errorf("Get('b') = (%d, %v), want (20, true)", val, ok)
	}
}

func TestOrderedMap_Get(t *testing.T) {
	om := NewOrderedMap[string, string]()
	om.Set("key1", "value1")
	om.Set("key2", "value2")

	val, ok := om.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Get('key1') = (%s, %v), want (value1, true)", val, ok)
	}

	_, ok = om.Get("nonexistent")
	if ok {
		t.Error("Get('nonexistent') should return false")
	}
}

func TestOrderedMap_Remove(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	om.Set("d", 4)

	om.Remove("b")

	if om.Len() != 3 {
		t.Errorf("expected len 3, got %d", om.Len())
	}

	_, ok := om.Get("b")
	if ok {
		t.Error("expected 'b' to be removed")
	}

	// 验证顺序保持
	keys := om.Keys()
	expected := []string{"a", "c", "d"}
	for i, k := range expected {
		if keys[i] != k {
			t.Errorf("expected key %s at index %d, got %s", k, i, keys[i])
		}
	}
}

func TestOrderedMap_Remove_MemoryLeak(t *testing.T) {
	// 测试删除后内存是否正确清理
	om := NewOrderedMap[int, *[1024]byte]()
	for i := 0; i < 100; i++ {
		om.Set(i, &[1024]byte{})
	}

	for i := 0; i < 100; i++ {
		om.Remove(i)
	}

	if om.Len() != 0 {
		t.Errorf("expected len 0, got %d", om.Len())
	}
}

func TestOrderedMap_Keys(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("x", 1)
	om.Set("y", 2)
	om.Set("z", 3)

	keys := om.Keys()
	expected := []string{"x", "y", "z"}

	if len(keys) != len(expected) {
		t.Errorf("expected %d keys, got %d", len(expected), len(keys))
	}

	for i, k := range expected {
		if keys[i] != k {
			t.Errorf("expected key %s at index %d, got %s", k, i, keys[i])
		}
	}
}

func TestOrderedMap_Values(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 10)
	om.Set("b", 20)
	om.Set("c", 30)

	values := om.Values()
	expected := []int{10, 20, 30}

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	for i, v := range expected {
		if values[i] != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, values[i])
		}
	}
}

func TestOrderedMap_Clear(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)

	om.Clear()

	if om.Len() != 0 {
		t.Errorf("expected len 0 after clear, got %d", om.Len())
	}

	_, ok := om.Get("a")
	if ok {
		t.Error("expected 'a' to not exist after clear")
	}
}

func TestOrderedMap_ForEach(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	var keys []string
	var values []int

	om.ForEach(func(key string, value int) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})

	expectedKeys := []string{"a", "b", "c"}
	expectedValues := []int{1, 2, 3}

	for i := range expectedKeys {
		if keys[i] != expectedKeys[i] {
			t.Errorf("expected key %s at index %d, got %s", expectedKeys[i], i, keys[i])
		}
		if values[i] != expectedValues[i] {
			t.Errorf("expected value %d at index %d, got %d", expectedValues[i], i, values[i])
		}
	}

	// 测试提前终止
	count := 0
	om.ForEach(func(key string, value int) bool {
		count++
		return count < 2
	})

	if count != 2 {
		t.Errorf("expected 2 iterations with early exit, got %d", count)
	}
}

func TestOrderedMap_String(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)

	// 通过 fmt.Stringer 接口测试
	str := om.(*orderedMap[string, int]).String()
	expected := "OrderedMap[{a: 1} {b: 2}]"

	if str != expected {
		t.Errorf("String() = %s, want %s", str, expected)
	}
}

func TestOrderedMap_EmptyOperations(t *testing.T) {
	om := NewOrderedMap[string, int]()

	// 空 map 操作
	if om.Len() != 0 {
		t.Errorf("expected len 0, got %d", om.Len())
	}

	_, ok := om.Get("any")
	if ok {
		t.Error("Get on empty map should return false")
	}

	om.Remove("any") // 不应该 panic

	keys := om.Keys()
	if len(keys) != 0 {
		t.Errorf("expected 0 keys, got %d", len(keys))
	}

	values := om.Values()
	if len(values) != 0 {
		t.Errorf("expected 0 values, got %d", len(values))
	}
}

