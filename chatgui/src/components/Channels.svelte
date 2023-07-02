<script lang="ts">
	import { channel } from '../store';

	let channels = ['general', 'spam'];
	let selectedChannel = 'general';

	const selectChannel = (e: any, channelName: string) => {
		e.preventDefault();
		console.log('selectedChannel: ' + channelName);
		channel.set(channelName);
	};
</script>

<div class="sidebar">
	<div class="channelItems">
		{#each channels as chan}
			<button class="button" on:click={(e) => selectChannel(e, chan)}>
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
</style>
