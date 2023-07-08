<script>
	import { server } from '../stores/server';
	import { status } from '../stores/status';
</script>

<div>
  <div class="badge {$status === 'connected' ? 'badge-succes' : 'badge-neutral'} w-40">{$server}</div>
</div>
