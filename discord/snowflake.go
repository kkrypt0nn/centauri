package discord

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Snowflake represents a unique ID
type Snowflake uint64

// MarshalJSON will take care to marshal a snowflake (discord.Snowflake) to a string
func (s *Snowflake) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(*s), 10)), nil
}

// UnmarshalJSON will take care to unmarshal a string to a snowflake (discord.Snowflake)
func (s *Snowflake) UnmarshalJSON(b []byte) error {
	var source string
	if err := json.Unmarshal(b, &source); err != nil {
		return err
	}

	parsed, err := ParseSnowflake(source)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

// String returns the string representation of a snowflake (discord.Snowflake)
func (s *Snowflake) String() string {
	return strconv.FormatUint(uint64(*s), 10)
}

// CreatedAt returns the creation time of a snowflake (discord.Snowflake)
func (s *Snowflake) CreatedAt() time.Time {
	timestamp := (*s >> 22) + 1420070400000
	return time.UnixMilli(int64(timestamp))
}

// ArraySnowflakes represents an array of snowflakes
type ArraySnowflakes []Snowflake

// String returns a comma-separated string of the array of strings
func (s *ArraySnowflakes) String() string {
	var snowflakes string
	for i, snowflake := range *s {
		if i == len(*s)-1 {
			snowflakes += snowflake.String()
			continue
		}
		snowflakes += fmt.Sprintf("%d,", snowflake)
	}
	return snowflakes
}

// MarshalJSON will take care to marshal an array of snowflakes (discord.Snowflake) to an array of strings
func (s *ArraySnowflakes) MarshalJSON() ([]byte, error) {
	var snowflakeStrings []string
	for _, snowflake := range *s {
		snowflakeStrings = append(snowflakeStrings, strconv.FormatUint(uint64(snowflake), 10))
	}
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(snowflakeStrings)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON will take care to unmarshal an array of strings to an array of snowflakes (discord.Snowflake)
func (s *ArraySnowflakes) UnmarshalJSON(b []byte) error {
	var snowflakes []Snowflake
	var source []string
	if err := json.Unmarshal(b, &source); err != nil {
		return err
	}

	for _, stringSnowflake := range source {
		parsed, err := ParseSnowflake(stringSnowflake)
		if err != nil {
			return err
		}
		snowflakes = append(snowflakes, parsed)
	}
	*s = snowflakes
	return nil
}

// ParseSnowflake parses a string as a snowflake (discord.Snowflake)
func ParseSnowflake(snowflake string) (Snowflake, error) {
	if snowflake == "" {
		return Snowflake(0), nil
	}

	sf, err := strconv.ParseUint(snowflake, 10, 64)
	return Snowflake(sf), err
}

// MustParseSnowflake parses a string as a snowflake (discord.Snowflake) - must be valid
func MustParseSnowflake(snowflake string) Snowflake {
	sf, _ := strconv.ParseUint(snowflake, 10, 64)
	return Snowflake(sf)
}
