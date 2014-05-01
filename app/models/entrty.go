package models

type User struct {
	Id       int    `db:"user_id"`
	NickName string `db:"nick_name"`
	Password string `db:"pass_word"`
	Email    string `db:"email"`
}

func (u *User) String() string {
	return "nickName:" + u.NickName + " password:" + u.Password + " email:" + u.Email
}

type ChatLog struct {
	Id           int    `db:"chat_log_id"`
	FromUserId   int    `db:"from_user_id"`
	FromUserName string `db:"from_user_name"`
	ToUserId     int    `db:"to_user_id"`
	ToUserName   string `db:"to_user_name"`
	MsgContent   string `db:"msg_content"`
	Read         bool   `db:"read"`
	ReadTime     int    `db:"read_time"`
	ToAll        bool   `db:"to_all"`
	SendTime     int    `db:"send_time"`
}
