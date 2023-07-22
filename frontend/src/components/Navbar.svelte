<script lang="ts">
	import { server } from '$lib/stores/server';
	import { status } from '$lib/stores/status';
	import { RectangleListSolid, UsersSolid } from 'flowbite-svelte-icons';
	import { showUserList } from '$lib/stores/users';
	import { showChannelList } from '$lib/stores/channel';
	import { currentUser } from '$lib/pocketbase';
	import Logout from './Logout.svelte';

	$: highlight = $status !== 'connected' ? 'bg-red-400' : 'bg-green-400';
</script>

<div class="navbar bg-base-300">
	<div class="flex-1 justify-start gap-2">
		<div class="avatar">
			<div class="w-2 rounded-full {highlight}" />
		</div>
		<div class="btn btn-ghost upper-case text-xl">{$server}</div>
		{#if $currentUser}
			<div class="btn btn-ghost upper-case text-sm">{$currentUser?.name}</div>
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
		{/if}
	</div>
	{#if $currentUser}
		<div class="flex-none">
			<div class="justify-end">
				<Logout />
			</div>
		</div>
	{/if}
</div>
