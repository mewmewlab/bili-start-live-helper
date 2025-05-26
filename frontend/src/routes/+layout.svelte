<script>
	import '../app.css';
	import { Toaster } from "$lib/components/ui/sonner/index.js";
	let { children } = $props();
</script>

<Toaster richColors  position="top-center"/>
<main class="flex-1 overflow-y-auto p-4 w-full mx-auto justify-center">
	{#if children} 
		{@render children()}
	{:else}
		<p>fallback content</p>
	{/if}
</main>
