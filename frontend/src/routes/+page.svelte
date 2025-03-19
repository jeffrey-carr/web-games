<script lang="ts">
	import { Board, Button, Dropdown, Spinner } from '$lib/components';
	import PopMessage from '$lib/components/PopMessage.svelte';
	import type { InvalidHint, BoardStyle, Coordinate, DropdownOption } from '$lib/types';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	const DEFAULT_SIZE = 6;
	const SIZES = [4, 6, 8];

	let size = $state(DEFAULT_SIZE);
	let err = $state<string | null>(null);

	let board = $state<number[][]>([]);
	let lockedCells = $state<Coordinate[]>([]);

	const HINT_DURATION_MS = 3500;
	let hint = $state<InvalidHint>();
	let hintTimer = $state<number | undefined>();
	const setHint = (newHint: InvalidHint) => {
		clearTimeout(hintTimer);
		hint = newHint;
		hintTimer = setTimeout(() => {
			hint = undefined;
		}, HINT_DURATION_MS);
	};

	let showCorrect = $state(false);

	let style = $state<BoardStyle>('colors');

	// 0 - Not generating
	// 1 - Generating
	// 2 - Generating (long)
	// 3 - Generation (very long)
	// 4 - Generation (very very long)
	let generatingLevel = $state(0);
	let generating = $derived(generatingLevel > 0);
	const GENERATING_MESSAGES = [
		'Generating puzzle...',
		'This is taking a while, sorry',
		'Doo doo do...',
		"Hold on, it's coming!!!",
		'Okay, well eventually'
	];

	let validating = $state(false);

	const styleOptions: DropdownOption[] = [
		{
			label: 'Numbers',
			onclick: () => {
				style = 'numbers';
			}
		},
		{
			label: 'Colors',
			onclick: () => {
				style = 'colors';
			}
		}
	];

	const getBoard = async () => {
		if (err != null) {
			return;
		}

		board = [];

		let generatingTimeout: number | undefined;

		generatingLevel = 1;
		generatingTimeout = setInterval(() => {
			generatingLevel = Math.min(GENERATING_MESSAGES.length, generatingLevel + 1);
		}, 10000);

		const boardRequest = await fetch(`${PUBLIC_BACKEND_URL}/new-game?size=${size}`);
		const data = await boardRequest.json();
		board = data['board'];
		lockedCells = getLockedCells(board);

		clearInterval(generatingTimeout);
		generatingLevel = 0;
	};

	const getLockedCells = (matrix: number[][]): Coordinate[] => {
		const coords: Coordinate[] = [];
		for (let row = 0; row < matrix.length; row++) {
			for (let col = 0; col < matrix.length; col++) {
				if (matrix[row][col] >= 0) {
					coords.push({ row, col });
				}
			}
		}

		return coords;
	};

	const checkSolution = async (): Promise<boolean> => {
		validating = true;
		const validateRequest = await fetch(`${PUBLIC_BACKEND_URL}/validate-game`, {
			method: 'POST',
			body: JSON.stringify({ board })
		});

		const response = await validateRequest.json();
		if (response['valid']) {
			showCorrect = true;
		} else if (response['hint']) {
			setHint(response['hint']);
		}

		validating = false;
		return true;
	};

	const newGame = () => {
		showCorrect = false;
		getBoard();
	};
</script>

<div class="container">
	{#snippet sizeButton(n: number)}
		<button class={`size-button ${n === size ? 'selected' : ''}`} onclick={() => (size = n)}>
			{n}
		</button>
	{/snippet}

	<h1 class="title">Binoku</h1>
	<div class="button-container">
		<Button size="medium" onclick={getBoard}>Generate puzzle</Button>
		<Button size="medium" onclick={checkSolution} disabled={validating}>
			{#if validating}
				<Spinner color="dark" />
			{:else}
				Check solution
			{/if}
		</Button>
		<Dropdown options={styleOptions}>Style</Dropdown>
	</div>
	<div class="buttons-container">
		{#each SIZES as size}
			{@render sizeButton(size)}
		{/each}
	</div>
	{#if generating}
		<div class="spinner-container">
			<Spinner />
		</div>
		{#each { length: generatingLevel } as _, i}
			<p class="loading-message">{GENERATING_MESSAGES[i]}</p>
		{/each}
	{:else if board.length > 0}
		<Board bind:board {lockedCells} {style} {hint} />

		{#if showCorrect}
			<PopMessage>
				<div class="correct-message">
					<h1 class="correct-message-header">Correct!</h1>
					<div class="correct-message-body">
						<p>You completed a {board.length}x{board.length} puzzle in {0}</p>
						<div class="buttons-container">
							{#each SIZES as size}
								{@render sizeButton(size)}
							{/each}
						</div>
						<div class="buttons-container">
							<Button onclick={newGame}>New Game</Button>
							<Button
								onclick={() => {
									showCorrect = false;
								}}>View board</Button
							>
						</div>
					</div>
				</div>
			</PopMessage>
		{/if}
	{/if}
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		justify-content: start;
		align-items: center;
		gap: 2rem;

		height: 100vh;
		width: 100vw;
	}

	.spinner-container {
		--size-rem: 2.5rem;
		height: var(--size-rem);
		width: var(--size-rem);
	}

	.loading-message {
		margin: 0;
		padding: 0;

		animation: slide-in 200ms ease-out;
	}

	.title {
		margin-top: 0;
		margin-bottom: 0;
		font-size: 5rem;
	}

	.button-container {
		display: flex;
		justify-content: center;
		gap: 1rem;

		width: 100%;
	}

	.size-button {
		--transition-ms: 200ms;

		background-color: var(--light);
		height: 3rem;
		width: 3rem;

		border: 1px solid var(--dark);
		border-radius: 5px;

		transition:
			color var(--transition-ms) linear,
			background-color var(--transition-ms) linear,
			border-color var(--transition-ms) linear,
			transform var(--transition-ms) linear,
			box-shadow var(--transition-ms) linear;
	}

	.size-button.selected {
		--move-x-rem: 0.1rem;
		--move-y-rem: calc(var(--move-x-rem) * -1);
		background-color: var(--gray);
		transform: translate(var(--move-x-rem), var(--move-y-rem));
	}
	.size-button.selected:nth-child(odd) {
		box-shadow: var(--move-y-rem) var(--move-x-rem) var(--red);
	}
	.size-button.selected:nth-child(even) {
		box-shadow: var(--move-y-rem) var(--move-x-rem) var(--blue);
	}

	.correct-message {
		display: flex;
		flex-direction: column;
		justify-content: space-evenly;
		align-items: center;

		height: 100%;
		width: 100%;
	}

	.correct-message-header {
		text-align: center;
		font-family: 'Spicy Rice', Impact;
		font-size: 5rem;

		margin: 0;
		padding: 0;
	}

	.correct-message-body {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		gap: 1rem;

		width: 100%;
	}

	.buttons-container {
		display: flex;
		gap: 0.5rem;
	}

	@keyframes slide-in {
		from {
			opacity: 0;
			transform: translateY(-1rem);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
</style>
