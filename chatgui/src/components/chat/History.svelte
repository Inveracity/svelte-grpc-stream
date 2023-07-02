<script lang="ts">
	import { messages, channel } from '../../store';
</script>

<div class="content">
	{#each $messages as msg}
    {#if msg.channel === $channel}
		<p class="chatline">
			user: {msg.msg}
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
