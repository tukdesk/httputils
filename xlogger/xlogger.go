package xlogger

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/qiniu/log"
)

var (
	normalLogger = log.New(os.Stderr, "", log.LstdFlags|log.Llevel|log.Lshortfile|log.Lmodule)
	simpleLogger = log.New(os.Stderr, "", log.LstdFlags|log.Llevel)
)

func SetNormalLogger(l *log.Logger) {
	if l == nil {
		return
	}
	normalLogger = l
}

func SetSimpleLogger(l *log.Logger) {
	if l == nil {
		return
	}
	simpleLogger = l
}

var pid = uint32(os.Getpid())

func genReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

type XLogger struct {
	reqId string
	r     *http.Request
}

func NewXLogger(w http.ResponseWriter, r *http.Request) *XLogger {
	reqId := genReqId()
	w.Header().Set("X-Req-Id", reqId)
	return &XLogger{
		reqId: reqId,
		r:     r,
	}
}

func (this *XLogger) Print(v ...interface{}) {
	normalLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprint(v...))
}

func (this *XLogger) Println(v ...interface{}) {
	normalLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Printf(format string, v ...interface{}) {
	normalLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Debug(v ...interface{}) {
	if normalLogger.Level > log.Ldebug {
		return
	}
	normalLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Debugf(format string, v ...interface{}) {
	if normalLogger.Level > log.Ldebug {
		return
	}
	normalLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Info(v ...interface{}) {
	if normalLogger.Level > log.Linfo {
		return
	}
	normalLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Infof(format string, v ...interface{}) {
	if normalLogger.Level > log.Linfo {
		return
	}
	normalLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Warn(v ...interface{}) {
	if normalLogger.Level > log.Lwarn {
		return
	}
	normalLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Warnf(format string, v ...interface{}) {
	if normalLogger.Level > log.Lwarn {
		return
	}
	normalLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Error(v ...interface{}) {
	normalLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Errorf(format string, v ...interface{}) {
	normalLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimplePrint(v ...interface{}) {
	simpleLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprint(v...))
}

func (this *XLogger) SimplePrintln(v ...interface{}) {
	simpleLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimplePrintf(format string, v ...interface{}) {
	simpleLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleDebug(v ...interface{}) {
	if simpleLogger.Level > log.Ldebug {
		return
	}
	simpleLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleDebugf(format string, v ...interface{}) {
	if simpleLogger.Level > log.Ldebug {
		return
	}
	simpleLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleInfo(v ...interface{}) {
	if simpleLogger.Level > log.Linfo {
		return
	}
	simpleLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleInfof(format string, v ...interface{}) {
	if simpleLogger.Level > log.Linfo {
		return
	}
	simpleLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleWarn(v ...interface{}) {
	if simpleLogger.Level > log.Lwarn {
		return
	}
	simpleLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleWarnf(format string, v ...interface{}) {
	if simpleLogger.Level > log.Lwarn {
		return
	}
	simpleLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleError(v ...interface{}) {
	simpleLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleErrorf(format string, v ...interface{}) {
	simpleLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Request() *http.Request {
	return this.r
}

func (this *XLogger) ResetForRequest(w http.ResponseWriter, r *http.Request) {
	if this.r == r {
		return
	}

	reqId := genReqId()
	w.Header().Set("X-Req-Id", reqId)
	this.reqId = reqId
	this.r = r
	return
}
