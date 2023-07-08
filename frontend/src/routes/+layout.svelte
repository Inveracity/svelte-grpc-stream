<script>
	import '../app.css';
	import { onDestroy, onMount } from 'svelte';
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

<div class="w-screen h-screen">
	<div class="flex flex-col">
		<div class="flex flex-row justify-between">
			<Status />
			<Login />
		</div>

		<div class="flex-1">
			<slot />
		</div>
	</div>
</div>
