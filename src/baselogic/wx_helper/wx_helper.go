package wx_helper

import (
	"github.com/imroc/req"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"utils"
)
type WxSessionInfo struct{
	OpenId string `json:"open_id"`
	SessionKey string `json:"session_key"`
	UnionId string `json:"union_id"`
}

type WxHelper struct{
	AppId string
	Secert string
}
func NewWxHelper()*WxHelper{
	return &WxHelper{
		AppId: utils.GetEnv("APP_ID", ""),
		Secert: utils.GetEnv("APP_SECERT", ""),
	}
}
func (helper *WxHelper)GetWxMiniSession(code string)(WxSessionInfo,error){
	js2SessionUrl := fmt.Sprintf(JSSESSION_URL,helper.AppId, helper.Secert, code)

	session := WxSessionInfo{}
	rsp, err :=req.New().Get(js2SessionUrl)
	if err!=nil{
		fmt.Println("req url=",js2SessionUrl,",err=",err)
		return session,nil
	}

	js, _ := simplejson.NewJson(rsp.Bytes())
	if js.Get("errcode").MustInt() != 0 {
		return session,errors.New(js.Get("errmsg").MustString())
	}

	session.OpenId = js.Get("openid").MustString()
	session.SessionKey = js.Get("session_key").MustString()
	session.UnionId = js.Get("unionid").MustString()
	return session,err
}
