<script>
	import '../app.css';
	import Channels from '../components/Channels.svelte';
	import Input from '../components/Input.svelte';
	import History from '../components/History.svelte';
	import { status } from '../stores/status';
	import Loader from '../components/Loader.svelte';
</script>

{#if $status === 'pending'}
	<Loader />
{:else if $status === 'connected'}
	<div class="row">
		<Channels />
		<div class="column">
			<History />
			<Input />
		</div>
	</div>
{/if}
