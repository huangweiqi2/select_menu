package errs

import "errors"

var AuthErr = errors.New("用户名或密码错误")
var NoArgumentErr = errors.New("传入参数为空")
var PhoneExist = errors.New("手机号已存在")
var CreatModelErr = errors.New("creat model error")
var UpdateModelErr = errors.New("update model error")
var DeleteModelErr = errors.New("delete model error")
var GenerateTokenErr = errors.New("generate token error")
