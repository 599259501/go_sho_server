package login

type MiniProgramLogin struct{}
func NewMiniProgramLogin()*MiniProgramLogin{
	return &MiniProgramLogin{}
}
func (loginService *MiniProgramLogin)CheckLogin(loginInfo interface{})(bool,error){
	return true,nil
}
