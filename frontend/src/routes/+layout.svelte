<script>
	import '../app.css';
	import { onDestroy, onMount } from 'svelte';
	import { currentUser, fetchChannels, fetchUsers, pb } from '$lib/pocketbase';
	import { Connect, Disconnect } from '$lib/grpc';
	import { server } from '$lib/stores/server';

	onMount(() => {
		if ($currentUser) {
			fetchChannels();
			fetchUsers();
			Connect($server, $currentUser.username, '0', pb.authStore.token);
		}
	});
	onDestroy(() => {
		Disconnect();
	});
</script>

<slot />
