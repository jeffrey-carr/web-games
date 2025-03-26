<script lang="ts">
	import leftArrow from '$lib/assets/arrow-left.svg';
	import rightArrow from '$lib/assets/arrow-right.svg';

	import type { Snippet } from 'svelte';
	import { Button } from '$lib/components/common';

	let { snippet, numPages }: { snippet: Snippet<[number]>; numPages: number } = $props();
	let currentSlide = $state(0);

	const nextSlide = () => {
		currentSlide = (currentSlide + 1) % numPages;
	};

	const previousSlide = () => {
		currentSlide = Math.max(0, currentSlide - 1);
	};
</script>

<div class="container">
	<div class="content-container">
		{@render snippet(currentSlide)}
	</div>
	<div class="controls-container">
		{#snippet arrow(direction: 'left' | 'right')}
			<div class="arrow">
				{#if direction === 'left'}
					<Button size="small" onclick={previousSlide}>
						<img class="arrow" src={leftArrow} alt="A left arrow" />
					</Button>
				{:else}
					<Button size="small" onclick={nextSlide}>
						<img class="arrow" src={rightArrow} alt="A right arrow" />
					</Button>
				{/if}
			</div>
		{/snippet}
		{#snippet dot(highlighted: boolean)}
			<div class={`dot ${highlighted ? 'highlighted' : ''}`}></div>
		{/snippet}
		{@render arrow('left')}
		{#each { length: numPages } as _, i}
			{@render dot(currentSlide === i)}
		{/each}
		{@render arrow('right')}
	</div>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		gap: 1rem;

		height: 100%;
	}
	.controls-container {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 1.5rem;

		width: 100%;
	}

	.dot {
		height: 1rem;
		width: 1rem;

		border: 1px solid black;
		border-radius: 100%;

		transition: background-color 200ms linear;
	}
	.dot.highlighted {
		background-color: black;
	}
</style>
