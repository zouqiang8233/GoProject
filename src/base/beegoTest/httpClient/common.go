package main

type MsgHeader struct {
	Version string `xml:"version,omitempty"`  // 版本号
	MsgCode string `xml:"msg_code,omitempty"` // 消息号
}
