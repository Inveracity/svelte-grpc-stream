<script>
	// Stores
	import { status } from '$lib/stores/status';

	// Components
	import Channels from '$lib/components/Channels.svelte';
	import Input from '$lib/components/Input.svelte';
	import History from '$lib/components/History.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import Loading from '$lib/components/Loading.svelte';
	import Users from '$lib/components/Users.svelte';
</script>

<div class="flex flex-col h-screen w-screen">
	<div class="flex bg-secondary-content">
		<Navbar />
	</div>

	{#if $status === 'pending'}
		<Loading />
	{:else if $status === 'connected'}
		<div class="flex flex-row h-full w-full overflow-y-auto">
			<Users />
			<History />
			<Channels />
		</div>
		<div class="flex flex-row w-full">
			<Input />
		</div>
	{/if}
</div>
