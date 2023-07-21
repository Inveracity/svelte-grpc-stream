<script lang="ts">
	let message = '';
	let alerttype = 'alert-error';

	// https://akashmittal.com/svelte-calling-function-child-parent/
	// call from parent as:
	// let toast: Toast;
	// toast.callToast(message, alerttype);
	// <Toast bind:this={toast} />

	export function callToast(msg: string, type?: string | undefined, ms?: number | undefined) {
		console.log('callToast()');
		let milliseconds: number = 8000;
		if (ms) {
			milliseconds = ms;
		}
		if (type) {
			alerttype = `alert-${type}`;
		}
		message = msg;

		setTimeout(() => {
			message = '';
		}, milliseconds);
	}
</script>

<div class="toast toast-center z-40">
	{#if message}
		<div class="alert {alerttype}">
			<div>
				<span>{message}</span>
			</div>
		</div>
	{/if}
</div>
