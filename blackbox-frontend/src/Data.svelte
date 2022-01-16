<script lang="ts">
  import Chart from "./lib/Chart.svelte";
  import type { ChartData } from "./lib/Chart.svelte";

  const labels = [];
  const x = [];
  const n = 100;
  for (let i = 0; i < n; i++) {
    labels.push(i);
    x.push(Math.random() * 10);
  }

  let data: ChartData = {
    datasets: [
      {
        label: "Dataset 1",
        data: x,
        borderColor: "#FF0000",
      },
    ],
  };
  const d2: ChartData = {
    datasets: [
      {
        label: "Dataset 1",
        data: [1, 2, 3, 4, 5, 6, 7],
        borderColor: "#FF0000",
      },
      {
        label: "Dataset 2",
        data: [7, 6, 5, 4, 3, 2, 1],
        borderColor: "#0000FF",
      },
    ],
  };

  setInterval(() => {
    data = {
      datasets: [
        {
          label: data.datasets[0].label,
          data: [...data.datasets[0].data.slice(1), Math.floor(Math.random() * 10)],
          borderColor: data.datasets[0].borderColor,
        },
      ],
    };

  }, 50);
</script>

<div class="bg-slate-900 min-h-screen flex items-center justify-center">
  <div class="grid grid-cols-4 gap-2 min-h-screen min-w-full">
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="top_left" {data} title="Acceleration"/>
    </div>

    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="1" data={d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="2" data={d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="3" data={d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="4" data={d2} />
    </div>
    <div
      class="bg-slate-800 text-white rounded-xl row-span-2 col-span-2 min-h-full"
    >
      6
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="5" data={d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="6" data={d2} />
    </div>
    <div
      class="bg-slate-800 text-white flex justify-center items-center rounded-xl"
    >
      <Chart id="7" data={d2} />
    </div>
  </div>
</div>
