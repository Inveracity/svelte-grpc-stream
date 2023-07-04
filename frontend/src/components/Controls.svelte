<script lang="ts">
	import { Connect } from '../grpc';
	import { server } from '../stores/server';
	import { status } from '../stores/status';
	import { username } from '../stores/username';
</script>

<div class="center">
	<input
		bind:value={$username}
		placeholder="user id"
		disabled={$status === 'connected' || $status === 'pending'}
	/>
	<input
		bind:value={$server}
		placeholder="server id"
		disabled={$status === 'connected' || $status === 'pending'}
	/>
	<button
		disabled={$status === 'connected' || $status === 'pending' || $username === ''}
		on:click={() => Connect($server, $username, '0')}
	>
		connect
	</button>
</div>

<style>
	button {
		margin-top: 1rem;
	}
</style>
