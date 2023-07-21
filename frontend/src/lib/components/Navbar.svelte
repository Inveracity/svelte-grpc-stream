<script>
	import { server } from '$lib/stores/server';
	import Login from './Login.svelte';
</script>

<div class="navbar bg-base-300">
	<div class="flex-1 justify-start">
		<div class="avatar">
			<div class="w-2 rounded-full bg-green-400" />
		</div>
		<div class="btn btn-ghost upper-case text-xl">{$server}</div>
	</div>
	<div class="justify-end">
		<Login />
	</div>
</div>
