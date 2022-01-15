<script lang="ts">
  import { onMount } from "svelte";
  import { Chart } from "chart.js";

  export let id: string;

  onMount(renderChart);

  function renderChart() {
    let ctx = document.getElementById(id);

    const labels = [1, 2, 3, 4, 5, 6, 7];
    const data = {
      labels: labels,
      datasets: [
        {
          label: "Dataset 1",
          data: [1, 2, 3, 4, 5, 6, 7],
          borderColor: "#FF0000",
          backgroundColor: "transparent",
        },
        {
          label: "Dataset 2",
          data: [7, 6, 5, 4, 3, 2, 1],
          borderColor: "#0000FF",
          backgroundColor: "transparent",
        },
      ],
    };

    const config = {
      type: "line",
      data: data,
      options: {
        responsive: true,
        plugins: {
          legend: {
            position: "bottom",
          },
        },
      },
    };

    new Chart(ctx as any, config as any);
  }
</script>

<canvas {id} class="object-cover max-h-full max-w-full" />
