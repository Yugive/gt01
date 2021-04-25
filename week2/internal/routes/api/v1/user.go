package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"week2/internal/dao"

	"week2/internal/errcode"
)

type UserApi struct{}

func NewUserApi() UserApi {
	return UserApi{}
}

func (u *UserApi) GetUserInfo(c *gin.Context) {

	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		errResp(c, errors.Wrapf(errcode.ErrParams, "bind params err: %+v", err))
		return
	}

	user, err := dao.GetUserById(userId)
	if err != nil {
		errResp(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "success", "message": "获取成功", "data": user})
}

func errResp(c *gin.Context, err error) {
	var code *errcode.MyError

	if !errors.As(err, &code) {
		code = errcode.ErrUnKnown
	}

	c.Error(err)
	c.JSON(code.Status, code)
}
