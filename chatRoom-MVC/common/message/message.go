package message

const (
	LoginType       = "LoginMsg"
	LoginResMsgType = "LoginResMsg"
	RegisterType    = "RegisterMsg"
	RegisterResType = "RegisterResMsg"

	NotifyUserStatusMsgType = "NotifyUserStatusMsg"

	SmsMsgType = "SmsMsg"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// Message MAIN msg transfer
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// LoginMsg login message(from client to server)
type LoginMsg struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

// LoginResMsg login respond message(from server to client)
type LoginResMsg struct {
	Code    int `json:"code"` // 200/500/300/404
	UserIds []int
	Error   string `json:"error"`
}

// RegisterMsg
type RegisterMsg struct {
	User User `json:"user"`
}

// RegisterResMsg
type RegisterResMsg struct {
	Code  int    `json:"code"` // 200/500/300/404
	Error string `json:"error"`
}

type NotifyUserStatusMsg struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMsg struct {
	Content string        `json:"content"`
	User    `json:"user"` // niming struct
}
