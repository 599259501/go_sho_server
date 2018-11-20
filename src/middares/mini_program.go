package middares

import
(
	"github.com/gin-gonic/gin"
	"baselogic/login"
	"fmt"
	"baselogic"
	"strconv"
)

func CheckMiniProgramLoginInfo()gin.HandlerFunc{
	return func(context *gin.Context) {
		userId,_ := strconv.ParseInt(context.DefaultPostForm("user_id", ""), 10, 64)
		loginHelper := login.NewMiniProgramLogin()
		sessionInfo := login.MiniProgramSessionInfo{
			UserId: int(userId),
			AccessToken: context.DefaultPostForm("access_token", ""),
		}

		isLogin,err := loginHelper.CheckLogin(sessionInfo)
		if err!=nil || !isLogin{
			fmt.Println("CheckLogin() has err=",err)
			// todo 这里要返回登录态校验失败的错误信息
			baselogic.JResponse(context, login.CHEK_LOGIN_INFO_FAIL, nil, "登录态校验失败")
			context.Abort()
			return
		}
		context.Next()
	}
}
