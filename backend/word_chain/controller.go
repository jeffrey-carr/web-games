package wordchain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"web_games/entities"
	"web_games/services"
	"web_games/utils"
)

// ErrInvalidPosition is returned when a player submits an invalid position
var ErrInvalidPosition = errors.New("Invalid position")

// TargetChainLength is the targeted length of each chain
// TODO - make this a game option rather than hard-coded
const TargetChainLength = 7

// Controller represents a WordLadder controller
type Controller interface {
	CreateGame(ctx context.Context) (Game, error)
	ValidateGuess(guess string, game Game) (bool, Game, error)
}

type controller struct {
	dictionary   Dictionary
	runningGames entities.AsyncMap[string, Game]

	encryption services.Encryption
}

// NewController creates a new word ladder Controller
func NewController(
	maxServers int,
	maxPlayersPerServer int,
	dictionary Dictionary,
	encryption services.Encryption,
) Controller {
	return controller{
		dictionary:   dictionary,
		runningGames: entities.NewAsyncMap(map[string]Game{}),
		encryption:   encryption,
	}
}

func (c controller) CreateMultiplayerLobby(ctx context.Context) (Game, error) {
	// TODO
	return Game{}, errors.New("not implemented")
}

func (c controller) CreateGame(ctx context.Context) (Game, error) {
	chain := c.generateLadder(entities.NewSetStack[string]())
	state := GameState{
		GeneratedChain: chain.Slice(),
		UserProgress:   1,
	}
	encryptedState, err := c.encryption.Encrypt(state)
	if err != nil {
		return Game{}, err
	}

	game := Game{
		GameState:      state,
		UUID:           utils.NewUUIDString(),
		EncryptedState: encryptedState,
	}

	fmt.Printf("Created game: %v\n", game)
	return game, nil
}

func (c controller) generateLobbyCode() string {
	// This is very brute-force, but we should only ever have like 2 activate
	// games so this isn't really a concern
	lobbyCode := ""
	exists := true
	for exists {
		lobbyCode = NewLobbyCode()
		_, exists = c.runningGames.Get(lobbyCode)
	}

	return lobbyCode
}

// generate ladder generates a feasibly solvable start and end word
// we love recursion here
func (c controller) generateLadder(currentLadder entities.SetStack[string]) entities.SetStack[string] {
	if currentLadder.Size() == 0 {
		word := utils.GetRandomItem(utils.GetKeys(c.dictionary))
		currentLadder.Push(word)
	}

	// Base case
	if currentLadder.Size() == TargetChainLength {
		return currentLadder
	}

	// Using the last word, find a word that works with the ladder
	previousWord, _ := currentLadder.Peek()
	possibleWords, ok := c.dictionary[previousWord]
	// It's possible that the chosen word does not have any associated words
	// If that happens, move on.
	// TODO - it's possible that this could create an infinite loop, we should fix that
	if !ok {
		currentLadder.Pop()
		return c.generateLadder(currentLadder)
	}

	nextWord := utils.GetRandomItem(possibleWords)
	currentLadder.Push(nextWord)
	return c.generateLadder(currentLadder)
}

func (c controller) ValidateGuess(guess string, game Game) (bool, Game, error) {
	var state GameState
	err := c.encryption.Decrypt(game.EncryptedState, &state)
	if err != nil {
		return false, game, err
	}

	if !strings.EqualFold(guess, state.GeneratedChain[state.UserProgress]) {
		return false, game, nil
	}

	updatedData := GameState{
		GeneratedChain: state.GeneratedChain,
		UserProgress:   state.UserProgress + 1,
	}
	encryptedData, err := c.encryption.Encrypt(updatedData)
	if err != nil {
		return false, game, err
	}
	updatedGame := Game{
		GameState:      updatedData,
		UUID:           game.UUID, // TODO - maybe encrypt this too?
		EncryptedState: encryptedData,
	}

	return true, updatedGame, nil
}
