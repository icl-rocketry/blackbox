<script lang="ts">
  import type { Readable } from "svelte/store";
  export let generator: Readable<number>;
  export let units: string;
  export let fullscreen = false;
</script>

<p class="text-white {fullscreen ? "text-9xl" : "text-6xl"} font-mono text-center overflow-hidden">{$generator} {units}</p>
