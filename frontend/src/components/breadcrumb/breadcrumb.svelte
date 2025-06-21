<script lang="ts">
	import type { Crumb } from "./types";

	interface Props {
		elements: Crumb[];
	}

	let { elements }: Props = $props();
</script>

<div class="breadcrumbs">
	{#each elements as elem, i (i)}
		<a href={elem.link}>{elem.display}</a>
		{#if i !== elements.length - 1}
			>
		{/if}
	{/each}
</div>

<style>
	.breadcrumbs > a:last-child {
		color: red;
	}
</style>
