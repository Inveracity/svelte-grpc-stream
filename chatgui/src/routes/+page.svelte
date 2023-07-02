<script>
	import '../app.css';
  import Channels from '../components/Channels.svelte';
  import Input from '../components/chat/Input.svelte';
	import History from '../components/chat/History.svelte';
</script>

<Channels />

<div class="box">
  <History />
  <Input />
</div>

<style>
	.box {
		flex-grow: 1;
		display: flex;
		flex-direction: column;
		min-height: 0;
	}
</style>
