<script lang="ts">
  import Slider from "@bulatdashiev/svelte-slider";
  let value = [50];
</script>

<div class="bg-slate-900 h-20 flex flex-row justify-center">
    <div class="flex flex-col justify-center w-full">
        <Slider min="0" max="100" step="1" bind:value/>
    </div>
</div>

<style>
  :root {
    --track-bg: #ebebeb;
    --progress-bg: red;
    --thumb-bg: red;
  }
</style>
