<script lang="ts">
	import { Subscribe, Unsubscribe, SendNotification } from '../grpc';
	import { status } from '../store';
	let userid = 'user1';
	let channelid = "channel1";
</script>

<div>
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
