package game

type GameState struct {
	Boards     [9][9]string `json:"boards"`      // 9 inner boards, each with 9 cells
	OuterBoard [9]string    `json:"outer_board"` // Status of the 9 inner boards ("X", "O", "Draw", or "")
	Turn       string       `json:"turn"`        // "X" or "O"
	Winner     string       `json:"winner"`      // "X", "O", "Draw", or ""
	IsGameOver bool         `json:"is_game_over"`
	NextBoard  int          `json:"next_board"` // Index of valid board for next move (-1 for any)
}

func NewGameState() *GameState {
	var boards [9][9]string

	return &GameState{
		Boards:     boards,
		OuterBoard: [9]string{},
		Turn:       "X", // X starts
		NextBoard:  -1,  // First move can be anywhere
	}
}

func (g *GameState) MakeMove(outerIndex, innerIndex int, player string) bool {
	if g.IsGameOver || player != g.Turn {
		return false
	}
	if outerIndex < 0 || outerIndex >= 9 {
		return false
	}
	if innerIndex < 0 || innerIndex >= 9 {
		return false
	}

	if g.NextBoard != -1 && outerIndex != g.NextBoard {
		return false
	}

	if g.OuterBoard[outerIndex] != "" {
		return false
	}
	if g.Boards[outerIndex][innerIndex] != "" {
		return false
	}

	g.Boards[outerIndex][innerIndex] = player

	g.CheckInnerStatus(outerIndex)

	g.CheckGlobalStatus()

	if !g.IsGameOver {
		if g.Turn == "X" {
			g.Turn = "O"
		} else {
			g.Turn = "X"
		}

		nextTarget := innerIndex

		if g.OuterBoard[nextTarget] != "" {
			g.NextBoard = -1
		} else {
			g.NextBoard = nextTarget
		}
	}

	return true
}

func (g *GameState) CheckInnerStatus(boardIndex int) {
	board := g.Boards[boardIndex]
	if g.OuterBoard[boardIndex] != "" {
		return
	}

	winner := checkWin(board[:])
	if winner != "" {
		g.OuterBoard[boardIndex] = winner
		return
	}

	isFull := true
	for _, cell := range board {
		if cell == "" {
			isFull = false
			break
		}
	}
	if isFull {
		g.OuterBoard[boardIndex] = "Draw"
	}
}

func (g *GameState) CheckGlobalStatus() {
	winner := checkWin(g.OuterBoard[:])
	if winner != "" {
		g.Winner = winner
		g.IsGameOver = true
		return
	}

	isFull := true
	for _, status := range g.OuterBoard {
		if status == "" {
			isFull = false
			break
		}
	}
	if isFull {
		g.Winner = "Draw"
		g.IsGameOver = true
	}
}

func checkWin(board []string) string {
	lines := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, line := range lines {
		if board[line[0]] != "" &&
			board[line[0]] != "Draw" &&
			board[line[0]] == board[line[1]] &&
			board[line[1]] == board[line[2]] {
			return board[line[0]]
		}
	}
	return ""
}
