<script lang="ts">
	import { currentUser, pb } from '$lib/pocketbase';
	import { server } from '../stores/server';
	import { Connect, Disconnect } from '../grpc';
	import { messages } from '../stores/messages';
	import { username } from '../stores/username';

	let password = '';

	async function login() {
		await pb.collection('users').authWithPassword($username, password);
		await Connect($server, $username, '0');
	}

	function logout() {
		pb.authStore.clear();
		Disconnect();
		messages.reset();
	}
</script>

<div class="login">
	{#if $currentUser}
		<p class="padding">Signed in as {$currentUser.username}</p>
		<button class="padding" on:click={logout}> Logout </button>
	{:else}
		<form class="padding" on:submit|preventDefault>
			<input type="text" bind:value={$username} placeholder="Email" />
			<input type="password" bind:value={password} placeholder="Password" />
			<input type="text" bind:value={$server} />
			<button on:click={login}> Login </button>
		</form>
	{/if}
</div>

<style>
	.login {
		display: flex;
		flex-direction: row;
		align-items: center;
	}
	.padding {
		padding: 10px;
	}
</style>
