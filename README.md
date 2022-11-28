# snowflake

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Project-Sparrow/snowflake)

Package snowflake provides a simple snowflake ID generator
along with interface implementations to make it easy to use
with database/sql and encoding/json.

## Functions

### func [Init](/snowflake.go#L47)

`func Init(e time.Time, w, p int)`

Init initializes the Snowflake generator.
This MUST be called before any calls to Generate.

## Types

### type [NullSnowflake](/null_snowflake.go#L12)

`type NullSnowflake struct { ... }`

NullSnowflake is a nullable Snowflake

#### func [NewNullSnowflake](/null_snowflake.go#L18)

`func NewNullSnowflake(s Snowflake, valid bool) NullSnowflake`

NewNullSnowflake creates a new NullSnowflake

### type [Snowflake](/snowflake.go#L55)

`type Snowflake uint64`

Snowflake represents a single Snowflake ID.

#### func [Generate](/snowflake.go#L59)

`func Generate() Snowflake`

Generate generates a new Snowflake.
This function is thread-safe.

```golang
package main

import (
	"fmt"
	"time"

	"github.com/Project-Sparrow/snowflake"
)

func main() {
	epoch := time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)

	snowflake.Init(epoch, 1, 1)

	fmt.Println(snowflake.Generate())
}

```

#### func [SnowflakeFromString](/snowflake.go#L76)

`func SnowflakeFromString(s string) (Snowflake, error)`

SnowflakeFromString attempts to parse a Snowflake from a string.
