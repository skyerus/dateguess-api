package rand

import (
	"math/rand"
	"time"
)

func Time() time.Time {
	var start int64
	start = -2208944868 // Jan 1st 1900
	randomUnix := rand.Int63n(time.Now().Unix()-start) + start
	return time.Unix(randomUnix, 0)
}
