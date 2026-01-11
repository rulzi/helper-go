# helper-go

[![Go Reference](https://pkg.go.dev/badge/github.com/rulzi/helper-go.svg)](https://pkg.go.dev/github.com/rulzi/helper-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/rulzi/helper-go)](https://goreportcard.com/report/github.com/rulzi/helper-go)
[![Tests](https://github.com/rulzi/helper-go/actions/workflows/tests.yml/badge.svg)](https://github.com/rulzi/helper-go/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/rulzi/helper-go/branch/main/graph/badge.svg)](https://codecov.io/gh/rulzi/helper-go)
[![Release](https://img.shields.io/github/release/rulzi/helper-go.svg)](https://github.com/rulzi/helper-go/releases/latest)

Kumpulan helper functions untuk Golang yang reusable dan production-ready.

**helper-go** adalah porting helper Laravel ke Golang, dibuat untuk kebutuhan author dan community.

## Fitur

- ✅ **Array Helpers** - Operasi array/map dengan dot notation, filtering, mapping, sorting
- ✅ **String Helpers** - Manipulasi string, case conversion, validation, encoding
- ✅ **Number Helpers** - Formatting, currency, percentage, file size, human-readable
- ✅ **Idiomatic Go** - Mengikuti best practices dan standar Go
- ✅ **Zero Dependencies** - Tidak menggunakan dependency eksternal yang berat
- ✅ **Well Tested** - Unit test untuk setiap public function
- ✅ **Production Ready** - Siap digunakan di production

## Instalasi

```bash
go get github.com/rulzi/helper-go
```

## Struktur Project

```
helper-go/
├── arr/           # Array helpers
├── number/        # Number helpers
├── str/           # String helpers
├── examples/      # Contoh penggunaan
├── go.mod
└── README.md
```

## Quick Start

### Array Helpers

```go
import "github.com/rulzi/helper-go/arr"

// Get value dari nested map menggunakan dot notation
data := map[string]interface{}{
    "user": map[string]interface{}{"name": "John"},
}
name := arr.Get(data, "user.name", "Unknown")

// Filter dan map array
numbers := []interface{}{1, 2, 3, 4, 5}
evens := arr.Where(numbers, func(n interface{}) bool {
    return n.(int)%2 == 0
})
doubled := arr.Map(numbers, func(n interface{}) interface{} {
    return n.(int) * 2
})
```

### String Helpers

```go
import "github.com/rulzi/helper-go/str"

// Case conversion
camel := str.Camel("hello world")      // helloWorld
snake := str.Snake("HelloWorld", "_")  // hello_world
kebab := str.Kebab("Hello World")      // hello-world

// Validation
isUrl := str.IsUrl("https://example.com", []string{"http", "https"})
isUuid := str.IsUuid("550e8400-e29b-41d4-a716-446655440000")
```

### Number Helpers

```go
import "github.com/rulzi/helper-go/number"

// Formatting
formatted := number.Format(1234.567, nil, nil, nil)
currency := number.Currency(1234.56, "USD", nil, nil)  // $1234.56
fileSize := number.FileSize(1024*1024*5, 2, nil)      // 5.00 MB
human := number.ForHumans(1500, 1, nil, true)          // 1.5K
ordinal := number.Ordinal(1, nil)                      // 1st
```

## Package Overview

### `arr` - Array Helpers
- Access & Manipulation: `Get`, `Set`, `Has`, `HasOne`, `HasAny`, `Exists`
- Filtering & Searching: `First`, `Last`, `Where`, `Reject`, `WhereNotNull`
- Transformation: `Map`, `MapWithKeys`, `Pluck`, `KeyBy`
- Sorting & Shuffling: `Sort`, `SortDesc`, `SortRecursive`, `Shuffle`
- Array Operations: `Flatten`, `Collapse`, `CrossJoin`, `Take`, `Random`
- Map Operations: `Only`, `Except`, `Forget`, `Dot`, `Undot`, `Divide`
- Utilities: `Join`, `Query`, `ToCssClasses`, `ToCssStyles`, `Wrap`

### `str` - String Helpers
- Case Conversion: `Camel`, `Snake`, `Kebab`, `Studly`, `Pascal`, `Upper`, `Lower`, `Title`, `Ucfirst`, `Lcfirst`
- String Extraction: `After`, `Before`, `Between`, `Substr`, `Take`, `CharAt`
- String Manipulation: `Replace`, `ReplaceFirst`, `ReplaceLast`, `Remove`, `Reverse`, `Repeat`
- Padding & Trimming: `PadLeft`, `PadRight`, `PadBoth`, `Trim`, `Ltrim`, `Rtrim`, `Squish`
- Validation: `Contains`, `StartsWith`, `EndsWith`, `IsAscii`, `IsJson`, `IsUrl`, `IsUuid`
- Formatting: `Limit`, `Words`, `Numbers`, `Slug`
- Encoding: `ToBase64`, `FromBase64`
- Regex: `Match`, `MatchAll`, `IsMatch`, `ReplaceMatches`

### `number` - Number Helpers
- Formatting: `Format`, `Currency`, `Percentage`, `FileSize`, `ForHumans`, `Abbreviate`, `Summarize`
- Conversion: `Ordinal`, `Spell`, `SpellOrdinal`
- Operations: `Clamp`, `Trim`, `Pairs`
- Locale & Currency: `UseLocale`, `UseCurrency`, `WithLocale`, `WithCurrency`, `DefaultLocale`, `DefaultCurrency`

## Contoh Penggunaan

Lihat folder `examples/` untuk contoh penggunaan lengkap:
- `examples/arr/` - Contoh Array helpers
- `examples/number/` - Contoh Number helpers
- `examples/str/` - Contoh String helpers

## Testing

```bash
# Jalankan semua tests
go test ./...

# Dengan coverage
go test -cover ./...
```

## Dokumentasi

Untuk dokumentasi lengkap, lihat:
- [GoDoc](https://pkg.go.dev/github.com/rulzi/helper-go) - Referensi API dan paket
- [Wiki](https://github.com/rulzi/helper-go/wiki) - Dokumentasi lengkap semua fungsi dengan contoh penggunaan

## License

This project is open source and available under the [MIT License](LICENSE).

Copyright (c) 2024 helper-go contributors

## Author

**Khoirul Afandi**

- Instagram: [@afandi_](https://instagram.com/afandi_)
- LinkedIn: [Khoirul Afandi](https://www.linkedin.com/in/khoirulafandi/)

---

Created with ❤️ for the Go community
