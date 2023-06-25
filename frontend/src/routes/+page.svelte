<script lang="ts">
	import { Subscribe, Unsubscribe, SendNotification } from '../grpc';
	import { notifier } from '../store';
	import { status } from '../store';
	import { BarLoader } from 'svelte-loading-spinners';
	let userid = 'user1';
	let channelid = "channel1";
	let timestamp = "";
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

	<div class="events">
		<button style="float: right;" on:click={notifier.reset}> clear </button>
		{#each $notifier as item}
			<p>{item}</p>
		{/each}
	</div>
</div>

<style>
	.events {
		border: 1px solid black;
		height: 200px;
		overflow: scroll;
	}
	div {
		margin: 10px;
	}
	input {
		padding: 5px;
		margin: 10px;
		width: 210px;
	}
	button {
		padding: 5px;
		margin: 10px;
		width: 100px;
	}
	p {
		margin: 10px;
		font-family: 'Courier New', Courier, monospace;
	}
</style>
