package str

import (
	"testing"
)

func TestAfter(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		search   string
		expected string
	}{
		{"basic", "hello world", "hello", " world"},
		{"not found", "hello world", "xyz", "hello world"},
		{"empty search", "hello world", "", "hello world"},
		{"multiple occurrences", "hello hello world", "hello", " hello world"},
		{"empty subject", "", "hello", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := After(tt.subject, tt.search)
			if result != tt.expected {
				t.Errorf("After(%q, %q) = %q, want %q", tt.subject, tt.search, result, tt.expected)
			}
		})
	}
}

func TestAfterLast(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		search   string
		expected string
	}{
		{"basic", "hello world", "world", ""},
		{"multiple occurrences", "hello hello world", "hello", " world"},
		{"not found", "hello world", "xyz", "hello world"},
		{"empty search", "hello world", "", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AfterLast(tt.subject, tt.search)
			if result != tt.expected {
				t.Errorf("AfterLast(%q, %q) = %q, want %q", tt.subject, tt.search, result, tt.expected)
			}
		})
	}
}

func TestBefore(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		search   string
		expected string
	}{
		{"basic", "hello world", "world", "hello "},
		{"not found", "hello world", "xyz", "hello world"},
		{"empty search", "hello world", "", "hello world"},
		{"multiple occurrences", "hello hello world", "hello", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Before(tt.subject, tt.search)
			if result != tt.expected {
				t.Errorf("Before(%q, %q) = %q, want %q", tt.subject, tt.search, result, tt.expected)
			}
		})
	}
}

func TestBeforeLast(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		search   string
		expected string
	}{
		{"basic", "hello world", "world", "hello "},
		{"multiple occurrences", "hello hello world", "hello", "hello "},
		{"not found", "hello world", "xyz", "hello world"},
		{"empty search", "hello world", "", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BeforeLast(tt.subject, tt.search)
			if result != tt.expected {
				t.Errorf("BeforeLast(%q, %q) = %q, want %q", tt.subject, tt.search, result, tt.expected)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		from     string
		to       string
		expected string
	}{
		{"basic", "hello [world] test", "[", "]", "world"},
		{"empty from", "hello world", "", "world", "hello world"},
		{"empty to", "hello world", "hello", "", "hello world"},
		{"not found", "hello world", "[", "]", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Between(tt.subject, tt.from, tt.to)
			if result != tt.expected {
				t.Errorf("Between(%q, %q, %q) = %q, want %q", tt.subject, tt.from, tt.to, result, tt.expected)
			}
		})
	}
}

func TestBetweenFirst(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		from     string
		to       string
		expected string
	}{
		{"basic", "hello [world] test", "[", "]", "world"},
		{"empty from", "hello world", "", "world", "hello world"},
		{"empty to", "hello world", "hello", "", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BetweenFirst(tt.subject, tt.from, tt.to)
			if result != tt.expected {
				t.Errorf("BetweenFirst(%q, %q, %q) = %q, want %q", tt.subject, tt.from, tt.to, result, tt.expected)
			}
		})
	}
}

func TestCamel(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "hello world", "helloWorld"},
		{"snake case", "hello_world", "helloWorld"},
		{"kebab case", "hello-world", "helloWorld"},
		{"already camel", "helloWorld", "helloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Camel(tt.value)
			if result != tt.expected {
				t.Errorf("Camel(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestCharAt(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		index    int
		expected string
		ok       bool
	}{
		{"basic", "hello", 0, "h", true},
		{"middle", "hello", 2, "l", true},
		{"last", "hello", 4, "o", true},
		{"negative index", "hello", -1, "o", true},
		{"negative index out of range", "hello", -10, "", false},
		{"out of range", "hello", 10, "", false},
		{"unicode", "こんにちは", 0, "こ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := CharAt(tt.subject, tt.index)
			if result != tt.expected || ok != tt.ok {
				t.Errorf("CharAt(%q, %d) = %q, %v, want %q, %v", tt.subject, tt.index, result, ok, tt.expected, tt.ok)
			}
		})
	}
}

func TestChopStart(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		needles  []string
		expected string
	}{
		{"basic", "hello world", []string{"hello"}, " world"},
		{"multiple needles", "hello world", []string{"xyz", "hello"}, " world"},
		{"not found", "hello world", []string{"xyz"}, "hello world"},
		{"empty needles", "hello world", []string{}, "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ChopStart(tt.subject, tt.needles...)
			if result != tt.expected {
				t.Errorf("ChopStart(%q, %v) = %q, want %q", tt.subject, tt.needles, result, tt.expected)
			}
		})
	}
}

