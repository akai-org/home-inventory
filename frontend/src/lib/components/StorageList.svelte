<script lang="ts">
  import { onMount } from "svelte";
  import type { Storage } from "$lib/types/models";
  import { listStorages } from "$lib/api/services";

  let storages: Storage[] = [];

  onMount(async () => {
    storages = await listStorages();
  });
</script>

<div class="storage-grid">
  {#each storages as storage}
    <a href="/storages/{storage.id}" class="storage-card">
      <h5>{storage.name}</h5>
    </a>
  {/each}
</div>

<style>
  .storage-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 1.5rem;
  }

  .storage-card {
    background-color: #ffffff;
    padding: 1.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
    text-decoration: none;
    color: inherit;
  }

  .storage-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.12);
  }

  h5 {
    margin: 0;
    font-size: 1.2rem;
    font-weight: 600;
  }
</style>
