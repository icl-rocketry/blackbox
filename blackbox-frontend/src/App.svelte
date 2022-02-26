<script lang="ts">
  import Data from "./Data.svelte";
import Footer from "./Footer.svelte";
</script>

<main>
  <div class="h-screen flex flex-col justify-between">
    <Data/>
    <Footer/>
  </div>
</main>

<style global>
  @tailwind utilities;
  @tailwind components;
  @tailwind base;
</style>
