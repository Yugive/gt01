### 题目

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码

### 解答

应该wrap这个错误，可以将业务层关系的业务错误信息包含进去，用户不需要关注底层的错误问题，只需要将包含的业务错误原因返回给用户即可，清晰明了，
此外，可以通过相关的堆栈信息具体定位到相关的错误根因，可以进一步查找问题所在

### 目录结构

```bash
.
├── main.go
└── internal
    ├── dao
    │   └── user.go
    ├── model
    │   └── mysql.go
    │   └── user.go
    │── errcode
    │   └── code.go
    │   └── error.go
    └── routes
        └── route.go
        └── api
            └── v1
                └── user.go
```

### 运行

```bash
go run main.go
```