<script lang="ts">
	import { Button, Confetti, Modal } from '$lib/components/common';
	const SIZES = [4, 6, 8] as const;
	type BoardSize = (typeof SIZES)[number];

	let showCorrect = $state(false);
</script>

<div class="container">
	{#snippet sizeButton(n: BoardSize)}
		<button class="size-button" onclick={() => {}}>
			{n}
		</button>
	{/snippet}
	<Modal bind:open={showCorrect}>
		<div class="correct-message">
			<h1>Correct!</h1>
			<!-- <p>You completed a {board.length}x{board.length} puzzle in {0}</p> -->
			<p>Would you like to play again?</p>
			<div class="buttons-container">
				{#each SIZES as size}
					{@render sizeButton(size)}
				{/each}
			</div>
			<Button
				onclick={() => {
					showCorrect = false;
				}}>View board</Button
			>
		</div>
	</Modal>
	{#if showCorrect}
		<Confetti />
	{/if}

	<Button onclick={() => (showCorrect = true)}>Show Correct</Button>
</div>

<style>
	.container {
		height: 100%;
		width: 100%;
	}

	.correct-message {
		display: inline-flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;

		height: 60vh;
		width: 85vw;
	}

	.correct-message h1 {
		margin-bottom: 0;
	}

	.buttons-container {
		display: flex;
		gap: 0.5rem;
	}

	.size-button {
		--button-size: 3rem;
		--transition-ms: 200ms;
		--move-x-rem: 0.1rem;
		--move-y-rem: calc(var(--move-x-rem) * -1);

		background-color: var(--light);
		height: var(--button-size);
		width: var(--button-size);

		border: 1px solid var(--dark);
		border-radius: 5px;

		transition:
			color var(--transition-ms) linear,
			background-color var(--transition-ms) linear,
			border-color var(--transition-ms) linear,
			transform var(--transition-ms) linear,
			box-shadow var(--transition-ms) linear;
	}
	.size-button:hover {
		background-color: var(--gray);
		transform: translate(var(--move-x-rem), var(--move-y-rem));
	}
	.size-button:hover:nth-child(odd) {
		box-shadow: var(--move-y-rem) var(--move-x-rem) var(--red);
	}
	.size-button:hover:nth-child(even) {
		box-shadow: var(--move-y-rem) var(--move-x-rem) var(--blue);
	}
	.size-button:active {
		--transition-ms: 120ms;
		transform: translate(0, 0);
		box-shadow: none !important;
	}
</style>
