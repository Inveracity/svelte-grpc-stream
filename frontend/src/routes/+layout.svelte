<script>
    import Circle from '../components/Circle.svelte'
    import { status } from '../store';

</script>
<nav>
    <div>
        <h2>GRPC Stream</h2>
    </div>
    <div style="display: flex">
            <Circle color={$status === "connected" ? "green" : "grey"} /> <p> {$status} </p>
    </div>
</nav>

<slot />
