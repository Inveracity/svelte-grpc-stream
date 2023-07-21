<script lang="ts">
	import { server } from '$lib/stores/server';
	import { currentUser, pb } from '$lib/pocketbase';
	import { Connect } from '$lib/grpc';
	import Toast from './Toast.svelte';
	let toast: Toast;
	let password = '';
	let username = '';

	async function login() {
		await pb
			.collection('users')
			.authWithPassword(username, password)
			.then((_) => {
				toast.callToast('Login successful', 'success');
				Connect($server, username, '0', pb.authStore.token);
			})
			.catch((err) => {
				toast.callToast(err.message, 'error');
			});
	}
</script>

<div>
	<Toast bind:this={toast} />
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
