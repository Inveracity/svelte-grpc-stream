<script lang="ts">
    import { Subscribe, Unsubscribe} from '../grpc'
    import { notifier } from '../store'
    import { status } from '../store'
    import { BarLoader } from 'svelte-loading-spinners';
    import { send_notification } from '../api'
    let subscriberId = ""
    let receiverId = ""
</script>

<div>
    <h1>GRPC Stream</h1>

    <div>
        <p> status: {$status ? $status : "disconnected"}</p>

        <div style="visibility: {$status !== 'pending' ? 'hidden' : 'visible'} ">
            <BarLoader color="#ffc800" />
        </div>
        
        <input bind:value={subscriberId} placeholder="subscriber id" disabled={$status === "connected" || $status === "pending"} />
        
        {#if $status === 'connected'}
            <button on:click={() => Unsubscribe()}> Unsubscribe </button>
        {:else}
            <button on:click={() => Subscribe(subscriberId)} disabled={subscriberId === "" && ($status === "connected" || $status === "pending")}> Subscribe </button>
        {/if}
        
        
    </div>

    <!-- TODO: Make the button a component -->
    <div>
        <input bind:value={receiverId} placeholder="who should receive this event?"/>
        <button 
            disabled={receiverId === ""} 
            on:click={() => send_notification({
                subid: receiverId, 
                text: "notification!", 
                sender: subscriberId
            })}> 
                notify 
        </button>
        <button 
            disabled={receiverId === ""} 
            on:click={() => send_notification({
                subid: receiverId, 
                text: "who would've thought?", 
                sender: subscriberId
            })}> 
                who? 
        </button>
        <button 
            disabled={receiverId === ""} 
            on:click={() => send_notification({
                subid: receiverId, 
                text: "not me!", 
                sender: subscriberId
            })}> 
                not me! 
        </button>
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