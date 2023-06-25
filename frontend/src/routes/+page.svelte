<script lang="ts">
	import '../app.css';
	import { beforeUpdate, afterUpdate } from 'svelte';
	import { Subscribe, Unsubscribe, SendNotification } from '../grpc';
	import { notifier } from '../store';
	import { status } from '../store';
	import { BarLoader } from 'svelte-loading-spinners';

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

	let userid = 'user1';
	let channelid = "channel1";
</script>

<div>
	<h1>GRPC Stream</h1>

	<div>
		<p>status: {$status ? $status : 'disconnected'}</p>

		<div style="visibility: {$status !== 'pending' ? 'hidden' : 'visible'} ">
			<BarLoader color="#ffc800" />
		</div>

		<input
			bind:value={userid}
			placeholder="user id"
			disabled={$status === 'connected' || $status === 'pending'}
		/>

		<input
			bind:value={channelid}
			placeholder="channel id"
			disabled={$status === 'connected' || $status === 'pending'}
		/>

		{#if $status === 'connected'}
			<button on:click={Unsubscribe}> Unsubscribe </button>
		{:else}
			<button on:click={() => Subscribe(channelid, userid, "")} disabled={userid === '' || $status === 'connected' || $status === 'pending'}> Subscribe </button>
		{/if}
	</div>

	<div>
		<input bind:value={channelid} placeholder="who should receive this event?" />
		<button disabled={channelid === ''} on:click={() => SendNotification(channelid, userid,'notification!')}> notify </button>
		<button disabled={channelid === ''} on:click={() => SendNotification(channelid, userid,`who would've thought?`)}> who </button>
		<button disabled={channelid === ''} on:click={() => SendNotification(channelid, userid,'not me!')}> not me </button>
	</div>
	<hr />

	<div class="notifications">
		<div>
			<button on:click={notifier.reset}> clear </button>
		</div>
		<div class="events" bind:this={eventDiv}>
			{#each $notifier as msg}
			<p>{msg}</p>
			{/each}
		</div>
	</div>
</div>

<style>
	.notifications {
		width: 80%;
	}
	.events {
		width: auto;
		height: 400px;
		overflow-y: scroll;
	}
	.events::-webkit-scrollbar {
		width: 0px !important; /*removes the scrollbar but still scrollable*/
	}
</style>
