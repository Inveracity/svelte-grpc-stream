<script lang="ts">
  import { SendMessage } from '$lib/grpc';
  import { pb, writeChannel } from '$lib/pocketbase';
  import { channel, channels } from '$lib/stores/channel';

  let newChannelActive = false;
  let newChannelName = '';

  const selectChannel = (channelName: string) => {
    channel.set(channelName);
  };

  const newChannelActivate = () => {
    if ($channels.length >= 10) {
      return;
    }
    newChannelActive = true;
  };

  const addChannel = (name: string) => {
    channels.add(name);
    writeChannel(name);
    SendMessage({
      channelId: 'system',
      text: `channel_add ${name}`,
      userId: pb.authStore.model?.name || ''
    });
    newChannelActive = false;
    newChannelName = '';
  };
</script>

<div class="flex w-40 bg-neutral-focus overflow-auto">
  <div class="flex flex-col w-full">
    {#each $channels as chan}
      <button
        class="btn m-2 {chan === $channel ? 'btn-accent' : 'btn-accent btn-outline'}"
        on:click|preventDefault={() => selectChannel(chan)}
      >
        <p>{chan}</p>
      </button>
    {/each}
    {#if newChannelActive}
      <div class="m-1">
        <form class="form-control" on:submit|preventDefault={() => addChannel(newChannelName)}>
          <!-- svelte-ignore a11y-autofocus -->
          <input
            autofocus={true}
            class="input input-bordered input-md max-w-sm input-accent w-full"
            type="text"
            placeholder="new channel"
            maxlength="12"
            bind:value={newChannelName}
            on:focusout={() => (newChannelActive = false)}
          />
        </form>
      </div>
    {:else}
      <button class="btn btn-ghost" on:click|preventDefault={newChannelActivate}>
        <p>+</p>
      </button>
    {/if}
  </div>
</div>
