<script lang="ts">
	import { Headline, FormattedContent } from '../../components/blocks/Blocks';
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
		{/if}
	{/each}

	<div class="block question">
		<div>Dies ist eine Frage. Wie ist die Antwort auf die Frage?</div>
		<div>
			<input type="text" />
			<button>Absenden</button>
		</div>
	</div>

	<div class="block hin">
		<h3>Tipp 1</h3>
		<p>
			Laboris velit eu qui sint nostrud in laboris. Nulla pariatur eiusmod ut dolore pariatur elit.
			Nostrud do enim laboris ex consequat anim aliqua magna sunt nostrud culpa elit aliquip esse.
		</p>
	</div>

	<div class="block hin">
		<h3>Tipp 2</h3>
		<p>
			Laboris velit eu qui sint nostrud in laboris. Nulla pariatur eiusmod ut dolore pariatur elit.
			Nostrud do enim laboris ex consequat anim aliqua magna sunt nostrud culpa elit aliquip esse.
		</p>
	</div>
</div>
