<script lang="ts">
	import { Accordion } from '../ui/UI';
	export let title = '';
	export let text = '';
</script>

<div class="block hint">
	<Accordion {title} isOpen={false}>
		{text}
	</Accordion>
</div>
