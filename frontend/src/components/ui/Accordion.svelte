<!-- BIG THANKS to el3um4s: https://github.com/el3um4s/svelte-component-info/blob/main/src/lib/components/helpers/Accordion.svelte -->
<script lang="ts">
	import { Arrow } from './icons/Icons';
	import { slide } from 'svelte/transition';
	export let title = 'Title';
	export let isOpen = true;
	let onClick = () => {
		isOpen = !isOpen;
	};
</script>

<section>
	<button on:click={onClick}>
		<header>
			<Arrow {isOpen} />
			{title}
		</header>
	</button>
	{#if isOpen}
		<div transition:slide={{ duration: 500 }} data-testid="accordion-open">
			<slot />
		</div>
	{/if}
</section>

<style lang="postcss">
	section {
		@apply flex flex-col py-2 my-2;
	}
	header {
		@apply flex flex-row items-center justify-start p-2 font-bold;
		background-color: var(--text-color, theme('colors.orange.500'));
		color: var(--background-color, theme('colors.gray.50'));
	}
	div {
		@apply flex flex-col p-2 border border-solid w-full;
		background-color: var(--background-color, theme('colors.gray.50'));
		color: var(--text-color, theme('colors.gray.800'));
		border-color: var(--text-color, theme('colors.orange.500'));
	}
</style>
