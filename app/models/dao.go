package models

import (
	r "github.com/revel/revel"
)

func CheckNickNameNotRepeat(nickName string) bool {
	if nickName == "" {
		return false
	}
	id, err := Dbm.SelectStr("select user_id from t_user where nick_name = ?", nickName)
	checkErr(err, "select error")
	r.INFO.Println(id)
	if id != "" {
		return false
	}
	return true
}

func CheckEmailNotRepeat(email string) bool {
	if email == "" {
		return false
	}
	id, err := Dbm.SelectStr("select user_id from t_user where email = ?", email)
	checkErr(err, "select error")
	r.INFO.Println(id)
	if id != "" {
		return false
	}
	return true
}

func InsertUser(user *User) bool {
	r.INFO.Println(user)
	err := Dbm.Insert(user)
	if err != nil {
		checkErr(err, "insert user failed")
		return false
	}
	return true
}

func InsertChatLog(chatLog *ChatLog) bool {
	r.INFO.Println(chatLog)
	err := Dbm.Insert(chatLog)
	if err != nil {
		checkErr(err, "chatLog user failed")
		return false
	}
	return true
}
func GetUser(nickName string) *User {
	if nickName == "" {
		return nil
	}
	user := new(User)
	Dbm.SelectOne(user, "select * from t_user where nick_name = ?", nickName)
	return user
}

func CheckLoginRight(user *User) bool {
	if user == nil || user.Email == "" || user.Password == "" {
		return false
	}
	err := Dbm.SelectOne(user, "select * from t_user where (email = ? or nick_name= ?) and pass_word = ? ", user.Email, user.Email, user.Password)
	if err != nil {
		return false
	}
	return true
}

//
func ListChatLog(fromUserId, toUserId int) []ChatLog {
	chatLog := make([]ChatLog, 0)
	_, err := Dbm.Select(&chatLog, "select * from t_chat_log where (from_user_id = ? or to_user_id = ?) and (to_user_id = ? or from_user_id = ?)", fromUserId, fromUserId, toUserId, toUserId)
	if err != nil {
		checkErr(err, "select chatLog failed")
	}
	r.INFO.Println(chatLog)
	return chatLog

}
