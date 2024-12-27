package config

import (
	"context"
	"fmt"
	"github.com/lmittmann/tint"
	"github.com/mdobak/go-xerrors"
	slogmulti "github.com/samber/slog-multi"
	"github.com/tienla0409/gofiber-boilerplate/constant"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}

type fiberHandler struct {
	slog.Handler
}

func InitLoggerConfig(logPath, logType string) (*os.File, error) {
	err := os.MkdirAll(logPath, 0755)
	if err != nil {
		return nil, err
	}

	logFile := fmt.Sprintf("%s/%s_log-%s.log", logPath, logType, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		return nil, err
	}

	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: replaceAttr,
	})
	fh := &fiberHandler{Handler: fileHandler}

	consoleHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.DateTime,
	})
	ch := &fiberHandler{Handler: consoleHandler}

	logger := slog.New(slogmulti.Fanout(
		fh,
		ch,
	))
	slog.SetDefault(logger)

	return file, nil
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = fmtErr(v)
		}
	}

	return a
}

// fmtErr returns a slog.Value with keys `msg` and `trace`. If the error
// does not implement interface { StackTrace() errors.StackTrace }, the `trace`
// key is omitted.
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	frames := marshalStack(err)

	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}

	return slog.GroupValue(groupValues...)
}

// marshalStack extracts stack frames from the error
func marshalStack(err error) []stackFrame {
	trace := xerrors.StackTrace(err)

	if len(trace) == 0 {
		return nil
	}

	frames := trace.Frames()

	s := make([]stackFrame, len(frames))

	for i, v := range frames {
		f := stackFrame{
			Source: filepath.Join(
				filepath.Base(filepath.Dir(v.File)),
				filepath.Base(v.File),
			),
			Func: filepath.Base(v.Function),
			Line: v.Line,
		}

		s[i] = f
	}

	return s
}

func (f *fiberHandler) Handle(ctx context.Context, r slog.Record) error {
	if ctx.Value(constant.RequestID) == nil {
		return f.Handler.Handle(ctx, r)
	}

	requestID := ctx.Value(constant.RequestID).(string)
	sourceIp := ctx.Value(constant.SourceIP).(string)
	method := ctx.Value(constant.Method).(string)
	path := ctx.Value(constant.Path).(string)

	requestGroup := slog.Group(
		constant.Request,
		slog.String(constant.RequestID, requestID),
		slog.String(constant.SourceIP, sourceIp),
		slog.String(constant.Method, method),
		slog.String(constant.Path, path),
	)

	r.AddAttrs(requestGroup)

	return f.Handler.Handle(ctx, r)
}
