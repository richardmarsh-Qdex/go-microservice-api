package metrics

import "sync/atomic"

var (
	requestCount  uint64
	errorCount    uint64
	activeConns   int64
)

func IncRequest() {
	atomic.AddUint64(&requestCount, 1)
}

func IncError() {
	atomic.AddUint64(&errorCount, 1)
}

func Snapshot() map[string]uint64 {
	return map[string]uint64{
		"requests_total": atomic.LoadUint64(&requestCount),
		"errors_total":   atomic.LoadUint64(&errorCount),
	}
}

func ActiveConnections() int64 {
	return atomic.LoadInt64(&activeConns)
}
