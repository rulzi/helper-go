package arr

import (
	"reflect"
	"testing"
)

func TestAccessible(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"slice", []int{1, 2, 3}, true},
		{"array", [3]int{1, 2, 3}, true},
		{"map", map[string]int{"a": 1}, true},
		{"nil", nil, false},
		{"string", "hello", false},
		{"int", 123, false},
		{"bool", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Accessible(tt.value)
			if result != tt.expected {
				t.Errorf("Accessible(%v) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		key      string
		value    interface{}
		expected map[string]interface{}
	}{
		{
			"add new key",
			map[string]interface{}{"a": 1},
			"b",
			2,
			map[string]interface{}{"a": 1, "b": 2},
		},
		{
			"key exists",
			map[string]interface{}{"a": 1},
			"a",
			2,
			map[string]interface{}{"a": 1},
		},
		{
			"empty map",
			map[string]interface{}{},
			"a",
			1,
			map[string]interface{}{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.array, tt.key, tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Add() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCollapse(t *testing.T) {
	tests := []struct {
		name     string
		arrays   interface{}
		expected []interface{}
	}{
		{
			"slice of slices",
			[][]int{{1, 2}, {3, 4}},
			[]interface{}{1, 2, 3, 4},
		},
		{
			"slice of arrays",
			[][2]int{{1, 2}, {3, 4}},
			[]interface{}{1, 2, 3, 4},
		},
		{
			"slice of maps",
			[]map[string]int{{"a": 1}, {"b": 2}},
			[]interface{}{1, 2},
		},
		{
			"nil",
			nil,
			[]interface{}{},
		},
		{
			"empty slice",
			[][]int{},
			[]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Collapse(tt.arrays)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Collapse() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCrossJoin(t *testing.T) {
	tests := []struct {
		name     string
		arrays   [][]interface{}
		expected [][]interface{}
	}{
		{
			"two arrays",
			[][]interface{}{{1, 2}, {3, 4}},
			[][]interface{}{{1, 3}, {1, 4}, {2, 3}, {2, 4}},
		},
		{
			"three arrays",
			[][]interface{}{{1}, {2}, {3}},
			[][]interface{}{{1, 2, 3}},
		},
		{
			"empty",
			[][]interface{}{},
			[][]interface{}{},
		},
		{
			"single array",
			[][]interface{}{{1, 2}},
			[][]interface{}{{1}, {2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CrossJoin(tt.arrays...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CrossJoin() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name           string
		array          map[string]interface{}
		expectedKeys   []string
		expectedValues []interface{}
	}{
		{
			"basic",
			map[string]interface{}{"a": 1, "b": 2},
			[]string{"a", "b"},
			[]interface{}{1, 2},
		},
		{
			"empty",
			map[string]interface{}{},
			[]string{},
			[]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, values := Divide(tt.array)
			// Sort keys for comparison since map iteration order is random
			if len(keys) != len(tt.expectedKeys) {
				t.Errorf("Divide() keys length = %d, want %d", len(keys), len(tt.expectedKeys))
			}
			if len(values) != len(tt.expectedValues) {
				t.Errorf("Divide() values length = %d, want %d", len(values), len(tt.expectedValues))
			}
		})
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		prepend  string
		expected map[string]interface{}
	}{
		{
			"nested map",
			map[string]interface{}{
				"user": map[string]interface{}{
					"name": "John",
					"age":  30,
				},
			},
			"",
			map[string]interface{}{
				"user.name": "John",
				"user.age":  30,
			},
		},
		{
			"with prepend",
			map[string]interface{}{"a": 1},
			"prefix.",
			map[string]interface{}{"prefix.a": 1},
		},
		{
			"empty",
			map[string]interface{}{},
			"",
			map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Dot(tt.array, tt.prepend)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Dot() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUndot(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			"dot notation",
			map[string]interface{}{
				"user.name": "John",
				"user.age":  30,
			},
			map[string]interface{}{
				"user": map[string]interface{}{
					"name": "John",
					"age":  30,
				},
			},
		},
		{
			"simple keys",
			map[string]interface{}{"a": 1, "b": 2},
			map[string]interface{}{"a": 1, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Undot(tt.array)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Undot() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExcept(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		keys     []string
		expected map[string]interface{}
	}{
		{
			"remove keys",
			map[string]interface{}{"a": 1, "b": 2, "c": 3},
			[]string{"a", "c"},
			map[string]interface{}{"b": 2},
		},
		{
			"empty keys",
			map[string]interface{}{"a": 1},
			[]string{},
			map[string]interface{}{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Except(tt.array, tt.keys)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Except() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		key      string
		expected bool
	}{
		{"exists", map[string]interface{}{"a": 1}, "a", true},
		{"not exists", map[string]interface{}{"a": 1}, "b", false},
		{"nil map", nil, "a", false},
		{"empty map", map[string]interface{}{}, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Exists(tt.array, tt.key)
			if result != tt.expected {
				t.Errorf("Exists() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		name         string
		array        interface{}
		callback     func(interface{}) bool
		defaultValue interface{}
		expected     interface{}
	}{
		{
			"slice without callback",
			[]int{1, 2, 3},
			nil,
			nil,
			1,
		},
		{
			"slice with callback",
			[]int{1, 2, 3},
			func(v interface{}) bool { return v.(int) > 2 },
			nil,
			3,
		},
		{
			"empty slice",
			[]int{},
			nil,
			"default",
			"default",
		},
		{
			"map without callback",
			map[string]int{"a": 1, "b": 2},
			nil,
			nil,
			1, // or 2, order is random
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := First(tt.array, tt.callback, tt.defaultValue)
			if tt.name == "map without callback" {
				// Map order is random, just check it's one of the values
				if result != 1 && result != 2 {
					t.Errorf("First() = %v, want 1 or 2", result)
				}
			} else if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("First() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name         string
		array        interface{}
		callback     func(interface{}) bool
		defaultValue interface{}
		expected     interface{}
	}{
		{
			"slice without callback",
			[]int{1, 2, 3},
			nil,
			nil,
			3,
		},
		{
			"slice with callback",
			[]int{1, 2, 3},
			func(v interface{}) bool { return v.(int) < 3 },
			nil,
			2,
		},
		{
			"empty slice",
			[]int{},
			nil,
			"default",
			"default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Last(tt.array, tt.callback, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Last() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		limit    int
		expected []interface{}
	}{
		{"positive limit", []interface{}{1, 2, 3, 4}, 2, []interface{}{1, 2}},
		{"negative limit", []interface{}{1, 2, 3, 4}, -2, []interface{}{3, 4}},
		{"limit larger than array", []interface{}{1, 2}, 5, []interface{}{1, 2}},
		{"zero limit", []interface{}{1, 2, 3}, 0, []interface{}{}},
		{"empty array", []interface{}{}, 2, []interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Take(tt.array, tt.limit)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Take() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		array    interface{}
		depth    int
		expected []interface{}
	}{
		{
			"nested slice unlimited",
			[][]int{{1, 2}, {3, 4}},
			0,
			[]interface{}{1, 2, 3, 4},
		},
		{
			"nested slice depth 1",
			[][]int{{1, 2}, {3, 4}},
			1,
			[]interface{}{[]int{1, 2}, []int{3, 4}},
		},
		{
			"triple nested",
			[][][]int{{{1, 2}}, {{3, 4}}},
			0,
			[]interface{}{1, 2, 3, 4},
		},
		{
			"map",
			map[string]interface{}{"a": 1, "b": 2},
			0,
			[]interface{}{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Flatten(tt.array, tt.depth)
			if len(result) != len(tt.expected) {
				t.Errorf("Flatten() length = %d, want %d", len(result), len(tt.expected))
			}
		})
	}
}

func TestForget(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		keys     []string
		expected map[string]interface{}
	}{
		{
			"simple key",
			map[string]interface{}{"a": 1, "b": 2},
			[]string{"a"},
			map[string]interface{}{"b": 2},
		},
		{
			"dot notation",
			map[string]interface{}{
				"user": map[string]interface{}{
					"name": "John",
					"age":  30,
				},
			},
			[]string{"user.name"},
			map[string]interface{}{
				"user": map[string]interface{}{
					"age": 30,
				},
			},
		},
		{
			"multiple keys",
			map[string]interface{}{"a": 1, "b": 2, "c": 3},
			[]string{"a", "c"},
			map[string]interface{}{"b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Forget(tt.array, tt.keys)
			if !reflect.DeepEqual(tt.array, tt.expected) {
				t.Errorf("Forget() = %v, want %v", tt.array, tt.expected)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name         string
		array        map[string]interface{}
		key          string
		defaultValue interface{}
		expected     interface{}
	}{
		{
			"simple key",
			map[string]interface{}{"a": 1},
			"a",
			nil,
			1,
		},
		{
			"dot notation",
			map[string]interface{}{
				"user": map[string]interface{}{
					"name": "John",
				},
			},
			"user.name",
			nil,
			"John",
		},
		{
			"not found",
			map[string]interface{}{"a": 1},
			"b",
			"default",
			"default",
		},
		{
			"empty key",
			map[string]interface{}{"a": 1},
			"",
			nil,
			map[string]interface{}{"a": 1},
		},
		{
			"nil map",
			nil,
			"a",
			"default",
			"default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Get(tt.array, tt.key, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Get() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		keys     []string
		expected bool
	}{
		{
			"all exist",
			map[string]interface{}{"a": 1, "b": 2},
			[]string{"a", "b"},
			true,
		},
		{
			"one missing",
			map[string]interface{}{"a": 1},
			[]string{"a", "b"},
			false,
		},
		{
			"empty keys",
			map[string]interface{}{"a": 1},
			[]string{},
			false,
		},
		{
			"nil map",
			nil,
			[]string{"a"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Has(tt.array, tt.keys)
			if result != tt.expected {
				t.Errorf("Has() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHasOne(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		key      string
		expected bool
	}{
		{"exists", map[string]interface{}{"a": 1}, "a", true},
		{"not exists", map[string]interface{}{"a": 1}, "b", false},
		{
			"dot notation exists",
			map[string]interface{}{
				"user": map[string]interface{}{"name": "John"},
			},
			"user.name",
			true,
		},
		{
			"dot notation not exists",
			map[string]interface{}{"a": 1},
			"b.c",
			false,
		},
		{"nil map", nil, "a", false},
		{"empty key", map[string]interface{}{"a": 1}, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasOne(tt.array, tt.key)
			if result != tt.expected {
				t.Errorf("HasOne() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHasAny(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		keys     []string
		expected bool
	}{
		{
			"one exists",
			map[string]interface{}{"a": 1},
			[]string{"a", "b"},
			true,
		},
		{
			"none exists",
			map[string]interface{}{"a": 1},
			[]string{"b", "c"},
			false,
		},
		{
			"empty keys",
			map[string]interface{}{"a": 1},
			[]string{},
			false,
		},
		{
			"nil map",
			nil,
			[]string{"a"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasAny(tt.array, tt.keys)
			if result != tt.expected {
				t.Errorf("HasAny() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsAssoc(t *testing.T) {
	tests := []struct {
		name     string
		array    interface{}
		expected bool
	}{
		{"map", map[string]int{"a": 1}, true},
		{"slice", []int{1, 2, 3}, false},
		{"array", [3]int{1, 2, 3}, false},
		{"string", "hello", false},
		{"int", 123, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAssoc(tt.array)
			if result != tt.expected {
				t.Errorf("IsAssoc() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsList(t *testing.T) {
	tests := []struct {
		name     string
		array    interface{}
		expected bool
	}{
		{"slice", []int{1, 2, 3}, true},
		{"array", [3]int{1, 2, 3}, true},
		{"map", map[string]int{"a": 1}, false},
		{"string", "hello", false},
		{"int", 123, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsList(tt.array)
			if result != tt.expected {
				t.Errorf("IsList() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name      string
		array     []string
		glue      string
		finalGlue string
		expected  string
	}{
		{"basic", []string{"a", "b", "c"}, ",", "", "a,b,c"},
		{"with final glue", []string{"a", "b", "c"}, ",", " and ", "a,b and c"},
		{"single item", []string{"a"}, ",", " and ", "a"},
		{"empty", []string{}, ",", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Join(tt.array, tt.glue, tt.finalGlue)
			if result != tt.expected {
				t.Errorf("Join() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestKeyBy(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		keyBy    interface{}
		expected map[string]interface{}
	}{
		{
			"by field",
			[]interface{}{
				map[string]interface{}{"id": "1", "name": "John"},
				map[string]interface{}{"id": "2", "name": "Jane"},
			},
			"id",
			map[string]interface{}{
				"1": map[string]interface{}{"id": "1", "name": "John"},
				"2": map[string]interface{}{"id": "2", "name": "Jane"},
			},
		},
		{
			"by callback",
			[]interface{}{1, 2, 3},
			func(v interface{}) string {
				return "key_" + toString(v)
			},
			map[string]interface{}{
				"key_1": 1,
				"key_2": 2,
				"key_3": 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyBy(tt.array, tt.keyBy)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyBy() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPrependKeysWith(t *testing.T) {
	tests := []struct {
		name        string
		array       map[string]interface{}
		prependWith string
		expected    map[string]interface{}
	}{
		{
			"basic",
			map[string]interface{}{"a": 1, "b": 2},
			"prefix_",
			map[string]interface{}{"prefix_a": 1, "prefix_b": 2},
		},
		{
			"empty prepend",
			map[string]interface{}{"a": 1},
			"",
			map[string]interface{}{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrependKeysWith(tt.array, tt.prependWith)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PrependKeysWith() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestOnly(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		keys     []string
		expected map[string]interface{}
	}{
		{
			"subset",
			map[string]interface{}{"a": 1, "b": 2, "c": 3},
			[]string{"a", "c"},
			map[string]interface{}{"a": 1, "c": 3},
		},
		{
			"non-existent key",
			map[string]interface{}{"a": 1},
			[]string{"a", "b"},
			map[string]interface{}{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Only(tt.array, tt.keys)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Only() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		keys     []string
		expected []map[string]interface{}
	}{
		{
			"select fields",
			[]interface{}{
				map[string]interface{}{"id": 1, "name": "John", "age": 30},
				map[string]interface{}{"id": 2, "name": "Jane", "age": 25},
			},
			[]string{"id", "name"},
			[]map[string]interface{}{
				{"id": 1, "name": "John"},
				{"id": 2, "name": "Jane"},
			},
		},
		{
			"non-map item",
			[]interface{}{
				map[string]interface{}{"a": 1},
				123,
			},
			[]string{"a"},
			[]map[string]interface{}{
				{"a": 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Select(tt.array, tt.keys)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Select() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPluck(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		value    string
		key      string
		expected map[string]interface{}
	}{
		{
			"with key",
			[]interface{}{
				map[string]interface{}{"id": 1, "name": "John"},
				map[string]interface{}{"id": 2, "name": "Jane"},
			},
			"name",
			"id",
			map[string]interface{}{
				"1": "John",
				"2": "Jane",
			},
		},
		{
			"without key",
			[]interface{}{
				map[string]interface{}{"name": "John"},
			},
			"name",
			"",
			map[string]interface{}{
				"0": "John",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pluck(tt.array, tt.value, tt.key)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Pluck() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		callback func(interface{}) interface{}
		expected []interface{}
	}{
		{
			"multiply by 2",
			[]interface{}{1, 2, 3},
			func(v interface{}) interface{} {
				return v.(int) * 2
			},
			[]interface{}{2, 4, 6},
		},
		{
			"to string",
			[]interface{}{1, 2, 3},
			func(v interface{}) interface{} {
				return toString(v)
			},
			[]interface{}{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMapWithKeys(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		callback func(interface{}) map[string]interface{}
		expected map[string]interface{}
	}{
		{
			"transform to map",
			[]interface{}{1, 2},
			func(v interface{}) map[string]interface{} {
				return map[string]interface{}{
					toString(v): v.(int) * 2,
				}
			},
			map[string]interface{}{
				"1": 2,
				"2": 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapWithKeys(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MapWithKeys() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMapSpread(t *testing.T) {
	tests := []struct {
		name     string
		array    [][]interface{}
		callback func(...interface{}) interface{}
		expected []interface{}
	}{
		{
			"sum",
			[][]interface{}{{1, 2}, {3, 4}},
			func(v ...interface{}) interface{} {
				sum := 0
				for _, val := range v {
					sum += val.(int)
				}
				return sum
			},
			[]interface{}{3, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapSpread(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MapSpread() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		value    interface{}
		key      []string
		expected []interface{}
	}{
		{"basic", []interface{}{2, 3}, 1, []string{}, []interface{}{1, 2, 3}},
		{"with key", []interface{}{2, 3}, 1, []string{"key"}, []interface{}{1, 2, 3}},
		{"empty array", []interface{}{}, 1, []string{}, []interface{}{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Prepend(tt.array, tt.value, tt.key...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Prepend() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPull(t *testing.T) {
	tests := []struct {
		name         string
		array        map[string]interface{}
		key          string
		defaultValue interface{}
		expectedVal  interface{}
		expectedMap  map[string]interface{}
	}{
		{
			"pull and remove",
			map[string]interface{}{"a": 1, "b": 2},
			"a",
			nil,
			1,
			map[string]interface{}{"b": 2},
		},
		{
			"not found",
			map[string]interface{}{"a": 1},
			"b",
			"default",
			"default",
			map[string]interface{}{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pull(tt.array, tt.key, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedVal) {
				t.Errorf("Pull() value = %v, want %v", result, tt.expectedVal)
			}
			if !reflect.DeepEqual(tt.array, tt.expectedMap) {
				t.Errorf("Pull() array = %v, want %v", tt.array, tt.expectedMap)
			}
		})
	}
}

func TestQuery(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		expected string
	}{
		{
			"basic",
			map[string]interface{}{"a": 1, "b": "hello"},
			"a=1&b=hello",
		},
		{
			"empty",
			map[string]interface{}{},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Query(tt.array)
			// Query string order is random, so we just check it contains the key-value pairs
			if tt.name == "basic" {
				if result != "a=1&b=hello" && result != "b=hello&a=1" {
					t.Errorf("Query() = %q, want %q or %q", result, "a=1&b=hello", "b=hello&a=1")
				}
			} else if result != tt.expected {
				t.Errorf("Query() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		name         string
		array        []interface{}
		number       int
		preserveKeys bool
		expectError  bool
	}{
		{"single item", []interface{}{1, 2, 3}, 1, false, false},
		{"multiple items", []interface{}{1, 2, 3}, 2, false, false},
		{"with keys", []interface{}{1, 2, 3}, 2, true, false},
		{"empty array", []interface{}{}, 1, false, false},
		{"zero number", []interface{}{1, 2, 3}, 0, false, false},
		{"too many", []interface{}{1, 2}, 5, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Random(tt.array, tt.number, tt.preserveKeys)
			if tt.expectError {
				if err == nil {
					t.Errorf("Random() expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Random() error = %v", err)
				}
				if tt.number == 1 && result == nil && len(tt.array) > 0 {
					t.Errorf("Random() returned nil for non-empty array")
				}
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		key      string
		value    interface{}
		expected map[string]interface{}
	}{
		{
			"simple key",
			map[string]interface{}{},
			"a",
			1,
			map[string]interface{}{"a": 1},
		},
		{
			"dot notation",
			map[string]interface{}{},
			"user.name",
			"John",
			map[string]interface{}{
				"user": map[string]interface{}{
					"name": "John",
				},
			},
		},
		{
			"empty key",
			map[string]interface{}{"a": 1},
			"",
			2,
			map[string]interface{}{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Set(tt.array, tt.key, tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Set() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
	}{
		{"basic", []interface{}{1, 2, 3, 4, 5}},
		{"empty", []interface{}{}},
		{"single", []interface{}{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Shuffle(tt.array)
			if len(result) != len(tt.array) {
				t.Errorf("Shuffle() length = %d, want %d", len(result), len(tt.array))
			}
			// Check all elements are present (order may differ)
			if len(tt.array) > 0 {
				originalMap := make(map[interface{}]bool)
				for _, v := range tt.array {
					originalMap[v] = true
				}
				for _, v := range result {
					if !originalMap[v] {
						t.Errorf("Shuffle() contains unexpected value %v", v)
					}
				}
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		callback func(interface{}, interface{}) bool
		expected []interface{}
	}{
		{
			"without callback",
			[]interface{}{3, 1, 2},
			nil,
			[]interface{}{1, 2, 3},
		},
		{
			"with callback",
			[]interface{}{3, 1, 2},
			func(a, b interface{}) bool {
				return a.(int) < b.(int)
			},
			[]interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sort(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Sort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSortDesc(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		callback func(interface{}, interface{}) bool
		expected []interface{}
	}{
		{
			"without callback",
			[]interface{}{1, 3, 2},
			nil,
			[]interface{}{3, 2, 1},
		},
		{
			"with callback",
			[]interface{}{1, 3, 2},
			func(a, b interface{}) bool {
				return a.(int) < b.(int)
			},
			[]interface{}{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortDesc(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortDesc() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSortRecursive(t *testing.T) {
	tests := []struct {
		name       string
		array      map[string]interface{}
		descending bool
		expected   map[string]interface{}
	}{
		{
			"ascending",
			map[string]interface{}{
				"c": 3,
				"a": 1,
				"b": map[string]interface{}{
					"z": 2,
					"x": 1,
				},
			},
			false,
			map[string]interface{}{
				"a": 1,
				"b": map[string]interface{}{
					"x": 1,
					"z": 2,
				},
				"c": 3,
			},
		},
		{
			"descending",
			map[string]interface{}{
				"a": 1,
				"c": 3,
				"b": 2,
			},
			true,
			map[string]interface{}{
				"c": 3,
				"b": 2,
				"a": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortRecursive(tt.array, tt.descending)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSortRecursiveDesc(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			"descending",
			map[string]interface{}{
				"a": 1,
				"c": 3,
				"b": 2,
			},
			map[string]interface{}{
				"c": 3,
				"b": 2,
				"a": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortRecursiveDesc(tt.array)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortRecursiveDesc() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToCssClasses(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		expected string
	}{
		{
			"true values",
			map[string]interface{}{
				"active":   true,
				"disabled": false,
				"hidden":   true,
			},
			"active hidden",
		},
		{
			"string values",
			map[string]interface{}{
				"class1": "value",
				"class2": "",
			},
			"class1",
		},
		{
			"nil values",
			map[string]interface{}{
				"class1": nil,
				"class2": true,
			},
			"class2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCssClasses(tt.array)
			// Order may vary, so check if all expected classes are present
			if tt.name == "true values" {
				if result != "active hidden" && result != "hidden active" {
					t.Errorf("ToCssClasses() = %q, want %q or %q", result, "active hidden", "hidden active")
				}
			} else if result != tt.expected {
				t.Errorf("ToCssClasses() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToCssStyles(t *testing.T) {
	tests := []struct {
		name     string
		array    map[string]interface{}
		expected string
	}{
		{
			"true values",
			map[string]interface{}{
				"color":   true,
				"display": false,
			},
			"color;",
		},
		{
			"string values",
			map[string]interface{}{
				"margin":  "10px",
				"padding": "",
			},
			"margin;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCssStyles(tt.array)
			if result != tt.expected {
				t.Errorf("ToCssStyles() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestWhere(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		callback func(interface{}) bool
		expected []interface{}
	}{
		{
			"filter even",
			[]interface{}{1, 2, 3, 4},
			func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			[]interface{}{2, 4},
		},
		{
			"filter greater than",
			[]interface{}{1, 2, 3, 4, 5},
			func(v interface{}) bool {
				return v.(int) > 3
			},
			[]interface{}{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Where(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Where() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReject(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		callback func(interface{}) bool
		expected []interface{}
	}{
		{
			"reject even",
			[]interface{}{1, 2, 3, 4},
			func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			[]interface{}{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reject(tt.array, tt.callback)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Reject() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestWhereNotNull(t *testing.T) {
	tests := []struct {
		name     string
		array    []interface{}
		expected []interface{}
	}{
		{
			"filter nil",
			[]interface{}{1, nil, 2, nil, 3},
			[]interface{}{1, 2, 3},
		},
		{
			"all nil",
			[]interface{}{nil, nil},
			[]interface{}{},
		},
		{
			"no nil",
			[]interface{}{1, 2, 3},
			[]interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WhereNotNull(tt.array)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WhereNotNull() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected []interface{}
	}{
		{"int", 1, []interface{}{1}},
		{"string", "hello", []interface{}{"hello"}},
		{"slice", []int{1, 2, 3}, []interface{}{1, 2, 3}},
		{"array", [3]int{1, 2, 3}, []interface{}{1, 2, 3}},
		{"nil", nil, []interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Wrap(tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Wrap() = %v, want %v", result, tt.expected)
			}
		})
	}
}
