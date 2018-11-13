package wx_helper

import (
	"github.com/imroc"
)
type WxSessionInfo struct{
}

type WxHelper struct{}
func NewWxHelper()*WxHelper{
	return &WxHelper{}
}
func (helper *WxHelper)GetWxMiniSession(code string)(WxSessionInfo,error){
	return WxSessionInfo{},nil
}
