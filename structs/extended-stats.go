package structs

import (
	"time"

	"github.com/thoas/stats"
)

// ExtendedData combines the stats data with the Transcoder specific data
type ExtendedData struct {
	*stats.Data
	Transcoder  *TranscoderData  `json:"transcoderData"`
	Config      *TranscodeConfig `json:"transcoderConfig"`
	StreamStats *StreamStats     `json:"streamStats"`
}

// ExtendedStats brings metrics from the executable to the table
type ExtendedStats struct {
	*stats.Stats
	LastErrorTime      time.Time
	LastRestartTime    time.Time
	ErrorCount         int64
	ErrorFrequency     float64
	PeakErrorFrequency float64
	RestartCount       int64
	RestartFrequency   float64
	LastError          string
	MaxRestartTries    int64
}

// NewExtendedStats constructs a new instance of ExtendedStats
func NewExtendedStats() *ExtendedStats {
	s := &ExtendedStats{Stats: stats.New()}
	s.LastErrorTime = time.Now()
	s.LastRestartTime = s.LastErrorTime
	s.ErrorCount = 0
	s.ErrorFrequency = 0.0
	s.PeakErrorFrequency = 0.0
	s.RestartCount = 0
	s.RestartFrequency = 0.0
	s.MaxRestartTries = 3
	return s
}

// Data returns the data serializable structure
func (es *ExtendedStats) Data(current *ServerData) *ExtendedData {
	avgErrFreq := 0.0
	if es.ErrorCount != 0 {
		avgErrFreq = es.ErrorFrequency / float64(es.ErrorCount)
	}

	lastErrorSec := time.Since(es.LastErrorTime)
	lastRestartSec := time.Since(es.LastRestartTime)

	e := &ExtendedData{
		Data: new(stats.Data),
		Transcoder: &TranscoderData{
			LastErrorSec:          lastErrorSec.Seconds(),
			LastRestartSec:        lastRestartSec.Seconds(),
			CurrentState:          current.currentState,
			TotalErrorCount:       es.ErrorCount,
			AverageErrorFrequency: avgErrFreq,
			PeakErrorFrequency:    es.PeakErrorFrequency,
			RestartCount:          es.RestartCount,
			RestartFrequency:      es.RestartFrequency,
			LastError:             es.LastError,
		},
		StreamStats: GetStreamAverages(),
	}
	e.Data = es.Stats.Data()
	e.Config = current.currentConfig
	return e
}

// RecordError increments the error count and associated metrics
func (es *ExtendedStats) RecordError(message string) {
	elapsed := time.Since(es.LastErrorTime)
	es.LastErrorTime = time.Now()
	es.ErrorCount++
	es.ErrorFrequency = 1 / elapsed.Seconds()
	if es.PeakErrorFrequency < es.ErrorFrequency {
		es.PeakErrorFrequency = es.ErrorFrequency
	}
	es.LastError = message
}

// RecordRestart increments the restart counter, returns false if number of restarts have exceeded MaxRestartTries
func (es *ExtendedStats) RecordRestart() bool {
	elapsed := time.Since(es.LastRestartTime)
	es.LastRestartTime = time.Now()
	es.RestartCount++
	es.RestartFrequency = 1 / elapsed.Seconds()
	return es.RestartCount > es.MaxRestartTries
}

// ResetAllStats zeros the clock on all the ExtendedStats
func (es *ExtendedStats) ResetAllStats() {
	es.LastError = "Stats cleared (check last error time)"
	es.LastErrorTime = time.Now()
	es.LastRestartTime = es.LastErrorTime
	es.ErrorCount = 0
	es.ErrorFrequency = 0.0
	es.PeakErrorFrequency = 0.0
	es.RestartCount = 0
	es.RestartFrequency = 0.0
}
