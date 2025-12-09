package game

type PlayRequest struct {
	Nickname string `form:"nickname"`
	RoomCode string `form:"room_code"`
}