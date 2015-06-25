package tools

import (
	"fmt"
	"time"
)

var (
	intSecond      = int64(time.Second)
	intMillisecond = int64(time.Millisecond)
	intMicrosecond = int64(time.Microsecond)
	intNanosecond  = int64(time.Nanosecond)

	floatSecond      = float64(time.Second)
	floatMillisecond = float64(time.Millisecond)
	floatMicrosecond = float64(time.Microsecond)
	floatNanosecond  = float64(time.Nanosecond)
)

func SinceStr(then time.Time) string {
	duration := time.Now().UnixNano() - then.UnixNano()
	switch {
	case duration > intSecond:
		return fmt.Sprintf("%.2f s", float64(duration)/floatSecond)
	case duration > intMillisecond:
		return fmt.Sprintf("%.2f ms", float64(duration)/floatMillisecond)
	case duration > intMicrosecond:
		return fmt.Sprintf("%.2f us", float64(duration)/floatMicrosecond)
	default:
		return fmt.Sprintf("%.2f ns", float64(duration)/floatNanosecond)
	}
}
