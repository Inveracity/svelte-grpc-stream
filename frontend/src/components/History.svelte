<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';
	import { channel } from '../stores/channel';
	import { messages } from '../stores/messages';
  import { currentUser } from '$lib/pocketbase';

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

<div class="flex flex-col bg-slate-900 w-full h-full" bind:this={eventDiv}>
	{#each $messages as msg}
		{#if msg.channel === $channel}
			<div class="chat {msg.user === $currentUser?.username ? 'chat-end' : 'chat-start'} w-auto">
        <div class="chat-bubble {msg.user === $currentUser?.username ? 'chat-bubble-primary' : 'chat-bubble-secondary'}">
					<p>{msg.timestamp}</p>
					<p>{msg.user}</p>
					<p>{msg.message}</p>
				</div>
			</div>
		{/if}
	{/each}
</div>

<style>
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
