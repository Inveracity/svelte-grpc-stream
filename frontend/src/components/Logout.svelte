<script>
	import { pb } from '$lib/pocketbase';
	import { Disconnect } from '$lib/grpc';
	import { ArrowRightFromBracketSolid } from 'flowbite-svelte-icons';

	function logout() {
		pb.authStore.clear();
		Disconnect();
	}
</script>

<div class="bg-neutral w-40">
	<ul class="menu menu-horizontal gap-3">
		<button class="flex-1 btn btn-ghost" on:click={logout}>
			<div class="h-5 w-5">
				<ArrowRightFromBracketSolid />
			</div>
		</button>
	</ul>
</div>
