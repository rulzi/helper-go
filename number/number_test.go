package number

import (
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name         string
		number       float64
		precision    *int
		maxPrecision *int
		locale       *string
		expected     string
	}{
		{"basic default", 123.456, nil, nil, nil, "123.456"},
		{"with precision", 123.456, intPtr(2), nil, nil, "123.46"},
		{"with max precision", 123.456, nil, intPtr(2), nil, "123.46"},
		{"max precision overrides precision", 123.456, intPtr(1), intPtr(2), nil, "123.46"},
		{"indonesian locale", 123.456, nil, nil, stringPtr("id"), "123,456"},
		{"indonesian with precision", 123.456, intPtr(2), nil, stringPtr("id"), "123,46"},
		{"indonesian prefix locale", 123.456, nil, nil, stringPtr("id_ID"), "123,456"},
		{"zero", 0.0, nil, nil, nil, "0"},
		{"negative", -123.456, nil, nil, nil, "-123.456"},
		{"large number", 1234567.89, nil, nil, nil, "1234567.89"},
		{"small number", 0.001, nil, nil, nil, "0.001"},
		{"no decimal", 123.0, nil, nil, nil, "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Format(tt.number, tt.precision, tt.maxPrecision, tt.locale)
			if result != tt.expected {
				t.Errorf("Format(%f, %v, %v, %v) = %q, want %q", tt.number, tt.precision, tt.maxPrecision, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestSpell(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   *string
		after    *int
		until    *int
		expected string
	}{
		{"basic", 123.0, nil, nil, nil, "123"},
		{"zero", 0.0, nil, nil, nil, "zero"},
		{"with after threshold", 5.0, nil, intPtr(10), nil, "5"},
		{"above after threshold", 15.0, nil, intPtr(10), nil, "15"},
		{"with until threshold", 5.0, nil, nil, intPtr(10), "5"},
		{"above until threshold", 15.0, nil, nil, intPtr(10), "15"},
		{"with locale", 123.0, stringPtr("id"), nil, nil, "123"},
		{"negative", -123.0, nil, nil, nil, "-123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Spell(tt.number, tt.locale, tt.after, tt.until)
			if result != tt.expected {
				t.Errorf("Spell(%f, %v, %v, %v) = %q, want %q", tt.number, tt.locale, tt.after, tt.until, result, tt.expected)
			}
		})
	}
}

func TestOrdinal(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   *string
		expected string
	}{
		{"first", 1.0, nil, "1st"},
		{"second", 2.0, nil, "2nd"},
		{"third", 3.0, nil, "3rd"},
		{"fourth", 4.0, nil, "4th"},
		{"eleventh", 11.0, nil, "11th"},
		{"twelfth", 12.0, nil, "12th"},
		{"thirteenth", 13.0, nil, "13th"},
		{"twenty-first", 21.0, nil, "21st"},
		{"twenty-second", 22.0, nil, "22nd"},
		{"twenty-third", 23.0, nil, "23rd"},
		{"hundredth", 100.0, nil, "100th"},
		{"negative", -1.0, nil, "-1st"},
		{"zero", 0.0, nil, "0th"},
		{"with locale", 1.0, stringPtr("id"), "1st"},
		{"decimal rounds down", 1.9, nil, "1st"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ordinal(tt.number, tt.locale)
			if result != tt.expected {
				t.Errorf("Ordinal(%f, %v) = %q, want %q", tt.number, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestSpellOrdinal(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   *string
		expected string
	}{
		{"basic", 1.0, nil, "1st ordinal"},
		{"second", 2.0, nil, "2nd ordinal"},
		{"with locale", 1.0, stringPtr("id"), "1st ordinal"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SpellOrdinal(tt.number, tt.locale)
			if result != tt.expected {
				t.Errorf("SpellOrdinal(%f, %v) = %q, want %q", tt.number, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name         string
		number       float64
		precision    int
		maxPrecision *int
		locale       *string
		expected     string
	}{
		{"basic", 50.0, 2, nil, nil, "0.50%"},
		{"with precision", 50.0, 1, nil, nil, "0.5%"},
		{"with max precision", 50.0, 3, intPtr(2), nil, "0.50%"},
		{"indonesian locale", 50.0, 2, nil, stringPtr("id"), "0,50%"},
		{"zero", 0.0, 2, nil, nil, "0.00%"},
		{"hundred", 100.0, 2, nil, nil, "1.00%"},
		{"negative", -50.0, 2, nil, nil, "-0.50%"},
		{"small value", 0.5, 2, nil, nil, "0.01%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Percentage(tt.number, tt.precision, tt.maxPrecision, tt.locale)
			if result != tt.expected {
				t.Errorf("Percentage(%f, %d, %v, %v) = %q, want %q", tt.number, tt.precision, tt.maxPrecision, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestCurrency(t *testing.T) {
	tests := []struct {
		name      string
		number    float64
		in        string
		locale    *string
		precision *int
		expected  string
	}{
		{"basic USD", 123.45, "USD", nil, nil, "$123.45"},
		{"EUR", 123.45, "EUR", nil, nil, "€123.45"},
		{"GBP", 123.45, "GBP", nil, nil, "£123.45"},
		{"IDR", 123.45, "IDR", nil, nil, "Rp123.45"},
		{"IDR indonesian locale", 123.45, "IDR", stringPtr("id"), nil, "Rp 123,45"},
		{"with precision", 123.456, "USD", nil, intPtr(3), "$123.456"},
		{"default currency", 123.45, "", nil, nil, "$123.45"},
		{"unknown currency", 123.45, "XYZ", nil, nil, "XYZ 123.45"},
		{"zero", 0.0, "USD", nil, nil, "$0.00"},
		{"negative", -123.45, "USD", nil, nil, "$-123.45"},
		{"indonesian locale USD", 123.45, "USD", stringPtr("id"), nil, "$ 123,45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Currency(tt.number, tt.in, tt.locale, tt.precision)
			if result != tt.expected {
				t.Errorf("Currency(%f, %q, %v, %v) = %q, want %q", tt.number, tt.in, tt.locale, tt.precision, result, tt.expected)
			}
		})
	}
}

func TestFileSize(t *testing.T) {
	tests := []struct {
		name         string
		bytes        float64
		precision    int
		maxPrecision *int
		expected     string
	}{
		{"bytes", 500.0, 0, nil, "500 B"},
		{"kilobytes", 1024.0, 2, nil, "1.00 KB"},
		{"megabytes", 1048576.0, 2, nil, "1.00 MB"},
		{"gigabytes", 1073741824.0, 2, nil, "1.00 GB"},
		{"terabytes", 1099511627776.0, 2, nil, "1.00 TB"},
		{"with precision", 1536.0, 1, nil, "1.5 KB"},
		{"with max precision", 1536.0, 3, intPtr(1), "1.5 KB"},
		{"large value", 1024.0 * 1024 * 1024 * 1024, 2, nil, "1.00 TB"},
		{"zero", 0.0, 0, nil, "0 B"},
		{"fractional KB", 512.0, 2, nil, "512.00 B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FileSize(tt.bytes, tt.precision, tt.maxPrecision)
			if result != tt.expected {
				t.Errorf("FileSize(%f, %d, %v) = %q, want %q", tt.bytes, tt.precision, tt.maxPrecision, result, tt.expected)
			}
		})
	}
}

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		name         string
		number       float64
		precision    int
		maxPrecision *int
		expected     string
	}{
		{"thousands", 1000.0, 0, nil, "1K"},
		{"millions", 1000000.0, 0, nil, "1M"},
		{"billions", 1000000000.0, 0, nil, "1B"},
		{"trillions", 1000000000000.0, 0, nil, "1T"},
		{"with precision", 1500.0, 1, nil, "1.5K"},
		{"with max precision", 1500.0, 3, intPtr(1), "1.5K"},
		{"zero", 0.0, 0, nil, "0"},
		{"negative", -1000.0, 0, nil, "-1K"},
		{"less than thousand", 500.0, 0, nil, "500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abbreviate(tt.number, tt.precision, tt.maxPrecision)
			if result != tt.expected {
				t.Errorf("Abbreviate(%f, %d, %v) = %q, want %q", tt.number, tt.precision, tt.maxPrecision, result, tt.expected)
			}
		})
	}
}

func TestForHumans(t *testing.T) {
	tests := []struct {
		name         string
		number       float64
		precision    int
		maxPrecision *int
		abbreviate   bool
		expected     string
	}{
		{"abbreviated thousands", 1000.0, 0, nil, true, "1K"},
		{"abbreviated millions", 1000000.0, 0, nil, true, "1M"},
		{"full thousands", 1000.0, 0, nil, false, "1 thousand"},
		{"full millions", 1000000.0, 0, nil, false, "1 million"},
		{"with precision abbreviated", 1500.0, 1, nil, true, "1.5K"},
		{"with precision full", 1500.0, 1, nil, false, "1.5 thousand"},
		{"zero", 0.0, 0, nil, true, "0"},
		{"negative", -1000.0, 0, nil, true, "-1K"},
		{"less than thousand", 500.0, 0, nil, true, "500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ForHumans(tt.number, tt.precision, tt.maxPrecision, tt.abbreviate)
			if result != tt.expected {
				t.Errorf("ForHumans(%f, %d, %v, %v) = %q, want %q", tt.number, tt.precision, tt.maxPrecision, tt.abbreviate, result, tt.expected)
			}
		})
	}
}

func TestSummarize(t *testing.T) {
	tests := []struct {
		name         string
		number       float64
		precision    int
		maxPrecision *int
		units        map[int]string
		expected     string
	}{
		{"basic with default units", 1000.0, 0, nil, nil, "1K"},
		{"custom units", 1000.0, 0, nil, map[int]string{3: "kilo"}, "1kilo"},
		{"zero", 0.0, 0, nil, nil, "0"},
		{"zero with precision", 0.0, 2, nil, nil, "0.00"},
		{"negative", -1000.0, 0, nil, nil, "-1K"},
		{"very large number", 1e15, 0, nil, nil, "1Q"},
		{"with precision", 1500.0, 1, nil, nil, "1.5K"},
		{"empty units map", 1000.0, 0, nil, map[int]string{}, "1K"},
		{"number without unit match", 500.0, 0, nil, nil, "500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Summarize(tt.number, tt.precision, tt.maxPrecision, tt.units)
			if result != tt.expected {
				t.Errorf("Summarize(%f, %d, %v, %v) = %q, want %q", tt.number, tt.precision, tt.maxPrecision, tt.units, result, tt.expected)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		min      float64
		max      float64
		expected float64
	}{
		{"within range", 5.0, 0.0, 10.0, 5.0},
		{"below min", -5.0, 0.0, 10.0, 0.0},
		{"above max", 15.0, 0.0, 10.0, 10.0},
		{"at min", 0.0, 0.0, 10.0, 0.0},
		{"at max", 10.0, 0.0, 10.0, 10.0},
		{"negative range", -5.0, -10.0, -1.0, -5.0},
		{"decimal", 5.5, 0.0, 10.0, 5.5},
		{"min greater than max", 5.0, 10.0, 0.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clamp(tt.number, tt.min, tt.max)
			if result != tt.expected {
				t.Errorf("Clamp(%f, %f, %f) = %f, want %f", tt.number, tt.min, tt.max, result, tt.expected)
			}
		})
	}
}

func TestPairs(t *testing.T) {
	tests := []struct {
		name     string
		to       float64
		by       float64
		offset   float64
		expected [][]float64
	}{
		{"basic", 10.0, 2.0, 0.0, [][]float64{{0.0, 2.0}, {2.0, 4.0}, {4.0, 6.0}, {6.0, 8.0}, {8.0, 10.0}}},
		{"with offset", 10.0, 2.0, 1.0, [][]float64{{1.0, 2.0}, {3.0, 4.0}, {5.0, 6.0}, {7.0, 8.0}, {9.0, 10.0}}},
		{"uneven division", 10.0, 3.0, 0.0, [][]float64{{0.0, 3.0}, {3.0, 6.0}, {6.0, 9.0}, {9.0, 10.0}}},
		{"small range", 2.0, 1.0, 0.0, [][]float64{{0.0, 1.0}, {1.0, 2.0}}},
		{"zero to", 0.0, 1.0, 0.0, [][]float64{}},
		{"decimal step", 5.0, 1.5, 0.0, [][]float64{{0.0, 1.5}, {1.5, 3.0}, {3.0, 4.5}, {4.5, 5.0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pairs(tt.to, tt.by, tt.offset)
			if len(result) != len(tt.expected) {
				t.Errorf("Pairs(%f, %f, %f) length = %d, want %d", tt.to, tt.by, tt.offset, len(result), len(tt.expected))
				return
			}
			for i, pair := range result {
				if len(pair) != 2 {
					t.Errorf("Pairs(%f, %f, %f)[%d] length = %d, want 2", tt.to, tt.by, tt.offset, i, len(pair))
					continue
				}
				if pair[0] != tt.expected[i][0] || pair[1] != tt.expected[i][1] {
					t.Errorf("Pairs(%f, %f, %f)[%d] = [%f, %f], want [%f, %f]", tt.to, tt.by, tt.offset, i, pair[0], pair[1], tt.expected[i][0], tt.expected[i][1])
				}
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		expected float64
	}{
		{"with trailing zeros", 123.4500, 123.45},
		{"no trailing zeros", 123.45, 123.45},
		{"integer", 123.0, 123.0},
		{"zero", 0.0, 0.0},
		{"negative", -123.4500, -123.45},
		{"small decimal", 0.001, 0.001},
		{"many trailing zeros", 123.000000, 123.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Trim(tt.number)
			if result != tt.expected {
				t.Errorf("Trim(%f) = %f, want %f", tt.number, result, tt.expected)
			}
		})
	}
}

func TestWithLocale(t *testing.T) {
	// Save original locale
	originalLocale := DefaultLocale()
	defer UseLocale(originalLocale)

	t.Run("changes locale in callback", func(t *testing.T) {
		UseLocale("en")
		WithLocale("id", func() {
			if DefaultLocale() != "id" {
				t.Errorf("WithLocale: locale should be 'id' inside callback, got %q", DefaultLocale())
			}
		})
		if DefaultLocale() != "en" {
			t.Errorf("WithLocale: locale should be restored to 'en' after callback, got %q", DefaultLocale())
		}
	})

	t.Run("restores locale after panic", func(t *testing.T) {
		UseLocale("en")
		func() {
			defer func() {
				_ = recover() // Expected panic
			}()
			WithLocale("id", func() {
				panic("test panic")
			})
		}()
		// Note: In Go, defer in WithLocale should still execute even after panic
		// This test verifies the locale is restored
		if DefaultLocale() != "en" {
			t.Errorf("WithLocale: locale should be restored after panic, got %q", DefaultLocale())
		}
	})
}

func TestWithCurrency(t *testing.T) {
	// Save original currency
	originalCurrency := DefaultCurrency()
	defer UseCurrency(originalCurrency)

	t.Run("changes currency in callback", func(t *testing.T) {
		UseCurrency("USD")
		WithCurrency("EUR", func() {
			if DefaultCurrency() != "EUR" {
				t.Errorf("WithCurrency: currency should be 'EUR' inside callback, got %q", DefaultCurrency())
			}
		})
		if DefaultCurrency() != "USD" {
			t.Errorf("WithCurrency: currency should be restored to 'USD' after callback, got %q", DefaultCurrency())
		}
	})
}

func TestUseLocale(t *testing.T) {
	// Save original locale
	originalLocale := DefaultLocale()
	defer UseLocale(originalLocale)

	tests := []struct {
		name   string
		locale string
	}{
		{"english", "en"},
		{"indonesian", "id"},
		{"indonesian with region", "id_ID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseLocale(tt.locale)
			if DefaultLocale() != tt.locale {
				t.Errorf("UseLocale(%q): DefaultLocale() = %q, want %q", tt.locale, DefaultLocale(), tt.locale)
			}
		})
	}
}

func TestUseCurrency(t *testing.T) {
	// Save original currency
	originalCurrency := DefaultCurrency()
	defer UseCurrency(originalCurrency)

	tests := []struct {
		name     string
		currency string
	}{
		{"USD", "USD"},
		{"EUR", "EUR"},
		{"IDR", "IDR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseCurrency(tt.currency)
			if DefaultCurrency() != tt.currency {
				t.Errorf("UseCurrency(%q): DefaultCurrency() = %q, want %q", tt.currency, DefaultCurrency(), tt.currency)
			}
		})
	}
}

func TestDefaultLocale(t *testing.T) {
	// Save original locale
	originalLocale := DefaultLocale()
	defer UseLocale(originalLocale)

	t.Run("returns current locale", func(t *testing.T) {
		UseLocale("id")
		result := DefaultLocale()
		if result != "id" {
			t.Errorf("DefaultLocale() = %q, want %q", result, "id")
		}
	})
}

func TestDefaultCurrency(t *testing.T) {
	// Save original currency
	originalCurrency := DefaultCurrency()
	defer UseCurrency(originalCurrency)

	t.Run("returns current currency", func(t *testing.T) {
		UseCurrency("EUR")
		result := DefaultCurrency()
		if result != "EUR" {
			t.Errorf("DefaultCurrency() = %q, want %q", result, "EUR")
		}
	})
}

// Helper functions for test
func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}
