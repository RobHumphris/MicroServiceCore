package structs

import "log"

const stopped = "Stopped"
const running = "Running"
const errored = "Errored"

// ServerData holds the structures used in the running and monitoring of this server.
type ServerData struct {
	currentConfig     *TranscodeConfig
	streamConfig      *StreamConfig
	currentState      string
	currentError      error
	executablePath    string
	Port              string
	metricsServerType string
	metricsEndpoint   string
	metricsEnabled    bool
	Debug             bool
}

// ConfigComplete returns true if all config items have been set to a value
func (tc *TranscodeConfig) ConfigComplete() bool {
	return tc.Mode != "" && tc.Source != "" && tc.Channel != "" && tc.Destination != "" && tc.SourceType != ""
}

// NewServerData constructs a new ServerData structure
func NewServerData() *ServerData {
	port, err := getPort()
	if err != nil {
		log.Printf("Warning: %v\n", err)
	}

	executablePath, err := getExecutablePath()
	if err != nil {
		log.Printf("Warning: %v\n", err)
	}

	mst, err := getMetricsServerType()
	if err != nil {
		log.Printf("Error with metrics server type %v\n", err)
	}

	mep, err := getMetricsServerEndpoint()
	if err != nil {
		log.Printf("Error with metrics server endpoint %v\n", err)
	}

	sd := new(ServerData)
	(*sd).currentConfig = &TranscodeConfig{Source: "", Channel: "", Destination: ""}
	(*sd).currentState = stopped
	(*sd).executablePath = executablePath
	(*sd).Port = port
	(*sd).metricsServerType = mst
	(*sd).metricsEndpoint = mep
	(*sd).metricsEnabled = (mst != "" && mep != "")
	(*sd).Debug = getDebugState()
	return sd
}

// TranscodeConfig contains the elements required to start the Transcoder streaming.
type TranscodeConfig struct {
	Ident       string `json:"transcoderId"`
	Mode        string `json:"mode"`
	Source      string `json:"source"`
	SourceType  string `json:"sourceType"`
	Channel     string `json:"channel"`
	Destination string `json:"destination"`
}

// StreamInfo is effectively the tuple: {index: 0, rate: "200k"}
type StreamInfo struct {
	Index int    `json:"index"`
	Rate  string `json:"rate"`
}

// SubtitleInfo defines the subtitle index - experimental.
type SubtitleInfo struct {
	Index int `json:"index"`
}

// StreamConfig contains arrays of both video and audio StreamInfo entries
type StreamConfig struct {
	HardwareAccel    string       `json:"hwAccel"`
	VideoDecoder     string       `json:"videoDecoder"`
	VideoDecoderOpts []string     `json:"videoDecoderOpts"`
	AudioDecoder     string       `json:"audioDecoder"`
	AudioDecoderOpts []string     `json:"audioDecoderOpts"`
	VideoEncoder     string       `json:"videoEncoder"`
	VideoEncoderOpts []string     `json:"videoEncoderOpts"`
	AudioEncoder     string       `json:"audioEncoder"`
	AudioEncoderOpts []string     `json:"audioEncoderOpts"`
	VideoStreams     []StreamInfo `json:"videoStreams"`
	AudioStreams     []StreamInfo `json:"audioStreams"`
	WindowSize       int          `json:"windowSize"`
	SubtitleCodec    string       `json:"subtitleCodec"`
	SubtitlesIndexes []int        `json:"subtitleIndexes"`
}
