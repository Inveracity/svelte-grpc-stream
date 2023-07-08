<script lang="ts">
	import { channel } from '../stores/channel';
	import { currentUser } from '$lib/pocketbase';
	import { status } from '../stores/status';
	import { SendMessage } from '../lib/grpc';
	import type { OutgoingMessage } from '../types';
	let message = '';

	const onKeyPress = (e: any) => {
		if (e.charCode === 13 && !e.shiftKey) {
			e.preventDefault();
			let msg: OutgoingMessage = {
				channelId: $channel,
				userId: $currentUser?.username,
				text: message
			};
			SendMessage(msg);
			message = '';
		}
	};
</script>

<div class="w-full">
	<textarea
    class="textarea textarea-secondary w-full h-12"
		placeholder={$status === 'connected' ? 'Message' : '⛔'}
		bind:value={message}
		on:keypress={onKeyPress}
    disabled={$status !== 'connected'}
	/>
</div>

<style>
	::-webkit-resizer {
		display: none; /* remove the resize handle on the bottom right */
	}
</style>
