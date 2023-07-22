<script>
	import '../app.css';
	import { onDestroy, onMount } from 'svelte';
	import { currentUser, fetchChannels, fetchUsers } from '$lib/pocketbase';
	import { Connect, Disconnect } from '$lib/grpc';
	import { server } from '$lib/stores/server';

	onMount(() => {
		if ($currentUser) {
			// if user is already logged in, then connect to the server immediately
			Connect($server, $currentUser.username, '0');
			fetchChannels();
			fetchUsers();
		}
	});
	onDestroy(() => {
		Disconnect();
	});
</script>

<slot />
