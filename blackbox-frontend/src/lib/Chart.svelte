<script context="module" lang="ts">
  export type ChartData = {
    datasets: {
      label: string;
      data: number[];
      colour: string;
    }[];
  };

  //LineChart is an abstract class that manages the data of a line chart.
  //It is also a Svelte store
  abstract class LineChart {
    protected data: ChartData;
    private readonly subscribers: ((data: ChartData) => void)[] = [];

    abstract add(...values: number[]): void;

    public subscribe(callback: (data: ChartData) => void): () => void {
      const last = this.subscribers.length;
      this.subscribers.push(callback);
      callback(this.data);

      return () => {
        this.subscribers.splice(last, 1);
      };
    }

    public set(v: ChartData): void {
      this.data = v;
      this.subscribers.forEach((fn) => fn(this.data));
    }
  }

  export class RollingWindowChart extends LineChart {
    private maxSize: number;

    constructor(maxSize: number, lines: { label: string; colour: string }[]) {
      super();
      this.maxSize = maxSize;
      this.data = {
        datasets: lines.map((line) => ({
          label: line.label,
          data: [],
          colour: line.colour,
        })),
      };
    }

    public add(...values: number[]): void {
      if (this.data.datasets.length !== values.length) {
        throw new Error(
          `RollingWindowChart: number of values (${values.length}) does not match the number of lines (${this.data.datasets.length})`
        );
      }
      for (let i = 0; i < this.data.datasets.length; i++) {
        const dataset = this.data.datasets[i];

        dataset.data.push(values[i]);
        if (dataset.data.length > this.maxSize) {
          dataset.data.shift();
        }
      }
      this.set(this.data);
    }
  }

  export class SimulatedChart extends LineChart {
    constructor(name: string, colour: string, simulated_data: number[]) {
      super();
      this.data = {
        datasets: [
          {
            label: name,
            data: [],
            colour: colour,
          },
          {
            label: "Simulated",
            data: simulated_data,
            colour: "#FFFFFF",
          },
        ],
      };
    }

    public add(...values: number[]): void {
      if (values.length !== 1) {
        throw new Error(
          `SimulatedChart: number of values (${values.length}) is not 1`
        );
      }
      if (
        this.data.datasets[0].data.length === this.data.datasets[1].data.length
      ) {
        console.warn(
          "SimulatedChart: data is being added but the simulation is complete"
        );
      }
      this.data.datasets[0].data.push(values[0]);
      this.set(this.data);
    }
  }
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from "svelte";
  import { Chart, registerables } from "chart.js";
  Chart.register(...registerables);

  export let id: string;
  export let line_chart: LineChart = undefined;
  export let title = "Title";
  export let xMax: number = 100;
  export let xMin: number = 0;
  export let yMax: number = 10;
  export let yMin: number = -10;

  export let data: ChartData;
  let chart: Chart;

  if (line_chart) {
    onDestroy(
      line_chart.subscribe((v) => {
        data = v;
      })
    );
  }

  function renderChart() {
    let ctx = document.getElementById(id);

    const chart_data = {
      labels: Array.from(Array(xMax - xMin)).map((_, i) => i + xMin),
      datasets: data.datasets.map(({ label, data, colour: borderColor }) => ({
        label,
        data,
        borderColor,
        backgroundColor: "transparent",
        radius: 0,
        tension: 0.1,
        borderWidth: 2,
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
    if (data.datasets.length > 0) {
      chart_data.datasets.forEach((dataset: { data: any[] }, index: number) => {
        dataset.data = data.datasets[index].data;
      });

      chart.update();
    }
  });
</script>

<canvas {id} class="object-cover max-h-full max-w-full" />
