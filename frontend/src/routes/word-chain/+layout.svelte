<script lang="ts">
	import './globals.css';

	let { children }: { children: () => any } = $props();
</script>

{@render children()}
