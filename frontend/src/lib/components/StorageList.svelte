<script lang="ts">
  import { onMount } from "svelte";
  import type { Storage } from "$lib/types/models";
  import { listStorages } from "$lib/api/services";

  let storages: Storage[] = [];

  onMount(async () => {
    storages = await listStorages();
  });
</script>

<div class="list-group">
  {#each storages as storage}
    <a href="/storages/{storage.id}" class="list-group-item list-group-item-action">
      {storage.name}
    </a>
  {/each}
</div>
