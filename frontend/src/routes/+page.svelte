<script>
	import Channels from '../components/Channels.svelte';
	import Input from '../components/Input.svelte';
	import History from '../components/History.svelte';
	import { status } from '../stores/status';
</script>

{#if $status === 'pending'}
<span class="loading loading-dots loading-lg"></span>
{:else if $status !== 'connected'}
	<div class="flex flex-row h-full w-full">
		<Channels />
		<div class="flex flex-col w-full">
			<History />
			<Input />
		</div>
	</div>
{/if}
