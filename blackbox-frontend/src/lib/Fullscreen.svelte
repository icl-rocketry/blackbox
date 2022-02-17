<script lang="ts">
  export let rows = 1;
  let expanded = true; //This is dumb, but the code doesn't work with it as false
  function expand() {
    const grid = this.parentNode;
    const items = grid.children;
    for (let i = 0; i < items.length; i++) {
      if (items[i] != this) {
        items[i].hidden = !items[i].hidden;
      }
    }
    expanded = grid.className.search("grid-cols-4") == -1;
    if (expanded) {
      grid.className = grid.className.replace(
        "lg:grid-cols-1",
        "lg:grid-cols-4"
      );
    } else {
      grid.className = grid.className.replace("grid-cols-4", "grid-cols-1");
    }
  }
</script>

<div on:click={expand} class="cursor-pointer row-span-{rows}">
  <slot fullscreen={!expanded}/>
</div>
