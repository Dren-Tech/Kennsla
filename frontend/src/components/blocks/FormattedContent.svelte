<script lang="ts">
	import MarkdownIt from 'markdown-it';

	export let text = '';

	const md = new MarkdownIt();

	const renderedText = md.render(text);
</script>

<div class="block formatted-content">
	<div class="my-2">{@html renderedText}</div>
</div>
