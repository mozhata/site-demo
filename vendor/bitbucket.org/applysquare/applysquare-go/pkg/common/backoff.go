package common

import (
	"math/rand"
	"time"
)

func BackOff(retries int, baseDelay, maxDelay time.Duration) time.Duration {
	const multiplier = 1.3
	const randRatio = 0.4

	backOff, maxDelayF := float64(baseDelay), float64(maxDelay)
	for backOff < maxDelayF && retries > 0 {
		retries--
		backOff *= multiplier
	}
	if backOff > maxDelayF {
		backOff = maxDelayF
	}

	backOff -= rand.Float64() * randRatio * backOff
	if backOff < 0 {
		backOff = 0
	}
	return time.Duration(backOff)
}
