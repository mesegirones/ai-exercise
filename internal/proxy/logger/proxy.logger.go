package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Config interface {
	IsLocal() bool
}

type LoggerProxy struct {
	IsLocal     bool
	Logger      *slog.Logger
	ErrorClient *slog.Logger
}

func NewLoggerProxy(ctx context.Context, config Config) *LoggerProxy {
	return &LoggerProxy{
		IsLocal: config.IsLocal(),
		Logger:  createLogger(),
	}
}

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

func (p *LoggerProxy) Debug(v ...interface{}) {
	p.Logger.Debug(fmt.Sprint(v...))
}

func (p *LoggerProxy) Info(v ...interface{}) {
	p.Logger.Info(fmt.Sprint(v...))
}

func (p *LoggerProxy) Warning(v ...interface{}) {
	p.Logger.Warn(fmt.Sprint(v...))
}

func (p *LoggerProxy) Error(v ...interface{}) {
	p.Logger.Error(fmt.Sprint(v...))
}

func (p *LoggerProxy) Get() *slog.Logger {
	return p.Logger
}

type LoggerDebugWriter struct {
	Logger *LoggerProxy
}

func (w *LoggerDebugWriter) Write(p []byte) (int, error) {
	logString := string(p)
	w.Logger.Debug(logString)
	return len(p), nil
}

func (l *LoggerProxy) GetDebugWriter() io.Writer {
	return &LoggerDebugWriter{
		Logger: l,
	}
}

type LoggerInfoWriter struct {
	Logger *LoggerProxy
}

func (w *LoggerInfoWriter) Write(p []byte) (int, error) {
	logString := string(p)
	w.Logger.Info(logString)
	return len(p), nil
}

func (l *LoggerProxy) GetInfoWriter() io.Writer {
	return &LoggerInfoWriter{
		Logger: l,
	}
}

type LoggerErrorWriter struct {
	Logger *LoggerProxy
}

func (w *LoggerErrorWriter) Write(p []byte) (int, error) {
	logString := string(p)
	w.Logger.Error(logString)
	return len(p), nil
}

func (l *LoggerProxy) GetErrorWriter() io.Writer {
	return &LoggerErrorWriter{
		Logger: l,
	}
}
