export enum Event {
  LOBBY_CODE = "lobby-code",
}

export type LobbyCodeEvent = {
  code: string;
};

export type Chain = string[];
export type WordChainState = {
  generatedChain: Chain;
  userProgress: number;
};
export type WordChainGame = WordChainState & {
  uuid: string;
  encryptedState: string;
};

export type ValidateAnswerRequest = {
  guess: string;
  gameState: WordChainGame;
};
export type ValidateAnswerResponse = {
  correct: boolean;
  updatedGame: WordChainGame;
};
