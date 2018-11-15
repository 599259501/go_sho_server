package middares

import
(
	"github.com/gin-gonic/gin"
	"baselogic/login"
	"fmt"
	"baselogic"
)

func CheckMiniProgramLoginInfo()gin.HandlerFunc{
	return func(context *gin.Context) {
		loginHelper := login.NewMiniProgramLogin()
		sessionInfo := login.MiniProgramSessionInfo{
			UserId: context.GetInt("user_id"),
			AccessToken: context.DefaultPostForm("access_token", ""),
		}

		isLogin,err := loginHelper.CheckLogin(sessionInfo)
		if err!=nil || !isLogin{
			fmt.Println("CheckLogin() has err=",err)
			// todo 这里要返回登录态校验失败的错误信息
			baselogic.JResponse(context, login.CHEK_LOGIN_INFO_FAIL, nil, "登录态校验失败")
			return
		}
		context.Next()
	}
}
