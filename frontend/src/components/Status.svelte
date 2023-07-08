<script>
	import { server } from '../stores/server';
	import { status } from '../stores/status';
</script>


<div class="badge {$status === 'connected' ? 'badge-success' : 'badge-neutral'} w-auto h-full m-1">
  {$server}
</div>

