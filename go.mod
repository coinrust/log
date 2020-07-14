module github.com/coinrust/log

go 1.13

require (
	github.com/coinrust/slog v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.15.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/coinrust/slog => ../../../github.com/coinrust/slog
