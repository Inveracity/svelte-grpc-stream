<script>
	import { onDestroy, onMount } from 'svelte';
	import '../app.css';
	import Login from '../components/Login.svelte';
	import Status from '../components/Status.svelte';
	import { currentUser, pb } from '$lib/pocketbase';
	import { Connect, Disconnect } from '$lib/grpc';
	import { server } from '../stores/server';

	onMount(() => {
		if ($currentUser) {
			Connect($server, $currentUser.username, '0', pb.authStore.token);
		}
	});
	onDestroy(() => {
		Disconnect();
	});
</script>

<div class="column">
	<div class="statusbar">
		<Status />
		<Login />
	</div>

	<slot />
</div>

<style>
	.statusbar {
		display: flex;
		background-color: #272727;
		border: 0px;
		border-radius: 4px;
		padding: 10px;
		margin: 0px 0;
		height: 30px;
		justify-content: space-between;
	}
</style>
