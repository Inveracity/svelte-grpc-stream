<script>
	import { pb } from '$lib/pocketbase';
	import { Disconnect } from '$lib/grpc';
	import { ArrowRightFromBracketSolid } from 'flowbite-svelte-icons';

	function logout() {
		pb.authStore.clear();
		Disconnect();
	}
</script>

<div>
	<ul class="menu menu-horizontal gap-3">
		<button class="flex gap-2" on:click={logout}>
			<p>sign out</p>
			<div class="h-5 w-5">
				<ArrowRightFromBracketSolid />
			</div>
		</button>
	</ul>
</div>
