package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RobUserSearch struct {
	request.PageInfo
}

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Verity   string `json:"verity"`
	VerityId string `json:"verityId"`
}
