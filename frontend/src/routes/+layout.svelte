<script>
	import '../app.css';
	import Circle from '../components/Circle.svelte';
	import { status } from '../store';
	import { BarLoader } from 'svelte-loading-spinners';
</script>

<nav>
	<div>
		<h2>GRPC Stream</h2>
	</div>
	<Circle color={$status === 'connected' ? 'green' : 'grey'} />
	<div style="visibility: {$status !== 'pending' ? 'hidden' : 'visible'} ">
		<BarLoader color="#ffc800" />
	</div>
</nav>

<slot />
