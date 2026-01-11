package number

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

var (
	// Default locale for number formatting
	defaultLocale = "en"
	// Default currency code
	defaultCurrency = "USD"
	// Mutex for thread-safe access to default settings
	mu sync.RWMutex

	// Currency symbols map - initialized once for better performance
	currencySymbols = map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
		"IDR": "Rp",
		"INR": "₹",
		"CNY": "¥",
		"AUD": "A$",
		"CAD": "C$",
	}

	// File size units - initialized once
	fileSizeUnits = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	// Abbreviated units for ForHumans
	abbreviatedUnits = map[int]string{
		3:  "K",
		6:  "M",
		9:  "B",
		12: "T",
		15: "Q",
	}

	// Full units for ForHumans
	fullUnits = map[int]string{
		3:  " thousand",
		6:  " million",
		9:  " billion",
		12: " trillion",
		15: " quadrillion",
	}

	// Default units for Summarize
	defaultUnits = map[int]string{
		3:  "K",
		6:  "M",
		9:  "B",
		12: "T",
		15: "Q",
	}

	// Replacer for Indonesian locale (comma as decimal separator)
	idReplacer = strings.NewReplacer(".", ",")
)

// Format formats the given number according to the current locale.
// For full locale support, consider using golang.org/x/text/message
func Format(number float64, precision *int, maxPrecision *int, locale *string) string {
	loc := defaultLocale
	if locale != nil {
		loc = *locale
	}

	// Basic formatting - for full locale support, use golang.org/x/text
	var prec int
	if maxPrecision != nil {
		prec = *maxPrecision
	} else if precision != nil {
		prec = *precision
	} else {
		prec = -1 // Use default precision
	}

	// Simple formatting based on locale
	if loc == "id" || strings.HasPrefix(loc, "id_") {
		// Indonesian locale uses comma as decimal separator
		return formatNumberID(number, prec)
	}

	// Default (English) formatting
	return formatNumberEN(number, prec)
}

func formatNumberEN(number float64, precision int) string {
	if precision < 0 {
		return strconv.FormatFloat(number, 'f', -1, 64)
	}
	return strconv.FormatFloat(number, 'f', precision, 64)
}

func formatNumberID(number float64, precision int) string {
	var formatted string
	if precision < 0 {
		formatted = strconv.FormatFloat(number, 'f', -1, 64)
	} else {
		formatted = strconv.FormatFloat(number, 'f', precision, 64)
	}
	return idReplacer.Replace(formatted)
}

// Spell spells out the given number in the given locale.
// Note: Full spellout support requires external library like golang.org/x/text
func Spell(number float64, locale *string, after *int, until *int) string {
	// Get locale once
	loc := defaultLocale
	if locale != nil {
		loc = *locale
	}

	if after != nil && number <= float64(*after) {
		return Format(number, nil, nil, &loc)
	}

	if until != nil && number >= float64(*until) {
		return Format(number, nil, nil, &loc)
	}

	// Basic spellout - for full support, use golang.org/x/text
	return spelloutBasic(number)
}

func spelloutBasic(number float64) string {
	// Very basic implementation
	// For production, use golang.org/x/text/message
	if number == 0 {
		return "zero"
	}
	return fmt.Sprintf("%.0f", number)
}

// Ordinal converts the given number to ordinal form (1st, 2nd, 3rd, etc.)
func Ordinal(number float64, locale *string) string {
	n := int64(number)
	suffix := getOrdinalSuffix(n)
	return fmt.Sprintf("%d%s", n, suffix)
}

