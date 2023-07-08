<script>
	import { currentUser, pb } from '$lib/pocketbase';
	import { Disconnect } from '../lib/grpc';

	function logout() {
		pb.authStore.clear();
		Disconnect();
	}
</script>

<div class="navbar bg-base-300">
	<div class="flex-1 justify-center">
		<div class="btn btn-ghost normal-case text-xl">chat</div>
	</div>
	<div class="flex-none gap-8">
		<div class="avatar">
			<div class="w-2 rounded-full bg-green-400" />
		</div>
		<button class="btn btn-neutral" on:click={logout}>Logout</button>
	</div>
</div>
