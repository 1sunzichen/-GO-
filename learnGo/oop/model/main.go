package model
type userinfo struct {
	Name string
	Age int
	Height float32
	Eduschool string
	Hobby []string
	MoreInfo map[string]interface{}
}
func NewUserInfo(name string,age int,height float32,
	eduschool string,hobby []string,moreinfo map[string] interface{}) *userinfo{
	return &userinfo{
		Name:name,
		Age:age,
		Height:height,
		Eduschool: eduschool,
		Hobby:hobby,
		MoreInfo: moreinfo,

	}
}