package snowflake_test

import (
	"testing"
	"time"

	"github.com/Project-Sparrow/snowflake"
)

func TestGenerate(t *testing.T) {
	epoch := time.Now()
	snowflake.Init(epoch, 1, 1)

	s := snowflake.Generate()

	if (int(s)&0x3E0000)>>17 != 1 {
		t.Fail()
	}

	if (int(s)&0x1F000)>>12 != 1 {
		t.Fail()
	}
}
