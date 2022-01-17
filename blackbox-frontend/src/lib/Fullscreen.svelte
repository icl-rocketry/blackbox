<script lang="ts">
  import FullscreenIcon from "svelte-material-icons/Fullscreen.svelte";

  import { getContext } from 'svelte';

  export let className: string;
  export let child: any;
  export let props: svelte.JSX.IntrinsicAttributes&{ [name: string]: any; };

  const { open } = getContext('simple-modal');

  function expand() {
    open(child, props)
  }
</script>


<div class={className}>
  <div
    class="flex flex-col object-cover min-h-full min-w-full items-start justify-center"
  >
    <div class="pr-3 self-end">
      <button on:click={expand}>
        <FullscreenIcon size="1.4em" />
      </button>
    </div>
    <svelte:component this={child} {...props}/>
  </div>
</div>
