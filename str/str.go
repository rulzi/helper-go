package str

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	snakeCache  = make(map[string]map[string]string)
	camelCache  = make(map[string]string)
	studlyCache = make(map[string]string)
	// Pre-compiled regex patterns for common use cases
	numbersRegex = regexp.MustCompile(`[^0-9]`)
	squishRegex  = regexp.MustCompile(`\s+`)
	uuidRegex    = regexp.MustCompile(`^[\da-fA-F]{8}-[\da-fA-F]{4}-[\da-fA-F]{4}-[\da-fA-F]{4}-[\da-fA-F]{12}$`)
)

// After returns the remainder of a string after the first occurrence of a given value.
func After(subject, search string) string {
	if search == "" {
		return subject
	}
	parts := strings.SplitN(subject, search, 2)
	if len(parts) < 2 {
		return subject
	}
	return parts[1]
}

// AfterLast returns the remainder of a string after the last occurrence of a given value.
func AfterLast(subject, search string) string {
	if search == "" {
		return subject
	}
	pos := strings.LastIndex(subject, search)
	if pos == -1 {
		return subject
	}
	return subject[pos+len(search):]
}

// Before returns the portion of a string before the first occurrence of a given value.
func Before(subject, search string) string {
	if search == "" {
		return subject
	}
	pos := strings.Index(subject, search)
	if pos == -1 {
		return subject
	}
	return subject[:pos]
}

// BeforeLast returns the portion of a string before the last occurrence of a given value.
func BeforeLast(subject, search string) string {
	if search == "" {
		return subject
	}
	pos := strings.LastIndex(subject, search)
	if pos == -1 {
		return subject
	}
	return subject[:pos]
}

// Between returns the portion of a string between two given values.
func Between(subject, from, to string) string {
	if from == "" || to == "" {
		return subject
	}
	return BeforeLast(After(subject, from), to)
}

// BetweenFirst returns the smallest possible portion of a string between two given values.
func BetweenFirst(subject, from, to string) string {
	if from == "" || to == "" {
		return subject
	}
	return Before(After(subject, from), to)
}

// Camel converts a value to camel case.
func Camel(value string) string {
	if cached, ok := camelCache[value]; ok {
		return cached
	}
	result := Lcfirst(Studly(value))
	camelCache[value] = result
	return result
}

// CharAt returns the character at the specified index.
func CharAt(subject string, index int) (string, bool) {
	if subject == "" {
		return "", false
	}

	// Fast path for ASCII strings
	if index >= 0 && index < len(subject) {
		// Check if it's a single byte character
		if subject[index] < 128 {
			return subject[index : index+1], true
		}
	}

	// Handle unicode and negative indices
	runes := []rune(subject)
	length := len(runes)

	if index < 0 {
		index = length + index
	}

	if index < 0 || index >= length {
		return "", false
	}

	return string(runes[index]), true
}

// ChopStart removes the given string(s) if it exists at the start of the haystack.
func ChopStart(subject string, needles ...string) string {
	for _, needle := range needles {
		if strings.HasPrefix(subject, needle) {
			return subject[len(needle):]
		}
	}
	return subject
}

// ChopEnd removes the given string(s) if it exists at the end of the haystack.
func ChopEnd(subject string, needles ...string) string {
	for _, needle := range needles {
		if strings.HasSuffix(subject, needle) {
			return subject[:len(subject)-len(needle)]
		}
	}
	return subject
}

// Contains determines if a given string contains a given substring.
func Contains(haystack string, needles []string, ignoreCase bool) bool {
	if ignoreCase {
		haystack = strings.ToLower(haystack)
	}

	for _, needle := range needles {
		if needle == "" {
			continue
		}
		search := needle
		if ignoreCase {
			search = strings.ToLower(needle)
		}
		if strings.Contains(haystack, search) {
			return true
		}
	}
	return false
}

