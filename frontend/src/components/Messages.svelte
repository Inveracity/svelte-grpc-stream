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
		/* Add shadows to create the "card" effect */
		box-shadow: 0 4px 8px 0 rgba(230, 230, 230, 0.042);
		transition: 0.3s;
		margin: 10px;
		padding: 10px;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: stretch;
	}

	/* On mouse-over, add a deeper shadow */
	.card:hover {
		box-shadow: 0 8px 16px 0 rgba(222, 222, 222, 0.2);
	}

	/* Add some padding inside the card container */
	.container {
		padding: 2px 16px 2px 16px;
		margin: 10px 10px 10px 10px;
		background-color: rgb(40, 56, 79);
	}
</style>
