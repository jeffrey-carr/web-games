<script lang="ts">
	let {
		onclick,
		size = 'fit',
		disabled = false,
		children
	}: {
		onclick: () => void;
		size?: 'fit' | 'small' | 'medium' | 'large' | 'fill' | 'fill-width';
		disabled?: boolean;
		children?: () => any;
	} = $props();
</script>

<button class={`button ${size}`} {onclick}>
	{@render children?.()}
</button>

<style>
	.button {
		--transition-ms: 200ms;

		background-color: var(--light);

		border: 1px solid var(--dark);
		border-radius: 5px;

		padding: 0.5rem;

		transition:
			color var(--transition-ms) linear,
			background-color var(--transition-ms) linear;
	}

	.button:hover {
		color: var(--light);
		background-color: var(--dark);
	}

	.button:disabled {
		background-color: var(--gray);
		color: var(--dark);
	}

	.button.small {
		padding: 0;
		font-size: 0.6rem;
		height: 2rem;
		width: 5rem;
	}
	.button.medium {
		height: 2.5rem;
		width: 8rem;
	}
	.button.large {
		width: 15rem;
	}
	.button.fill {
		width: 100%;
		height: 100%;
	}
	.button.fill-width {
		width: 100%;
	}
</style>
