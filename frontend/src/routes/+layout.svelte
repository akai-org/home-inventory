<script>
  import Layout from '$lib/components/Layout.svelte';
  import WelcomeMessage from '$lib/components/WelcomeMessage.svelte';
</script>

<Layout>
  <slot />
</Layout>

<WelcomeMessage />
