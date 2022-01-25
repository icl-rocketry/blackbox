<script lang="ts">
  import * as Leaflet from "leaflet";
  import "leaflet/dist/leaflet.css";
  let map: { remove: () => void };
  export let rocket_long: number;
  export let rocket_lat: number;
  let us_long = 0;
  let us_lat = 0;

  function emojiIcon(emoji: string): any {
    return Leaflet.divIcon({
      html: `<svg xmlns="http://www.w3.org/2000/svg">
            <text y="30" font-size="3em">${emoji}</text>
          </svg>
          `,
      className: "dummy",
    });
  }
  const rocket = emojiIcon("🚀");
  const us_icon = emojiIcon("🔭");

  let rocket_marker = Leaflet.marker([rocket_lat, rocket_long], {
    icon: rocket,
  });
  let us_marker = Leaflet.marker([us_lat, us_long], { icon: us_icon });
  
  $: us_marker.setLatLng([us_lat, us_long]);
  $: rocket_marker.setLatLng([rocket_lat, rocket_long]);

  function createMap(container: HTMLDivElement) {
    let m = Leaflet.map(container, {
      zoomControl: false,
    }).setView([51.505, -0.09], 13);
    Leaflet.tileLayer(
      "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
      {
        attribution: `&copy;<a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a>,
            &copy;<a href="https://carto.com/attributions" target="_blank">CARTO</a>`,
        subdomains: "abcd",
        maxZoom: 20,
      }
    ).addTo(m);
    navigator.geolocation.watchPosition(function (position) {
        us_lat = position.coords.latitude;
        us_long = position.coords.longitude;
    });

    return m;
  }

  function mapAction(container: HTMLDivElement) {
    map = createMap(container);
    rocket_marker.addTo(map);
    us_marker.addTo(map);
    return {
      destroy: () => {
        map.remove();
      },
    };
  }
</script>

<div class="min-h-full min-w-full h-full rounded-xl" use:mapAction />
