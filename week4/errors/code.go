package errors
var (

DBQuery = &ErrorCode{Status: 500, Code: "db.query.fail", Message: "数据库查询失败"}

ErrParams = &ErrorCode{Status: 400, Code: "err.params", Message: "参数错误"}

ErrUnKnown = &ErrorCode{Status: 500, Code: "err.unknown", Message: "未知错误"}

UserNotFound = &ErrorCode{Status: 404, Code: "err.unknown", Message: "用户不存在"}
)