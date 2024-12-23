// app/serializer/hype_test.go
package serializer

import (
	"math"
	"testing"
	"time"

	"nevissGo/ent"
)

func TestNewHype(t *testing.T) {
	tests := []struct {
		name            string
		hype            *ent.Hype
		elapsedTime     time.Duration
		expectedSeconds int
	}{
		{
			name: "Half a second elapsed",
			hype: &ent.Hype{
				HypePerMinute:   60,
				AmountRemaining: 5,
				MaxHype:         10,
				LastUpdatedAt:   time.Now(),
			},
			elapsedTime:     500 * time.Millisecond,
			expectedSeconds: 1,
		},,
	{
		name: "Partial time elapsed",
		hype: &ent.Hype{
		HypePerMinute:   2,
		AmountRemaining: 5,
		MaxHype:         10,
		LastUpdatedAt:   time.Now(),
	},
		elapsedTime:     15 * time.Second,
		expectedSeconds: 15,
	},
	{
		name: "Multiple hypes elapsed",
		hype: &ent.Hype{
		HypePerMinute:   2,
		AmountRemaining: 5,
		MaxHype:         10,
		LastUpdatedAt:   time.Now(),
	},
		elapsedTime:     75 * time.Second,
		expectedSeconds: 15,
	},
	{
		name: "AmountRemaining equals MaxHype",
		hype: &ent.Hype{
		HypePerMinute:   60,
		AmountRemaining: 10,
		MaxHype:         10,
		LastUpdatedAt:   time.Now(),
	},
		elapsedTime:     30 * time.Second,
		expectedSeconds: 0,
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentTime := time.Now()
			hype := tt.hype
			hype.LastUpdatedAt = currentTime.Add(-tt.elapsedTime)

			serializer := NewHype(hype)

			if serializer.TimeUntilNextHype != tt.expectedSeconds {
				secondsPerHype := 60.0 / float64(hype.HypePerMinute)
				secondsSinceUpdate := tt.elapsedTime.Seconds()
				remainingSecondsFloat := secondsPerHype - math.Mod(secondsSinceUpdate, secondsPerHype)
				calculatedSeconds := int(math.Ceil(remainingSecondsFloat))

				t.Errorf("expected TimeUntilNextHype %d, got %d. Calculated: %d", tt.expectedSeconds, serializer.TimeUntilNextHype, calculatedSeconds)
			}
		})
	}
}
