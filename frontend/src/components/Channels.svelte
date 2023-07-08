<script lang="ts">
	import { channel } from '../stores/channel';

	let channels = ['general', 'spam'];

	const selectChannel = (e: any, channelName: string) => {
		e.preventDefault();
		console.log('selectedChannel: ' + channelName);
		channel.set(channelName);
	};
</script>

<div class="flex w-40">
	<div class="flex flex-col w-full">
		{#each channels as chan}
			<!-- disabled is set to true when channel is selected -->
			<button
				class="btn {chan === $channel ? 'btn-secondary' : 'btn-secondary btn-outline  '} m-2"
				on:click={(e) => selectChannel(e, chan)}
			>
				<p>{chan}</p>
			</button>
		{/each}
	</div>
</div>
