<script>
    import { server } from '../stores/server';
    import { status } from '../stores/status';
	import Circle from "./Circle.svelte";
</script>

<div class="statusbar">
    <Circle color={$status === 'connected' ? 'green' : 'grey'}>
        <p> connected to server: {$server} </p>
    </Circle>

</div>

<style>
    .statusbar {
        background-color: #272727;
        border: 0px;
        border-radius: 4px;
        padding: 10px;
        margin: 0px 0;
    }
</style>
