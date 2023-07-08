<script lang="ts">
	import { currentUser, pb } from '$lib/pocketbase';
	import { server } from '../stores/server';
	import { Connect } from '../lib/grpc';
	import Logout from './Logout.svelte';

	let password = '';
	let username = '';

	let mockUser = 'testuser';

	async function login() {
		await pb.collection('users').authWithPassword(username, password);
		await Connect($server, username, '0', pb.authStore.token);
	}
</script>

<div>
	{#if mockUser}
		<Logout />
	{:else}
		<form on:submit|preventDefault>
			<input class="input input-secondary" type="text" bind:value={username} placeholder="Email" />
			<input
				class="input input-secondary"
				type="password"
				bind:value={password}
				placeholder="Password"
			/>
			<input class="input input-accent" type="text" bind:value={$server} placeholder="myserver" />
			<button class="btn btn-secondary" on:click={login}> Login </button>
		</form>
	{/if}
</div>
