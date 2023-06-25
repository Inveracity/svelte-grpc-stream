<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';
	import { notifier } from '../store';
	import Messages from './Messages.svelte';

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

<div class="notifications">
	<div>
		<button on:click={notifier.reset}> clear </button>
	</div>
	<div class="events" bind:this={eventDiv}>
		{#each $notifier as msg}
			<Messages {msg} />
		{/each}
	</div>
</div>

<style>
	.events {
		height: 400px;
		overflow-y: scroll;
	}
	.events::-webkit-scrollbar {
		width: 0px !important; /*removes the scrollbar but still scrollable*/
	}
</style>
