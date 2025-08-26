<script lang="ts">
  import { onMount } from "svelte";
  import type { Storage } from "$lib/types/models";
  import { listStorages, createStorage } from "$lib/api/services";

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

<h1 class="mb-4">Storages</h1>

<div class="row">
  <div class="col-md-4">
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">Add New Storage</h5>
        <form on:submit|preventDefault={handleAddStorage}>
          <div class="mb-3">
            <label for="storageName" class="form-label">Name</label>
            <input id="storageName" type="text" class="form-control" placeholder="Name" bind:value={newStorage.name} />
          </div>
          <button type="submit" class="btn btn-primary">Add Storage</button>
        </form>
      </div>
    </div>
  </div>
  <div class="col-md-8">
    <div class="list-group">
      {#each storages as storage}
        <a href="/storages/{storage.id}" class="list-group-item list-group-item-action">
          {storage.name}
        </a>
      {/each}
    </div>
  </div>
</div>
