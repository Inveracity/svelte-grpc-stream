<script>
	import '../app.css';
	import Channels from '../components/Channels.svelte';
	import Input from '../components/Input.svelte';
	import History from '../components/History.svelte';
	import Circle from '../components/Circle.svelte';
	import Controls from '../components/Controls.svelte';
	import { status } from '../store';
	import { BarLoader } from 'svelte-loading-spinners';
</script>

{#if $status === 'pending'}
	<div class="loader">
		<BarLoader color="#fff" />
	</div>
{:else if $status === 'connected'}
	<Channels />

	<div class="box">
		<History />
		<Circle color={$status === 'connected' ? 'green' : 'grey'} />
		<Input />
	</div>
{:else}
	<Controls />
{/if}

<style>
	.box {
		flex-grow: 1;
		display: flex;
		flex-direction: column;
		min-height: 0;
	}
</style>
