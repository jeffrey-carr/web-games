<script lang="ts">
	import lock from '$lib/assets/lock.svg';
	import type { Coordinate } from '$lib/types';

	const buildClassStr = (value: number, isHint?: boolean): string => {
		let str = '';

		if (value >= 0) {
			str = `${str} ${value === 0 ? 'zero' : 'one'}`;
		}

		if (isHint) {
			str = `${str} hint`;
		}

		return str.trim();
	};

	// const getPosition = (rowIndex: number, colIndex: number): string => {
	// 	const top = rowIndex === 0;
	// 	const bottom = rowIndex === board.length - 1;
	// 	const left = colIndex === 0;
	// 	const right = colIndex === board.length - 1;

	// 	switch (true) {
	// 		case top && left:
	// 			return 'top-left';
	// 		case top && right:
	// 			return 'top-right';
	// 		case bottom && left:
	// 			return 'bottom-left';
	// 		case bottom && right:
	// 			return 'bottom-right';
	// 	}

	// 	return '';
	// };

	let {
		value = $bindable(),
		updateValue,
		coord,
		isHint = false,
		locked = false,
		zeroColor,
		oneColor,
		children
	}: {
		value: number;
		updateValue: (row: number, col: number, v: number) => void;
		coord: Coordinate;
		isHint?: boolean;
		zeroColor?: string;
		oneColor?: string;
		locked?: boolean;
		children?: () => any;
	} = $props();

	// let position = $derived(getPosition(coord.row, coord.col));
	let containerClassStr = $derived(buildClassStr(value, isHint));

	const change = () => {
		let newValue: number;
		if (value === 1) {
			newValue = -1;
		} else {
			newValue = value + 1;
		}

		updateValue(coord.row, coord.col, newValue);
	};
</script>

<div
	class={`container ${containerClassStr}`}
	style={`--zeroColor: ${zeroColor ?? 'var(--blue)'}; --oneColor: ${oneColor ?? 'var(--red)'}`}
>
	<button aria-label="board cell" class="button" onclick={change} disabled={locked}>
		{@render children?.()}
	</button>
	{#if locked}
		<img class="lock" src={lock} alt="lock icon" />
	{/if}
</div>

<style>
	.container {
		--animate-ms: 75ms;
		--border-radius-px: 5px;
		--zeroColor: --blue;
		--oneColor: --red;

		position: relative;

		border: 1px solid var(--dark);
		background-color: var(--gray);

		height: 100%;
		width: 100%;

		transition:
			background-color var(--animate-ms) linear,
			border-color var(--animate-ms) linear;
	}
	.container.top-left {
		border-top-left-radius: var(--border-radius-px);
	}
	.container.top-right {
		border-top-right-radius: var(--border-radius-px);
	}
	.container.bottom-left {
		border-bottom-left-radius: var(--border-radius-px);
	}
	.container.bottom-right {
		border-bottom-right-radius: var(--border-radius-px);
	}
	.container:disabled:hover {
		border-color: black;
	}

	.zero {
		background-color: var(--zeroColor);
	}

	.one {
		background-color: var(--oneColor);
	}

	.hint {
		border: 2px solid var(--red);
	}

	.button {
		position: relative;
		z-index: 1;

		height: 100%;
		width: 100%;

		border: none;
		background-color: transparent;

		font-size: 1.2rem;
	}
	.button:disabled {
		color: black;
	}

	.lock {
		--gap: 0.3rem;
		--size: 0.8rem;

		position: absolute;
		top: var(--gap);
		right: var(--gap);

		height: var(--size);
		width: var(--size);
	}
</style>
