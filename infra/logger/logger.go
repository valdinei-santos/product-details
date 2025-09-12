package logger

import (
	"context"
	"log/slog"
	"os"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}
type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() *SlogLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		//Level: slog.LevelDebug,
	})
	log := slog.New(handler)
	return &SlogLogger{logger: log}
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.DebugContext(context.Background(), msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.InfoContext(context.Background(), msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.WarnContext(context.Background(), msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.ErrorContext(context.Background(), msg, args...)
}

func (l *SlogLogger) With(args ...any) Logger {
	return &SlogLogger{logger: l.logger.With(args...)}
}

func (l *SlogLogger) WithContext(ctx context.Context) Logger {
	// A forma correta é criar um novo logger *sem* associar o contexto diretamente aqui.
	// A associação e extração de valores do contexto são feitas através do Handler's Attrer.
	// Retornamos uma nova instância do SlogLogger com o logger existente.
	// A presença do contexto será tratada pelo Handler quando os logs forem formatados.
	return &SlogLogger{logger: l.logger}
}

func (l *SlogLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *SlogLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *SlogLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *SlogLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}
