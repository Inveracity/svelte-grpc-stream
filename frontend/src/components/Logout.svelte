<script>
	import { currentUser, pb } from '$lib/pocketbase';
	import { Disconnect } from '../lib/grpc';
	function logout() {
		pb.authStore.clear();
		Disconnect();
	}
</script>

<!-- <p>Signed in as {$currentUser?.username}</p>
<button on:click={logout}> Logout </button> -->

<div>
	<p>testuser</p>
</div>