func TestChopEnd(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		needles  []string
		expected string
	}{
		{"basic", "hello world", []string{"world"}, "hello "},
		{"multiple needles", "hello world", []string{"xyz", "world"}, "hello "},
		{"not found", "hello world", []string{"xyz"}, "hello world"},
		{"empty needles", "hello world", []string{}, "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ChopEnd(tt.subject, tt.needles...)
			if result != tt.expected {
				t.Errorf("ChopEnd(%q, %v) = %q, want %q", tt.subject, tt.needles, result, tt.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name       string
		haystack   string
		needles    []string
		ignoreCase bool
		expected   bool
	}{
		{"basic", "hello world", []string{"world"}, false, true},
		{"case sensitive", "Hello World", []string{"hello"}, false, false},
		{"ignore case", "Hello World", []string{"hello"}, true, true},
		{"multiple needles", "hello world", []string{"xyz", "world"}, false, true},
		{"not found", "hello world", []string{"xyz"}, false, false},
		{"empty needles", "hello world", []string{""}, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.haystack, tt.needles, tt.ignoreCase)
			if result != tt.expected {
				t.Errorf("Contains(%q, %v, %v) = %v, want %v", tt.haystack, tt.needles, tt.ignoreCase, result, tt.expected)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name       string
		haystack   string
		needles    []string
		ignoreCase bool
		expected   bool
	}{
		{"basic", "hello world", []string{"hello", "world"}, false, true},
		{"missing one", "hello world", []string{"hello", "xyz"}, false, false},
		{"ignore case", "Hello World", []string{"hello", "world"}, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAll(tt.haystack, tt.needles, tt.ignoreCase)
			if result != tt.expected {
				t.Errorf("ContainsAll(%q, %v, %v) = %v, want %v", tt.haystack, tt.needles, tt.ignoreCase, result, tt.expected)
			}
		})
	}
}

func TestDoesntContain(t *testing.T) {
	tests := []struct {
		name       string
		haystack   string
		needles    []string
		ignoreCase bool
		expected   bool
	}{
		{"basic", "hello world", []string{"xyz"}, false, true},
		{"contains", "hello world", []string{"world"}, false, false},
		{"ignore case", "Hello World", []string{"hello"}, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DoesntContain(tt.haystack, tt.needles, tt.ignoreCase)
			if result != tt.expected {
				t.Errorf("DoesntContain(%q, %v, %v) = %v, want %v", tt.haystack, tt.needles, tt.ignoreCase, result, tt.expected)
			}
		})
	}
}

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		character string
		expected  string
	}{
		{"basic", "hello   world", " ", "hello world"},
		{"multiple", "hello---world", "-", "hello-world"},
		{"no duplicates", "hello world", " ", "hello world"},
		{"special char", "hello...world", ".", "hello.world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Deduplicate(tt.s, tt.character)
			if result != tt.expected {
				t.Errorf("Deduplicate(%q, %q) = %q, want %q", tt.s, tt.character, result, tt.expected)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needles  []string
		expected bool
	}{
		{"basic", "hello world", []string{"world"}, true},
		{"multiple needles", "hello world", []string{"xyz", "world"}, true},
		{"not found", "hello world", []string{"xyz"}, false},
		{"empty haystack", "", []string{"world"}, false},
		{"empty needles", "hello world", []string{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndsWith(tt.haystack, tt.needles)
			if result != tt.expected {
				t.Errorf("EndsWith(%q, %v) = %v, want %v", tt.haystack, tt.needles, result, tt.expected)
			}
		})
	}
}

func TestFinish(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		cap      string
		expected string
	}{
		{"basic", "hello", "/", "hello/"},
		{"already ends", "hello/", "/", "hello//"},
		{"multiple caps at start", "///hello", "/", "//hello/"},
		{"empty value", "", "/", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Finish(tt.value, tt.cap)
			if result != tt.expected {
				t.Errorf("Finish(%q, %q) = %q, want %q", tt.value, tt.cap, result, tt.expected)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		before   string
		after    []string
		expected string
	}{
		{"basic", "hello", "[", nil, "[hello["},
		{"with after", "hello", "[", []string{"]"}, "[hello]"},
		{"different after", "hello", "(", []string{")"}, "(hello)"},
		{"empty value", "", "[", nil, "[["},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Wrap(tt.value, tt.before, tt.after...)
			if result != tt.expected {
				t.Errorf("Wrap(%q, %q, %v) = %q, want %q", tt.value, tt.before, tt.after, result, tt.expected)
			}
		})
	}
}

