import { SSEClient } from "$lib/objects";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { Event, type LobbyCodeEvent } from "$lib/types/wordchain";

export const newSSEClient = (lobbyCodeCallback: (e: LobbyCodeEvent) => void): SSEClient => {
  const client = new SSEClient(`${PUBLIC_BACKEND_URL}/word-ladder/create-lobby`);

  client.connect();
  
  // TODO - add the rest of the listeners
  client.addEventListener(Event.LOBBY_CODE, lobbyCodeCallback);

  return client;
};