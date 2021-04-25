package errcode

var (
	UserNotExist = &MyError{Status: 404, Code: "get.userinfo.fail", Message: "用户不存在"}

	DBQuery = &MyError{Status: 500, Code: "db.query.fail", Message: "数据库查询失败"}

	ErrParams = &MyError{Status: 400, Code: "err.params", Message: "参数错误"}

	ErrUnKnown = &MyError{Status: 500, Code: "err.unknown", Message: "未知错误"}
)
