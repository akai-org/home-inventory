<script lang="ts">
  import type { Item } from "$lib/types/models";
  import QRCode from "./QRCode.svelte";
  import Barcode from "./Barcode.svelte";

  export let item: Item;
  let showQRCode = false;
  let showBarcode = false;
</script>

<div class="card">
  <h3>{item.name}</h3>
  <p>{item.description}</p>
  <p>Type: {item.type}</p>
  {#if item.tags && item.tags.length > 0}
    <div>
      <strong>Tags:</strong>
      {#each item.tags as tag}
        <span>{tag}</span>
      {/each}
    </div>
  {/if}
  <button on:click={() => showQRCode = !showQRCode}>
    {showQRCode ? "Hide" : "Show"} QR Code
  </button>
  <button on:click={() => showBarcode = !showBarcode}>
    {showBarcode ? "Hide" : "Show"} Barcode
  </button>
  {#if showQRCode}
    <QRCode entityId={item.id} entityType="item" />
  {/if}
  {#if showBarcode}
    <Barcode data={item.id} />
  {/if}
</div>

<style>
  .card {
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 10px;
    margin-bottom: 10px;
  }

  span {
    display: inline-block;
    background-color: #eee;
    padding: 2px 5px;
    border-radius: 3px;
    margin-right: 5px;
  }
</style>