// ContainsAll determines if a given string contains all array values.
func ContainsAll(haystack string, needles []string, ignoreCase bool) bool {
	if len(needles) == 0 {
		return true
	}

	var haystackLower string
	if ignoreCase {
		haystackLower = strings.ToLower(haystack)
	}

	for _, needle := range needles {
		if needle == "" {
			continue
		}
		search := needle
		if ignoreCase {
			search = strings.ToLower(needle)
			if !strings.Contains(haystackLower, search) {
				return false
			}
		} else {
			if !strings.Contains(haystack, search) {
				return false
			}
		}
	}
	return true
}

// DoesntContain determines if a given string doesn't contain a given substring.
func DoesntContain(haystack string, needles []string, ignoreCase bool) bool {
	return !Contains(haystack, needles, ignoreCase)
}

// Deduplicate replaces consecutive instances of a given character with a single character.
func Deduplicate(s, character string) string {
	if character == "" || len(s) == 0 {
		return s
	}
	// Use strings.Replace for single character (faster than regex)
	if len(character) == 1 {
		char := character[0]
		var result strings.Builder
		result.Grow(len(s))
		prev := byte(0)
		for i := 0; i < len(s); i++ {
			if s[i] == char {
				if prev != char {
					result.WriteByte(char)
				}
			} else {
				result.WriteByte(s[i])
			}
			prev = s[i]
		}
		return result.String()
	}
	// Fallback to regex for multi-character
	re := regexp.MustCompile(regexp.QuoteMeta(character) + "+")
	return re.ReplaceAllString(s, character)
}

// EndsWith determines if a given string ends with a given substring.
func EndsWith(haystack string, needles []string) bool {
	if haystack == "" {
		return false
	}
	for _, needle := range needles {
		if needle != "" && strings.HasSuffix(haystack, needle) {
			return true
		}
	}
	return false
}

// Finish caps a string with a single instance of a given value.
func Finish(value, cap string) string {
	quoted := regexp.QuoteMeta(cap)
	re := regexp.MustCompile("(?:^" + quoted + ")+")
	value = re.ReplaceAllString(value, "")
	return value + cap
}

// Wrap wraps the string with the given strings.
func Wrap(value, before string, after ...string) string {
	afterStr := before
	if len(after) > 0 {
		afterStr = after[0]
	}
	return before + value + afterStr
}

// Unwrap unwraps the string with the given strings.
func Unwrap(value, before string, after ...string) string {
	afterStr := before
	if len(after) > 0 {
		afterStr = after[0]
	}

	value = strings.TrimPrefix(value, before)
	value = strings.TrimSuffix(value, afterStr)

	return value
}

// IsAscii determines if a given string is 7 bit ASCII.
func IsAscii(value string) bool {
	for _, r := range value {
		if r > 127 {
			return false
		}
	}
	return true
}

// IsJson determines if a given value is valid JSON.
func IsJson(value string) bool {
	var js interface{}
	return json.Unmarshal([]byte(value), &js) == nil
}

