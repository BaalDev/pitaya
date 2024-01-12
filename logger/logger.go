// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logger

import (
	// "fmt"
	// "path"
	// "runtime"

	"github.com/sirupsen/logrus"
	"github.com/BaalDev/pitaya/v2/logger/interfaces"
	logruswrapper "github.com/BaalDev/pitaya/v2/logger/logrus"
)

// Log is the default logger
var Log = initLogger()

func initLogger() interfaces.Logger {
	plog := logrus.New()
	// plog.SetReportCaller(true)
	plog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		// CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		// 	filename := path.Base(f.File)
		// 	return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		// },
	})
	plog.Level = logrus.DebugLevel

	// log := plog.WithFields(logrus.Fields{
	// 	"source": "pitaya",
	// })
	log := plog.WithFields(nil)
	return logruswrapper.NewWithFieldLogger(log)
}

// SetLogger rewrites the default logger
func SetLogger(l interfaces.Logger) {
	if l != nil {
		Log = l
	}
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	Log.Fatalln(args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	Log.Debugln(args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	Log.Errorln(args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	Log.Infoln(args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	Log.Warnln(args...)
}

func Panicf(format string, args ...interface{}) {
	Log.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	Log.Panicln(args...)
}
