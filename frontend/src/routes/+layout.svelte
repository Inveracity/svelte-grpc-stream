<script>
	import '../app.css';
	import { onDestroy, onMount } from 'svelte';
	import Login from '../components/Login.svelte';
	import Status from '../components/Status.svelte';
	import { currentUser, pb } from '$lib/pocketbase';
	import { Connect, Disconnect } from '$lib/grpc';
	import { server } from '../stores/server';
	import Navbar from '../components/Navbar.svelte';

	onMount(() => {
		if ($currentUser) {
			Connect($server, $currentUser.username, '0', pb.authStore.token);
		}
	});
	onDestroy(() => {
		Disconnect();
	});
</script>

<div class="flex flex-col h-screen">
	<div class="flex flex-row justify-between bg-secondary-content">
		<Navbar />
	</div>

	<slot />
</div>
