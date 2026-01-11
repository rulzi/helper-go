package arr

import (
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Accessible determines whether the given value is array accessible.
func Accessible(value interface{}) bool {
	if value == nil {
		return false
	}
	kind := reflect.TypeOf(value).Kind()
	return kind == reflect.Slice || kind == reflect.Array || kind == reflect.Map
}

// Add adds an element to an array using "dot" notation if it doesn't exist.
func Add(array map[string]interface{}, key string, value interface{}) map[string]interface{} {
	if Get(array, key, nil) == nil {
		Set(array, key, value)
	}
	return array
}

// Collapse collapses an array of arrays into a single array.
func Collapse(arrays interface{}) []interface{} {
	val := reflect.ValueOf(arrays)
	if !val.IsValid() {
		return []interface{}{}
	}

	var estimatedSize int
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		estimatedSize = val.Len() * 2 // Estimate average 2 items per sub-array
	default:
		return []interface{}{}
	}

	result := make([]interface{}, 0, estimatedSize)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i).Interface()
			if item == nil {
				continue
			}

			itemVal := reflect.ValueOf(item)
			if itemVal.Kind() == reflect.Slice || itemVal.Kind() == reflect.Array {
				for j := 0; j < itemVal.Len(); j++ {
					result = append(result, itemVal.Index(j).Interface())
				}
			} else if itemVal.Kind() == reflect.Map {
				for _, key := range itemVal.MapKeys() {
					result = append(result, itemVal.MapIndex(key).Interface())
				}
			}
		}
	}

	return result
}

// CrossJoin cross joins the given arrays, returning all possible permutations.
func CrossJoin(arrays ...[]interface{}) [][]interface{} {
	if len(arrays) == 0 {
		return [][]interface{}{}
	}

	results := [][]interface{}{{}}

	for _, arr := range arrays {
		newResults := [][]interface{}{}
		for _, product := range results {
			for _, item := range arr {
				newProduct := make([]interface{}, len(product), len(product)+1)
				copy(newProduct, product)
				newProduct = append(newProduct, item)
				newResults = append(newResults, newProduct)
			}
		}
		results = newResults
	}

	return results
}

// Divide divides an array into two arrays. One with keys and the other with values.
func Divide(array map[string]interface{}) ([]string, []interface{}) {
	keys := make([]string, 0, len(array))
	values := make([]interface{}, 0, len(array))

	for k, v := range array {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// Dot flattens a multi-dimensional associative array with dots.
func Dot(array map[string]interface{}, prepend string) map[string]interface{} {
	results := make(map[string]interface{}, len(array))

	for key, value := range array {
		prefixedKey := prepend + key

		if valueMap, ok := value.(map[string]interface{}); ok && len(valueMap) > 0 {
			nested := Dot(valueMap, prefixedKey+".")
			for k, v := range nested {
				results[k] = v
			}
		} else {
			results[prefixedKey] = value
		}
	}

	return results
}

// Undot converts a flatten "dot" notation array into an expanded array.
func Undot(array map[string]interface{}) map[string]interface{} {
	results := make(map[string]interface{})

	for key, value := range array {
		Set(results, key, value)
	}

	return results
}

// Except gets all of the given array except for a specified array of keys.
func Except(array map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{}, len(array))
	for k, v := range array {
		result[k] = v
	}
	Forget(result, keys)
	return result
}

// Exists determines if the given key exists in the provided array.
func Exists(array map[string]interface{}, key string) bool {
	if array == nil {
		return false
	}
	_, exists := array[key]
	return exists
}

// First returns the first element in an array passing a given truth test.
func First(array interface{}, callback func(interface{}) bool, defaultValue interface{}) interface{} {
	val := reflect.ValueOf(array)
	if !val.IsValid() {
		return defaultValue
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return defaultValue
		}

		if callback == nil {
			return val.Index(0).Interface()
		}

		for i := 0; i < val.Len(); i++ {
			item := val.Index(i).Interface()
			if callback(item) {
				return item
			}
		}
	case reflect.Map:
		if val.Len() == 0 {
			return defaultValue
		}

		if callback == nil {
			for _, key := range val.MapKeys() {
				return val.MapIndex(key).Interface()
			}
		}

		for _, key := range val.MapKeys() {
			item := val.MapIndex(key).Interface()
			if callback(item) {
				return item
			}
		}
	}

	return defaultValue
}

