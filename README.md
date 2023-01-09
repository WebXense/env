# env

This package provides methods to get environment variables with type conversion.

## Installation

```bash
go get -u github.com/WebXense/env
```

## Get value

```go
// string
s := env.String("STR_VAR")

// int
i := env.Int("INT_VAR")

// int64
i64 := env.Int64("INT64_VAR")

// uint
u := env.Uint("UINT_VAR")

// uint64
u64 := env.Uint64("UINT64_VAR")

// float64
f64 := env.Float64("FLOAT64_VAR")

// bool
b := env.Bool("BOOL_VAR")
```

## Get list of values

```go
// []string
ss := env.StringList("STR_VAR")

// []int
is := env.IntList("INT_VAR")

// []int64
i64s := env.Int64List("INT64_VAR")

// []uint
us := env.UintList("UINT_VAR")

// []uint64
u64s := env.Uint64List("UINT64_VAR")

// []float64
f64s := env.Float64List("FLOAT64_VAR")

// []bool
bs := env.BoolList("BOOL_VAR")
```

## Allow missing environment variable

Missing variable is an error case by default. Sometime, you may want to allow missing environment variable:

```go

s := env.String("STR_VAR", true) // will not error if missing

i := env.Int("INT_VAR", false) // same as default case

```

