<script lang="ts">
	import { currentUser, pb } from '$lib/pocketbase';
	import { server } from '../stores/server';
	import { Connect, Disconnect } from '../lib/grpc';

	let password = '';
	let username = '';

	async function login() {
		await pb.collection('users').authWithPassword(username, password);
		await Connect($server, username, '0', pb.authStore.token);
	}

	function logout() {
		pb.authStore.clear();
		Disconnect();
	}
</script>

<div>
	{#if $currentUser}
		<p>Signed in as {$currentUser.username}</p>
		<button on:click={logout}> Logout </button>
	{:else}
		<form on:submit|preventDefault>
			<input type="text" bind:value={username} placeholder="Email" />
			<input type="password" bind:value={password} placeholder="Password" />
			<input type="text" bind:value={$server} />
			<button class="btn" on:click={login}> Login </button>
		</form>
	{/if}
</div>
