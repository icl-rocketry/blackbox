<script context="module" lang="ts">
  export type ChartData = {
    datasets: {
      label: string;
      data: number[];
      borderColor: string;
    }[];
  };

  export interface LineChart {
    add(values: number[]): void;
    getData(): ChartData;
  }

  export class RollingWindowChart implements LineChart {
    constructor(maxSize: number, lines: { label: string; colour: string }[]) {
      //@ts-ignore
      this.maxSize = maxSize;
      //@ts-ignore
      this.data = {
        datasets: lines.map((line) => ({
          label: line.label,
          data: [],
          borderColor: line.colour,
        })),
      };
    }


    public add(values: number[]): void {
      //@ts-ignore
      for (let i = 0; i < this.data.datasets.length; i++) {
        //@ts-ignore
        const dataset = this.data.datasets[i];

        dataset.data.push(values[i]);
        //@ts-ignore
        if (dataset.data.length > this.maxSize) {
          dataset.data.shift();
        }
      }
    }

    public getData(): ChartData {
      //@ts-ignore
      return this.data;
    }
  }
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
<script lang="ts">
  import { afterUpdate, onMount } from "svelte";
  import { Chart, registerables } from "chart.js";
  Chart.register(...registerables);

  export let id: string;
  export let data: LineChart;
  export let title = "Title";
  export const xMax: number = 100;
  export const xMin: number = 0;
  export const yMax: number = 10;
  export const yMin: number = -10;

  let chart: Chart;

  function renderChart() {
    let ctx = document.getElementById(id);

    const chart_data = {
      labels: Array.from(Array(xMax - xMin)).map((_, i) => i + xMin),
      datasets: data.getData().datasets.map(({ label, data, borderColor }) => ({
        label,
        data,
        borderColor,
        backgroundColor: "transparent",
        radius: 0,
        tension: 0.1,
        borderWidth: 1,
      })),
    };

    const config = {
      type: "line",
      data: chart_data,
      options: {
        toolTip: false,
        animation: false,
        scales: {
          xAxis: {
            max: xMax,
            min: xMin,
            grid: {
              drawBorder: false,
              display: false,
            },
            type: "linear",
          },
          yAxis: {
            max: yMax,
            min: yMin,
            grid: {
              drawBorder: false,
              display: false,
            },
          },
        },
        plugins: {
          legend: {
            position: "bottom",
          },
          title: {
            display: true,
            text: title,
            color: "#FFFFFF",
          },
        },
      },
    };
    return new Chart(ctx as any, config as any);
  }

  onMount(() => {
    chart = renderChart();
  });
  afterUpdate(() => {
    const chart_data = chart.data;
    if (data.getData().datasets.length > 0) {
      chart_data.datasets.forEach((dataset: { data: any[] }, index: number) => {
        dataset.data = data.getData().datasets[index].data;
      });

      chart.update();
    }
  });
</script>

<canvas {id} class="object-cover max-h-full max-w-full" />
