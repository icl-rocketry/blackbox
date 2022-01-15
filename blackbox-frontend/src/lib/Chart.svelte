<script lang="ts">
  import { onMount } from "svelte";
  import Chart from "chart.js";

  export let id: string;

  onMount(renderChart);
  function renderChart() {
    var ctx = document.getElementById(id).getContext("2d");
    var chart = new Chart(ctx, {
      type: "line",
      data: {
        labels: ["January", "February", "March", "April", "May", "June", "July", "August"],
        datasets: [
          {
            label: "My First dataset",
            borderColor: "rgb(255, 99, 132)",
            data: [0, 10, 5, 2, 20, 30, 45, 80]
          }
        ]
      },
      options: {}
    });
  }
</script>
<canvas id={id} class="object-cover max-h-full max-w-full"></canvas>