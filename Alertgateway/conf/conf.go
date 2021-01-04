package conf

const (
	CsrfToken = "11122233444"
)


type Config struct {
	Run RunStat
	LogInfo  LogConf
}

type LogConf struct {
	LogName string
	LogPath string
	LogStat string
	LogLevel int32
}

type RunStat struct {
	StartPort string
}

