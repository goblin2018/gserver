package api

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Dept     *Dept  `json:"dept"`
	Wechat   string `json:"wechat"`
	Password string `json:"password"`
	Desc     string `json:"desc"`
}

type ListUserOpt struct {
	*ListOpt
}
