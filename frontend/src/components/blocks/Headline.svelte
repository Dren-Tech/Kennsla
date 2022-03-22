<script lang="ts">
	type HeadlineType = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';

	export let text: string = 'Text';
	export let headlineType: HeadlineType = 'h1';
</script>

<div class="block headline">
	{#if headlineType == 'h1'}
		<h1 class="text-2xl font-semibold">{text}</h1>
	{:else if headlineType == 'h2'}
		<h2 class="text-xl font-semibold">{text}</h2>
	{/if}
</div>
