<script lang="ts">
  import Modal from "svelte-simple-modal";
  import { writable } from "svelte/store";
  import Chart, { SimulatedChart } from "./lib/Chart.svelte";
  import { RollingWindowChart } from "./lib/Chart.svelte";
  import Fullscreen from "./lib/Fullscreen.svelte";
  import Text from "./lib/Text.svelte";

  const data = new RollingWindowChart(100, [
    {
      label: "X",
      colour: "#FF0000",
    },
  ]);

  const simulation = new SimulatedChart(
    "Altitude",
    "green",
    [1, 2, 3, 4, 5, 6, 7, 8, 9]
  );

  const d2 = new RollingWindowChart(7, [
    {
      label: "Dataset 1",
      colour: "#FF0000",
    },
    {
      label: "Dataset 2",
      colour: "#0000FF",
    },
  ]);

  let distance = writable(0);

  setInterval(() => {
    data.add(Math.random() * 10);
    distance.update((v) => v + Math.floor(Math.random() * 100));
    simulation.add(Math.random() * 10);
  }, 50);
</script>

<Modal
  styleWindow={{
    width: "90%",
    height: "90%",
    backgroundColor: "#1e293b",
    color: "white",
  }}
  closeButton={false}
>
  <div class="bg-slate-900 min-h-screen flex items-center justify-center">
    <div
      class="grid grid-cols-4 gap-2 min-h-screen min-w-full overflow-y-hidden overflow-x-hidden"
    >
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "top_left",
          title: "Acceleration",
          line_chart: data,
        }}
      />
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Text}
        props={{
          generator: distance,
          units: "meters",
        }}
      />
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "2",
          title: "Altitude",
          line_chart: simulation,
        }}
      />
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "3",
          title: "Dummy",
          line_chart: d2,
        }}
      />
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "4",
          title: "Dummy",
          line_chart: d2,
        }}
      />
      <div
        class="bg-slate-800 text-white rounded-xl row-span-2 col-span-2 min-h-full"
      >
        6
      </div>
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "5",
          title: "Dummy",
          line_chart: d2,
        }}
      />
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "6",
          title: "Dummy",
          line_chart: d2,
        }}
      />
      <Fullscreen
        className="bg-slate-800 text-white flex justify-center items-center rounded-xl"
        child={Chart}
        props={{
          id: "7",
          title: "Dummy",
          line_chart: d2,
        }}
      />
    </div>
  </div>
</Modal>
