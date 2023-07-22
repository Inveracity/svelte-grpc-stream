<script>
	import { server } from '$lib/stores/server';
	import { status } from '$lib/stores/status';
	import { currentUser } from '$lib/pocketbase';
	import { showUserList } from '$lib/stores/users';
	import { showChannelList } from '$lib/stores/channel';

	import { BarsSolid, RectangleListSolid, UsersSolid } from 'flowbite-svelte-icons';

	import Logout from './Logout.svelte';

	$: highlight = $status !== 'connected' ? 'bg-gray-400' : 'bg-green-400';
</script>

<div class="navbar bg-base-200">
	<div class="navbar-start">
		<div class="dropdown">
			<div class="m-2 lg:hidden">
				<BarsSolid />
			</div>

			<div
				class="menu menu-sm dropdown-content w-60 z-[1] p-2 shadow bg-neutral-300 rounded-box w-52"
			>
				{#if $currentUser}
					<div class="upper-case text-sm">{$currentUser?.name}</div>
					<div>
						<button
							class="flex w-full justify-start btn btn-secondary"
							on:click|preventDefault={() => showUserList.toggle()}
						>
							<UsersSolid /> user list
						</button>
					</div>
					<div>
						<button
							class="flex w-full justify-start btn btn-secondary"
							on:click|preventDefault={() => showChannelList.toggle()}
						>
							<RectangleListSolid /> channel list
						</button>
					</div>
				{/if}
			</div>
		</div>
		<div class="justify-center uppercase text-xl">
			{$server}
			<div class="avatar">
				<div class="w-2 rounded-full {highlight}" />
			</div>
		</div>
	</div>
	<div class="navbar-center hidden lg:flex">
		<ul class="menu menu-horizontal gap-2 px-1">
			{#if $currentUser}
				<div>
					<button class="btn btn-secondary" on:click|preventDefault={() => showUserList.toggle()}>
						<UsersSolid />
					</button>
				</div>
				<div>
					<button
						class="btn btn-secondary"
						on:click|preventDefault={() => showChannelList.toggle()}
					>
						<RectangleListSolid />
					</button>
				</div>
			{/if}
		</ul>
	</div>
	<div class="navbar-end">
		<div class="btn btn-ghost text-sm">{$currentUser?.name || ''}</div>
		<div class="divider divider-horizontal"></div>
		{#if $currentUser}
			<Logout />
		{/if}
	</div>
</div>
