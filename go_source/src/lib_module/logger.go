package lib_module

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/robfig/cron/v3"
)

type custom_log_TextFormatter struct {
	logrus.TextFormatter
}

//logrus  : https://codesk.tistory.com/155
//crontab : https://programmer.ink/think/cron-of-go-daily.html

func (f *custom_log_TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var ct_loglevel string
	var ct_funcname string
	var ct_logtime string
	var ct_logmsg string
	var log_msg_res []byte

	const (
		//Entry Struct Level Const
		PanicLevel = 0
		InfoLevel  = 4
		TraceLevel = 6
	)

	if entry != nil {
		if entry.Caller != nil {
			c_func := strings.Split(entry.Caller.Function, ".")
			ct_funcname = c_func[len(c_func)-1]
		} else {
			ct_funcname = "Empty Entry Caller !!! , Check SetReportCaller(true) "
		}

		if 0 <= len(f.TimestampFormat) {
			ct_logtime = entry.Time.Format(f.TimestampFormat)
		} else {
			ct_logtime = "Empty TimestampFormat !!!  , Check TimestampFormat in TextFormatter struc  t"
		}

		if 0 <= len(entry.Message) {
			ct_logmsg = entry.Message
		} else {
			ct_logmsg = "Empty Entry Message !!!"
		}

		if PanicLevel <= entry.Level && TraceLevel <= 6 {
			//ct_loglevel := entry.Level.String()
			ct_loglevel = strings.ToUpper(entry.Level.String())

			if InfoLevel == entry.Level {
				log_msg_res = []byte(fmt.Sprintf("[ %7s :: %s ] ::: %s\n", ct_loglevel, ct_logtime, ct_logmsg))
			} else {
				log_msg_res = []byte(fmt.Sprintf("[ %7s :: %s :: %s ] ::: %s\n", ct_loglevel, ct_logtime, ct_funcname, ct_logmsg))
			}

		} else {
			ct_loglevel = "Empty Entry LogLevel !!!"
		}

	}

	return log_msg_res, nil
}

func logger_init(out_flag bool, log_path string) (*logrus.Logger, *lumberjack.Logger) {

	lum := &lumberjack.Logger{
		Filename:   log_path,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	local_logger := &logrus.Logger{
		Level: logrus.InfoLevel,
		Formatter: &custom_log_TextFormatter{logrus.TextFormatter{
			FullTimestamp:          false,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
		},
	}

	if !out_flag {
		multi := io.MultiWriter(os.Stdout)
		local_logger.SetOutput(multi)
	} else {
		multi := io.MultiWriter(lum, os.Stdout)
		local_logger.SetOutput(multi)
	}

	local_logger.SetReportCaller(true)

	return local_logger, lum

}

func GetLogInstance(out_flag bool, log_path string) *logrus.Logger {

	log_inst_once.Do(func() {

		Logger, Lum = logger_init(out_flag, log_path)

		log_crontab := cron.New()
		log_crontab.AddFunc("@daily", func() {
			Lum.Rotate()
		})
		log_crontab.Start()

	})

	return Logger
}
