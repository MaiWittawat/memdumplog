package memdumplog

type Driver string

const (
	Logrus  Driver = "logrus"
	Zap     Driver = "zap"
	Zerolog Driver = "zerolog"
	Slog    Driver = "slog"
)

type Config struct {
	Driver     Driver
	BufferSize int
}
