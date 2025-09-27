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
  <div class="item-details-page">
    <div class="breadcrumb-container">
      <a href="/" class="breadcrumb-link">Home</a>
      <span>/</span>
      <a href="/storages" class="breadcrumb-link">Storages</a>
      <span>/</span>
      <a href={`/storages/${storage.id}`} class="breadcrumb-link">{storage.name}</a>
      <span>/</span>
      <span class="breadcrumb-active">{item.name}</span>
    </div>

    <div class="details-layout">
      <div class="image-column">
        <img src={item.image_url || 'https://via.placeholder.com/500'} alt={item.name}>
      </div>
      <div class="info-column">
        <h1 class="item-name">{item.name}</h1>
        <p class="item-description">{item.description}</p>
        <div class="item-meta">
          <p><strong>Type:</strong> {item.type}</p>
          <p><strong>Location:</strong> {storage.name}</p>
        </div>
        {#if item.tags && item.tags.length > 0}
          <div class="tags-container">
            <strong>Tags:</strong>
            {#each item.tags as tag}
              <span class="tag">{tag}</span>
            {/each}
          </div>
        {/if}
        <a href={`/storages/${storage.id}`} class="back-button">Back to Storage</a>
      </div>
    </div>
  </div>
{:else}
  <p>Loading...</p>
{/if}

<style>
  .item-details-page {
    padding: 2rem;
  }

  .breadcrumb-container {
    margin-bottom: 2rem;
    font-size: 0.9rem;
    color: #555;
  }

  .breadcrumb-link {
    color: #4a90e2;
    text-decoration: none;
  }

  .breadcrumb-link:hover {
    text-decoration: underline;
  }

  .breadcrumb-active {
    font-weight: 500;
    color: #333;
  }

  .details-layout {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 3rem;
    align-items: start;
  }

  .image-column img {
    width: 100%;
    border-radius: 12px;
    box-shadow: 0 8px 25px rgba(0,0,0,0.15);
  }

  .info-column {
    background-color: #fff;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0,0,0,0.08);
  }

  .item-name {
    font-size: 2.5rem;
    font-weight: 700;
    margin-bottom: 1rem;
  }

  .item-description {
    font-size: 1.1rem;
    line-height: 1.6;
    color: #555;
    margin-bottom: 2rem;
  }

  .item-meta p {
    margin-bottom: 0.5rem;
  }

  .tags-container {
    margin-top: 1.5rem;
  }

  .tag {
    background-color: #e9ecef;
    color: #333;
    padding: 0.4rem 0.8rem;
    border-radius: 6px;
    font-size: 0.9rem;
    margin-right: 0.5rem;
  }

  .back-button {
    display: inline-block;
    margin-top: 2rem;
    background-color: #4a90e2;
    color: white;
    text-decoration: none;
    padding: 0.8rem 1.5rem;
    border-radius: 8px;
    font-weight: 500;
    transition: background-color 0.2s ease;
  }

  .back-button:hover {
    background-color: #357abd;
  }
</style>
