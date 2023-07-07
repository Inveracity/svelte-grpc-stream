<script>
	import { onDestroy, onMount } from 'svelte';
	import '../app.css';
	import Login from '../components/Login.svelte';
	import Status from '../components/Status.svelte';
	import { currentUser } from '$lib/pocketbase';
	import { Connect, Disconnect } from '../grpc';
	import { server } from '../stores/server';
	import { username } from '../stores/username';
	import { messages } from '../stores/messages';

	onMount(() => {
		if ($currentUser) {
			Connect($server, $username, '0');
		}
	});
	onDestroy(() => {
		messages.reset();
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
