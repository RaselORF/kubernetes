/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logs

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/wait"
)

const logFlushFreqFlagName = "log-flush-frequency"

var logFlushFreq = pflag.Duration(logFlushFreqFlagName, 5*time.Second, "Maximum number of seconds between log flushes")

// AddFlags registers this package's flags on arbitrary FlagSets, such that they point to the
// same value as the global flags.
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(logFlushFreqFlagName))
}

// SysLogWriter serves as a bridge between the standard log package and the glog package.
type SysLogWriter struct{ w io.Writer }

// Write implements the io.Writer interface.
func (writer SysLogWriter) Write(data []byte) (n int, err error) {
	return writer.w.Write(data)
}

// InitLogs initializes logs the way we want for kubernetes.
func InitLogs() {
	w, _ := syslog.New(syslog.LOG_NOTICE, os.Args[0])
	errWriter := SysLogWriter{w: w}
	log.SetOutput(errWriter)
	log.SetFlags(0)
	glog.SetOutput(errWriter)
	// The default glog flush interval is 5 seconds.
	go wait.Forever(glog.Flush, *logFlushFreq)
}

// FlushLogs flushes logs immediately.
func FlushLogs() {
	glog.Flush()
}

// NewLogger creates a new log.Logger which sends logs to glog.Info.
func NewLogger(prefix string) *log.Logger {
	return log.New(SysLogWriter{}, prefix, 0)
}

// GlogSetter is a setter to set glog level.
func GlogSetter(val string) (string, error) {
	var level glog.Level
	if err := level.Set(val); err != nil {
		return "", fmt.Errorf("failed set glog.logging.verbosity %s: %v", val, err)
	}
	return fmt.Sprintf("successfully set glog.logging.verbosity to %s", val), nil
}
