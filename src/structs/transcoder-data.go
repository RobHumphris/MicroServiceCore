package structs

// TranscoderData contains transcoder specific data
type TranscoderData struct {
	LastErrorSec          float64 `json:"lastErrorSec"`
	LastRestartSec        float64 `json:"lastRestartSec"`
	CurrentState          string  `json:"currentState"`
	TotalErrorCount       int64   `json:"totalErrorCount"`
	AverageErrorFrequency float64 `json:"avgErrorFreq"`
	PeakErrorFrequency    float64 `json:"peakErrorFreq"`
	RestartCount          int64   `json:"restartCount"`
	RestartFrequency      float64 `json:"restartFreq"`
	LastError             string  `json:"lastError"`
}
