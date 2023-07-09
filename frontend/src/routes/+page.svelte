<script>
	import Channels from '../components/Channels.svelte';
	import Input from '../components/Input.svelte';
	import History from '../components/History.svelte';
	import { status } from '../stores/status';
	import Navbar from '../components/Navbar.svelte';
	import Logout from '../components/Logout.svelte';
</script>

<div class="flex flex-row justify-between bg-secondary-content">
	<Navbar />
</div>

{#if $status === 'pending'}
	<span class="loading loading-dots loading-lg" />
{:else if $status === 'connected'}
	<div class="flex flex-row h-full w-full">
		<History />
		<Channels />
	</div>
	<div class="flex flex-row w-full">
		<Input />
		<Logout />
	</div>
{/if}
