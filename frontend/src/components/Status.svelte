<script>
	import { server } from '../stores/server';
	import { status } from '../stores/status';
	import Circle from './Circle.svelte';
</script>

<div>
	<Circle color={$status === 'connected' ? 'green' : 'grey'}>
		<p>connected to server: {$server}</p>
	</Circle>
</div>
