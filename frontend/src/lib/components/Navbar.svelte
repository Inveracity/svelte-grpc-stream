<script>
	import { server } from '$lib/stores/server';
	import { showUserList } from '$lib/stores/users';
	import { RectangleListSolid, UsersSolid } from 'flowbite-svelte-icons';
	import Login from './Login.svelte';
	import { showChannelList } from '$lib/stores/channel';
</script>

<div class="navbar bg-base-300">
	<div class="flex-1 justify-start gap-2">
		<div class="avatar">
			<div class="w-2 rounded-full bg-green-400" />
		</div>
		<div class="btn btn-ghost upper-case text-xl">{$server}</div>
		<div>
			<button class="btn btn-secondary" on:click|preventDefault={() => showUserList.toggle()}>
				<UsersSolid />
			</button>
		</div>
		<div>
			<button class="btn btn-secondary" on:click|preventDefault={() => showChannelList.toggle()}>
				<RectangleListSolid />
			</button>
		</div>
	</div>
	<div class="flex-none">
		<div class="justify-end">
			<Login />
		</div>
	</div>
</div>
