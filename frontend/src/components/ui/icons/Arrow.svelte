<script lang="ts">
	export let isOpen = false;
	export let onClick = () => {
		isOpen = !isOpen;
	};
</script>

<button on:click={onClick}>
	<svg
		class:isOpen
		xmlns="http://www.w3.org/2000/svg"
		viewBox="0 0 24 24"
		fill="none"
		stroke-linecap="round"
		stroke-linejoin="round"
		stroke-width="3"
	>
		<path d="M19 9l-7 7-7-7" />
	</svg>
</button>

<style lang="postcss">
	:root {
		--transition-duration: 500ms;
	}
	button {
		cursor: pointer;
		display: flex;
		align-items: center;
		overflow: hidden;
	}
	svg {
		transition: transform var(--transition-duration);
	}
	svg {
		@apply h-6 w-6;
		stroke: var(--background-color, theme('colors.gray.50'));
		background-color: var(--text-color, theme('colors.orange.500'));
	}
	.isOpen {
		transform: rotate(180deg);
	}
</style>
