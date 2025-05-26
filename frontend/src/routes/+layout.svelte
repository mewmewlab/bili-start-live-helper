<script>
    import { onMount } from 'svelte';
	import '../app.css';
    import { goto } from '$app/navigation';
	import { Toaster } from "$lib/components/ui/sonner/index.js";
	onMount(async () => {
		// 检查是否登录
		const res = true
		if (!res) {
			goto('/login')
		}
	})	
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
