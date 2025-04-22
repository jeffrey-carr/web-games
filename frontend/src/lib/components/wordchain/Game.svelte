<script lang="ts">
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import type {
		ValidateAnswerRequest,
		ValidateAnswerResponse,
		WordChainGame
	} from '$lib/types/wordchain';
	import Word from './Word.svelte';

	const TIMEOUT_PENALTY = 5000;

	let { game }: { game: WordChainGame } = $props();
	let gameUUID = $state('');
	$inspect(game);
	let revealedWords = $state<string[]>([]);
	let guesses = $state<string[]>([]);
	let timeouts = $state<number[]>([]);
	let revealedLetters = $state<number[]>([]);

	$effect(() => {
		if (!game) {
			gameUUID = '';
			return;
		}

		revealedWords = [...game.generatedChain.slice(0, game.userProgress)];
		const newGuesses = Array(game.generatedChain.length)
			.fill('')
			.map((_, i) => {
				if (i === 0 || game.userProgress > i) {
					return game.generatedChain[i];
				}

				return game.generatedChain[i].slice(0, revealedLetters[i]);
			});

		guesses = newGuesses;

		if (gameUUID !== game.uuid) {
			gameUUID = game.uuid;

			const initialTimeouts = Array(game.generatedChain.length).fill(0);
			timeouts = initialTimeouts;

			const initialRevealedLetters = Array(game.generatedChain.length).fill(1);
			// First word is completely revealed
			initialRevealedLetters[0] = game.generatedChain[0].length;
			revealedLetters = initialRevealedLetters;
		}
	});

	const updateGuess = (guess: string, index: number) => {
		if (timeouts[index] > new Date().getTime()) {
			return;
		}

		if (guess.length === 0) {
			guess = game.generatedChain[index].substring(0, revealedLetters[index]);
			guesses[index] = guess;
			return;
		}

		const revealed = game.generatedChain[index].substring(0, revealedLetters[index]);
		if (!guess.startsWith(revealed)) {
			guess = revealed + guess.substring(revealedLetters[index]);
		}

		guess = guess.toUpperCase();
		if (guess === guesses[index]) {
			return;
		}

		guesses[index] = guess;

		if (guesses[index].length === game.generatedChain[index].length) {
			submitGuess(guess, index);
		}
	};

	const submitGuess = async (guess: string, index: number) => {
		console.log(`Submitting guess {${guess}}`);

		const request = await fetch(`${PUBLIC_BACKEND_URL}/word-chain/validate-answer`, {
			method: 'POST',
			body: JSON.stringify({
				guess,
				gameState: game
			} as ValidateAnswerRequest)
		});

		if (request.status !== 200) {
			const data = request.json();
			console.error(data);
			return;
		}

		const response: ValidateAnswerResponse = await request.json();
		if (response.correct) {
			game = response.updatedGame;
			document.getElementById(`word-${game.userProgress}`)?.scrollIntoView();
		} else {
			timeouts[index] = new Date().getTime() + TIMEOUT_PENALTY;

			const targetWordLength = game.generatedChain[index].length;
			if (revealedLetters[index] < targetWordLength - 1) {
				revealedLetters[index]++;
			}

			guesses[index] = game.generatedChain[index].substring(0, revealedLetters[index]);
		}
	};
</script>

<div class="container">
	<div class="words-column">
		{#each game.generatedChain as word, i}
			<div class="word" id={`word-${i}`}>
				<Word
					word={guesses[i]}
					targetWord={word}
					onUpdate={(newGuess: string) => updateGuess(newGuess, i)}
					correct={i !== 0 && i < game.userProgress}
					locked={i !== game.userProgress}
					timedOutUntil={timeouts[i]}
					revealedLetters={revealedLetters[i]}
				/>
			</div>
		{/each}
	</div>
</div>

<style>
	.container {
		display: flex;
		justify-content: center;

		width: 100%;
	}

	.words-column {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		gap: 1.3rem;

		padding: 1rem;
	}

	.word {
		display: flex;

		height: 4rem;
	}
</style>
