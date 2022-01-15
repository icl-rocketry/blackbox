<script lang="ts">
  $: counter = 0;
</script>

<main>
  <div class="bg-red-500">
    Count: {counter}
  </div>
  <button on:click={(_) => counter++}>Increment</button>
</main>

<style global>
  @tailwind utilities;
  @tailwind components;
  @tailwind base;
</style>
