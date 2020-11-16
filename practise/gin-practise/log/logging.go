package log

type Configuration struct {
	LogFile  string
	LogLevel string

	RotateMaxSize    int
	RotateMaxAge     int
	RotateMaxBackups int
	Compress         bool
}

type Logger interface {
	Info(args ...interface{})
	Infof(f string, args ...interface{})
}

var Glog Logger
var AccessLog Logger

func init() {
	fileLog := "/Users/caoyuan/workstation/golang-learning/practise/gin-practise/ginlog.log"
	Glog, _ = NewZapLogger(Configuration{
		LogFile:          fileLog,
		LogLevel:         "INFO",
		RotateMaxSize:    500,
		RotateMaxAge:     7,
		RotateMaxBackups: 3,
	})

	afileLog := "/Users/caoyuan/workstation/golang-learning/practise/gin-practise/gin-access.log"
	AccessLog, _ = NewZapLogger(Configuration{
		LogFile:          afileLog,
		LogLevel:         "INFO",
		RotateMaxSize:    500,
		RotateMaxAge:     7,
		RotateMaxBackups: 3,
	})
}
