<script>
    import {Subscribe} from '../grpc'
    import {notifier} from '../store'

    let subscriberId = "joe"

</script>

<div>
    <h1>GRPC Stream</h1>

    <div>
        <p>Click the Subscribe button receive realtime events</p>
    </div>

    <div>
        <input bind:value={subscriberId} placeholder="subscriber id" />
        <br />
        <button on:click={Subscribe}> Subscribe </button>
        <button on:click={notifier.reset}> clear </button>
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
        padding: 2px;
        margin: 10px;
    }
    button {
        padding: 2px;
        margin: 10px;
    }
    p {
        margin: 10px;
        font-family: 'Courier New', Courier, monospace;
    }
</style>