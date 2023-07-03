<script lang="ts">
	import { channel, username } from '../store';
	import { SendMessage } from '../grpc';
	import type { OutgoingMessage } from '../types';
	let message = '';

	const onKeyPress = (e: any) => {
		if (e.charCode === 13 && !e.shiftKey) {
			e.preventDefault();
			let msg: OutgoingMessage = { channelId: $channel, userId: $username, text: message };
			SendMessage(msg);
			message = '';
		}
	};
</script>

<div class="userInput">
	<textarea id="userinput" placeholder="Message" bind:value={message} on:keypress={onKeyPress} />
</div>

<style>
	.userInput {
		resize: vertical;
		flex: none;
	}
	::-webkit-resizer {
		display: none;
	}
	textarea#userinput {
		background-color: #333;
		color: #fff;
		width: 100%;
		border-radius: 0;
		border-style: none;
		margin-bottom: -3px;
		padding: 0;
		border: none;
		resize: none;
	}
	textarea#userinput:focus {
		outline: none;
	}
</style>
