<script>
	// Stores
	import { status } from '$lib/stores/status';

	// Components
	import Channels from '../components/Channels.svelte';
	import Input from '../components/Input.svelte';
	import History from '../components/History.svelte';
	import Navbar from '../components/Navbar.svelte';
	import Loading from '../components/Loading.svelte';
	import Users from '../components/Users.svelte';
	import { showUserList } from '$lib/stores/users';
	import { showChannelList } from '$lib/stores/channel';
</script>

<div class="flex flex-col h-screen w-screen">
	<div class="flex bg-secondary-content">
		<Navbar />
	</div>

	{#if $status === 'pending'}
		<Loading />
	{:else if $status === 'connected'}
		<div class="flex flex-row h-full w-full overflow-y-auto">
			{#if $showUserList}
				<Users />
			{/if}
			<History />
			{#if $showChannelList}
				<Channels />
			{/if}
		</div>
		<div class="flex flex-row w-full">
			<Input />
		</div>
	{/if}
</div>
