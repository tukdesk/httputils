package gojimiddleware

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/qiniu/log"
)

var pid = uint32(os.Getpid())

func genReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

type xlogger struct {
	logger             *log.Logger
	reqId, reqIdLogStr string
}

func newXLogger(w http.ResponseWriter, r *http.Request) *xlogger {
	reqId := genReqId()
	w.Header().Set("X-Req-Id", reqId)
	return &xlogger{
		logger: log.Std,
		reqId:  reqId,
	}
}

func (this *xlogger) Print(v ...interface{}) {
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprint(v...))
}

func (this *xlogger) Println(v ...interface{}) {
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *xlogger) Printf(format string, v ...interface{}) {
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *xlogger) Debug(v ...interface{}) {
	if this.logger.Level > log.Ldebug {
		return
	}
	this.logger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintln(v...))
}

func (this *xlogger) Debugf(format string, v ...interface{}) {
	if this.logger.Level > log.Ldebug {
		return
	}
	this.logger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintf(format, v...))
}

func (this *xlogger) Info(v ...interface{}) {
	if this.logger.Level > log.Linfo {
		return
	}
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *xlogger) Infof(format string, v ...interface{}) {
	if this.logger.Level > log.Linfo {
		return
	}
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *xlogger) Warn(v ...interface{}) {
	if this.logger.Level > log.Lwarn {
		return
	}
	this.logger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintln(v...))
}

func (this *xlogger) Warnf(format string, v ...interface{}) {
	if this.logger.Level > log.Lwarn {
		return
	}
	this.logger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintf(format, v...))
}

func (this *xlogger) Error(v ...interface{}) {
	this.logger.Output(this.reqId, log.Lerror, 2, fmt.Sprintln(v...))
}

func (this *xlogger) Errorf(format string, v ...interface{}) {
	this.logger.Output(this.reqId, log.Lerror, 2, fmt.Sprintf(format, v...))
}
