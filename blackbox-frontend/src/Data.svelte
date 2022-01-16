<script lang="ts">
  import Chart, { SimulatedChart } from "./lib/Chart.svelte";
  import { RollingWindowChart } from "./lib/Chart.svelte";

  const data = new RollingWindowChart(100, [
    {
      label: "X",
      colour: "#FF0000",
    },
  ]);

  const simulation = new SimulatedChart("Altitude", "green", [1, 2, 3, 4, 5, 6, 7, 8, 9])

  const d2 = new RollingWindowChart(7, [
    {
        label: "Dataset 1",
        colour: "#FF0000",
      },
      {
        label: "Dataset 2",
        colour: "#0000FF",
      },
  ])

  setInterval(() => {
    data.add(Math.random() * 10);
    simulation.add(Math.random() * 10);
  }, 50);
</script>

<div class="bg-slate-900 min-h-screen flex items-center justify-center">
  <div class="grid grid-cols-4 gap-2 min-h-screen min-w-full">
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="top_left" bind:data={$data} title="Acceleration" />
    </div>

    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="1" bind:data={$simulation}/>
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="2" bind:data={$d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="3" bind:data={$d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="4" bind:data={$d2} />
    </div>
    <div
      class="bg-slate-800 text-white rounded-xl row-span-2 col-span-2 min-h-full"
    >
      6
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="5" bind:data={$d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="6" bind:data={$d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="7" bind:data={$d2} />
    </div>
  </div>
</div>
