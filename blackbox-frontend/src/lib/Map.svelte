<script lang="ts">
  import * as L from "leaflet";
  // If you're playing with this in the Svelte REPL, import the CSS using the
  // syntax in svelte:head instead. For normal development, this is better.
  import "leaflet/dist/leaflet.css";
  let map: { remove: () => void; };
  export let long = 51.505;
  export let lat = -0.09;

  let marker = L.marker([long, lat], {});
  $: marker.setLatLng([long, lat]);
  function createMap(container: HTMLDivElement) {
    let m = L.map(container, {
      zoomControl: false
    }).setView([51.505, -0.09], 13);
    L.tileLayer(
      "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
      {
        attribution: `&copy;<a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a>,
            &copy;<a href="https://carto.com/attributions" target="_blank">CARTO</a>`,
        subdomains: "abcd",
        maxZoom: 20,
      }
    ).addTo(m);
    m.locate({setView: true, watch:true});
    return m;
  }

  function mapAction(container: HTMLDivElement) {
    map = createMap(container);
    marker.addTo(map);
    return {
      destroy: () => {
        map.remove();
      },
    };
  }
</script>

<div class="min-h-full min-w-full h-full rounded-xl" use:mapAction />
