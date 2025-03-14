<script lang="ts">
	import type { BoardStyle, Coordinate } from '$lib/types';
	import Cell from './Cell.svelte';

	let {
		board = $bindable(),
		style = 'colors',
		lockedCells
	}: {
		board: number[][];
		style?: BoardStyle;
		lockedCells: Coordinate[];
	} = $props();

	const getPosition = (rowIndex: number, colIndex: number): string => {
		const top = rowIndex === 0;
		const bottom = rowIndex === board.length - 1;
		const left = colIndex === 0;
		const right = colIndex === board.length - 1;

		switch (true) {
			case top && left:
				return 'top-left';
			case top && right:
				return 'top-right';
			case bottom && left:
				return 'bottom-left';
			case bottom && right:
				return 'bottom-right';
		}

		return '';
	};

	const isLocked = (row: number, col: number): boolean => {
		return lockedCells.findIndex((cell) => cell.row === row && cell.col === col) >= 0;
	};
</script>

<div class="container" style={`--size: ${board.length}`}>
	{#each board as row, rowI}
		{#each row as cell, colI}
			<Cell
				bind:value={board[rowI][colI]}
				locked={isLocked(rowI, colI)}
				position={getPosition(rowI, colI)}
				zeroColor={style === 'numbers' ? 'var(--gray)' : undefined}
				oneColor={style === 'numbers' ? 'var(--gray)' : undefined}
			>
				{#if style === 'numbers' && cell >= 0}
					{cell}
				{/if}
			</Cell>
		{/each}
	{/each}
</div>

<style>
	.container {
		--container-size: min(50vw, 50vh);

		display: grid;
		grid-template-columns: repeat(var(--size, 0), 1fr);

		height: var(--container-size);
		width: var(--container-size);
	}

	@media only screen and (max-width: 600px) {
		.container {
			--container-size: min(80vw, 80vh);
		}
	}
</style>
