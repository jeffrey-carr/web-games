<script lang="ts">
	import Button from './Button.svelte';

	let {
		open = $bindable(),
		children
	}: {
		open: boolean;
		children?: () => any;
	} = $props();

	const close = () => {
		open = false;
	};
</script>

<div class={`container ${open ? 'open' : ''}`}>
	<div class="close-button">
		<Button onclick={close} size="fill">X</Button>
	</div>
	{@render children?.()}
</div>

<style>
	.container {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 10;

		margin: 0;
		padding: 0;

		border: 1px solid var(--dark);
		border-radius: 5px;

		color: var(--dark);
		background-color: var(--light);

		transition: opacity 100ms linear;

		overflow: auto;

		/* Hidden settings */
		opacity: 0;
		pointer-events: none;
	}

	.open {
		opacity: 1;
		pointer-events: all;
	}

	.close-button {
		--gap: 0.5rem;
		--size: 2rem;

		position: absolute;
		top: var(--gap);
		left: calc(100% - var(--size) - var(--gap));
		z-index: 2;

		height: var(--size);
		width: var(--size);
	}
</style>
