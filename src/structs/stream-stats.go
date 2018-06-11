package structs

// StreamStatReport presents the details for the stream
type StreamStatReport struct {
	StreamType    string `json:"streamType"`
	BitRate       string `json:"bitRate"`
	AverageSizeKb int    `json:"averageSizeKb"`
}

// StreamStats holds a map of streams and their report data
type StreamStats struct {
	report map[string]StreamStatReport
}

// GetStreamAverages returns the latest averages for the transcoded streams
func GetStreamAverages() *StreamStats {
	retObject := StreamStats{report: make(map[string]StreamStatReport)}

	/**
	streamStats.RLock()
	for key, stat := range streamStats.m {
		a := retMap[key]
		a.StreamType = stat.streamType
		a.BitRate = stat.bitRate
		if stat.samples > 0 {
			a.AverageSizeKb = int(stat.totalKB) / stat.samples
		} else {
			a.AverageSizeKb = -1
		}

		retMap[key] = a
	}
	streamStats.RUnlock(
	*/
	return &retObject
}