// IsUrl determines if a given value is a valid URL.
func IsUrl(value string, protocols []string) bool {
	u, err := url.Parse(value)
	if err != nil {
		return false
	}

	if len(protocols) > 0 {
		found := false
		for _, protocol := range protocols {
			if u.Scheme == protocol {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return u.Scheme != "" && u.Host != ""
}

// IsUuid determines if a given value is a valid UUID.
func IsUuid(value string) bool {
	return uuidRegex.MatchString(value)
}

// Kebab converts a string to kebab case.
func Kebab(value string) string {
	return Snake(value, "-")
}

// Length returns the length of the given string.
func Length(value string) int {
	return utf8.RuneCountInString(value)
}

// Limit limits the number of characters in a string.
func Limit(value string, limit int, end string, preserveWords bool) string {
	if limit <= 0 {
		return value
	}

	runes := []rune(value)
	if len(runes) <= limit {
		return value
	}

	if !preserveWords {
		return string(runes[:limit]) + end
	}

	trimmed := string(runes[:limit])
	if len(runes) > limit && runes[limit] == ' ' {
		return trimmed + end
	}

	lastSpace := strings.LastIndex(trimmed, " ")
	if lastSpace > 0 {
		return trimmed[:lastSpace] + end
	}

	return trimmed + end
}

// Lower converts the given string to lower-case.
func Lower(value string) string {
	return strings.ToLower(value)
}

// Words limits the number of words in a string.
func Words(value string, words int, end string) string {
	parts := strings.Fields(value)
	if len(parts) <= words {
		return value
	}
	return strings.Join(parts[:words], " ") + end
}

// Numbers removes all non-numeric characters from a string.
func Numbers(value string) string {
	return numbersRegex.ReplaceAllString(value, "")
}

// PadBoth pads both sides of a string with another.
func PadBoth(value string, length int, pad string) string {
	if len(value) >= length || pad == "" {
		return value
	}

	short := length - len(value)
	shortLeft := short / 2
	shortRight := short - shortLeft

	return strings.Repeat(pad, shortLeft) + value + strings.Repeat(pad, shortRight)
}

// PadLeft pads the left side of a string with another.
func PadLeft(value string, length int, pad string) string {
	if len(value) >= length || pad == "" {
		return value
	}
	short := length - len(value)
	return strings.Repeat(pad, short) + value
}

// PadRight pads the right side of a string with another.
func PadRight(value string, length int, pad string) string {
	if len(value) >= length || pad == "" {
		return value
	}
	short := length - len(value)
	return value + strings.Repeat(pad, short)
}

// ParseCallback parses a Class@method style callback into class and method.
func ParseCallback(callback, defaultMethod string) (string, string) {
	if strings.Contains(callback, "@") {
		parts := strings.SplitN(callback, "@", 2)
		return parts[0], parts[1]
	}
	return callback, defaultMethod
}

// Position finds the position of the first occurrence of a given substring in a string.
func Position(haystack, needle string, offset int) int {
	if offset < 0 {
		offset = 0
	}
	if offset >= len(haystack) || needle == "" {
		return -1
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return offset + pos
}

// Random generates a more truly "random" alpha-numeric string.
func Random(length int) string {
	if length <= 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(length)

	for result.Len() < length {
		size := length - result.Len()
		bytesSize := (size/3 + 1) * 3

		bytes := make([]byte, bytesSize)
		if _, err := rand.Read(bytes); err != nil {
			panic("failed to generate random bytes: " + err.Error())
		}

		encoded := base64.URLEncoding.EncodeToString(bytes)
		// Remove URL-unsafe characters more efficiently using strings.Replacer
		encoded = strings.ReplaceAll(encoded, "/", "")
		encoded = strings.ReplaceAll(encoded, "+", "")
		encoded = strings.ReplaceAll(encoded, "=", "")

		if len(encoded) > size {
			encoded = encoded[:size]
		}

		result.WriteString(encoded)
	}

	return result.String()[:length]
}

// Repeat repeats the given string.
func Repeat(s string, times int) string {
	return strings.Repeat(s, times)
}

// Replace replaces the given value in the given string.
func Replace(search, replace, subject string, caseSensitive bool) string {
	if search == "" {
		return subject
	}
	if caseSensitive {
		return strings.ReplaceAll(subject, search, replace)
	}
	// Case-insensitive: convert to lowercase for matching
	// Note: This changes the case of the result, which matches original behavior
	return strings.ReplaceAll(strings.ToLower(subject), strings.ToLower(search), replace)
}

// ReplaceFirst replaces the first occurrence of a given value in the string.
func ReplaceFirst(search, replace, subject string) string {
	if search == "" {
		return subject
	}
	pos := strings.Index(subject, search)
	if pos == -1 {
		return subject
	}
	return subject[:pos] + replace + subject[pos+len(search):]
}

// ReplaceStart replaces the first occurrence of the given value if it appears at the start of the string.
func ReplaceStart(search, replace, subject string) string {
	if search == "" {
		return subject
	}
	if strings.HasPrefix(subject, search) {
		return ReplaceFirst(search, replace, subject)
	}
	return subject
}

// ReplaceLast replaces the last occurrence of a given value in the string.
func ReplaceLast(search, replace, subject string) string {
	if search == "" {
		return subject
	}
	pos := strings.LastIndex(subject, search)
	if pos == -1 {
		return subject
	}
	return subject[:pos] + replace + subject[pos+len(search):]
}

// ReplaceEnd replaces the last occurrence of a given value if it appears at the end of the string.
func ReplaceEnd(search, replace, subject string) string {
	if search == "" {
		return subject
	}
	if strings.HasSuffix(subject, search) {
		return ReplaceLast(search, replace, subject)
	}
	return subject
}

// Remove removes any occurrence of the given string in the subject.
func Remove(search []string, subject string, caseSensitive bool) string {
	if len(search) == 0 {
		return subject
	}
	if caseSensitive {
		for _, s := range search {
			if s != "" {
				subject = strings.ReplaceAll(subject, s, "")
			}
		}
		return subject
	}
	// Case-insensitive: convert to lowercase for matching
	// Note: This changes the case of the result, which matches original behavior
	result := strings.ToLower(subject)
	for _, s := range search {
		if s != "" {
			result = strings.ReplaceAll(result, strings.ToLower(s), "")
		}
	}
	return result
}

// Reverse reverses the given string.
func Reverse(value string) string {
	runes := []rune(value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Start begins a string with a single instance of a given value.
func Start(value, prefix string) string {
	quoted := regexp.QuoteMeta(prefix)
	re := regexp.MustCompile("^(?:" + quoted + ")+")
	value = re.ReplaceAllString(value, "")
	return prefix + value
}

// Upper converts the given string to upper-case.
func Upper(value string) string {
	return strings.ToUpper(value)
}

// Title converts the given string to proper case.
func Title(value string) string {
	// Use cases.Title which properly handles Unicode
	t := cases.Title(language.Und, cases.NoLower)
	return t.String(strings.ToLower(value))
}

// Snake converts a string to snake case.
func Snake(value, delimiter string) string {
	if cache, ok := snakeCache[value]; ok {
		if cached, ok := cache[delimiter]; ok {
			return cached
		}
	}

	// Check if already lowercase
	allLower := true
	for _, r := range value {
		if unicode.IsUpper(r) {
			allLower = false
			break
		}
	}

	if allLower {
		if snakeCache[value] == nil {
			snakeCache[value] = make(map[string]string)
		}
		snakeCache[value][delimiter] = value
		return value
	}

	// Convert to snake case
	var result strings.Builder
	runes := []rune(value)

	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteString(delimiter)
			}
			result.WriteRune(unicode.ToLower(r))
		} else if r == ' ' || r == '-' || r == '_' {
			result.WriteString(delimiter)
		} else {
			result.WriteRune(r)
		}
	}

	resultStr := result.String()
	if snakeCache[value] == nil {
		snakeCache[value] = make(map[string]string)
	}
	snakeCache[value][delimiter] = resultStr
	return resultStr
}

// Studly converts a value to studly caps case.
func Studly(value string) string {
	// Use original value as cache key
	if cached, ok := studlyCache[value]; ok {
		return cached
	}

	originalValue := value
	// Replace hyphens and underscores with spaces
	value = strings.ReplaceAll(value, "-", " ")
	value = strings.ReplaceAll(value, "_", " ")

	words := strings.Fields(value)
	var result strings.Builder
	result.Grow(len(originalValue))

	for _, word := range words {
		if len(word) > 0 {
			result.WriteString(Ucfirst(word))
		}
	}

	resultStr := result.String()
	studlyCache[originalValue] = resultStr
	return resultStr
}

// Pascal converts a value to Pascal case.
func Pascal(value string) string {
	return Studly(value)
}

// Substr returns the portion of the string specified by the start and length parameters.
func Substr(s string, start int, length ...int) string {
	runes := []rune(s)
	runesLen := len(runes)

	if start < 0 {
		start = runesLen + start
	}
	if start < 0 {
		start = 0
	}
	if start >= runesLen {
		return ""
	}

	end := runesLen
	if len(length) > 0 && length[0] > 0 {
		end = start + length[0]
		if end > runesLen {
			end = runesLen
		}
	}

	return string(runes[start:end])
}

// SubstrCount returns the number of substring occurrences.
func SubstrCount(haystack, needle string, offset, length int) int {
	if needle == "" {
		return 0
	}

	if offset < 0 {
		offset = 0
	}

	substr := haystack
	if offset > 0 || length > 0 {
		runes := []rune(haystack)
		runesLen := len(runes)

		if offset >= runesLen {
			return 0
		}

		end := runesLen
		if length > 0 && offset+length < runesLen {
			end = offset + length
		}

		substr = string(runes[offset:end])
	}

	return strings.Count(substr, needle)
}

// Swap swaps multiple keywords in a string with other keywords simultaneously.
func Swap(m map[string]string, subject string) string {
	// Handle simultaneous replacement by using a regex approach
	// Collect all keys and sort by length (longest first) to avoid partial matches
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort keys by length descending, then by key name ascending for deterministic order
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			// First compare by length (longer keys first)
			if len(keys[i]) < len(keys[j]) {
				keys[i], keys[j] = keys[j], keys[i]
			} else if len(keys[i]) == len(keys[j]) {
				// For same length keys, sort by key name (alphabetically)
				if keys[i] > keys[j] {
					keys[i], keys[j] = keys[j], keys[i]
				}
			}
		}
	}

	result := subject
	// Apply replacements in order (longest first) to avoid conflicts
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result = strings.ReplaceAll(result, key, value)
		}
	}

	return result
}

