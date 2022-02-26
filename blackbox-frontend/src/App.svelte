<script lang="ts">
  import Data from "./Data.svelte";
</script>

<main>
  <div class="h-screen flex flex-col justify-between">
    <Data/>
    <div class="bg-blue-400 h-10">HELLO</div>
  </div>
</main>

<style global>
  @tailwind utilities;
  @tailwind components;
  @tailwind base;
</style>
