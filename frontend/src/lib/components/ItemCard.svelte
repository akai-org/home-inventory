<script lang="ts">
  import type { Item } from "$lib/types/models";
  import QRCode from "./QRCode.svelte";
  import Barcode from "./Barcode.svelte";

  export let item: Item;
  let showQRCode = false;
  let showBarcode = false;
</script>

<div class="item-card">
  <a href="/storages/{item.storage_id}/items/{item.id}" class="image-link">
    <img src={item.image_url || 'https://via.placeholder.com/300'} alt={item.name}>
  </a>
  <div class="card-content">
    <h5 class="card-title">{item.name}</h5>
    <p class="card-description">{item.description}</p>
    <p class="card-type"><small>Type: {item.type}</small></p>
    {#if item.tags && item.tags.length > 0}
      <div class="tags-container">
        {#each item.tags as tag}
          <span class="tag">{tag}</span>
        {/each}
      </div>
    {/if}
  </div>
  <div class="card-footer">
    <div class="button-group">
      <button class="icon-button" on:click={() => showQRCode = !showQRCode} title="{showQRCode ? 'Hide' : 'Show'} QR Code">
        <i class="bi bi-qr-code"></i>
      </button>
      <button class="icon-button" on:click={() => showBarcode = !showBarcode} title="{showBarcode ? 'Hide' : 'Show'} Barcode">
        <i class="bi bi-upc-scan"></i>
      </button>
    </div>
    <a href="/storages/{item.storage_id}/items/{item.id}" class="details-button">View Details</a>
  </div>
  {#if showQRCode}
    <QRCode entityId={item.id} show={showQRCode} entityType="item" on:close={() => (showQRCode = false)} />
  {/if}
  {#if showBarcode}
    <Barcode data={item.id} show={showBarcode} on:close={() => (showBarcode = false)} />
  {/if}
</div>

<style>
  .item-card {
    background-color: #ffffff;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    overflow: hidden;
    display: flex;
    flex-direction: column;
    transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
  }

  .item-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.12);
  }

  .image-link img {
    width: 100%;
    height: 200px;
    object-fit: contain;
  }

  .card-content {
    padding: 1.5rem;
    flex-grow: 1;
  }

  .card-title {
    font-size: 1.25rem;
    font-weight: 600;
    margin-bottom: 0.5rem;
  }

  .card-description {
    color: #555;
    margin-bottom: 1rem;
  }

  .card-type {
    color: #777;
    font-size: 0.9rem;
  }

  .tags-container {
    margin-top: 1rem;
  }

  .tag {
    background-color: #e9ecef;
    color: #333;
    padding: 0.3rem 0.6rem;
    border-radius: 6px;
    font-size: 0.8rem;
    margin-right: 0.5rem;
  }

  .card-footer {
    padding: 1rem 1.5rem;
    border-top: 1px solid #f0f2f5;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .button-group .icon-button {
    background: none;
    border: none;
    font-size: 1.2rem;
    color: #777;
    cursor: pointer;
    transition: color 0.2s ease;
  }

  .button-group .icon-button:hover {
    color: #333;
  }

  .details-button {
    background-color: #4a90e2;
    color: white;
    text-decoration: none;
    padding: 0.6rem 1.2rem;
    border-radius: 8px;
    font-weight: 500;
    transition: background-color 0.2s ease;
  }

  .details-button:hover {
    background-color: #357abd;
  }
</style>
