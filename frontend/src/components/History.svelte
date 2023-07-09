<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';
	import { channel } from '../stores/channel';
	import { messages } from '../stores/messages';
	import { currentUser } from '$lib/pocketbase';
	import { FaceLaughSolid } from 'flowbite-svelte-icons';

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

<div class="flex flex-col w-full h-full overflow-y-auto" bind:this={eventDiv}>
	{#each $messages as msg}
		{#if msg.channel === $channel}
			<div
				class="chat {msg.user === $currentUser?.username ? 'chat-end' : 'chat-start'} m-2 w-auto"
			>
				<div class="chat-image avatar">
					<div class="w-10 rounded-full">
						<div class="justify-center items-center flex overflow-hidden h-full bg-accent-focus">
							<FaceLaughSolid />
						</div>
					</div>
				</div>
				<div class="chat-header">
					{msg.user}
					<time class="text-xs opacity-50">{msg.timestamp}</time>
				</div>
				<div
					class="chat-bubble whitespace-pre-line {msg.user === $currentUser?.username
						? 'chat-bubble-primary'
						: 'chat-bubble-secondary'}"
				>
					{msg.message}
				</div>
			</div>
		{/if}
	{/each}
</div>
