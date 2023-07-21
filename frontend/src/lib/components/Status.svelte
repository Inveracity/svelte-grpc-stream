<script>
	import { server } from '$lib/stores/server';
	import { status } from '$lib/stores/status';
</script>

<div class="badge {$status === 'connected' ? 'badge-success' : 'badge-neutral'} w-auto h-full">
	{$server}
</div>
