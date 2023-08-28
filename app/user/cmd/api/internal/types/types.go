// Code generated by goctl. DO NOT EDIT.
package types

type GetMyInfoReq struct {
	Authorization string `header:"Authorization"`
}

type GetMyInfoResp struct {
	Avatar    string `json:"avatar"`
	NickName  string `json:"nickname"`
	Name      string `json:"name"`
	School    string `json:"school"`
	Group     string `json:"group"`
	Email     string `json:"email"`
	StudentID string `json:"student_id"`
	QQ        string `json:"qq"`
}

type GetUserInfoReq struct {
	Authorization string `header:"Authorization"`
	UserId        string `path:"id"`
}

type GetUserInfoResp struct {
	Avatar    string `json:"avatar"`
	NickName  string `json:"nickname"`
	Name      string `json:"name"`
	School    string `json:"school"`
	Group     string `json:"group"`
	Email     string `json:"email"`
	StudentID string `json:"student_id"`
	QQ        string `json:"qq"`
}

type SetUserInfoReq struct {
	Authorization string `header:"Authorization"`
	Avatar        string `json:"avatar"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	School        string `json:"school"`
	QQ            string `json:"qq"`
}

type SetUserInfoResp struct {
	Flag bool `json:"flag"`
}

type SetUserTypeReq struct {
	Authorization string `header:"Authorization"`
	Email         string `json:"email"`
	UserType      string `json:"user_type,options=[freshman,normal,admin,super_admin]"`
}

type SetUserTypeResp struct {
	Flag bool `json:"flag"`
}

type GetAdminListReq struct {
	Authorization string `header:"Authorization"`
	UserType      string `form:"user_type,options=[super_admin,admin,normal]"`
}

type One struct {
	UserId   string `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type GetAdminListResp struct {
	List []One `json:"list"`
}