<script lang="ts">
  import { createEventDispatcher, onMount } from "svelte";
  import type { Item, Storage } from "$lib/types/models";

  export let storages: Storage[] = [];

  const dispatch = createEventDispatcher();

  let newItem: Partial<Item> = {};
  let tags = "";
  let imagePreview: string | ArrayBuffer | null = null;

  function handleSubmit() {
    if (tags) {
      newItem.tags = tags.split(",").map(t => t.trim());
    }
    dispatch("addItem", newItem);
    newItem = {};
    tags = "";
    imagePreview = null;
  }

  function handleImageUpload(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const reader = new FileReader();
      reader.onload = (e) => {
        imagePreview = e.target?.result ?? null;
        newItem.image_url = imagePreview as string;
      };
      reader.readAsDataURL(input.files[0]);
    }
  }
</script>

<form on:submit|preventDefault={handleSubmit}>
  <div class="mb-3">
    <label for="itemName" class="form-label">Name</label>
    <input id="itemName" type="text" class="form-control" placeholder="Name" bind:value={newItem.name} required />
  </div>
  <div class="mb-3">
    <label for="itemType" class="form-label">Type</label>
    <input id="itemType" type="text" class="form-control" placeholder="Type" bind:value={newItem.type} required />
  </div>
  <div class="mb-3">
    <label for="itemDescription" class="form-label">Description</label>
    <textarea id="itemDescription" class="form-control" placeholder="Description" bind:value={newItem.description}></textarea>
  </div>
  <div class="mb-3">
    <label for="itemStorage" class="form-label">Storage</label>
    <select id="itemStorage" class="form-select" bind:value={newItem.storage_id} required>
      <option value="" disabled selected>Select a storage</option>
      {#each storages as storage}
        <option value={storage.id}>{storage.name}</option>
      {/each}
    </select>
  </div>
  <div class="mb-3">
    <label for="itemTags" class="form-label">Tags (comma separated)</label>
    <input id="itemTags" type="text" class="form-control" placeholder="Tags (comma separated)" bind:value={tags} />
  </div>
  <div class="mb-3">
    <label for="itemImage" class="form-label">Image</label>
    <input id="itemImage" type="file" class="form-control" on:change={handleImageUpload} accept="image/*" />
  </div>
  {#if imagePreview}
    <div class="mb-3">
      <img src={imagePreview} alt="Image preview" class="img-thumbnail" />
    </div>
  {/if}
  <button type="submit" class="btn btn-primary">Add Item</button>
</form>
