package verify_code

type ImageVerifyCode struct{
}

func NewImageVerifyVode()*ImageVerifyCode{
	return &ImageVerifyCode{}
}

func (tool *ImageVerifyCode)CheckVerifyCode(userId,code string)bool{
	return true
}
