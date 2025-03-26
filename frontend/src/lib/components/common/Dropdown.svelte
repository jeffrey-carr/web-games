<script lang="ts">
	import type { DropdownOption } from '$lib/types';
	import Button from './Button.svelte';

	let {
		options,
		direction = 'down',
		children
	}: {
		options: DropdownOption[];
		direction?: 'up' | 'down';
		children?: () => any;
	} = $props();

	let dropped = $state(false);

	const drop = () => {
		dropped = !dropped;
	};
</script>

<div class="container">
	<div class="main button-wrapper">
		<Button onclick={drop} size="fill">
			{@render children?.()}
		</Button>
	</div>

	{#snippet Option(option: DropdownOption, index: number)}
		<div
			class={`dropdown-item button-wrapper ${dropped ? 'dropped' : ''} ${direction}`}
			style={`--option-num: ${index}`}
		>
			<Button onclick={option.onclick} size="fill">{option.label}</Button>
		</div>
	{/snippet}
	{#each options as option, i}
		{@render Option(option, i)}
	{/each}
</div>

<style>
	.container {
		--height-rem: 2.5rem;
		--width-rem: 8rem;
		--option-height-rem: calc(var(--height-rem) * 0.8);
		--option-width-rem: calc(var(--width-rem) * 0.8);
		--option-num: 0;
		--multiplier: 1;

		position: relative;
		z-index: 1;
	}

	.dropdown-item {
		--gap: 0.2rem;

		position: absolute;
		top: 0;
		left: calc(50% - var(--option-width-rem) / 2);
		z-index: 2;

		transition: transform 100ms ease-out;
	}
	.dropdown-item.up {
		--multiplier: -1;
	}
	.dropdown-item.dropped {
		transition-delay: calc(var(--option-num) * 75ms);
		transform: translateY(
			calc(
				(
						(var(--option-height-rem) + var(--gap)) * var(--option-num) + var(--height-rem) +
							var(--gap) * 2
					) *
					var(--multiplier)
			)
		);
	}

	.button-wrapper {
		height: var(--option-height-rem);
		width: var(--option-width-rem);
	}
	.button-wrapper.main {
		z-index: 3;
		height: var(--height-rem);
		width: var(--width-rem);
	}
</style>
