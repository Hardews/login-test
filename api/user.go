package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"login-test/middleware"
	"login-test/model"
	"login-test/service"
	"login-test/tool"
)

func Login(ctx *gin.Context) {
	var user model.User
	var res bool
	user.Username, res = ctx.GetPostForm("username")
	if !res {
		tool.RespErrorWithDate(ctx, "输入的账号为空")
		return
	}

	user.Password, res = ctx.GetPostForm("password")
	if !res {
		tool.RespErrorWithDate(ctx, "输入的密码为空")
		return
	}
	err, flag := service.CheckPassword(user.Username, user.Password)
	if err != nil {
		tool.RespInternetError(ctx)
		fmt.Println("check password failed,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(ctx, "密码错误")
		return
	} else {
		token, flag := middleware.SetToken(user.Username, "")
		if !flag {
			tool.RespInternetError(ctx)
			return
		}
		tool.RespSuccessfulWithDate(ctx, token)
	}
}
func Register(ctx *gin.Context) {
	var user model.User
	user.Username, _ = ctx.GetPostForm("username")
	user.Password, _ = ctx.GetPostForm("password")
	user.Nickname, _ = ctx.GetPostForm("nickName")
	if user.Username == "" {
		tool.RespErrorWithDate(ctx, "用户名为空")
		return
	}
	if user.Password == "" {
		tool.RespErrorWithDate(ctx, "密码为空")
		return
	}
	err, flag := service.CheckUsername(user)
	if err != nil {
		tool.RespInternetError(ctx)
		fmt.Println("check username failed,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(ctx, "账号已存在")
		return
	}

	err, user.Password = service.Encryption(user.Password)
	if err != nil {
		err = errors.New("internet error")
		fmt.Println(err)
		tool.RespErrorWithDate(ctx, err)
		return
	}

	err = service.WriteIn(user)
	if err != nil {
		tool.RespInternetError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