// Last returns the last element in an array passing a given truth test.
func Last(array interface{}, callback func(interface{}) bool, defaultValue interface{}) interface{} {
	val := reflect.ValueOf(array)
	if !val.IsValid() {
		return defaultValue
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return defaultValue
		}

		if callback == nil {
			return val.Index(val.Len() - 1).Interface()
		}

		for i := val.Len() - 1; i >= 0; i-- {
			item := val.Index(i).Interface()
			if callback(item) {
				return item
			}
		}
	case reflect.Map:
		if val.Len() == 0 {
			return defaultValue
		}

		keys := val.MapKeys()
		if callback == nil {
			return val.MapIndex(keys[len(keys)-1]).Interface()
		}

		for i := len(keys) - 1; i >= 0; i-- {
			item := val.MapIndex(keys[i]).Interface()
			if callback(item) {
				return item
			}
		}
	}

	return defaultValue
}

// Take takes the first or last {$limit} items from an array.
func Take(array []interface{}, limit int) []interface{} {
	if limit < 0 {
		start := len(array) + limit
		if start < 0 {
			start = 0
		}
		return array[start:]
	}

	if limit > len(array) {
		limit = len(array)
	}

	return array[:limit]
}

// Flatten flattens a multi-dimensional array into a single level.
func Flatten(array interface{}, depth int) []interface{} {
	if depth == 0 {
		depth = -1 // unlimited depth
	}

	result := make([]interface{}, 0, 16) // Pre-allocate with initial capacity
	flattenRecursive(array, depth, &result)

	return result
}

func flattenRecursive(item interface{}, depth int, result *[]interface{}) {
	if depth == 0 {
		*result = append(*result, item)
		return
	}

	val := reflect.ValueOf(item)
	if !val.IsValid() {
		return
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			flattenRecursive(val.Index(i).Interface(), depth-1, result)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			flattenRecursive(val.MapIndex(key).Interface(), depth-1, result)
		}
	default:
		*result = append(*result, item)
	}
}

// Forget removes one or many array items from a given array using "dot" notation.
func Forget(array map[string]interface{}, keys []string) {
	if array == nil || len(keys) == 0 {
		return
	}

	for _, key := range keys {
		if dotIdx := strings.IndexByte(key, '.'); dotIdx != -1 {
			parts := strings.Split(key, ".")
			current := array

			for i := 0; i < len(parts)-1; i++ {
				if val, ok := current[parts[i]].(map[string]interface{}); ok {
					current = val
				} else {
					goto nextKey
				}
			}

			delete(current, parts[len(parts)-1])
		} else {
			delete(array, key)
		}
	nextKey:
	}
}

// Get gets an item from an array using "dot" notation.
func Get(array map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if array == nil {
		return defaultValue
	}

	if key == "" {
		return array
	}

	if strings.IndexByte(key, '.') == -1 {
		if val, exists := array[key]; exists {
			return val
		}
		return defaultValue
	}

	parts := strings.Split(key, ".")
	current := array

	for i, part := range parts {
		if val, ok := current[part].(map[string]interface{}); ok {
			if i == len(parts)-1 {
				return val
			}
			current = val
		} else {
			if i == len(parts)-1 {
				if val, exists := current[part]; exists {
					return val
				}
			}
			return defaultValue
		}
	}

	return defaultValue
}

// Has checks if an item or items exist in an array using "dot" notation.
func Has(array map[string]interface{}, keys []string) bool {
	if array == nil || len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		if !HasOne(array, key) {
			return false
		}
	}

	return true
}

// HasOne checks if a single key exists in an array using "dot" notation.
func HasOne(array map[string]interface{}, key string) bool {
	if array == nil || key == "" {
		return false
	}

	if strings.IndexByte(key, '.') == -1 {
		_, exists := array[key]
		return exists
	}

	parts := strings.Split(key, ".")
	current := array

	for i, part := range parts {
		if val, ok := current[part].(map[string]interface{}); ok {
			if i == len(parts)-1 {
				return true
			}
			current = val
		} else {
			if i == len(parts)-1 {
				_, exists := current[part]
				return exists
			}
			return false
		}
	}

	return false
}

// HasAny determines if any of the keys exist in an array using "dot" notation.
func HasAny(array map[string]interface{}, keys []string) bool {
	if array == nil || len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		if HasOne(array, key) {
			return true
		}
	}

	return false
}

// IsAssoc determines if an array is associative.
func IsAssoc(array interface{}) bool {
	val := reflect.ValueOf(array)
	if !val.IsValid() {
		return false
	}

	if val.Kind() == reflect.Map {
		return true
	}

	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return false
	}

	// In Go, slices/arrays are always indexed, so we check if it's a map[string]interface{}
	// For slices, we consider them non-associative
	return false
}

