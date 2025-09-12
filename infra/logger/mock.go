package logger

import (
	"context"
	"fmt"
	"sync"
)

// MockLogger é um mock com a implementação da interface Logger
/* type MockLogger struct {
	mu    sync.Mutex
	logs  []string
	level string
} */
type MockLogger struct {
	mu      sync.Mutex
	logs    map[string][]string // Armazena logs por contexto
	context string              // Contexto atual
}

// NewMockLogger creates a new instance of MockLogger
/* func NewMockLogger() *MockLogger {
	return &MockLogger{
		logs: make([]string, 0),
	}
} */
func NewMockLogger() *MockLogger {
	return &MockLogger{
		logs: make(map[string][]string),
	}
}

// SetContext define o contexto atual
func (m *MockLogger) SetContext(context string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.context = context
}

// Debug mocks the Debug logging method
/* func (m *MockLogger) Debug(msg string, args ...any) {
	m.log("DEBUG", msg, args...)
} */
func (m *MockLogger) Debug(msg string, args ...any) {
	m.log("DEBUG", msg, args...)
}

// Info mocks the Info logging method
/* func (m *MockLogger) Info(msg string, args ...any) {
	m.log("INFO", msg, args...)
} */
func (m *MockLogger) Info(msg string, args ...any) {
	m.log("INFO", msg, args...)
}

// Warn mocks the Warn logging method
/* func (m *MockLogger) Warn(msg string, args ...any) {
	m.log("WARN", msg, args...)
} */
func (m *MockLogger) Warn(msg string, args ...any) {
	m.log("WARN", msg, args...)
}

// Error mocks the Error logging method
/* func (m *MockLogger) Error(msg string, args ...any) {
	m.log("ERROR", msg, args...)
} */
func (m *MockLogger) Error(msg string, args ...any) {
	m.log("ERROR", msg, args...)
}

// With mocks the With method, returning the same logger for simplicity
func (m *MockLogger) With(args ...any) Logger {
	m.log("WITH", "Adding context", args...)
	return m
}

// WithContext mocks the WithContext method, returning the same logger for simplicity
func (m *MockLogger) WithContext(ctx context.Context) Logger {
	m.log("WITH_CONTEXT", "Adding context with context.Context", ctx)
	return m
}

// DebugContext mocks the DebugContext logging method
func (m *MockLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	m.log("DEBUG_CONTEXT", msg, args...)
}

// InfoContext mocks the InfoContext logging method
func (m *MockLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	m.log("INFO_CONTEXT", msg, args...)
}

// WarnContext mocks the WarnContext logging method
func (m *MockLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	m.log("WARN_CONTEXT", msg, args...)
}

// ErrorContext mocks the ErrorContext logging method
func (m *MockLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	m.log("ERROR_CONTEXT", msg, args...)
}

// log stores the log message in memory
/* func (m *MockLogger) log(level, msg string, args ...any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	formatted := fmt.Sprintf("[%s] %s", level, fmt.Sprintf(msg, args...))
	m.logs = append(m.logs, formatted+"\n")
} */
func (m *MockLogger) log(level, msg string, args ...any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.context == "" {
		m.context = "default"
	}
	formatted := fmt.Sprintf("[%s] %s", level, fmt.Sprintf(msg, args...))
	m.logs[m.context] = append(m.logs[m.context], formatted)
}

// GetLogs returns all logs
/* func (m *MockLogger) GetLogs() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.logs
} */
func (m *MockLogger) GetLogs(context string) []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.logs[context]
}
