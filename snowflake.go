// MIT License

// Copyright (c) 2022 Project-Sparrow

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package snowflake provides a simple snowflake ID generator
// along with interface implementations to make it easy to use
// with database/sql and encoding/json.
package snowflake

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"
)

var (
	workerID  int
	processID int
	epoch     time.Time
	increment int
	mtx       sync.Mutex
)

// Init initializes the Snowflake generator.
// This MUST be called before any calls to Generate.
func Init(e time.Time, w, p int) {
	epoch = e
	workerID = w
	processID = p
	increment = 0
}

// Snowflake represents a single Snowflake ID.
type Snowflake uint64

// Generate generates a new Snowflake.
// This function is thread-safe.
func Generate() Snowflake {
	mtx.Lock()
	defer mtx.Unlock()
	s := Snowflake(0)

	timeComp := time.Since(epoch).Milliseconds()
	s |= Snowflake(timeComp << 22)
	s |= Snowflake(workerID << 17)
	s |= Snowflake(processID << 12)
	s |= Snowflake(increment)

	increment++

	return s
}

// SnowflakeFromString attempts to parse a Snowflake from a string.
func SnowflakeFromString(s string) (Snowflake, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return Snowflake(i), nil
}

// MarshalJSON implements json.Marshaler interface
func (s Snowflake) MarshalJSON() ([]byte, error) {
	// Needs to be a string or later snowflakes will be truncated
	return json.Marshal(s.String())
}

// MarshalJSON implements json.Unmarshaler interface
func (s *Snowflake) UnmarshalJSON(bytes []byte) error {
	var snowflake string
	err := json.Unmarshal(bytes, &snowflake)
	if err != nil {
		return err
	}

	if snowflake == "" || snowflake == "null" {
		*s = 0
		return nil
	}

	snowInt, err := strconv.ParseInt(snowflake, 10, 64)
	if err != nil {
		return err
	}

	*s = Snowflake(snowInt)

	return nil
}

// String implements fmt.Stringer interface
func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// Value implements driver.Valuer interface
func (s Snowflake) Value() (driver.Value, error) {
	return int64(s), nil
}

// Scan implements sql.Scanner interface
func (s *Snowflake) Scan(value interface{}) error {
	if value == nil {
		*s = Snowflake(0)
		return nil
	}

	switch v := value.(type) {
	case int64:
		*s = Snowflake(v)
	case string:
		iv, err := SnowflakeFromString(v)
		if err != nil {
			return err
		}

		*s = iv
	default:
		return errors.New("not a valid snowflake type")
	}
	return nil
}

// CreatedAt returns the time component of the Snowflake as a time.Time
func (s Snowflake) CreatedAt() time.Time {
	unixMilis := (s >> 22) + Snowflake(epoch.UnixMilli())

	return time.UnixMilli(int64(unixMilis))
}
