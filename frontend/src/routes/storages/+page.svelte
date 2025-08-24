<script lang="ts">
  import { onMount } from "svelte";
  import type { Storage } from "$lib/types/models";
  import { listStorages, createStorage } from "$lib/api/services";
  import Layout from "$lib/components/Layout.svelte";

  let storages: Storage[] = [];
  let newStorage: Partial<Storage> = {};

  onMount(async () => {
    storages = await listStorages();
  });

  async function handleAddStorage() {
    if (newStorage.name) {
      const createdStorage = await createStorage(newStorage as Storage);
      storages = [...storages, createdStorage];
      newStorage = {};
    }
  }
</script>

<Layout>
  <h1>Storages</h1>

  <form on:submit|preventDefault={handleAddStorage}>
    <input type="text" placeholder="Name" bind:value={newStorage.name} />
    <button type="submit">Add Storage</button>
  </form>

  <ul>
    {#each storages as storage}
      <li><a href="/storages/{storage.id}">{storage.name}</a></li>
    {/each}
  </ul>
</Layout>