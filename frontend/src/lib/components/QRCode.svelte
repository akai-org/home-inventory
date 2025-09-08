<script lang="ts">
  import { onMount } from "svelte";
  import { generateQRCode } from "$lib/api/services";

  export let entityId: string;
  export let entityType: string;
  export let show: boolean = false;

  let qrCodeUrl: string;
  let dialog: HTMLDialogElement;

  function close() {
    show = false;
  }

  const handleClick = (event: MouseEvent) => {
    if (event.target === dialog) {
      close();
    }
  };

  onMount(async () => {
    const blob = await generateQRCode(entityId, entityType);
    qrCodeUrl = URL.createObjectURL(blob);
  });

  $: if (dialog && show) {
    dialog.showModal();
  } else if (dialog && !show) {
    dialog.close();
  }
</script>

{#if qrCodeUrl && show}
<dialog bind:this={dialog} class="barcode-dialog" on:click={handleClick}>
  <div>
    <button class="close-button" on:click={close}>&times;</button>
    <img src={qrCodeUrl} alt="QR Code" />
  </div>
</dialog>
{/if}

<style>
  .barcode-dialog {
    border: none;
    padding: 0;
    text-align: center;
    position: fixed;
    inset: 0;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    max-width: 100vw;
    max-height: 100vh;
    margin: 0;
    backdrop-filter: blur(10px);
    background: rgba(0, 0, 0, 0.3);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
  }

  .barcode-dialog::backdrop {
    background: rgba(0, 0, 0, 0.3);
    backdrop-filter: blur(10px);
  }

  .barcode-dialog img {
    max-width: 90%;
    max-height: 90%;
    height: auto;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
    border-radius: 8px;
    background: white;
    padding: 1rem;
  }

  .close-button {
    position: absolute;
    top: 1rem;
    right: 1rem;
    background: none;
    border: none;
    font-size: 2rem;
    color: white;
    cursor: pointer;
  }
</style>