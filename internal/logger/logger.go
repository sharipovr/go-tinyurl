package logger

import "go.uber.org/zap"

func New() *zap.Logger {
	l, _ := zap.NewProduction() // для краткости игнорируем err
	return l
}
