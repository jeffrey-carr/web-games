<script lang="ts">
	import type { BoardStyle, Coordinate, InvalidHint } from '$lib/types';
	import { Stack } from '$lib/utils';
	import Button from './Button.svelte';
	import Cell from './Cell.svelte';

	let {
		board = $bindable(),
		hint,
		style = 'colors',
		lockedCells
	}: {
		board: number[][];
		hint?: InvalidHint;
		style?: BoardStyle;
		lockedCells: Coordinate[];
	} = $props();

	type Move = {
		row: number;
		col: number;
		value: number;
	};
	let moveStack = $state(new Stack<Move>());
	let noMoves = $state(true);

	const fillCell = (row: number, col: number, value: number) => {
		const previousValue = board[row][col];
		moveStack.push({ row, col, value: previousValue });
		noMoves = false;
		board[row][col] = value;
	};

	const isLocked = (row: number, col: number): boolean => {
		return lockedCells.findIndex((cell) => cell.row === row && cell.col === col) >= 0;
	};

	const isHintCell = (row: number, col: number): boolean => {
		if (hint?.rows?.includes(row)) {
			return true;
		}

		if (hint?.cols?.includes(col)) {
			return true;
		}

		return false;
	};

	const undoMove = () => {
		if (moveStack.size() === 0) {
			return;
		}

		const lastMove = moveStack.pop();
		if (lastMove == null) {
			return;
		}

		if (moveStack.size() === 0) {
			noMoves = true;
		}

		board[lastMove.row][lastMove.col] = lastMove.value;
	};

	const resetBoard = () => {
		while (moveStack.size() > 0) {
			undoMove();
		}
	};
</script>

<div class="container" style={`--size: ${board.length}`}>
	{#each board as row, rowI}
		{#each row as cell, colI}
			<Cell
				value={board[rowI][colI]}
				updateValue={fillCell}
				coord={{ row: rowI, col: colI }}
				locked={isLocked(rowI, colI)}
				zeroColor={style === 'numbers' ? 'var(--gray)' : undefined}
				oneColor={style === 'numbers' ? 'var(--gray)' : undefined}
				isHint={isHintCell(rowI, colI)}
			>
				{#if style === 'numbers' && cell >= 0}
					{cell}
				{/if}
			</Cell>
		{/each}
	{/each}
</div>
<div class="buttons-container">
	<Button onclick={undoMove} disabled={noMoves}>Undo last move</Button>
	<Button onclick={resetBoard} disabled={noMoves}>Reset board</Button>
</div>

<style>
	.container {
		--container-size: min(50vw, 50vh);

		display: grid;
		gap: 0.5rem;
		grid-template-columns: repeat(var(--size, 0), 1fr);

		height: var(--container-size);
		width: var(--container-size);
	}

	.buttons-container {
		display: flex;
		gap: 1rem;
		margin-bottom: 1rem;
	}

	@media only screen and (max-width: 600px) {
		.container {
			--container-size: min(80vw, 80vh);
		}
	}
</style>
