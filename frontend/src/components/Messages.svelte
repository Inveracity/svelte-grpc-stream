<script lang="ts">
	import type { Message } from '../types';

	export let msg: Message;
</script>

<div class="card">
	<div class="header">
		<p>{msg.timestamp}</p>
		<p>{msg.channel}</p>
		<p>{msg.user}</p>
	</div>
	<div class="container">
		<h4><b>{msg.message}</b></h4>
	</div>
</div>

<style>
	.card {
		box-shadow: 0 4px 8px 0 rgba(230, 230, 230, 0.042);
		transition: 0.3s;
		background-color: rgb(26, 32, 41);
		margin: 10px;
		padding: 10px;
		border-radius: 10px;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: stretch;
	}

	.card:hover {
		box-shadow: 0 8px 16px 0 rgba(222, 222, 222, 0.2);
	}

	.container {
		padding: 0px 5px 0px 5px;
	}
</style>
