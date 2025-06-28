package main

import (
	"bytes"
	"io"
	"log/slog"
	"strings"
	"testing"
)

// Первый способ мокирования логгера

func someFunc(logger *slog.Logger) {
	logger.Info("this is info")
}

func TestLoggerMock(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	someFunc(logger)
}

// Второй способ мокирования логгера
func someFunc2(logger *slog.Logger) {
	logger.Info("this is info")
}

func TestLoggerMock2(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := slog.New(slog.NewJSONHandler(buf, nil))

	// Вызов функции
	someFunc2(logger)

	// Проверяем значение логгера
	if s := buf.String(); !strings.Contains(s, "this is info") {
		t.Errorf("got %s, want %s", s, "this is info")
	}
}

// Третий способ мокирования логгера
func someFunc3(logger *slog.Logger) {
	logger.Info("this is info")
}

func TestLoggerMock3(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := slog.New(slog.NewJSONHandler(buf, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	// Вызов функции
	someFunc2(logger)

	// Проверяем значение логгера
	const wantLog = `level=INFO msg="this is info"`
	if s := strings.TrimSpace(buf.String()); wantLog != s {
		t.Errorf("got %s, want %s", s, "this is info")
	}
}
