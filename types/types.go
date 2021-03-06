package types

const (
	ProbeMetricsServerPort = 9299
	ConfigServerPort       = 9300
	OriginServerPort       = 9301
)

type OriginResponse struct {
	ServerTime string `json:"server_time"`
	Token      string `json:"token"`
}
