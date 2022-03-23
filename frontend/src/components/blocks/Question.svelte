<script lang="ts">
	interface QuestionItem {
		text: string;
		inputName: string;
	}

	export let questions: Array<QuestionItem>;
	export let submitButtonText: string;
</script>

<div class="block question">
	<div>
		{#each questions as question}
			<div>{question.text}</div>
			<div>
				<input
					class="border border-gray-500 rounded-lg p-1 my-2 w-1/3"
					type="text"
					name={question.inputName}
				/>
			</div>
		{/each}
	</div>

	<div>
		<button class="bg-blue-400 hover:bg-blue-600 text-white p-2 rounded-lg"
			>{submitButtonText}</button
		>
	</div>
</div>
