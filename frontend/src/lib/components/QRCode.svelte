<script lang="ts">
  import { onMount } from "svelte";
  import { generateQRCode } from "$lib/api/services";

  export let entityId: string;
  export let entityType: string;

  let qrCodeUrl: string;

  onMount(async () => {
    const blob = await generateQRCode(entityId, entityType);
    qrCodeUrl = URL.createObjectURL(blob);
  });
</script>

{#if qrCodeUrl}
  <img src={qrCodeUrl} alt="QR Code" />
{/if}