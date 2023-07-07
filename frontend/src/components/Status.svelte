<script>
	import { server } from '../stores/server';
	import { status } from '../stores/status';
	import Circle from './Circle.svelte';
</script>

<div class="status">
	<Circle color={$status === 'connected' ? 'green' : 'grey'}>
		<p>connected to server: {$server}</p>
	</Circle>
</div>

<style>
	.status {
		display: flex;
		flex-direction: row;
		align-items: center;
	}
</style>
