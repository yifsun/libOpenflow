package log

import (
	"flag"
	"log/syslog"
	"strconv"

	"k8s.io/klog/v2"
)

// Usage example:
//    log.SetLoggingOutputType(log.LOG_OUTPUT_TYPE_SYSLOG)
//    log.SetLoggingOutputType(log.LOG_OUTPUT_TYPE_CONSOLE)
//    log.SetLoggingOutputType(log.LOG_OUTPUT_TYPE_FILE)
//    log.SetLoggingOutputFile("myfile.log")
//    log.SetLoggingVerbosity(4)
//
//    klog.V(4).Info("messages")
//    klog.Flush()

func SetLoggingVerbosity(level klog.Level) {
	flag.Set("v", strconv.Itoa(int(level)))
}

func SetLoggingOutputFile(file string) {
	flag.Set("log_file", file)
}

type LoggingOutputType int32

const (
	LOG_OUTPUT_TYPE_CONSOLE		LoggingOutputType = iota
	LOG_OUTPUT_TYPE_FILE
	LOG_OUTPUT_TYPE_SYSLOG
)

func SetLoggingOutputType(t LoggingOutputType) error {
	switch t {
	case LOG_OUTPUT_TYPE_CONSOLE:
		flag.Set("logtostderr", "true")
	case LOG_OUTPUT_TYPE_FILE:
		flag.Set("logtostderr", "false")
	case LOG_OUTPUT_TYPE_SYSLOG:
		flag.Set("logtostderr", "false")
		logwriter, err := syslog.New(syslog.LOG_NOTICE, "libopenflow")
		if err != nil {
			return err
		}
		klog.SetOutput(logwriter)
	}
	return nil
}

func init() {
	klog.InitFlags(nil)
}
