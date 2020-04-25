package blog

import (
	"time"
)

func getCurrentUTCTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}
