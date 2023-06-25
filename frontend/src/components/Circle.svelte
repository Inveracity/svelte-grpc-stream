<script lang="ts">
    export let color: string = "red";
</script>

<div class="circle" id={color}></div>

<style>
.circle{
    -webkit-border-radius:8px;
    -moz-border-radius:8px;
    border-radius:8px;
    width:8px;
    height:8px;
    justify-content: center;
    text-align: center; /* This is not aligning the circle horizontally */
}

#red {
    background: rgb(189, 41, 41);
}

#grey {
    background: rgb(72, 72, 72);
}

#green {
    background: rgb(64, 168, 64);
}
</style>
