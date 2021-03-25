package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type TestHook struct {
}

func (t *TestHook) Levels() []log.Level {
	/* return []log.Level{
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	} */
	return log.AllLevels
}

func (t *TestHook) Fire(entry *log.Entry) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	line, _ := entry.String()

	file.Write([]byte(line))
	return nil
}