// Take takes the first or last {$limit} characters of a string.
func Take(s string, limit int) string {
	if limit < 0 {
		return Substr(s, limit)
	}
	return Substr(s, 0, limit)
}

// ToBase64 converts the given string to Base64 encoding.
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// FromBase64 decodes the given Base64 encoded string.
func FromBase64(s string, strict bool) (string, error) {
	var encoding *base64.Encoding
	if strict {
		encoding = base64.StdEncoding
	} else {
		encoding = base64.RawStdEncoding
	}

	data, err := encoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Lcfirst makes a string's first character lowercase.
func Lcfirst(s string) string {
	if s == "" {
		return s
	}
	// Fast path for ASCII
	if s[0] < 128 {
		if s[0] >= 'A' && s[0] <= 'Z' {
			return string(s[0]+32) + s[1:]
		}
		return s
	}
	// Unicode path
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// Ucfirst makes a string's first character uppercase.
func Ucfirst(s string) string {
	if s == "" {
		return s
	}
	// Fast path for ASCII
	if s[0] < 128 {
		if s[0] >= 'a' && s[0] <= 'z' {
			return string(s[0]-32) + s[1:]
		}
		return s
	}
	// Unicode path
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Ucsplit splits a string into pieces by uppercase characters.
func Ucsplit(s string) []string {
	var result []string
	var current strings.Builder

	for _, r := range s {
		if unicode.IsUpper(r) && current.Len() > 0 {
			result = append(result, current.String())
			current.Reset()
		}
		current.WriteRune(r)
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

// WordCount returns the number of words a string contains.
func WordCount(s string) int {
	return len(strings.Fields(s))
}

// Trim removes all whitespace from both ends of a string.
func Trim(value string, charlist ...string) string {
	if len(charlist) > 0 && charlist[0] != "" {
		return strings.Trim(value, charlist[0])
	}
	return strings.TrimSpace(value)
}

// Ltrim removes all whitespace from the beginning of a string.
func Ltrim(value string, charlist ...string) string {
	if len(charlist) > 0 && charlist[0] != "" {
		return strings.TrimLeft(value, charlist[0])
	}
	return strings.TrimLeft(value, " \t\n\r\v\f")
}

// Rtrim removes all whitespace from the end of a string.
func Rtrim(value string, charlist ...string) string {
	if len(charlist) > 0 && charlist[0] != "" {
		return strings.TrimRight(value, charlist[0])
	}
	return strings.TrimRight(value, " \t\n\r\v\f")
}

// Squish removes all "extra" blank space from the given string.
func Squish(value string) string {
	value = strings.TrimSpace(value)
	return squishRegex.ReplaceAllString(value, " ")
}

// StartsWith determines if a given string starts with a given substring.
func StartsWith(haystack string, needles []string) bool {
	if haystack == "" {
		return false
	}
	for _, needle := range needles {
		if needle != "" && strings.HasPrefix(haystack, needle) {
			return true
		}
	}
	return false
}

// Match gets the string matching the given pattern.
func Match(pattern, subject string) string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}

	matches := re.FindStringSubmatch(subject)
	if len(matches) == 0 {
		return ""
	}

	if len(matches) > 1 {
		return matches[1]
	}
	return matches[0]
}

// IsMatch determines if a given string matches a given pattern.
func IsMatch(pattern, value string) bool {
	matched, err := regexp.MatchString(pattern, value)
	return err == nil && matched
}

// MatchAll gets all strings matching the given pattern.
func MatchAll(pattern, subject string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(subject, -1)
	if len(matches) == 0 {
		return []string{}, nil
	}

	var result []string
	for _, match := range matches {
		if len(match) > 1 {
			result = append(result, match[1])
		} else {
			result = append(result, match[0])
		}
	}

	return result, nil
}

// ReplaceArray replaces a given value in the string sequentially with an array.
func ReplaceArray(search string, replace []string, subject string) string {
	if search == "" {
		return subject
	}
	parts := strings.Split(subject, search)
	if len(parts) == 1 {
		return subject
	}

	// Pre-calculate capacity for better performance
	capacity := len(subject)
	for i := 0; i < len(parts)-1 && i < len(replace); i++ {
		capacity += len(replace[i]) - len(search)
	}

	var result strings.Builder
	result.Grow(capacity)
	result.WriteString(parts[0])

	for i := 1; i < len(parts); i++ {
		replacement := search
		if i-1 < len(replace) {
			replacement = replace[i-1]
		}
		result.WriteString(replacement)
		result.WriteString(parts[i])
	}

	return result.String()
}

// ReplaceMatches replaces the patterns matching the given regular expression.
func ReplaceMatches(pattern string, replace interface{}, subject string, limit int) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	switch r := replace.(type) {
	case string:
		if limit > 0 {
			count := 0
			result := re.ReplaceAllStringFunc(subject, func(match string) string {
				if count >= limit {
					return match
				}
				count++
				return r
			})
			return result, nil
		}
		return re.ReplaceAllString(subject, r), nil
	case func(string) string:
		if limit > 0 {
			count := 0
			result := re.ReplaceAllStringFunc(subject, func(match string) string {
				if count >= limit {
					return match
				}
				count++
				return r(match)
			})
			return result, nil
		}
		return re.ReplaceAllStringFunc(subject, r), nil
	default:
		return "", fmt.Errorf("unsupported replace type")
	}
}

// Slug generates a URL friendly "slug" from a given string.
func Slug(title, separator, language string, dictionary map[string]string) string {
	// Convert to lowercase
	title = strings.ToLower(title)

	// Replace dictionary words
	for k, v := range dictionary {
		title = strings.ReplaceAll(title, k, separator+v+separator)
	}

	// Remove all characters that are not the separator, letters, numbers, or whitespace
	re := regexp.MustCompile(`[^` + regexp.QuoteMeta(separator) + `a-z0-9\s]+`)
	title = re.ReplaceAllString(title, "")

	// Replace all separator characters and whitespace by a single separator
	re = regexp.MustCompile(`[` + regexp.QuoteMeta(separator) + `\s]+`)
	title = re.ReplaceAllString(title, separator)

	// Trim separators
	title = strings.Trim(title, separator)

	return title
}

// FlushCache removes all strings from the casing caches.
func FlushCache() {
	snakeCache = make(map[string]map[string]string)
	camelCache = make(map[string]string)
	studlyCache = make(map[string]string)
}
