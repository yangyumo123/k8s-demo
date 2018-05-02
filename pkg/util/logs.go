package util

import (
	"flag"
	"time"
)

var logFlushFreq = flag.Duration("log_flush_frequency", 5*time.Second, "")

func init() {
	flag.Set("logtostderr", "true")
}