// IsList determines if an array is a list.
func IsList(array interface{}) bool {
	val := reflect.ValueOf(array)
	if !val.IsValid() {
		return false
	}

	return val.Kind() == reflect.Slice || val.Kind() == reflect.Array
}

// Join joins all items using a string. The final items can use a separate glue string.
func Join(array []string, glue, finalGlue string) string {
	if len(array) == 0 {
		return ""
	}

	if finalGlue == "" {
		return strings.Join(array, glue)
	}

	if len(array) == 1 {
		return array[0]
	}

	return strings.Join(array[:len(array)-1], glue) + finalGlue + array[len(array)-1]
}

// KeyBy keys an associative array by a field or using a callback.
func KeyBy(array []interface{}, keyBy interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(array))

	switch k := keyBy.(type) {
	case string:
		for _, item := range array {
			if itemMap, ok := item.(map[string]interface{}); ok {
				if key, exists := itemMap[k]; exists {
					keyStr := toString(key)
					result[keyStr] = item
				}
			}
		}
	case func(interface{}) string:
		for _, item := range array {
			key := k(item)
			result[key] = item
		}
	}

	return result
}

// PrependKeysWith prepends the key names of an associative array.
func PrependKeysWith(array map[string]interface{}, prependWith string) map[string]interface{} {
	result := make(map[string]interface{}, len(array))
	for k, v := range array {
		result[prependWith+k] = v
	}
	return result
}

// Only gets a subset of the items from the given array.
func Only(array map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		if val, exists := array[key]; exists {
			result[key] = val
		}
	}
	return result
}

// Select selects an array of values from an array.
func Select(array []interface{}, keys []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(array))

	for _, item := range array {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		selected := make(map[string]interface{})
		for _, key := range keys {
			if val, exists := itemMap[key]; exists {
				selected[key] = val
			}
		}
		result = append(result, selected)
	}

	return result
}

// Pluck plucks an array of values from an array.
func Pluck(array []interface{}, value string, key string) map[string]interface{} {
	result := make(map[string]interface{}, len(array))

	for _, item := range array {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		itemValue := Get(itemMap, value, nil)

		if key == "" {
			// If no key specified, we can't use map, return as slice would require different signature
			// For now, we'll use index as key
			result[toString(len(result))] = itemValue
		} else {
			itemKey := Get(itemMap, key, nil)
			keyStr := toString(itemKey)
			result[keyStr] = itemValue
		}
	}

	return result
}

// Map runs a map over each of the items in the array.
func Map(array []interface{}, callback func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(array))
	for i, item := range array {
		result[i] = callback(item)
	}
	return result
}

// MapWithKeys runs an associative map over each of the items.
func MapWithKeys(array []interface{}, callback func(interface{}) map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(array))
	for _, item := range array {
		assoc := callback(item)
		for k, v := range assoc {
			result[k] = v
		}
	}
	return result
}

// MapSpread runs a map over each nested chunk of items.
func MapSpread(array [][]interface{}, callback func(...interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(array))
	for i, chunk := range array {
		result[i] = callback(chunk...)
	}
	return result
}

// Prepend pushes an item onto the beginning of an array.
func Prepend(array []interface{}, value interface{}, key ...string) []interface{} {
	// For map-like behavior, we'd need a different return type
	// This implementation treats it as slice prepend
	return append([]interface{}{value}, array...)
}

// Pull gets a value from the array, and removes it.
func Pull(array map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value := Get(array, key, defaultValue)
	Forget(array, []string{key})
	return value
}

// Query converts the array into a query string.
func Query(array map[string]interface{}) string {
	values := url.Values{}
	for k, v := range array {
		values.Set(k, toString(v))
	}
	return values.Encode()
}

// Random gets one or a specified number of random values from an array.
func Random(array []interface{}, number int, preserveKeys bool) (interface{}, error) {
	if len(array) == 0 {
		if number > 0 {
			return []interface{}{}, nil
		}
		return nil, nil
	}

	if number <= 0 {
		number = 1
	}

	if number > len(array) {
		return nil, fmt.Errorf("you requested %d items, but there are only %d items available", number, len(array))
	}

	indices := rand.Perm(len(array))[:number]

	if number == 1 {
		return array[indices[0]], nil
	}

	if preserveKeys {
		result := make(map[string]interface{})
		for _, idx := range indices {
			key := toString(idx)
			result[key] = array[idx]
		}
		return result, nil
	}

	result := make([]interface{}, number)
	for i, idx := range indices {
		result[i] = array[idx]
	}
	return result, nil
}

