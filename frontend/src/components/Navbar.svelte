<script>
	import { currentUser, pb } from '$lib/pocketbase';
	import { Disconnect } from '../lib/grpc';
	import { server } from '../stores/server';

	function logout() {
		pb.authStore.clear();
		Disconnect();
	}
</script>

<div class="navbar bg-base-300">
	<div class="flex-1 justify-start">
		<div class="avatar">
			<div class="w-2 rounded-full bg-green-400" />
		</div>
		<div class="btn btn-ghost upper-case text-xl">{$server}</div>
	</div>
</div>
