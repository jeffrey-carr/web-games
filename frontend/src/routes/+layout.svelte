<script lang="ts">
	import '$lib/globals.css';

	let { children }: { children: () => any } = $props();
</script>

{@render children()}
