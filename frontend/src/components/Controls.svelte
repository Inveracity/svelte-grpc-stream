<script lang="ts">
	import { Connect } from '../grpc';
	import { status, username } from '../store';
</script>

<div>
	<input
		bind:value={$username}
		placeholder="user id"
		disabled={$status === 'connected' || $status === 'pending'}
	/>
	<button
		disabled={$status === 'connected' || $status === 'pending' || $username === ''}
		on:click={() => Connect('myserver', $username, '0')}
	>
		connect
	</button>
</div>
