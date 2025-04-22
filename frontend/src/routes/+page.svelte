<script lang="ts">
	import chainIcon from '$lib/assets/chain.svg';
	import binokuIcon from '$lib/assets/binoku.svg';
</script>

<div class="container">
	<h1>Welcome!</h1>

	<div class="games-container">
		<a class="game-link" href="/binoku">
			<img class="icon" src={binokuIcon} alt="Binoku icon" />
			<span class="label">Binoku</span>
		</a>
		<a class="game-link" href="/word-chain">
			<img class="icon" src={chainIcon} alt="Word Chain icon" />
			<span class="label">Word Chain <span class="subtext">BETA</span></span>
		</a>
	</div>
</div>

<style>
	.subtext {
		font-size: 1.5rem;
		color: var(--wc-highlight);
	}

	.games-container {
		display: flex;
		justify-content: space-evenly;
		align-items: center;

		height: 100%;
		width: 100%;
	}

	.game-link {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;

		color: var(--wc-text);
		text-decoration: none;

		height: 90vh;
		width: 100%;
	}
	.game-link:hover {
		background-color: rgba(255, 255, 255, 0.3);
	}

	.game-link .icon {
		object-fit: fill;
		max-height: 10rem;
	}

	.game-link .label {
		font-size: 5rem;
	}
</style>
