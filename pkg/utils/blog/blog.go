/**
func main() {
	dir, _ := os.Getwd()

	cfg := conf.LogConfig{
		LogDir:       dir + "/blog",
		LogMaxSize:   500,
		ToStdErr:     true,
		AlsoToStdErr: true,
		Verbosity:    6,
	}
	blog.InitLogs(cfg)
	...
}
*/
package blog

import (
	"fmt"
	"log"
	"regexp"
	"sync"
	"time"

	"go-netdisk/pkg/utils/blog/glog"
)

// GlogWriter serves as a bridge between the standard log package and the glog package.
type GlogWriter struct{}

// Write implements the io.Writer interface.
func (writer GlogWriter) Write(data []byte) (n int, err error) {
	glog.Info(string(data))
	return len(data), nil
}

var once sync.Once

// Log configuration
type LogConfig struct {
	LogDir     string `json:"log_dir" value:"./logs" usage:"If non-empty, write log files in this directory" mapstructure:"log_dir"`
	LogMaxSize uint64 `json:"log_max_size" value:"500" usage:"Max size (MB) per log file." mapstructure:"log_max_size"`
	LogMaxNum  int    `json:"log_max_num" value:"10" usage:"Max num of log file. The oldest will be removed if there is a extra file created." mapstructure:"log_max_num"`

	ToStdErr        bool   `json:"logtostderr" value:"false" usage:"log to standard error instead of files" mapstructure:"logtostderr"`
	AlsoToStdErr    bool   `json:"alsologtostderr" value:"false" usage:"log to standard error as well as files" mapstructure:"alsologtostderr"`
	Verbosity       int32  `json:"v" value:"0" usage:"log level for V logs" mapstructure:"v"`
	StdErrThreshold string `json:"stderrthreshold" value:"2" usage:"logs at or above this threshold go to stderr" mapstructure:"stderrthreshold"`
	VModule         string `json:"vmodule" value:"" usage:"comma-separated list of pattern=N settings for file-filtered logging" mapstructure:"vmodule"`
	TraceLocation   string `json:"log_backtrace_at" value:"" usage:"when logging hits line file:N, emit a stack trace" mapstructure:"log_backtrace_at"`
}

// InitLogs initializes logs the way we want for blog.
func InitLogs(logConfig LogConfig) {
	glog.InitLogs(logConfig.ToStdErr,
		logConfig.AlsoToStdErr,
		logConfig.Verbosity,
		logConfig.StdErrThreshold,
		logConfig.VModule,
		logConfig.TraceLocation,
		logConfig.LogDir,
		logConfig.LogMaxSize,
		logConfig.LogMaxNum,
	)
	once.Do(func() {
		log.SetOutput(GlogWriter{})
		log.SetFlags(0)
		// The default glog flush interval is 30 seconds, which is frighteningly long.
		go func() {
			d := time.Duration(5 * time.Second)
			tick := time.Tick(d)
			for {
				select {
				case <-tick:
					glog.Flush()
				}
			}
		}()
	})
}

func CloseLogs() {
	glog.Flush()
}

var (
	Info  = glog.Infof
	Infof = glog.Infof

	Warn  = glog.Warningf
	Warnf = glog.Warningf

	Error  = glog.Errorf
	Errorf = glog.Errorf

	Fatal  = glog.Fatal
	Fatalf = glog.Fatalf

	V = glog.V
)

func Debug(args ...interface{}) {
	if format, ok := (args[0]).(string); ok {
		glog.V(3).Infof(format, args[1:]...)
	} else {
		glog.V(3).Info(args...)
	}
}

func SetV(level int32) {
	glog.SetV(glog.Level(level))
}

// defaultRe and defaultHandler is for bcs-dns wrap its extra time tag in log.
// the extra time tag of bcs-dns: [04/Jan/2018:09:44:27 +0800]
var defaultRe = regexp.MustCompile(`\[\d{2}/\w+/\d{4}:\d{2}:\d{2}:\d{2} \+\d{4}\] `)
var defaultHandler WrapFunc = func(format string, args ...interface{}) string {
	src := fmt.Sprintf(format, args...)
	return defaultRe.ReplaceAllString(src, "")
}

// WrapFunc take the param the same as glog.Infof, and return string.
type WrapFunc func(string, ...interface{}) string

// Wrapper use WrapFunc to handle the log message before send it to glog.
// Can be use as:
//      var handler blog.WrapFunc = func(format string, args ...interface{}) string {
//          src := fmt.Sprintf(format, args...)
//          dst := regexp.MustCompile("boy").ReplaceAllString(src, "man")
//      }
//      blog.Wrapper(handler).V(2).Info("hello boy")
// And it will flush as:
//      I0104 09:44:27.796409   16233 blog.go:21] hello man
type Wrapper struct {
	Handler WrapFunc
	verbose glog.Verbose
}

// Info implementation
func (w *Wrapper) Info(format string, args ...interface{}) {
	if w.verbose {
		Info(w.Handler(format, args...))
	}
}

// Warn implementation
func (w *Wrapper) Warn(format string, args ...interface{}) {
	if w.verbose {
		Warn(w.Handler(format, args...))
	}
}

// Error implementation
func (w *Wrapper) Error(format string, args ...interface{}) {
	if w.verbose {
		Error(w.Handler(format, args...))
	}
}

// Fatal implementation
func (w *Wrapper) Fatal(format string, args ...interface{}) {
	if w.verbose {
		Fatal(w.Handler(format, args...))
	}
}

// V implementation
func (w *Wrapper) V(level glog.Level) *Wrapper {
	w.verbose = V(level)
	return w
}

// Wrap Wrapper function
func Wrap(handler WrapFunc) *Wrapper {
	if handler == nil {
		handler = defaultHandler
	}
	return &Wrapper{verbose: true, Handler: handler}
}
