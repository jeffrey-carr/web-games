package main

import (
	"web_games/binoku"
	wordchain "web_games/word_chain"
)

// DependencyContainer contains all server dependencies for injection
type DependencyContainer struct {
	BinokuController     binoku.GameManager
	WordLadderController wordchain.Controller
}
