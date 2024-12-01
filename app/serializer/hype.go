// app/serializer/hype.go
package serializer

import (
	"math"
	"nevissGo/ent"
	"time"
)

type HypeSerializer struct {
	AmountRemaining   int     `json:"amount_remaining"`
	MaxHype           int     `json:"max_hype"`
	HypePerSecond     float64 `json:"hype_per_second"`
	TimeUntilNextHype int     `json:"time_until_next_hype"`
	LastUpdatedAt     string  `json:"last_updated_at"`
}

func NewHype(hype *ent.Hype) *HypeSerializer {
	hypePerSecond := float64(hype.HypePerMinute) / 60.0

	timeSinceUpdate := time.Since(hype.LastUpdatedAt)
	secondsSinceUpdate := timeSinceUpdate.Seconds()
	secondsPerHype := 60.0 / float64(hype.HypePerMinute)

	remainingSecondsFloat := secondsPerHype - math.Mod(secondsSinceUpdate, secondsPerHype)
	if remainingSecondsFloat == secondsPerHype {
		remainingSecondsFloat = 0
	}
	remainingSeconds := int(math.Ceil(remainingSecondsFloat))

	// Set TimeUntilNextHype to 0 if AmountRemaining is at or above MaxHype
	if hype.AmountRemaining >= hype.MaxHype {
		remainingSeconds = 0
	}

	return &HypeSerializer{
		AmountRemaining:   hype.AmountRemaining,
		MaxHype:           hype.MaxHype,
		HypePerSecond:     hypePerSecond,
		TimeUntilNextHype: remainingSeconds,
		LastUpdatedAt:     hype.LastUpdatedAt.Format(time.RFC3339),
	}
}
