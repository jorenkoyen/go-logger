package log

import (
	"github.com/jorenkoyen/go-logger"
	golog "log"
	"strings"
)

type Adapter struct {
	logger *logger.Logger
	level  logger.Level
}

func NewAdapter(name string, level logger.Level) *Adapter {
	return &Adapter{
		logger: WithName(name),
		level:  level,
	}
}

func (a *Adapter) Write(p []byte) (n int, err error) {
	// raw byte message includes newline character
	a.logger.Log(a.level, strings.TrimSpace(string(p)))
	return len(p), nil
}

func (a *Adapter) CreateStdLogger() *golog.Logger {
	return golog.New(a, "", 0)
}
