<script lang="ts">
  import { onMount } from "svelte";

  import * as THREE from "three";
  import { STLLoader } from "three/examples/jsm/loaders/STLLoader";

  export let x: number;
  export let y: number;
  export let z: number;
  let updatePositionFunc = (_, __, ___) => {};

  $: updatePositionFunc(x, y, z);

  onMount(() => {
    let canvas = document.getElementById("canvas");

    // Setup
    let scene = new THREE.Scene();
    scene.background = new THREE.Color(0x1e293b);
    let cHeight = canvas.clientHeight;
    let cWidth = canvas.clientWidth;
    console.log(cHeight, cWidth);
    let camera = new THREE.PerspectiveCamera(70, cWidth / cHeight, 0.1, 10000);
    let renderer = new THREE.WebGLRenderer({ canvas: canvas, antialias: true });

    renderer.setSize(cWidth, cHeight);

    // Resize after viewport-size-change
    window.addEventListener("resize", function () {
      let height = cHeight;
      let width = cWidth;
      renderer.setSize(width, height);
      camera.aspect = width / height;
      camera.updateProjectionMatrix();
    });

    let loader = new STLLoader();
    loader.load(
      "/src/assets/apex.STL",
      function (geometry) {
        console.log(geometry);
        let material = new THREE.MeshNormalMaterial({});
        let mesh = new THREE.Mesh(geometry, material);
        scene.add(mesh);
        geometry.center();
        geometry.translate(1, 0, 0); //Account for bounding box center error
        camera.lookAt(mesh.position);

        updatePositionFunc = function set_euler(x: number, y: number, z: number) {
          mesh.rotation.x = x;
          mesh.rotation.y = y;
          mesh.rotation.z = z;
        };
      }
    );

    const axesHelper = new THREE.AxesHelper(50);
    scene.add(axesHelper);

    // Camera positioning 
    camera.position.x = 72 * 0.7;
    camera.position.y = 25 * 0.7;
    camera.position.z = 100 * 0.7;


    // Draw scene
    let render = function () {
      renderer.render(scene, camera);
    };

    // Run game loop (render,repeat)
    let GameLoop = function () {
      requestAnimationFrame(GameLoop);
      render();
    };

    GameLoop();
  });
</script>

<canvas id="canvas" class="object-cover min-h-full min-w-full rounded-xl"/>
