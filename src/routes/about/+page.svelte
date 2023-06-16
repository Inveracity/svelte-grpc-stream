<script>
    import {Subscribe} from '../../client'
    import {notifier} from '../../store'
</script>

<div>
    <p>This is the About page</p>
</div>

<button on:click={Subscribe}>
    Subscribe
</button>
<div>
    <p>{$notifier}</p>
</div>
<div>
    <p><a href="/">Home</a></p>
</div>

