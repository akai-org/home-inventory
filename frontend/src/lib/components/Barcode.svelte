<script lang="ts">
  import { onMount } from "svelte";
  import { generateBarcode } from "$lib/api/services";

  export let data: string;

  let barcodeUrl: string;

  onMount(async () => {
    const blob = await generateBarcode(data);
    barcodeUrl = URL.createObjectURL(blob);
  });
</script>

{#if barcodeUrl}
  <img src={barcodeUrl} alt="Barcode" />
{/if}