package tools

import (
	"net/http"
	"strings"
)

func GetRealIp(req *http.Request) string {
	if ips := req.Header.Get("X-Forwarded-For"); ips != "" {
		if ip := CutProxyAddrs(ips); ip != "" {
			return ip
		}
	}
	return CutProxyAddrs(req.RemoteAddr)
}

func CutRemoteAddr(addr string) string {
	pieces := strings.Split(addr, ":")
	if len(pieces) > 0 && pieces[0] != "[" {
		return pieces[0]
	}
	return "127.0.0.1"
}

func CutProxyAddrs(addrs string) string {
	pieces := strings.Split(addrs, ",")
	if len(pieces) > 0 && pieces[0] != "" {
		return CutRemoteAddr(pieces[0])
	}
	return ""
}
