<script lang="ts">
	import Headline from '../../components/blocks/Headline.svelte';
	import type Task from '../../domain/model/Task';

	export let task: Task;

	console.log(task);
</script>

<div class="task">
	{#each task.blocks as block}
		{#if block.type == 'headline'}
			<Headline text={block.payload.value} headlineType={block.payload.headlineType} />
		{/if}
	{/each}

	<div class="block formatted-content">
		<p>
			Consectetur exercitation voluptate laborum voluptate. Qui magna laborum deserunt do consequat
			irure anim enim cupidatat aliqua officia non tempor proident. Sit aliquip eiusmod fugiat
			incididunt do. Incididunt labore excepteur consequat sit aliquip labore ullamco consequat et
			officia aute tempor. Laborum minim consectetur eu labore cupidatat laborum cillum irure labore
			do laborum reprehenderit. Ullamco est id quis excepteur enim officia et ad veniam cillum
			eiusmod excepteur magna deserunt. Deserunt incididunt veniam incididunt velit nulla. Cupidatat
			qui quis nisi proident sunt consectetur sunt amet fugiat eiusmod sint proident. Veniam non
			sunt id est ullamco ad fugiat laboris pariatur incididunt ad. Ut occaecat nulla nostrud aliqua
			exercitation Lorem adipisicing amet proident culpa esse. Aliqua eu deserunt id sunt anim nisi
			officia Lorem irure ex. Consequat id amet velit elit reprehenderit. Ut qui fugiat dolor
			pariatur ipsum deserunt cillum.
		</p>
	</div>

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
