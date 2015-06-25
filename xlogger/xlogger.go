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
	nLogger            *log.Logger
	sLogger            *log.Logger
	reqId, reqIdLogStr string
}

func NewXLogger(w http.ResponseWriter, r *http.Request) *XLogger {
	reqId := genReqId()
	w.Header().Set("X-Req-Id", reqId)
	return &XLogger{
		nLogger: normalLogger,
		sLogger: simpleLogger,
		reqId:   reqId,
	}
}

func (this *XLogger) Print(v ...interface{}) {
	this.nLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprint(v...))
}

func (this *XLogger) Println(v ...interface{}) {
	this.nLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Printf(format string, v ...interface{}) {
	this.nLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Debug(v ...interface{}) {
	if this.nLogger.Level > log.Ldebug {
		return
	}
	this.nLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Debugf(format string, v ...interface{}) {
	if this.nLogger.Level > log.Ldebug {
		return
	}
	this.nLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Info(v ...interface{}) {
	if this.nLogger.Level > log.Linfo {
		return
	}
	this.nLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Infof(format string, v ...interface{}) {
	if this.nLogger.Level > log.Linfo {
		return
	}
	this.nLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Warn(v ...interface{}) {
	if this.nLogger.Level > log.Lwarn {
		return
	}
	this.nLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Warnf(format string, v ...interface{}) {
	if this.nLogger.Level > log.Lwarn {
		return
	}
	this.nLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Error(v ...interface{}) {
	this.nLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Errorf(format string, v ...interface{}) {
	this.nLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimplePrint(v ...interface{}) {
	this.sLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprint(v...))
}

func (this *XLogger) SimplePrintln(v ...interface{}) {
	this.sLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimplePrintf(format string, v ...interface{}) {
	this.sLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleDebug(v ...interface{}) {
	if this.sLogger.Level > log.Ldebug {
		return
	}
	this.sLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleDebugf(format string, v ...interface{}) {
	if this.sLogger.Level > log.Ldebug {
		return
	}
	this.sLogger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleInfo(v ...interface{}) {
	if this.sLogger.Level > log.Linfo {
		return
	}
	this.sLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleInfof(format string, v ...interface{}) {
	if this.sLogger.Level > log.Linfo {
		return
	}
	this.sLogger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleWarn(v ...interface{}) {
	if this.sLogger.Level > log.Lwarn {
		return
	}
	this.sLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleWarnf(format string, v ...interface{}) {
	if this.sLogger.Level > log.Lwarn {
		return
	}
	this.sLogger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) SimpleError(v ...interface{}) {
	this.sLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintln(v...))
}

func (this *XLogger) SimpleErrorf(format string, v ...interface{}) {
	this.sLogger.Output(this.reqId, log.Lerror, 2, fmt.Sprintf(format, v...))
}
