<script>
    import {Subscribe, Unsubscribe} from '../grpc'
    import {notifier} from '../store'
    import {status} from '../store'

    let subscriberId = "joe"

</script>

<div>
    <h1>GRPC Stream</h1>

    <div>
        <p>Click the Subscribe button receive realtime events</p>
    </div>

    <div>
        <input bind:value={subscriberId} placeholder="subscriber id" disabled={$status === "connected"} />
        <br />
        {#if $status === 'connected'}
            <button on:click={() => Unsubscribe()}> Unsubscribe </button>
        {:else}
            <button on:click={() => Subscribe(subscriberId)}> Subscribe </button>
        {/if}
        <button on:click={notifier.reset}> clear </button>
        <p> status: {$status}</p>
    </div>

    <div>
        {#each $notifier as item}
            <p>{item}</p>
        {/each}
    </div>
</div>


<style>
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