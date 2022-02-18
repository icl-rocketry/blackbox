<script lang="ts">
  import { onMount } from "svelte";

  import * as THREE from "three";
  import { Scene } from "three";
  import { STLLoader } from "three/examples/jsm/loaders/STLLoader";

  export let x: number;
  export let y: number;
  export let z: number;
  export let fullscreen = false;

  const xAxisRotationOffset = Math.PI / 8;
  const yAxisRotationOffset = -Math.PI / 4;

  let updatePositionFunc = (_: number, __: number, ___: number) => {};
  let renderRocket = (_: THREE.Scene, __: THREE.Camera) => {};

  $: updatePositionFunc(x, y, z);

  let loader = new STLLoader();

  loader.load("/src/assets/apex.STL", function (geometry) {
    let material = new THREE.MeshNormalMaterial({});
    let rocketMesh = new THREE.Mesh(geometry, material);
    geometry.center();
    geometry.translate(1, 0, 0); //Account for bounding box center error

    renderRocket = (scene, camera) => {
      scene.add(rocketMesh);
      camera.lookAt(rocketMesh.position);
    };

    updatePositionFunc = function set_euler(x: number, y: number, z: number) {
      rocketMesh.rotation.x = x + xAxisRotationOffset;
      rocketMesh.rotation.y = y + yAxisRotationOffset;
      rocketMesh.rotation.z = z;
    };
  });

  async function renderOnCanvas() {
    async function waitForCanvas(resolve, reject) {
      let canvas = document.getElementById(
        fullscreen ? "canvas-full" : "canvas-mini"
      );
      if (canvas === null || canvas === undefined) {
        setTimeout(() => waitForCanvas(resolve, reject), 250);
      } else {
        resolve(canvas);
      }
    }

    let canvas: HTMLCanvasElement = await new Promise(waitForCanvas);
    // Setup
    let height = canvas.clientHeight;
    let width = canvas.clientWidth;
    let camera = new THREE.PerspectiveCamera(70, width / height, 0.1, 10000);
    let scene = new THREE.Scene();
    scene.background = new THREE.Color(0x1e293b); //bg-slate-800
    let renderer = new THREE.WebGLRenderer({ canvas: canvas, antialias: true });

    renderer.setSize(width, height);
    renderRocket(scene, camera);

    const axesHelper = new THREE.AxesHelper(40);
    axesHelper.rotateX(xAxisRotationOffset);
    axesHelper.rotateY(yAxisRotationOffset);
    scene.add(axesHelper);

    // Camera positioning - empirical position, seems to work
    camera.position.x = 0;
    camera.position.y = 0;
    camera.position.z = 90;

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
  }

  $: fullscreen, renderOnCanvas();
</script>

{#if fullscreen}
  <canvas
    id="canvas-full"
    class="object-cover min-h-full min-w-full rounded-xl row-span-2"
  />
{:else}
  <canvas
    id="canvas-mini"
    class="object-cover min-h-full max-w-full rounded-xl row-span-2"
  />
{/if}
