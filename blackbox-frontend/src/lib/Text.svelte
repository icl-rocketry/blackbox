<script lang="ts">
  import type { Readable } from "svelte/store";
  export let generator: Readable<number>;
  export let units: string;
</script>

<p class="text-white text-6xl font-mono text-center overflow-hidden">{$generator} {units}</p>
