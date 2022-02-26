<script lang="ts">
  import Slider from "@bulatdashiev/svelte-slider";
  let value = 50;
</script>

<div class="bg-slate-900 h-20 flex flex-row justify-center">
  <div class="slidecontainer w-full flex flex-col justify-center">
    <input type="range" min="1" max="100" {value} class="slider" id="myRange" />
  </div>
</div>

<style>
  .slider {
    -webkit-appearance: none;
    width: 100%;
    height: 5px;
    border-radius: 5px;
    background: #d3d3d3;
    outline: none;
    opacity: 0.7;
    -webkit-transition: 0.2s;
    transition: opacity 0.2s;
  }

  .slider::-webkit-slider-thumb {
    -webkit-appearance: none;
    appearance: none;
    width: 15px;
    height: 15px;
    border-radius: 50%;
    background: red;
    cursor: pointer;
  }

  .slider::-moz-range-thumb {
    width: 15px;
    height: 15px;
    border-radius: 50%;
    background: red;
    cursor: pointer;
  }
</style>
