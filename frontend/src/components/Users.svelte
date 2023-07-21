<script>
	import { users } from '$lib/stores/users';
</script>

<div class="flex w-40 bg-neutral-focus overflow-auto">
	<div class="flex flex-col w-full">
		{#each $users as user}
			<div class="flex justify-center m-1 p-1">
				<div class="avatar {user.presence ? 'online' : 'offline'} justify-center">
					<div class="flex w-16 ring ring-neutral ring-offset-base-100 ring-offset-2">
						{user.name}
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
