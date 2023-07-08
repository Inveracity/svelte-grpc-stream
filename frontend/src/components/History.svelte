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

<div class="flex w-full" bind:this={eventDiv}>
	{#each $messages as msg}
		{#if msg.channel === $channel}
			<div class="chat chat-start">
				<div class="chat-bubble">
					<p>{msg.timestamp}</p>
					<p>{msg.user}</p>
					<p>{msg.message}</p>
				</div>
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
</style>
