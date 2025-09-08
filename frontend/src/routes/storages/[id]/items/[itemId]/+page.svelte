<script lang="ts">
  import { onMount } from "svelte";
  import type { Item, Storage } from "$lib/types/models";
  import { getItem, getStorage } from "$lib/api/services";
  import { page } from "$app/stores";

  let item: Item | null = null;
  let storage: Storage | null = null;

  onMount(async () => {
    const { id, itemId } = $page.params;
    item = await getItem(itemId);
    storage = await getStorage(id);
  });
</script>

{#if item && storage}
  <div class="container">
    <div class="row">
      <div class="col-12">
        <nav aria-label="breadcrumb">
          <ol class="breadcrumb">
            <li class="breadcrumb-item"><a href="/">Home</a></li>
            <li class="breadcrumb-item"><a href="/storages">Storages</a></li>
            <li class="breadcrumb-item"><a href={`/storages/${storage.id}`}>{storage.name}</a></li>
            <li class="breadcrumb-item active" aria-current="page">{item.name}</li>
          </ol>
        </nav>
      </div>
    </div>
    <div class="row">
      <div class="col-md-6">
        <img src={item.image_url || 'https://via.placeholder.com/500'} class="img-fluid" alt={item.name}>
      </div>
      <div class="col-md-6">
        <h1>{item.name}</h1>
        <p class="lead">{item.description}</p>
        <p><strong>Type:</strong> {item.type}</p>
        <p><strong>Location:</strong> {storage.name}</p>
        {#if item.tags && item.tags.length > 0}
          <div>
            <strong>Tags:</strong>
            {#each item.tags as tag}
              <span class="badge bg-secondary me-1">{tag}</span>
            {/each}
          </div>
        {/if}
      </div>
    </div>
  </div>
{:else}
  <p>Loading...</p>
{/if}
