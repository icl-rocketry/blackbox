<script lang="ts">

  import { getContext } from "svelte";

  export let className: string;
  export let child: any;
  export let props: svelte.JSX.IntrinsicAttributes & { [name: string]: any };

  const { open } = getContext("simple-modal");

  function expand() {
    open(child, props);
  }
</script>

<div class={className}>
  <div
    class="object-cover h-full min-h-full max-h-full min-w-full flex flex-1 justify-center items-center cursor-pointer"
    on:click={expand}
  >
    <svelte:component this={child} {...props} />
  </div>
</div>
