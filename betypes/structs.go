package betypes

import "database/sql"

type BotMessage struct {
	Message struct {
		Message_id int
		From       struct {
			Username string
			Id       int
		}
		Chat struct {
			Id int
		}
		Date int
		Text string
	}
}

type SelectChatID struct {
	ID        int64
	Chat      int64
	Username  string
	Timestamp int64
	Message   sql.NullString
	Command   sql.NullString
}
