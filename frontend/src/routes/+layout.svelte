<script>
	import '../app.css';
	import { onDestroy, onMount } from 'svelte';
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

<slot />
