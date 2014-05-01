package controllers

import (
	"chat/app/models"
	"chat/app/routes"
	"fmt"
	"github.com/revel/revel"
	"strconv"
)

type JsonStatus struct {
	Message string
	Success bool
}

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) Register() revel.Result {
	return c.Render()
}

func (c Application) ConfirmRegister(user models.User) revel.Result {
	b := models.InsertUser(&user)
	if b {
		c.Flash.Success("注册成功，请登录")
	}
	return c.Redirect(routes.Application.Index())
}

func (c Application) CheckEmailNotRepeat(user models.User) revel.Result {
	b := models.CheckEmailNotRepeat(user.Email)
	return c.RenderJson(b)
}

func (c Application) CheckNickNameNotRepeat(user models.User) revel.Result {
	b := models.CheckNickNameNotRepeat(user.NickName)
	return c.RenderJson(b)
}

func (c *Application) checkUser() revel.Result {
	user := c.user()
	if user != nil {
		c.RenderArgs["userO"] = user
	}
	return nil
}
func (c *Application) user() *models.User {
	//此处存在问题，每次都是c.RenderArgs["userO"] =nil 导致每次都要从数据库中查，后续从redis中查
	/*revel.INFO.Println(c.RenderArgs["userO"])
	if c.RenderArgs["userO"] != nil {
		return c.RenderArgs["userO"].(*models.User)
	}*/
	//session中放的是nickName,类似于java中的cookie
	if nickName, ok := c.Session["nickName"]; ok {
		return models.GetUser(nickName)
	}
	return nil
}

func (c *Application) Logout() revel.Result {
	c.RenderArgs["userO"] = nil
	c.Session["nickName"] = ""
	//session中放的是nickName
	if nickName, ok := c.Session["nickName"]; ok {
		fmt.Println("nickName__________" + nickName)
	}
	return c.Redirect(routes.Application.Index())
}

func (c Application) Enter(agaent string, user models.User) revel.Result {
	revel.INFO.Println(user)
	b := models.CheckLoginRight(&user)
	if !b {
		revel.INFO.Println("用户名或密码错误")
		c.Flash.Error("用户名或密码错误")
		return c.Redirect(routes.Application.Index())
	} else {
		c.Session["nickName"] = user.NickName
		c.Session["userId"] = strconv.Itoa(user.Id)
		revel.INFO.Println("login sucess")
	}
	switch agaent {
	case "refresh":
		return c.Redirect(routes.Refresh.Index())
	case "longpolling":
		return c.Redirect("/longpolling/room?user=%s", "李四")
	case "websocket":
		return c.Redirect("/websocket/room?user=%s", "王五")
	}
	return nil
}
