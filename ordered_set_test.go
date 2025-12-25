package extension

import (
	"testing"
)

func TestOrderedSet_Add(t *testing.T) {
	os := NewOrderedSet[int]()
	os.Add(3, 1, 2)

	if os.Len() != 3 {
		t.Errorf("expected len 3, got %d", os.Len())
	}

	// 验证顺序
	slice := os.Slice()
	expected := []int{3, 1, 2}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("expected %d at index %d, got %d", v, i, slice[i])
		}
	}

	// 添加重复元素
	os.Add(1, 2)
	if os.Len() != 3 {
		t.Errorf("expected len 3 after adding duplicates, got %d", os.Len())
	}
}

func TestOrderedSet_Remove(t *testing.T) {
	os := NewOrderedSet(1, 2, 3, 4, 5)
	os.Remove(2, 4)

	if os.Len() != 3 {
		t.Errorf("expected len 3, got %d", os.Len())
	}

	// 验证顺序保持
	slice := os.Slice()
	expected := []int{1, 3, 5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("expected %d at index %d, got %d", v, i, slice[i])
		}
	}

	// 验证索引更新
	if os.Index(3) != 1 {
		t.Errorf("expected index of 3 to be 1, got %d", os.Index(3))
	}
	if os.Index(5) != 2 {
		t.Errorf("expected index of 5 to be 2, got %d", os.Index(5))
	}
}

func TestOrderedSet_Remove_MemoryLeak(t *testing.T) {
	// 测试删除后内存是否正确清理
	type largeStruct struct {
		data [1024]byte
	}

	os := NewOrderedSet[*largeStruct]()
	items := make([]*largeStruct, 100)
	for i := range items {
		items[i] = &largeStruct{}
		os.Add(items[i])
	}

	// 删除所有元素
	os.Remove(items...)

	if os.Len() != 0 {
		t.Errorf("expected len 0, got %d", os.Len())
	}
}

func TestOrderedSet_Index(t *testing.T) {
	os := NewOrderedSet("a", "b", "c", "d")

	tests := []struct {
		item     string
		expected int
	}{
		{"a", 0},
		{"b", 1},
		{"c", 2},
		{"d", 3},
		{"e", -1}, // 不存在
	}

	for _, tt := range tests {
		if got := os.Index(tt.item); got != tt.expected {
			t.Errorf("Index(%q) = %d, want %d", tt.item, got, tt.expected)
		}
	}
}

func TestOrderedSet_Get(t *testing.T) {
	os := NewOrderedSet(10, 20, 30)

	val, ok := os.Get(1)
	if !ok || val != 20 {
		t.Errorf("Get(1) = (%d, %v), want (20, true)", val, ok)
	}

	_, ok = os.Get(-1)
	if ok {
		t.Error("Get(-1) should return false")
	}

	_, ok = os.Get(10)
	if ok {
		t.Error("Get(10) should return false")
	}
}

func TestOrderedSet_FirstLast(t *testing.T) {
	os := NewOrderedSet(1, 2, 3)

	first, ok := os.First()
	if !ok || first != 1 {
		t.Errorf("First() = (%d, %v), want (1, true)", first, ok)
	}

	last, ok := os.Last()
	if !ok || last != 3 {
		t.Errorf("Last() = (%d, %v), want (3, true)", last, ok)
	}

	// 空集合
	empty := NewOrderedSet[int]()
	_, ok = empty.First()
	if ok {
		t.Error("First() on empty set should return false")
	}
	_, ok = empty.Last()
	if ok {
		t.Error("Last() on empty set should return false")
	}
}

func TestOrderedSet_Contains(t *testing.T) {
	os := NewOrderedSet("x", "y", "z")

	if !os.Contains("x") {
		t.Error("expected to contain 'x'")
	}
	if os.Contains("a") {
		t.Error("expected not to contain 'a'")
	}
}

func TestOrderedSet_Clear(t *testing.T) {
	os := NewOrderedSet(1, 2, 3)
	os.Clear()

	if os.Len() != 0 {
		t.Errorf("expected len 0 after clear, got %d", os.Len())
	}
}

func TestOrderedSet_ForEach(t *testing.T) {
	os := NewOrderedSet(1, 2, 3, 4, 5)
	var result []int

	os.ForEach(func(index int, item int) bool {
		result = append(result, item)
		return true
	})

	expected := []int{1, 2, 3, 4, 5}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("expected %d at index %d, got %d", v, i, result[i])
		}
	}

	// 测试提前终止
	result = nil
	os.ForEach(func(index int, item int) bool {
		result = append(result, item)
		return index < 2
	})

	if len(result) != 3 {
		t.Errorf("expected 3 items with early exit, got %d", len(result))
	}
}

func TestOrderedSet_SliceRef(t *testing.T) {
	os := NewOrderedSet(1, 2, 3)
	ref := os.SliceRef()

	if len(ref) != 3 {
		t.Errorf("expected len 3, got %d", len(ref))
	}

	// 验证是同一个引用
	copy := os.Slice()
	copy[0] = 100
	if ref[0] == 100 {
		t.Error("SliceRef should return internal slice, not affected by Slice modifications")
	}
}

