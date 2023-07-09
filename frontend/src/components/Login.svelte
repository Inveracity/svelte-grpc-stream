<script lang="ts">
	import { server } from '$lib/stores/server';
	import { currentUser, pb } from '$lib/pocketbase';
	import { Connect } from '$lib/grpc';

	let password = '';
	let username = '';

	async function login() {
		await pb.collection('users').authWithPassword(username, password);
		await Connect($server, username, '0', pb.authStore.token);
	}
</script>

<div>
	{#if !$currentUser}
		<form on:submit|preventDefault>
			<input class="input input-secondary" type="text" bind:value={username} placeholder="Email" />
			<input
				class="input input-secondary"
				type="password"
				placeholder="Password"
				bind:value={password}
			/>
			<input class="input input-accent" type="text" bind:value={$server} placeholder="myserver" />
			<button class="btn btn-secondary" on:click={login}> Login </button>
		</form>
	{/if}
</div>
