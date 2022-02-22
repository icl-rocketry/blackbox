<script lang="ts">
  import Modal from "svelte-simple-modal";
  import { writable } from "svelte/store";
  import Chart, {
    SimulatedChart,
    RollingWindowChart,
  } from "./lib/Chart.svelte";
  import Fullscreen from "./lib/Fullscreen.svelte";
  import Map from "./lib/Map.svelte";
  import Text from "./lib/Text.svelte";
  import RocketViewer from "./lib/RocketViewer.svelte";

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

  let rocket_lat = 51.5007;
  let rocket_long = -0.1246;

  let x = 0;
  let y = 0;
  let z = 0;

  const socket = new WebSocket("wss://live.imperialrocketry.com/ws");

  let distance = writable(0);
  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    rocket_lat = msg["Location"]["Latitude"];
    rocket_long = msg["Location"]["Longitude"];
    data.add(msg["Pressure"]);
  };

  setInterval(() => {
    distance.update((v) => v + Math.floor(Math.random() * 100));
    // data.add(Math.random() * 10);
    simulation.add(Math.random() * 10);
    x += (Math.random() - 0.5) * 0.05;
    y += (Math.random() - 0.5) * 1;
    z += (Math.random() - 0.5) * 0.05;
  }, 100);
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
      class="grid sm:grid-cols-1 lg:grid-cols-4 gap-2 min-h-screen min-w-full overflow-y-visible"
    >
      <Fullscreen>
        <div
          class="bg-slate-800 text-white flex justify-center items-center rounded-xl h-full h-full"
        >
          <Chart id="top_left" title="Acceleration" bind:data={$data} />
        </div>
      </Fullscreen>
      <Fullscreen let:fullscreen>
        <div
          class="bg-slate-800 text-white flex justify-center items-center rounded-xl h-full h-full"
        >
          <Text generator={distance} units="meters" {fullscreen} />
        </div>
      </Fullscreen>
      <Fullscreen>
        <div
          class="bg-slate-800 text-white flex justify-center items-center rounded-xl h-full h-full"
        >
          <Chart id="2" title="Altitude" bind:data={$simulation} />
        </div>
      </Fullscreen>
      <Fullscreen>
        <div
          class="bg-slate-800 text-white flex justify-center items-center rounded-xl h-full"
        >
          <Chart id="3" title="Dummy" bind:data={$d2} />
        </div>
      </Fullscreen>
      <Fullscreen>
        <div
          class="bg-slate-800 text-white flex justify-center items-center rounded-xl h-full"
        >
          <Chart id="4" title="Dummy" bind:data={$d2} />
        </div>
      </Fullscreen>
      <div
        class="rounded-xl lg:row-span-2 sm:col-span-1 lg:col-span-2 min-h-full h-96"
      >
        <Map {rocket_lat} {rocket_long} />
      </div>

      <Fullscreen rows={2} let:fullscreen>
        <RocketViewer {x} {y} {z} {fullscreen}/>
      </Fullscreen>
      <Fullscreen>
        <div
          class="bg-slate-800 text-white flex justify-center items-center rounded-xl h-full"
        >
          <Chart id="6" title="Dummy" bind:data={$d2} />
        </div>
      </Fullscreen>
    </div>
  </div>
</Modal>
