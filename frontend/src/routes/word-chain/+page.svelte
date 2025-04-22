<script lang="ts">
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import type { WordChainGame } from '$lib/types/wordchain';
	import { Game } from '$lib/components/wordchain';
	import { Button } from '$lib/components/common';
	import { goto } from '$app/navigation';

	let game = $state<WordChainGame | null>(null);

	const newGame = async () => {
		const response = await fetch(`${PUBLIC_BACKEND_URL}/word-chain/new-game`);
		const data = await response.json();

		game = data;
	};

	const back = () => {
		goto('/');
	};
</script>

<div class="container">
	<h1 class="title">Word Chain <span class="subtext">BETA</span></h1>
	<div class="buttons-container">
		<Button onclick={back}>Back to Main</Button>
		<Button onclick={newGame}>New Game</Button>
	</div>
	{#if game}
		<Game {game} />
	{/if}
</div>

<style>
	.title {
		text-align: center;
	}
	.subtext {
		font-size: 0.7rem;
		color: var(--wc-highlight);
	}

	.buttons-container {
		display: flex;
		justify-content: center;
		gap: 0.3rem;

		width: 100%;

		margin-bottom: 2rem;
	}
</style>
