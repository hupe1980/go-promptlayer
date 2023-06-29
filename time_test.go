package promptlayer

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		testTime := Now()
		expectedJSON := []byte(strconv.FormatInt(time.Time(testTime).Unix(), 10))

		jsonData, err := json.Marshal(testTime)
		assert.NoError(t, err)
		assert.Equal(t, expectedJSON, jsonData)
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		jsonData := []byte("1688059574")

		var testTime Time
		err := json.Unmarshal(jsonData, &testTime)
		assert.NoError(t, err)

		expectedTime := time.Unix(1688059574, 0)
		assert.Equal(t, expectedTime, time.Time(testTime))
	})

	t.Run("Unix", func(t *testing.T) {
		testTime := Now()
		expectedUnix := time.Time(testTime).Unix()

		unixTime := testTime.Unix()
		assert.Equal(t, expectedUnix, unixTime)
	})

	t.Run("Add", func(t *testing.T) {
		testTime := Now()
		duration := 5 * time.Second

		expectedTime := time.Time(testTime).Add(duration)
		newTime := testTime.Add(duration)
		assert.Equal(t, expectedTime, time.Time(newTime))
	})

	t.Run("Time", func(t *testing.T) {
		testTime := Now()
		expectedTime := time.Time(testTime).UTC()

		resultTime := testTime.Time()
		assert.Equal(t, expectedTime, resultTime)
	})
}
