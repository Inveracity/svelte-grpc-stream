<script lang="ts">
	import Controls from '../components/Controls.svelte';
	import Notifications from '../components/Notifications.svelte';
</script>

<div class="layout">
	<div style="width: 60%">
		<Controls />
	</div>
	<hr />
	<div style="width: 40%">
		<Notifications />
	</div>
</div>

<style>
	.layout {
		display: flex;
		justify-content: space-between;
	}
</style>
