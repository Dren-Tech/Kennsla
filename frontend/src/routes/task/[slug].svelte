<script lang="ts">
	import { Headline, FormattedContent, Hint, Question } from '../../components/blocks/Blocks';
	import type Task from '../../domain/model/Task';

	export let task: Task;

	console.log(task);
</script>

<div class="task">
	{#each task.blocks as block}
		{#if block.type == 'headline'}
			<Headline text={block.payload.value} headlineType={block.payload.headlineType} />
		{:else if block.type == 'formatted-content'}
			<FormattedContent text={block.payload.value} />
		{:else if block.type == 'hint'}
			<Hint title={block.payload.title} text={block.payload.text} />
		{:else if block.type == 'question'}
			<Question
				submitButtonText={block.payload.submitButtonText}
				questions={block.payload.questions}
			/>
		{/if}
	{/each}
</div>
