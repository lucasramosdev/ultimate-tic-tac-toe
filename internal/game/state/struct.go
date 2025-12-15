package state

type GameState struct {
	Boards     [9][9]string `json:"boards"`
	OuterBoard [9]string    `json:"outer_board"`
	Turn       string       `json:"turn"`
	Winner     string       `json:"winner"`
	IsGameOver bool         `json:"is_game_over"`
	NextBoard  int          `json:"next_board"`
}
