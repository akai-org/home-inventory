<script lang="ts">
  import { createEventDispatcher, onMount } from "svelte";
  import { debounce } from "$lib/debounce";

  const dispatch = createEventDispatcher();

  let name = "";
  let type = "";
  let tags = "";

  // Hardcoded options for dropdowns
  const typeOptions = ["Electronics", "Furniture", "Clothing", "Books"];
  const tagOptions = ["Work", "Personal", "Home", "Office"];

  const debouncedFilter = debounce(() => {
    dispatch("filter", { name, type, tags });
  }, 300);

  onMount(() => {
    const form = document.querySelector("form");
    form?.addEventListener("input", debouncedFilter);
    form?.addEventListener("change", debouncedFilter);
  });
</script>

<form>
  <div class="mb-3">
    <label for="filterName" class="form-label">Name</label>
    <input id="filterName" type="text" class="form-control" placeholder="Name" bind:value={name} />
  </div>
  <div class="mb-3">
    <label for="filterType" class="form-label">Type</label>
    <select id="filterType" class="form-select" bind:value={type}>
      <option value="">All</option>
      {#each typeOptions as option}
        <option value={option}>{option}</option>
      {/each}
    </select>
  </div>
  <div class="mb-3">
    <label for="filterTags" class="form-label">Tags</label>
    <select id="filterTags" class="form-select" bind:value={tags}>
      <option value="">All</option>
      {#each tagOptions as option}
        <option value={option}>{option}</option>
      {/each}
    </select>
  </div>
</form>
