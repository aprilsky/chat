package controllers

import (
	"chat/app/chatroom"
	"chat/app/models"
	"fmt"
	"github.com/revel/revel"
)

type Refresh struct {
	Application
}

//first Join --> redirect --> Subscribe --> defer --> Cancle

func (c Refresh) Index() revel.Result {
	user, ok := c.RenderArgs["userO"].(*models.User)
	revel.INFO.Println(ok)
	if ok {
		OnlineList := chatroom.Join(*user)
		revel.INFO.Println(OnlineList)
		for i, _ := range OnlineList {
			fmt.Println(OnlineList[i])
		}
		return c.Render(OnlineList)
	}
	return nil

}
func (c Refresh) ChatLog(fromUserId, toUserId int) revel.Result {
	chatLogs := models.ListChatLog(fromUserId, toUserId)
	return c.RenderJson(chatLogs)
}
func (c Refresh) SendMessage(chatLog models.ChatLog) revel.Result {
	revel.INFO.Println(chatLog)
	models.InsertChatLog(&chatLog)
	return nil
}

/*func (c Refresh) Room(uId int) revel.Result {
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()
	//查询出来当前所有的消息
	events := subscription.Archive
	for i, _ := range events {
		if events[i].UId == uId {
			events[i].UId = 0
		}
	}
	return c.Render()
}*/

func (c Refresh) Leave(uId int) revel.Result {
	//chatroom.Leave(uId)
	return c.Redirect(Application.Index)
}
