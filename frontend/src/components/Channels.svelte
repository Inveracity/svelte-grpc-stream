<script lang="ts">
	import { channel } from '../store';

	let channels = ['general', 'spam'];

	const selectChannel = (e: any, channelName: string) => {
		e.preventDefault();
		console.log('selectedChannel: ' + channelName);
		channel.set(channelName);
	};
</script>

<div class="sidebar">
	<div class="channelItems">
		{#each channels as chan}
			<!-- disabled is set to true when channel is selected -->
			<button class="button" disabled={chan === $channel} on:click={(e) => selectChannel(e, chan)}>
				<p>{chan}</p>
			</button>
		{/each}
	</div>
</div>

<style>
	.sidebar {
		background-color: #333;
		width: 150px;
		height: 100%;
		flex: none;
	}
	.channelItems {
		padding: 10px;
	}
	.button {
		cursor: pointer;
		width: 100%;
		background-color: #333;
		color: white;
		border: none;
	}
	button:hover {
		background-color: #a5a5a5;
	}
	button:disabled {
		background-color: #636363;
	}
</style>
