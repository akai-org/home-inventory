<script lang="ts">
	import { onMount } from 'svelte';
	import type { Item, Storage } from '$lib/types/models';
	import { listItems, createItem, listStorages } from '$lib/api/services';
  import ItemCard from '$lib/components/ItemCard.svelte';
  import ItemForm from '$lib/components/ItemForm.svelte';
  import ItemFilter from '$lib/components/ItemFilter.svelte';
  import StorageList from '$lib/components/StorageList.svelte';

	let items: Item[] = [];
  let storages: Storage[] = [];
  let showAddItemModal = false;

	onMount(async () => {
		items = await listItems();
    storages = await listStorages();
	});

	async function handleAddItem(event: CustomEvent<Item>) {
    const newItem = event.detail;
		if (newItem.name && newItem.type && newItem.description && newItem.storage_id) {
			const createdItem = await createItem(newItem as Item);
			items = [...items, createdItem];
      showAddItemModal = false;
		}
	}

  async function handleFilter(event: CustomEvent<{ name: string, type: string, tags: string }>) {
    const { name, type, tags } = event.detail;
    const filterParams = new URLSearchParams();
    if (name) {
      filterParams.append("name", name);
    }
    if (type) {
      filterParams.append("type", type);
    }
    if (tags) {
      filterParams.append("tags", tags);
    }
    items = await listItems(filterParams);
  }
</script>

<div class="container-fluid py-4">
  <div class="row">
    <div class="col-md-3">
      <div class="card mb-4">
        <div class="card-body">
          <h5 class="card-title mb-3"><i class="bi bi-funnel-fill me-2"></i>Filters</h5>
          <ItemFilter on:filter={handleFilter} />
        </div>
      </div>
      <div class="card">
        <div class="card-body">
          <h5 class="card-title mb-3"><i class="bi bi-hdd-stack-fill me-2"></i>Storages</h5>
          <StorageList />
        </div>
      </div>
    </div>
    <div class="col-md-9">
      <div class="d-flex justify-content-between align-items-center mb-3">
        <h1 class="h3 mb-0">Inventory</h1>
        <div class="d-flex">
          <div class="input-group me-2">
            <span class="input-group-text"><i class="bi bi-search"></i></span>
            <input type="text" class="form-control" placeholder="Search by name..." on:input={(e) => handleFilter({ detail: { name: e.currentTarget.value, type: '', tags: '' }, bubbles: false, cancelable: false, composed: false,timeStamp: 0, isTrusted: false, type: ''})}>
          </div>
          <button class="btn btn-primary" on:click={() => showAddItemModal = true}><i class="bi bi-plus-lg me-2"></i>Add Item</button>
        </div>
      </div>
      <div class="row">
        {#each items as item}
          <div class="col-md-6 col-lg-4 mb-4">
            <ItemCard {item} />
          </div>
        {/each}
      </div>
    </div>
  </div>
</div>

{#if showAddItemModal}
  <div class="modal fade show d-block" tabindex="-1">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Add New Item</h5>
          <button type="button" class="btn-close" on:click={() => showAddItemModal = false}></button>
        </div>
        <div class="modal-body">
          <ItemForm on:addItem={handleAddItem} {storages} />
        </div>
      </div>
    </div>
  </div>
  <div class="modal-backdrop fade show"></div>
{/if}