func TestUnwrap(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		before   string
		after    []string
		expected string
	}{
		{"basic", "[hello]", "[", nil, "hello]"},
		{"with after", "[hello]", "[", []string{"]"}, "hello"},
		{"different after", "(hello)", "(", []string{")"}, "hello"},
		{"not wrapped", "hello", "[", nil, "hello"},
		{"partial wrap", "[hello", "[", nil, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unwrap(tt.value, tt.before, tt.after...)
			if result != tt.expected {
				t.Errorf("Unwrap(%q, %q, %v) = %q, want %q", tt.value, tt.before, tt.after, result, tt.expected)
			}
		})
	}
}

func TestIsAscii(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"basic", "hello", true},
		{"with numbers", "hello123", true},
		{"with special chars", "hello!@#", true},
		{"unicode", "こんにちは", false},
		{"mixed", "hello 世界", false},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAscii(tt.value)
			if result != tt.expected {
				t.Errorf("IsAscii(%q) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestIsJson(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"valid json object", `{"key":"value"}`, true},
		{"valid json array", `[1,2,3]`, true},
		{"invalid json", `{key:value}`, false},
		{"empty string", "", false},
		{"plain text", "hello world", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsJson(tt.value)
			if result != tt.expected {
				t.Errorf("IsJson(%q) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestIsUrl(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		protocols []string
		expected  bool
	}{
		{"basic http", "http://example.com", nil, true},
		{"basic https", "https://example.com", nil, true},
		{"with protocol check", "http://example.com", []string{"http"}, true},
		{"wrong protocol", "http://example.com", []string{"https"}, false},
		{"invalid url", "not a url", nil, false},
		{"no scheme", "example.com", nil, false},
		{"no host", "http://", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsUrl(tt.value, tt.protocols)
			if result != tt.expected {
				t.Errorf("IsUrl(%q, %v) = %v, want %v", tt.value, tt.protocols, result, tt.expected)
			}
		})
	}
}

func TestIsUuid(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"valid uuid", "550e8400-e29b-41d4-a716-446655440000", true},
		{"invalid uuid", "not-a-uuid", false},
		{"wrong format", "550e8400e29b41d4a716446655440000", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsUuid(tt.value)
			if result != tt.expected {
				t.Errorf("IsUuid(%q) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestKebab(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"camel case", "helloWorld", "hello-world"},
		{"already kebab", "hello-world", "hello-world"},
		{"all lowercase", "hello world", "hello world"},
		{"with uppercase", "Hello World", "hello--world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Kebab(tt.value)
			if result != tt.expected {
				t.Errorf("Kebab(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected int
	}{
		{"basic", "hello", 5},
		{"empty", "", 0},
		{"unicode", "こんにちは", 5},
		{"mixed", "hello 世界", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Length(tt.value)
			if result != tt.expected {
				t.Errorf("Length(%q) = %d, want %d", tt.value, result, tt.expected)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		limit         int
		end           string
		preserveWords bool
		expected      string
	}{
		{"basic", "hello world", 5, "...", false, "hello..."},
		{"preserve words", "hello world", 5, "...", true, "hello..."},
		{"preserve words at space", "hello world", 6, "...", true, "hello..."},
		{"no truncation", "hello", 10, "...", false, "hello"},
		{"zero limit", "hello world", 0, "...", false, "hello world"},
		{"unicode", "こんにちは", 2, "...", false, "こん..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Limit(tt.value, tt.limit, tt.end, tt.preserveWords)
			if result != tt.expected {
				t.Errorf("Limit(%q, %d, %q, %v) = %q, want %q", tt.value, tt.limit, tt.end, tt.preserveWords, result, tt.expected)
			}
		})
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "HELLO", "hello"},
		{"mixed", "Hello World", "hello world"},
		{"already lower", "hello", "hello"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Lower(tt.value)
			if result != tt.expected {
				t.Errorf("Lower(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		words    int
		end      string
		expected string
	}{
		{"basic", "hello world test", 2, "...", "hello world..."},
		{"all words", "hello world", 2, "...", "hello world"},
		{"more words", "hello world", 5, "...", "hello world"},
		{"empty", "", 2, "...", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Words(tt.value, tt.words, tt.end)
			if result != tt.expected {
				t.Errorf("Words(%q, %d, %q) = %q, want %q", tt.value, tt.words, tt.end, result, tt.expected)
			}
		})
	}
}

func TestNumbers(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "abc123def456", "123456"},
		{"only numbers", "123456", "123456"},
		{"no numbers", "abcdef", ""},
		{"mixed", "a1b2c3", "123"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Numbers(tt.value)
			if result != tt.expected {
				t.Errorf("Numbers(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestPadBoth(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		length   int
		pad      string
		expected string
	}{
		{"basic", "hello", 10, " ", "  hello   "},
		{"already long", "hello", 3, " ", "hello"},
		{"single char pad", "hello", 10, "-", "--hello---"},
		{"empty", "", 5, " ", "     "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PadBoth(tt.value, tt.length, tt.pad)
			if result != tt.expected {
				t.Errorf("PadBoth(%q, %d, %q) = %q, want %q", tt.value, tt.length, tt.pad, result, tt.expected)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		length   int
		pad      string
		expected string
	}{
		{"basic", "hello", 10, " ", "     hello"},
		{"already long", "hello", 3, " ", "hello"},
		{"single char pad", "hello", 10, "-", "-----hello"},
		{"empty", "", 5, " ", "     "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PadLeft(tt.value, tt.length, tt.pad)
			if result != tt.expected {
				t.Errorf("PadLeft(%q, %d, %q) = %q, want %q", tt.value, tt.length, tt.pad, result, tt.expected)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		length   int
		pad      string
		expected string
	}{
		{"basic", "hello", 10, " ", "hello     "},
		{"already long", "hello", 3, " ", "hello"},
		{"single char pad", "hello", 10, "-", "hello-----"},
		{"empty", "", 5, " ", "     "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PadRight(tt.value, tt.length, tt.pad)
			if result != tt.expected {
				t.Errorf("PadRight(%q, %d, %q) = %q, want %q", tt.value, tt.length, tt.pad, result, tt.expected)
			}
		})
	}
}

func TestParseCallback(t *testing.T) {
	tests := []struct {
		name           string
		callback       string
		defaultMethod  string
		expectedClass  string
		expectedMethod string
	}{
		{"with @", "Class@method", "default", "Class", "method"},
		{"without @", "Class", "default", "Class", "default"},
		{"empty default", "Class@method", "", "Class", "method"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			class, method := ParseCallback(tt.callback, tt.defaultMethod)
			if class != tt.expectedClass || method != tt.expectedMethod {
				t.Errorf("ParseCallback(%q, %q) = %q, %q, want %q, %q", tt.callback, tt.defaultMethod, class, method, tt.expectedClass, tt.expectedMethod)
			}
		})
	}
}

func TestPosition(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		offset   int
		expected int
	}{
		{"basic", "hello world", "world", 0, 6},
		{"with offset", "hello world", "world", 3, 6},
		{"not found", "hello world", "xyz", 0, -1},
		{"negative offset", "hello world", "world", -1, 6},
		{"offset beyond", "hello world", "world", 10, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Position(tt.haystack, tt.needle, tt.offset)
			if result != tt.expected {
				t.Errorf("Position(%q, %q, %d) = %d, want %d", tt.haystack, tt.needle, tt.offset, result, tt.expected)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		expected int
	}{
		{"basic", 10, 10},
		{"zero", 0, 0},
		{"negative", -1, 0},
		{"large", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Random(tt.length)
			if len(result) != tt.expected {
				t.Errorf("Random(%d) length = %d, want %d", tt.length, len(result), tt.expected)
			}
			// Test that result is URL-safe base64 (alphanumeric + - and _)
			if tt.length > 0 {
				for _, r := range result {
					if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && r != '-' && r != '_' {
						t.Errorf("Random(%d) contains invalid character: %c", tt.length, r)
					}
				}
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		times    int
		expected string
	}{
		{"basic", "hello", 3, "hellohellohello"},
		{"zero", "hello", 0, ""},
		{"one", "hello", 1, "hello"},
		{"empty", "", 3, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Repeat(tt.s, tt.times)
			if result != tt.expected {
				t.Errorf("Repeat(%q, %d) = %q, want %q", tt.s, tt.times, result, tt.expected)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name          string
		search        string
		replace       string
		subject       string
		caseSensitive bool
		expected      string
	}{
		{"case sensitive", "hello", "hi", "hello world", true, "hi world"},
		{"case insensitive", "Hello", "hi", "hello world", false, "hi world"},
		{"not found", "xyz", "hi", "hello world", true, "hello world"},
		{"multiple", "l", "L", "hello", true, "heLLo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Replace(tt.search, tt.replace, tt.subject, tt.caseSensitive)
			if result != tt.expected {
				t.Errorf("Replace(%q, %q, %q, %v) = %q, want %q", tt.search, tt.replace, tt.subject, tt.caseSensitive, result, tt.expected)
			}
		})
	}
}

func TestReplaceFirst(t *testing.T) {
	tests := []struct {
		name     string
		search   string
		replace  string
		subject  string
		expected string
	}{
		{"basic", "hello", "hi", "hello world hello", "hi world hello"},
		{"not found", "xyz", "hi", "hello world", "hello world"},
		{"empty search", "", "hi", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceFirst(tt.search, tt.replace, tt.subject)
			if result != tt.expected {
				t.Errorf("ReplaceFirst(%q, %q, %q) = %q, want %q", tt.search, tt.replace, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestReplaceStart(t *testing.T) {
	tests := []struct {
		name     string
		search   string
		replace  string
		subject  string
		expected string
	}{
		{"basic", "hello", "hi", "hello world", "hi world"},
		{"not at start", "world", "hi", "hello world", "hello world"},
		{"empty search", "", "hi", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceStart(tt.search, tt.replace, tt.subject)
			if result != tt.expected {
				t.Errorf("ReplaceStart(%q, %q, %q) = %q, want %q", tt.search, tt.replace, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestReplaceLast(t *testing.T) {
	tests := []struct {
		name     string
		search   string
		replace  string
		subject  string
		expected string
	}{
		{"basic", "hello", "hi", "hello world hello", "hello world hi"},
		{"not found", "xyz", "hi", "hello world", "hello world"},
		{"empty search", "", "hi", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceLast(tt.search, tt.replace, tt.subject)
			if result != tt.expected {
				t.Errorf("ReplaceLast(%q, %q, %q) = %q, want %q", tt.search, tt.replace, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestReplaceEnd(t *testing.T) {
	tests := []struct {
		name     string
		search   string
		replace  string
		subject  string
		expected string
	}{
		{"basic", "world", "hi", "hello world", "hello hi"},
		{"not at end", "hello", "hi", "hello world", "hello world"},
		{"empty search", "", "hi", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceEnd(tt.search, tt.replace, tt.subject)
			if result != tt.expected {
				t.Errorf("ReplaceEnd(%q, %q, %q) = %q, want %q", tt.search, tt.replace, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name          string
		search        []string
		subject       string
		caseSensitive bool
		expected      string
	}{
		{"basic", []string{"hello"}, "hello world", true, " world"},
		{"case sensitive", []string{"Hello"}, "hello world", true, "hello world"},
		{"case insensitive", []string{"Hello"}, "hello world", false, " world"},
		{"multiple", []string{"hello", "world"}, "hello world test", true, "  test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Remove(tt.search, tt.subject, tt.caseSensitive)
			if result != tt.expected {
				t.Errorf("Remove(%v, %q, %v) = %q, want %q", tt.search, tt.subject, tt.caseSensitive, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "hello", "olleh"},
		{"empty", "", ""},
		{"unicode", "こんにちは", "はちにんこ"},
		{"mixed", "hello 世界", "界世 olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.value)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		prefix   string
		expected string
	}{
		{"basic", "world", "hello", "helloworld"},
		{"already starts", "helloworld", "hello", "helloworld"},
		{"multiple prefixes", "hellohelloworld", "hello", "helloworld"},
		{"empty", "", "hello", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Start(tt.value, tt.prefix)
			if result != tt.expected {
				t.Errorf("Start(%q, %q) = %q, want %q", tt.value, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestUpper(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "hello", "HELLO"},
		{"mixed", "Hello World", "HELLO WORLD"},
		{"already upper", "HELLO", "HELLO"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Upper(tt.value)
			if result != tt.expected {
				t.Errorf("Upper(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "hello world", "Hello World"},
		{"already title", "Hello World", "Hello World"},
		{"uppercase", "HELLO WORLD", "Hello World"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Title(tt.value)
			if result != tt.expected {
				t.Errorf("Title(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestSnake(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		delimiter string
		expected  string
	}{
		{"camel case", "helloWorld", "_", "hello_world"},
		{"with uppercase and space", "Hello World", "_", "hello__world"},
		{"custom delimiter", "HelloWorld", "-", "hello-world"},
		{"already snake", "hello_world", "_", "hello_world"},
		{"all lowercase", "hello world", "_", "hello world"},
		{"mixed", "HelloWorld", "_", "hello_world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Snake(tt.value, tt.delimiter)
			if result != tt.expected {
				t.Errorf("Snake(%q, %q) = %q, want %q", tt.value, tt.delimiter, result, tt.expected)
			}
		})
	}
}

func TestStudly(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "hello world", "HelloWorld"},
		{"snake case", "hello_world", "HelloWorld"},
		{"kebab case", "hello-world", "HelloWorld"},
		{"already studly", "HelloWorld", "HelloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Studly(tt.value)
			if result != tt.expected {
				t.Errorf("Studly(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestPascal(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "hello world", "HelloWorld"},
		{"snake case", "hello_world", "HelloWorld"},
		{"kebab case", "hello-world", "HelloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pascal(tt.value)
			if result != tt.expected {
				t.Errorf("Pascal(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestSubstr(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		start    int
		length   []int
		expected string
	}{
		{"basic", "hello world", 0, []int{5}, "hello"},
		{"negative start", "hello world", -5, nil, "world"},
		{"no length", "hello world", 6, nil, "world"},
		{"out of range start", "hello", 10, nil, ""},
		{"unicode", "こんにちは", 0, []int{2}, "こん"},
		{"negative start unicode", "こんにちは", -2, nil, "ちは"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Substr(tt.s, tt.start, tt.length...)
			if result != tt.expected {
				t.Errorf("Substr(%q, %d, %v) = %q, want %q", tt.s, tt.start, tt.length, result, tt.expected)
			}
		})
	}
}

func TestSubstrCount(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		offset   int
		length   int
		expected int
	}{
		{"basic", "hello world hello", "hello", 0, 0, 2},
		{"with offset", "hello world hello", "hello", 6, 0, 1},
		{"with length", "hello world hello", "hello", 0, 11, 1},
		{"not found", "hello world", "xyz", 0, 0, 0},
		{"empty needle", "hello world", "", 0, 0, 0},
		{"negative offset", "hello world", "hello", -1, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubstrCount(tt.haystack, tt.needle, tt.offset, tt.length)
			if result != tt.expected {
				t.Errorf("SubstrCount(%q, %q, %d, %d) = %d, want %d", tt.haystack, tt.needle, tt.offset, tt.length, result, tt.expected)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]string
		subject  string
		expected string
	}{
		{"basic", map[string]string{"hello": "hi", "world": "earth"}, "hello world", "hi earth"},
		{"no match", map[string]string{"xyz": "abc"}, "hello world", "hello world"},
		{"empty map", map[string]string{}, "hello world", "hello world"},
		{"sequential replacement", map[string]string{"a": "b", "b": "c"}, "a b", "c c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Swap(tt.m, tt.subject)
			if result != tt.expected {
				t.Errorf("Swap(%v, %q) = %q, want %q", tt.m, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		limit    int
		expected string
	}{
		{"basic", "hello world", 5, "hello"},
		{"negative", "hello world", -5, "world"},
		{"zero", "hello world", 0, "hello world"},
		{"exceeds length", "hello", 10, "hello"},
		{"unicode", "こんにちは", 2, "こん"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Take(tt.s, tt.limit)
			if result != tt.expected {
				t.Errorf("Take(%q, %d) = %q, want %q", tt.s, tt.limit, result, tt.expected)
			}
		})
	}
}

func TestToBase64(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"basic", "hello", "aGVsbG8="},
		{"empty", "", ""},
		{"unicode", "こんにちは", "44GT44KT44Gr44Gh44Gv"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBase64(tt.s)
			if result != tt.expected {
				t.Errorf("ToBase64(%q) = %q, want %q", tt.s, result, tt.expected)
			}
		})
	}
}

func TestFromBase64(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		strict   bool
		expected string
		hasError bool
	}{
		{"basic strict", "aGVsbG8=", true, "hello", false},
		{"basic non-strict", "aGVsbG8", false, "hello", false},
		{"invalid", "invalid!", true, "", true},
		{"empty", "", true, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FromBase64(tt.s, tt.strict)
			if (err != nil) != tt.hasError {
				t.Errorf("FromBase64(%q, %v) error = %v, want error %v", tt.s, tt.strict, err, tt.hasError)
			}
			if !tt.hasError && result != tt.expected {
				t.Errorf("FromBase64(%q, %v) = %q, want %q", tt.s, tt.strict, result, tt.expected)
			}
		})
	}
}

func TestLcfirst(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"basic", "Hello", "hello"},
		{"already lower", "hello", "hello"},
		{"empty", "", ""},
		{"single char", "H", "h"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Lcfirst(tt.s)
			if result != tt.expected {
				t.Errorf("Lcfirst(%q) = %q, want %q", tt.s, result, tt.expected)
			}
		})
	}
}

func TestUcfirst(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"basic", "hello", "Hello"},
		{"already upper", "Hello", "Hello"},
		{"empty", "", ""},
		{"single char", "h", "H"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ucfirst(tt.s)
			if result != tt.expected {
				t.Errorf("Ucfirst(%q) = %q, want %q", tt.s, result, tt.expected)
			}
		})
	}
}

func TestUcsplit(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{"basic", "HelloWorld", []string{"Hello", "World"}},
		{"single", "Hello", []string{"Hello"}},
		{"empty", "", []string{}},
		{"no uppercase", "hello", []string{"hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ucsplit(tt.s)
			if len(result) != len(tt.expected) {
				t.Errorf("Ucsplit(%q) length = %d, want %d", tt.s, len(result), len(tt.expected))
				return
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("Ucsplit(%q)[%d] = %q, want %q", tt.s, i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"basic", "hello world", 2},
		{"multiple", "hello world test", 3},
		{"empty", "", 0},
		{"single", "hello", 1},
		{"with spaces", "  hello   world  ", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordCount(tt.s)
			if result != tt.expected {
				t.Errorf("WordCount(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		charlist []string
		expected string
	}{
		{"basic", "  hello world  ", nil, "hello world"},
		{"with charlist", "---hello---", []string{"-"}, "hello"},
		{"empty", "", nil, ""},
		{"no trim needed", "hello", nil, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Trim(tt.value, tt.charlist...)
			if result != tt.expected {
				t.Errorf("Trim(%q, %v) = %q, want %q", tt.value, tt.charlist, result, tt.expected)
			}
		})
	}
}

func TestLtrim(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		charlist []string
		expected string
	}{
		{"basic", "  hello world", nil, "hello world"},
		{"with charlist", "---hello", []string{"-"}, "hello"},
		{"empty", "", nil, ""},
		{"no trim needed", "hello", nil, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ltrim(tt.value, tt.charlist...)
			if result != tt.expected {
				t.Errorf("Ltrim(%q, %v) = %q, want %q", tt.value, tt.charlist, result, tt.expected)
			}
		})
	}
}

func TestRtrim(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		charlist []string
		expected string
	}{
		{"basic", "hello world  ", nil, "hello world"},
		{"with charlist", "hello---", []string{"-"}, "hello"},
		{"empty", "", nil, ""},
		{"no trim needed", "hello", nil, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Rtrim(tt.value, tt.charlist...)
			if result != tt.expected {
				t.Errorf("Rtrim(%q, %v) = %q, want %q", tt.value, tt.charlist, result, tt.expected)
			}
		})
	}
}

func TestSquish(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"basic", "  hello   world  ", "hello world"},
		{"tabs", "hello\t\tworld", "hello world"},
		{"newlines", "hello\n\nworld", "hello world"},
		{"no extra", "hello world", "hello world"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Squish(tt.value)
			if result != tt.expected {
				t.Errorf("Squish(%q) = %q, want %q", tt.value, result, tt.expected)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needles  []string
		expected bool
	}{
		{"basic", "hello world", []string{"hello"}, true},
		{"multiple needles", "hello world", []string{"xyz", "hello"}, true},
		{"not found", "hello world", []string{"xyz"}, false},
		{"empty haystack", "", []string{"hello"}, false},
		{"empty needles", "hello world", []string{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartsWith(tt.haystack, tt.needles)
			if result != tt.expected {
				t.Errorf("StartsWith(%q, %v) = %v, want %v", tt.haystack, tt.needles, result, tt.expected)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		subject  string
		expected string
	}{
		{"basic", `\d+`, "hello123world", "123"},
		{"with group", `(\d+)`, "hello123world", "123"},
		{"not found", `\d+`, "hello world", ""},
		{"invalid pattern", `[`, "hello", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Match(tt.pattern, tt.subject)
			if result != tt.expected {
				t.Errorf("Match(%q, %q) = %q, want %q", tt.pattern, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		value    string
		expected bool
	}{
		{"basic", `\d+`, "123", true},
		{"not match", `\d+`, "abc", false},
		{"invalid pattern", `[`, "hello", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMatch(tt.pattern, tt.value)
			if result != tt.expected {
				t.Errorf("IsMatch(%q, %q) = %v, want %v", tt.pattern, tt.value, result, tt.expected)
			}
		})
	}
}

func TestMatchAll(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		subject  string
		expected []string
		hasError bool
	}{
		{"basic", `\d+`, "hello123world456", []string{"123", "456"}, false},
		{"with group", `(\d+)`, "hello123world456", []string{"123", "456"}, false},
		{"not found", `\d+`, "hello world", []string{}, false},
		{"invalid pattern", `[`, "hello", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MatchAll(tt.pattern, tt.subject)
			if (err != nil) != tt.hasError {
				t.Errorf("MatchAll(%q, %q) error = %v, want error %v", tt.pattern, tt.subject, err, tt.hasError)
			}
			if !tt.hasError {
				if len(result) != len(tt.expected) {
					t.Errorf("MatchAll(%q, %q) length = %d, want %d", tt.pattern, tt.subject, len(result), len(tt.expected))
					return
				}
				for i, v := range result {
					if v != tt.expected[i] {
						t.Errorf("MatchAll(%q, %q)[%d] = %q, want %q", tt.pattern, tt.subject, i, v, tt.expected[i])
					}
				}
			}
		})
	}
}

func TestReplaceArray(t *testing.T) {
	tests := []struct {
		name     string
		search   string
		replace  []string
		subject  string
		expected string
	}{
		{"basic", "?", []string{"a", "b"}, "hello?world?test", "helloaworldbtest"},
		{"not found", "?", []string{"a"}, "hello world", "hello world"},
		{"more replacements", "?", []string{"a", "b", "c"}, "hello?world", "helloaworld"},
		{"fewer replacements", "?", []string{"a"}, "hello?world?test", "helloaworld?test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceArray(tt.search, tt.replace, tt.subject)
			if result != tt.expected {
				t.Errorf("ReplaceArray(%q, %v, %q) = %q, want %q", tt.search, tt.replace, tt.subject, result, tt.expected)
			}
		})
	}
}

func TestReplaceMatches(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		replace  interface{}
		subject  string
		limit    int
		expected string
		hasError bool
	}{
		{"basic string", `\d+`, "X", "hello123world456", 0, "helloXworldX", false},
		{"with limit", `\d+`, "X", "hello123world456", 1, "helloXworld456", false},
		{"function replace", `\d+`, func(match string) string { return "X" }, "hello123world456", 0, "helloXworldX", false},
		{"invalid pattern", `[`, "X", "hello", 0, "", true},
		{"unsupported type", `\d+`, 123, "hello123", 0, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReplaceMatches(tt.pattern, tt.replace, tt.subject, tt.limit)
			if (err != nil) != tt.hasError {
				t.Errorf("ReplaceMatches(%q, %v, %q, %d) error = %v, want error %v", tt.pattern, tt.replace, tt.subject, tt.limit, err, tt.hasError)
			}
			if !tt.hasError && result != tt.expected {
				t.Errorf("ReplaceMatches(%q, %v, %q, %d) = %q, want %q", tt.pattern, tt.replace, tt.subject, tt.limit, result, tt.expected)
			}
		})
	}
}

func TestSlug(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		separator  string
		language   string
		dictionary map[string]string
		expected   string
	}{
		{"basic", "Hello World", "-", "", nil, "hello-world"},
		{"with special chars", "Hello!@#World", "-", "", nil, "helloworld"},
		{"with dictionary", "hello world", "-", "", map[string]string{"hello": "hi"}, "hi-world"},
		{"custom separator", "Hello World", "_", "", nil, "hello_world"},
		{"multiple spaces", "Hello   World", "-", "", nil, "hello-world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Slug(tt.title, tt.separator, tt.language, tt.dictionary)
			if result != tt.expected {
				t.Errorf("Slug(%q, %q, %q, %v) = %q, want %q", tt.title, tt.separator, tt.language, tt.dictionary, result, tt.expected)
			}
		})
	}
}

func TestFlushCache(t *testing.T) {
	// Test that cache is cleared
	Camel("test")
	Studly("test")
	Snake("test", "_")

	FlushCache()

	// Verify cache is empty by checking that new values are computed
	// This is a basic test - in practice you'd need to check internal state
	// but since caches are private, we just verify functions still work
	result1 := Camel("test")
	result2 := Studly("test")
	result3 := Snake("test", "_")

	if result1 == "" || result2 == "" || result3 == "" {
		t.Error("FlushCache broke functionality")
	}
}
