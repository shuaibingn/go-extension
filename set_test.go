package extension

import (
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)

	if s.Len() != 3 {
		t.Errorf("expected len 3, got %d", s.Len())
	}

	// 添加重复元素
	s.Add(1, 2)
	if s.Len() != 3 {
		t.Errorf("expected len 3 after adding duplicates, got %d", s.Len())
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)
	s.Remove(2, 4)

	if s.Len() != 3 {
		t.Errorf("expected len 3, got %d", s.Len())
	}
	if s.Contains(2) || s.Contains(4) {
		t.Error("expected 2 and 4 to be removed")
	}

	// 删除不存在的元素
	s.Remove(100)
	if s.Len() != 3 {
		t.Errorf("expected len 3 after removing non-existent, got %d", s.Len())
	}
}

func TestSet_Contains(t *testing.T) {
	s := NewSet("a", "b", "c")

	if !s.Contains("a") {
		t.Error("expected to contain 'a'")
	}
	if s.Contains("d") {
		t.Error("expected not to contain 'd'")
	}
}

func TestSet_Clear(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Clear()

	if s.Len() != 0 {
		t.Errorf("expected len 0 after clear, got %d", s.Len())
	}
}

func TestSet_Equal(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 2, 1)
	s3 := NewSet(1, 2, 4)

	if !s1.Equal(s2) {
		t.Error("expected s1 and s2 to be equal")
	}
	if s1.Equal(s3) {
		t.Error("expected s1 and s3 to not be equal")
	}
}

func TestSet_Slice(t *testing.T) {
	s := NewSet(1, 2, 3)
	slice := s.Slice()

	if len(slice) != 3 {
		t.Errorf("expected slice len 3, got %d", len(slice))
	}
}

func TestSet_ForEach(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)
	count := 0

	s.ForEach(func(item int) bool {
		count++
		return true
	})

	if count != 5 {
		t.Errorf("expected to iterate 5 times, got %d", count)
	}

	// 测试提前终止
	count = 0
	s.ForEach(func(item int) bool {
		count++
		return count < 3
	})

	if count != 3 {
		t.Errorf("expected to iterate 3 times with early exit, got %d", count)
	}
}

func TestSet_Join(t *testing.T) {
	s := NewSet[string]()
	if s.Join(",") != "" {
		t.Error("expected empty string for empty set")
	}

	s.Add("a")
	if s.Join(",") != "a" {
		t.Errorf("expected 'a', got '%s'", s.Join(","))
	}
}