// Set sets an array item to a given value using "dot" notation.
func Set(array map[string]interface{}, key string, value interface{}) map[string]interface{} {
	if key == "" {
		return array
	}

	if strings.IndexByte(key, '.') == -1 {
		array[key] = value
		return array
	}

	parts := strings.Split(key, ".")
	current := array

	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if val, ok := current[part].(map[string]interface{}); ok {
			current = val
		} else {
			current[part] = make(map[string]interface{})
			current = current[part].(map[string]interface{})
		}
	}

	current[parts[len(parts)-1]] = value
	return array
}

// Shuffle shuffles the given array and returns the result.
func Shuffle(array []interface{}) []interface{} {
	result := make([]interface{}, len(array))
	copy(result, array)

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

// Sort sorts the array using the given callback.
func Sort(array []interface{}, callback func(interface{}, interface{}) bool) []interface{} {
	result := make([]interface{}, len(array))
	copy(result, array)

	if callback != nil {
		sort.Slice(result, func(i, j int) bool {
			return callback(result[i], result[j])
		})
	} else {
		// Default string comparison
		sort.Slice(result, func(i, j int) bool {
			return toString(result[i]) < toString(result[j])
		})
	}

	return result
}

// SortDesc sorts the array in descending order using the given callback.
func SortDesc(array []interface{}, callback func(interface{}, interface{}) bool) []interface{} {
	result := make([]interface{}, len(array))
	copy(result, array)

	if callback != nil {
		sort.Slice(result, func(i, j int) bool {
			return !callback(result[i], result[j])
		})
	} else {
		sort.Slice(result, func(i, j int) bool {
			return toString(result[i]) > toString(result[j])
		})
	}

	return result
}

// SortRecursive recursively sorts an array by keys and values.
func SortRecursive(array map[string]interface{}, descending bool) map[string]interface{} {
	result := make(map[string]interface{})

	// Sort keys
	keys := make([]string, 0, len(array))
	for k := range array {
		keys = append(keys, k)
	}

	if descending {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}

	// Copy values and recursively sort nested maps
	for _, k := range keys {
		val := array[k]
		if nestedMap, ok := val.(map[string]interface{}); ok {
			result[k] = SortRecursive(nestedMap, descending)
		} else {
			result[k] = val
		}
	}

	return result
}

// SortRecursiveDesc recursively sorts an array by keys and values in descending order.
func SortRecursiveDesc(array map[string]interface{}) map[string]interface{} {
	return SortRecursive(array, true)
}

// shouldInclude checks if a constraint value should be included.
func shouldInclude(constraint interface{}) bool {
	if constraint == nil {
		return false
	}

	constraintVal := reflect.ValueOf(constraint)
	switch constraintVal.Kind() {
	case reflect.Bool:
		return constraintVal.Bool()
	case reflect.String:
		return constraintVal.String() != ""
	default:
		return !constraintVal.IsZero()
	}
}

// ToCssClasses conditionally compiles classes from an array into a CSS class list.
func ToCssClasses(array map[string]interface{}) string {
	classes := make([]string, 0, len(array))

	for class, constraint := range array {
		if shouldInclude(constraint) {
			classes = append(classes, class)
		}
	}

	return strings.Join(classes, " ")
}

// ToCssStyles conditionally compiles styles from an array into a style list.
func ToCssStyles(array map[string]interface{}) string {
	styles := make([]string, 0, len(array))

	for style, constraint := range array {
		if shouldInclude(constraint) {
			styleStr := toString(style)
			if !strings.HasSuffix(styleStr, ";") {
				styleStr += ";"
			}
			styles = append(styles, styleStr)
		}
	}

	return strings.Join(styles, " ")
}

// Where filters the array using the given callback.
func Where(array []interface{}, callback func(interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(array)/2) // Pre-allocate with estimated capacity
	for _, item := range array {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

// Reject filters the array using the negation of the given callback.
func Reject(array []interface{}, callback func(interface{}) bool) []interface{} {
	return Where(array, func(item interface{}) bool {
		return !callback(item)
	})
}

// WhereNotNull filters items where the value is not null.
func WhereNotNull(array []interface{}) []interface{} {
	return Where(array, func(item interface{}) bool {
		return item != nil
	})
}

// Wrap if the given value is not an array and not null, wrap it in one.
func Wrap(value interface{}) []interface{} {
	if value == nil {
		return []interface{}{}
	}

	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		result := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			result[i] = val.Index(i).Interface()
		}
		return result
	}

	return []interface{}{value}
}

// Helper function to convert interface{} to string
func toString(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%v", v)
	}
}
