<script lang="ts">
  import type { Item } from "$lib/types/models";
  import QRCode from "./QRCode.svelte";
  import Barcode from "./Barcode.svelte";

  export let item: Item;
  let showQRCode = false;
  let showBarcode = false;
</script>

<div class="card h-100">
  <a href="/storages/{item.storage_id}/items/{item.id}">
    <img src={item.image_url || 'https://via.placeholder.com/150'} class="card-img-top" alt={item.name}>
  </a>
  <div class="card-body">
    <h5 class="card-title">{item.name}</h5>
    <p class="card-text">{item.description}</p>
    <p class="card-text"><small class="text-muted">Type: {item.type}</small></p>
    {#if item.tags && item.tags.length > 0}
      <div>
        <strong>Tags:</strong>
        {#each item.tags as tag}
          <span class="badge bg-secondary me-1">{tag}</span>
        {/each}
      </div>
    {/if}
  </div>
  <div class="card-footer d-flex justify-content-between">
    <div>
      <button class="btn btn-outline-primary btn-sm me-2" on:click={() => showQRCode = !showQRCode} title="{showQRCode ? 'Hide' : 'Show'} QR Code">
        <i class="bi bi-qr-code"></i>
      </button>
      <button class="btn btn-outline-secondary btn-sm" on:click={() => showBarcode = !showBarcode} title="{showBarcode ? 'Hide' : 'Show'} Barcode">
        <i class="bi bi-upc-scan"></i>
      </button>
    </div>
    <a href="/storages/{item.storage_id}/items/{item.id}" class="btn btn-primary btn-sm">View Details</a>
  </div>
  {#if showQRCode}
    <div class="p-3">
      <QRCode entityId={item.id} entityType="item" />
    </div>
  {/if}
  {#if showBarcode}
    <div class="p-3">
      <Barcode data={item.id} />
    </div>
  {/if}
</div>
