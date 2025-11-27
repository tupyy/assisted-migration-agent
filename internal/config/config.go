package config

const (
	ConnectedMode    string = "connected"
	DisconnectedMode string = "disconnected"
)

//go:generate go run github.com/ecordell/optgen -output zz_generated.configuration.go . Configuration
type Configuration struct {
	ServerMode    string         `debugmap:"visible" default:"dev"`
	Mode          string         `debugmap:"visible" default:"disconnected"`
	HTTPPort      int            `debugmap:"visible" default:"8080"`
	StaticsFolder string         `debugmap:"visible"`
	DataFolder    string         `debugmap:"visible"`
	Auth          Authentication `debugmap:"visible"`

	// Log
	LogFormat string `debugmap:"visible"`
	LogLevel  string `debugmap:"visible"`
}

type Authentication struct {
	Enabled     bool   `debugmap:"visible" default:"true"`
	JWTFilePath string `debugmap:"visible"`
}
