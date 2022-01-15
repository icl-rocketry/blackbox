<script>
  import Chart from "./lib/Chart.svelte";
</script>

<div class="bg-slate-900 min-h-screen flex items-center justify-center">
  <div class="grid grid-cols-4 gap-2 min-h-screen min-w-full">
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl">
      <Chart id="top_left"/>
    </div>

    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="1"/></div>
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="2"/></div>
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="3"/></div>
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="4"/></div>
    <div class="bg-slate-800 text-white rounded-xl row-span-2 col-span-2 min-h-full">
      6
    </div>
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="5"/></div>
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="6"/></div>
    <div class="bg-slate-800 text-white flex justify-center items-center rounded-xl"><Chart id="7"/></div>
  </div>
</div>
