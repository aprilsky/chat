// chatroom
package chatroom

import (
	"chat/app/models"
	//"container/list"
)

var (
	//当前的在线人员
	users = make([]models.User, 0)
)

func Join(user models.User) []models.User {
	if user.Id == 0 {
		return users
	}
	for _, u := range users {
		if u == user {
			return users
		}
	}
	users = append(users, user)
	return users
}
