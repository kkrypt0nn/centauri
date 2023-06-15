package discord

import (
	"testing"
)

func TestSnowflakeParse(t *testing.T) {
	var expect Snowflake = 1118794921445433344
	sf := "1118794921445433344"
	snowflake, err := ParseSnowflake(sf)
	if err == nil && snowflake == expect {
		t.Log("Successfully parsed a string into a snowflake")
		return
	}
	t.Errorf("Failed parsing a string into a snowflake (got: %d, expected: %d)", snowflake, expect)
}

func TestSnowflakeString(t *testing.T) {
	expect := "1118794921445433344"
	var snowflake Snowflake = 1118794921445433344
	snowflakeString := snowflake.String()
	if snowflakeString == expect {
		t.Log("Successfully got the string representation of a snowflake")
		return
	}
	t.Errorf("Failed getting the string representation of a snowflake (got: %s, expected: %s)", snowflakeString, expect)
}

func TestSnowflakeCreatedAt(t *testing.T) {
	expect := "2023-06-15 06:51:35.477 +0000 UTC"
	var snowflake Snowflake = 1118794921445433344
	utcCreated := snowflake.CreatedAt().UTC().String()
	if utcCreated == expect {
		t.Log("Successfully got the right creation date of a snowflake")
		return
	}
	t.Errorf("Failed getting the right creation date of a snowflake (got: %s, expected: %s)", utcCreated, expect)
}

func TestArraySnowflakeStringOne(t *testing.T) {
	expect := "1118794921445433344"
	var snowflakes = ArraySnowflakes{
		Snowflake(1118794921445433344),
	}
	snowflakesString := snowflakes.String()
	if snowflakesString == expect {
		t.Log("Successfully got the string representation of an array of snowflakes")
		return
	}
	t.Errorf("Failed getting the string representation of an array of snowflakes (got: %s, expected: %s)", snowflakesString, expect)
}

func TestArraySnowflakeStringMultiple(t *testing.T) {
	expect := "1118794921445433344,1118828871706480721,1118828906611478578"
	var snowflakes = ArraySnowflakes{
		Snowflake(1118794921445433344),
		Snowflake(1118828871706480721),
		Snowflake(1118828906611478578),
	}
	snowflakesString := snowflakes.String()
	if snowflakesString == expect {
		t.Log("Successfully got the string representation of an array of snowflakes")
		return
	}
	t.Errorf("Failed getting the string representation of an array of snowflakes (got: %s, expected: %s)", snowflakesString, expect)
}
