package `uuid` provides an RFC 4122 UUID generator.

Each supported UUID version provides two functions: `GenerateVn`, where n
is the version number, and `GenerateVnString`. The `GenerateVn` function
returns the UUID as a `Uuid` type (currently implemented as a `[]byte`), and
the `GenerateVnString` returns the UUID as a `string`.

Tests for each exported function are also provided, which include checks
to ensure that generated UUIDs are in the proper format.

## FUNCTIONS
```go
// GenerateV4String returns a version 4 UUID as a string.
func GenerateV4String() (u string, err error)
// GenerateV4 returns a version 4 UUID as a byte slice.
func GenerateV4() (u Uuid, err error)
```

## TYPES
```go
type Uuid []byte
```
