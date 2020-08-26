package log

import (
	log "github.com/sirupsen/logrus"
	"bytes"
	"fmt"
)

type TestFormatter struct {
	TimestampFormat string
}

func (t *TestFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestampFormat := t.TimestampFormat
	t.appendKeyValue(b, "time", entry.Time.Format(timestampFormat))
	t.appendKeyValue(b, "level", entry.Level.String())
	t.appendKeyValue(b, "msg", entry.Message)
	t.appendKeyValue(b, "Format", "test")

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *TestFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

func (f *TestFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(fmt.Sprintf("%q", stringVal))
}