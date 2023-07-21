<script lang="ts">
	import { channel } from '$lib/stores/channel';
	import { status } from '$lib/stores/status';
	import { currentUser, pb } from '$lib/pocketbase';
	import { SendMessage } from '$lib/grpc';
	import type { OutgoingMessage } from '$lib/types';

	let message = '';

	const onKeyPress = (e: any) => {
		if (e.charCode === 13 && !e.shiftKey) {
			e.preventDefault();
			let msg: OutgoingMessage = {
				channelId: $channel,
				userId: $currentUser?.username,
				text: message,
				jwt: pb.authStore.token
			};
			SendMessage(msg);
			message = '';
		}
	};
</script>

<div class="w-full">
	<div class="m-2">
		<textarea
			class="textarea textarea-secondary w-full h-12"
			placeholder={$status === 'connected' ? 'Message' : '⛔'}
			bind:value={message}
			on:keypress={onKeyPress}
			disabled={$status !== 'connected'}
		/>
	</div>
</div>

<style>
	::-webkit-resizer {
		display: none; /* remove the resize handle on the bottom right */
	}
</style>
