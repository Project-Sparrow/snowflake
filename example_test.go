package snowflake_test

import (
	"fmt"
	"time"

	"github.com/Project-Sparrow/snowflake"
)

func ExampleGenerate() {
	epoch := time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)

	snowflake.Init(epoch, 1, 1)

	fmt.Println(snowflake.Generate())
}
