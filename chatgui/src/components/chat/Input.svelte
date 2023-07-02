<script lang="ts">
	import { messages } from '../../store';
	import type { IMessage } from '../../types';
  import { channel } from '../../store';
	let message = '';

	const onKeyPress = (e: any) => {
		if (e.charCode === 13 && !e.shiftKey) {
			e.preventDefault();
			var msg: IMessage = { msg: message, channel: $channel };
			messages.add(msg);
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
