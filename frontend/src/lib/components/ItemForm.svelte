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

<form class="item-form" on:submit|preventDefault={handleSubmit}>
  <div class="form-group">
    <label for="itemName">Name</label>
    <input id="itemName" type="text" placeholder="Enter item name" bind:value={newItem.name} required />
  </div>
  <div class="form-group">
    <label for="itemType">Type</label>
    <input id="itemType" type="text" placeholder="Enter item type" bind:value={newItem.type} required />
  </div>
  <div class="form-group">
    <label for="itemDescription">Description</label>
    <textarea id="itemDescription" placeholder="Enter item description" bind:value={newItem.description}></textarea>
  </div>
  <div class="form-group">
    <label for="itemStorage">Storage</label>
    <select id="itemStorage" bind:value={newItem.storage_id} required>
      <option value="" disabled selected>Select a storage</option>
      {#each storages as storage}
        <option value={storage.id}>{storage.name}</option>
      {/each}
    </select>
  </div>
  <div class="form-group">
    <label for="itemTags">Tags (comma separated)</label>
    <input id="itemTags" type="text" placeholder="e.g., electronics, kitchen" bind:value={tags} />
  </div>
  <div class="form-group">
    <label for="itemImage">Image</label>
    <input id="itemImage" type="file" on:change={handleImageUpload} accept="image/*" />
  </div>
  {#if imagePreview}
    <div class="image-preview">
      <img src={imagePreview} alt="Image preview" />
    </div>
  {/if}
  <button type="submit" class="submit-button">Add Item</button>
</form>

<style>
  .item-form {
    background-color: #ffffff;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  }

  .form-group {
    margin-bottom: 1.5rem;
  }

  label {
    display: block;
    font-weight: 500;
    margin-bottom: 0.5rem;
    color: #333;
  }

  input[type="text"],
  textarea,
  select {
    width: 100%;
    padding: 0.8rem 1rem;
    border: 1px solid #ccc;
    border-radius: 8px;
    font-size: 1rem;
    transition: border-color 0.2s ease;
  }

  input[type="text"]:focus,
  textarea:focus,
  select:focus {
    outline: none;
    border-color: #4a90e2;
  }

  .image-preview {
    margin-top: 1rem;
    text-align: center;
  }

  .image-preview img {
    max-width: 100%;
    max-height: 200px;
    border-radius: 8px;
  }

  .submit-button {
    width: 100%;
    padding: 1rem;
    background-color: #4a90e2;
    color: white;
    border: none;
    border-radius: 8px;
    font-size: 1.1rem;
    font-weight: 500;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .submit-button:hover {
    background-color: #357abd;
  }
</style>
