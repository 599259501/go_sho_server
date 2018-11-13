package verify_code

type IBaseVerifyCode interface{
	CheckVerifyCode(userId,code string)bool
}