func getOrdinalSuffix(n int64) string {
	if n < 0 {
		n = -n
	}

	lastDigit := n % 10
	lastTwoDigits := n % 100

	if lastTwoDigits >= 11 && lastTwoDigits <= 13 {
		return "th"
	}

	switch lastDigit {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// SpellOrdinal spells out the given number in ordinal form
func SpellOrdinal(number float64, locale *string) string {
	// Basic implementation
	ord := Ordinal(number, locale)
	return fmt.Sprintf("%s ordinal", ord)
}

// Percentage converts the given number to its percentage equivalent
func Percentage(number float64, precision int, maxPrecision *int, locale *string) string {
	percentage := number / 100.0

	var prec int
	if maxPrecision != nil {
		prec = *maxPrecision
	} else {
		prec = precision
	}

	formatted := Format(percentage, &prec, nil, locale)
	return formatted + "%"
}

// Currency converts the given number to its currency equivalent
func Currency(number float64, in string, locale *string, precision *int) string {
	loc := defaultLocale
	if locale != nil {
		loc = *locale
	}

	curr := defaultCurrency
	if in != "" {
		curr = in
	}

	var prec int
	if precision != nil {
		prec = *precision
	} else {
		prec = 2 // Default currency precision
	}

	formatted := Format(number, &prec, nil, &loc)

	// Add currency symbol
	symbol := getCurrencySymbol(curr)
	// Indonesian locale uses space between symbol and number
	if len(loc) >= 2 && loc[:2] == "id" {
		return symbol + " " + formatted
	}
	return symbol + formatted
}

func getCurrencySymbol(currency string) string {
	if symbol, ok := currencySymbols[currency]; ok {
		return symbol
	}
	return currency + " "
}

// FileSize converts the given number to its file size equivalent
func FileSize(bytes float64, precision int, maxPrecision *int) string {
	i := 0
	size := bytes
	for size >= 1024 && i < len(fileSizeUnits)-1 {
		size /= 1024
		i++
	}

	var prec int
	if maxPrecision != nil {
		prec = *maxPrecision
	} else {
		prec = precision
	}

	formatted := Format(size, &prec, nil, nil)
	return formatted + " " + fileSizeUnits[i]
}

// Abbreviate converts the number to its human-readable abbreviated equivalent
func Abbreviate(number float64, precision int, maxPrecision *int) string {
	return ForHumans(number, precision, maxPrecision, true)
}

// ForHumans converts the number to its human-readable equivalent
func ForHumans(number float64, precision int, maxPrecision *int, abbreviate bool) string {
	if abbreviate {
		return Summarize(number, precision, maxPrecision, abbreviatedUnits)
	}
	return Summarize(number, precision, maxPrecision, fullUnits)
}

// Summarize converts the number to its human-readable equivalent with custom units
func Summarize(number float64, precision int, maxPrecision *int, units map[int]string) string {
	if len(units) == 0 {
		units = defaultUnits
	}

	if number == 0.0 {
		if precision > 0 {
			prec := precision
			return Format(0, &prec, maxPrecision, nil)
		}
		return "0"
	}

	if number < 0 {
		return "-" + Summarize(-number, precision, maxPrecision, units)
	}

	if number >= 1e15 {
		// Find the largest unit
		maxExp := 15
		for exp := range units {
			if exp > maxExp {
				maxExp = exp
			}
		}
		unit, exists := units[maxExp]
		if !exists {
			// Fallback to default unit if maxExp not in units
			unit = "Q"
		}
		formatted := Summarize(number/1e15, precision, maxPrecision, units)
		return strings.TrimSpace(formatted + unit)
	}

	numberExponent := int(math.Floor(math.Log10(number)))
	displayExponent := numberExponent - (numberExponent % 3)
	number /= math.Pow(10, float64(displayExponent))

	var prec int
	if maxPrecision != nil {
		prec = *maxPrecision
	} else {
		prec = precision
	}

	formatted := Format(number, &prec, nil, nil)
	unit, exists := units[displayExponent]
	if !exists || unit == "" {
		return formatted
	}

	return strings.TrimSpace(formatted + unit)
}

// Clamp clamps the given number between the given minimum and maximum
func Clamp(number, min, max float64) float64 {
	if number < min {
		return min
	}
	if number > max {
		return max
	}
	return number
}

// Pairs splits the given number into pairs of min/max values
func Pairs(to, by, offset float64) [][]float64 {
	var output [][]float64

	for lower := 0.0; lower < to; lower += by {
		upper := lower + by

		if upper > to {
			upper = to
		}

		output = append(output, []float64{lower + offset, upper})
	}

	return output
}

// Trim removes any trailing zero digits after the decimal point
func Trim(number float64) float64 {
	// Convert to string and back to remove trailing zeros
	str := strconv.FormatFloat(number, 'f', -1, 64)
	trimmed, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return number
	}
	return trimmed
}

// WithLocale executes the given callback using the given locale
func WithLocale(locale string, callback func()) {
	mu.Lock()
	previousLocale := defaultLocale
	defaultLocale = locale
	mu.Unlock()

	defer func() {
		mu.Lock()
		defaultLocale = previousLocale
		mu.Unlock()
	}()

	callback()
}

// WithCurrency executes the given callback using the given currency
func WithCurrency(currency string, callback func()) {
	mu.Lock()
	previousCurrency := defaultCurrency
	defaultCurrency = currency
	mu.Unlock()

	defer func() {
		mu.Lock()
		defaultCurrency = previousCurrency
		mu.Unlock()
	}()

	callback()
}

// UseLocale sets the default locale
func UseLocale(locale string) {
	mu.Lock()
	defer mu.Unlock()
	defaultLocale = locale
}

// UseCurrency sets the default currency
func UseCurrency(currency string) {
	mu.Lock()
	defer mu.Unlock()
	defaultCurrency = currency
}

// DefaultLocale returns the default locale
func DefaultLocale() string {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLocale
}

// DefaultCurrency returns the default currency
func DefaultCurrency() string {
	mu.RLock()
	defer mu.RUnlock()
	return defaultCurrency
}
