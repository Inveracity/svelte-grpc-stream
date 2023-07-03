<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';
	import { messages, channel } from '../store';

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
			<p class="chatline">
				{msg.user}: {msg.message}
			</p>
		{/if}
	{/each}
</div>

<style>
	.content {
		flex: 1;
		padding: 10px;
		overflow-y: scroll;
	}
	.chatline {
		padding: 0px;
		white-space: pre-line;
	}
</style>
