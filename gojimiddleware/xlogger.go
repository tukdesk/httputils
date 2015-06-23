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

var logger = log.New(os.Stderr, "", log.LstdFlags|log.Llevel)

var pid = uint32(os.Getpid())

func genReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

type XLogger struct {
	logger             *log.Logger
	reqId, reqIdLogStr string
}

func newXLogger(w http.ResponseWriter, r *http.Request) *XLogger {
	reqId := genReqId()
	w.Header().Set("X-Req-Id", reqId)
	return &XLogger{
		logger: logger,
		reqId:  reqId,
	}
}

func (this *XLogger) Print(v ...interface{}) {
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprint(v...))
}

func (this *XLogger) Println(v ...interface{}) {
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Printf(format string, v ...interface{}) {
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Debug(v ...interface{}) {
	if this.logger.Level > log.Ldebug {
		return
	}
	this.logger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Debugf(format string, v ...interface{}) {
	if this.logger.Level > log.Ldebug {
		return
	}
	this.logger.Output(this.reqId, log.Ldebug, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Info(v ...interface{}) {
	if this.logger.Level > log.Linfo {
		return
	}
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Infof(format string, v ...interface{}) {
	if this.logger.Level > log.Linfo {
		return
	}
	this.logger.Output(this.reqId, log.Linfo, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Warn(v ...interface{}) {
	if this.logger.Level > log.Lwarn {
		return
	}
	this.logger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Warnf(format string, v ...interface{}) {
	if this.logger.Level > log.Lwarn {
		return
	}
	this.logger.Output(this.reqId, log.Lwarn, 2, fmt.Sprintf(format, v...))
}

func (this *XLogger) Error(v ...interface{}) {
	this.logger.Output(this.reqId, log.Lerror, 2, fmt.Sprintln(v...))
}

func (this *XLogger) Errorf(format string, v ...interface{}) {
	this.logger.Output(this.reqId, log.Lerror, 2, fmt.Sprintf(format, v...))
}
