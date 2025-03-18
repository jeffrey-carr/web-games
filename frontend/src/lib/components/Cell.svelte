<script lang="ts">
	import lock from '$lib/assets/lock.svg';

	const buildClassStr = (value: number, position?: string, isHint?: boolean): string => {
		let str = position ?? '';

		if (value >= 0) {
			str = `${str} ${value === 0 ? 'zero' : 'one'}`;
		}

		if (isHint) {
			str = `${str} hint`;
		}

		return str.trim();
	};

	let {
		value = $bindable(),
		isHint = false,
		locked = false,
		position = '',
		zeroColor,
		oneColor,
		children
	}: {
		value: number;
		isHint?: boolean;
		zeroColor?: string;
		oneColor?: string;
		locked?: boolean;
		position?: string;
		children?: () => any;
	} = $props();

	let containerClassStr = $derived(buildClassStr(value, position, isHint));

	const change = () => {
		if (value === 1) {
			value = -1;
		} else {
			value++;
		}
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
