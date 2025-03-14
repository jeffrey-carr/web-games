<script lang="ts">
	import { Board, Button, Dropdown, Spinner } from '$lib/components';
	import type { BoardStyle, Coordinate, DropdownOption } from '$lib/types';

	const DEFAULT_SIZE = 6;

	let size = $state(DEFAULT_SIZE);
	let err = $state<string | null>(null);

	let board = $state<number[][]>([[]]);
	let lockedCells = $state<Coordinate[]>([]);

	let style = $state<BoardStyle>('colors');

	// 0 - Not generating
	// 1 - Generating
	// 2 - Generating (long)
	// 3 - Generation (very long)
	// 4 - Generation (very very long)
	let generatingLevel = $state(0);
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

		let generatingTimeout: number | undefined;
		let longGeneratingTimeout: number | undefined;

		generatingLevel = 1;
		generatingTimeout = setInterval(() => {
			generatingLevel = (generatingLevel + 1) % GENERATING_MESSAGES.length;
		}, 5000);

		const boardRequest = await fetch(`http://localhost:3001/new-game?size=${size}`);
		const data = await boardRequest.json();
		board = data['board'];
		lockedCells = getLockedCells(board);

		// clearTimeout(generatingTimeout);
		// clearTimeout(longGeneratingTimeout);
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
		// Check that every space is filled in
		if (board.some((row) => row.some((cell) => cell < 0))) {
			validating = false;
			err = 'Invalid!';
			return false;
		}

		const validateRequest = await fetch(`http://localhost:3001/validate-game`, {
			method: 'POST',
			body: JSON.stringify({ board })
		});
		const response = await validateRequest.json();

		if (response['valid']) {
			err = 'Valid!';
		} else {
			err = 'Invalid!';
		}

		validating = false;
		return true;
	};
</script>

<div class="container">
	<h1 class="title">Binoku</h1>
	<div class="button-container">
		<Button onclick={getBoard}>Generate puzzle</Button>
		<Button onclick={checkSolution}>Check solution</Button>
		<Dropdown options={styleOptions}>Style</Dropdown>
	</div>
	<div>
		{#snippet sizeButton(n: number)}
			<button class={`size-button ${n === size ? 'selected' : ''}`} onclick={() => (size = n)}>
				{n}
			</button>
		{/snippet}
		{@render sizeButton(4)}
		{@render sizeButton(6)}
		{@render sizeButton(8)}
	</div>
	{#if generatingLevel > 0}
		<Spinner />
		{#each { length: generatingLevel } as _, i}
			<p class="loading-message">{GENERATING_MESSAGES[i]}</p>
		{/each}
	{:else if board.length > 0}
		<Board bind:board {lockedCells} {style} />
	{/if}
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		justify-content: start;
		align-items: center;
		gap: 1rem;

		height: 100vh;
		width: 100vw;
	}

	.loading-message {
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
		color: var(--light);
		background-color: var(--dark);
		border-color: var(--light);
		transform: translate(var(--move-x-rem), var(--move-y-rem));
	}
	.size-button.selected:nth-child(odd) {
		box-shadow: var(--move-y-rem) var(--move-x-rem) var(--red);
	}
	.size-button.selected:nth-child(even) {
		box-shadow: var(--move-y-rem) var(--move-x-rem) var(--blue);
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
