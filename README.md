# Snowflake

A simple, dependency-free snowflake ID generation library.

## Usage

```go
import (
    time
    github.com/Project-Sparrow/snowflake
)

func main() {
    epoch := time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)

    snowflake.Init(epoch, 1, 1)

    fmt.Println(snowflake.Generate())
}
```
