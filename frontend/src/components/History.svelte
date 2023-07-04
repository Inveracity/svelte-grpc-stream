<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';
	import { channel } from '../stores/channel';
	import { messages } from '../stores/messages';

	let eventDiv: HTMLDivElement;
	let autoscroll = false;

	beforeUpdate(() => {
		if (eventDiv) {
			const scrollableDistance = eventDiv.scrollHeight - eventDiv.offsetHeight;
			autoscroll = eventDiv.scrollTop > scrollableDistance - 20;
		}
	});

	afterUpdate(() => {
		if (autoscroll) {
			eventDiv.scrollTo(0, eventDiv.scrollHeight);
		}
	});
</script>

<div class="content" bind:this={eventDiv}>
	{#each $messages as msg}
		{#if msg.channel === $channel}
			<div class="chatline">
				<p id="timestamp"> {msg.timestamp} </p>
				<p id="username"> {msg.user} </p>
				<p> {msg.message} </p>
			</div>
		{/if}
	{/each}
</div>

<style>
	.content {
		flex: 1;
		padding: 10px;
		overflow-y: scroll;
	}
	::-webkit-scrollbar {
		width: 12px;
	}
	::-webkit-scrollbar-track {
		background-color: rgba(0, 0, 0, 0.2);

	}
	::-webkit-scrollbar-thumb {
		border-radius: 10px;
		background-color: rgba(183, 110, 255, 0.491);
	}
	.chatline {
		display: flex;
		flex-direction: row;
		padding: 0px;
		margin: 0px;
		white-space: pre-line;
		align-items: baseline;
	}
	#timestamp {
		display: flex;
		font-size: 0.8em;
		margin-right: 10px;
		color: #999;
		align-items: center;
		width: 150px;
	}
	#username {
		display: inline-block;
		width: auto;
		max-width: 400px;
		min-width: 150px;
		margin-right: 10px;
		color: #9d87e3;
		align-items: center;
	}
</style>
