package routeros

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type logLevel byte

const (
	TRACE logLevel = 1 + iota
	DEBUG
	INFO
	WARN
	ERROR
)

// ColorizedDebug Used to display provider log color messages.
// Please set the environment variable
func ColorizedDebug(ctx context.Context, msg string, args ...map[string]interface{}) {
	if _, set := os.LookupEnv("ROS_LOG_COLOR"); set {
		color.NoColor = false
	}
	tflog.Debug(ctx, color.GreenString(msg), args...)
}

func ColorizedMessage(ctx context.Context, level logLevel, msg string, args ...map[string]interface{}) {
	if _, set := os.LookupEnv("ROS_LOG_COLOR"); set {
		color.NoColor = false
	}
	switch level {
	case TRACE:
		tflog.Trace(ctx, color.GreenString(msg), args...)
	case DEBUG:
		tflog.Debug(ctx, color.GreenString(msg), args...)
	case INFO:
		tflog.Info(ctx, color.HiBlueString(msg), args...)
	case WARN:
		tflog.Warn(ctx, color.HiRedString(msg), args...)
	case ERROR:
		tflog.Error(ctx, color.HiRedString(msg), args...)
	}
}
