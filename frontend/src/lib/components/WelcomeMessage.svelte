
<script lang="ts">
  import { onMount } from 'svelte';

  let showWelcome = false;

  onMount(() => {
    const visited = localStorage.getItem('visited');
    if (!visited) {
      showWelcome = true;
      localStorage.setItem('visited', 'true');
    }
  });

  function closeWelcome() {
    showWelcome = false;
  }
</script>

{#if showWelcome}
  <div class="welcome-overlay">
    <div class="welcome-message">
      <h2>Welcome to the Home Inventory App!</h2>
      <p>This is a web-based application for cataloging and managing items in a house/apartment or any other place.</p>
      <p>The system allows you to create a structured database of items with the ability to easily search and identify them.</p>
      <p>Enjoy!</p>
      <button on:click={closeWelcome}>Got it!</button>
    </div>
  </div>
{/if}

<style>
  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @keyframes slideIn {
    from {
      transform: translateY(-20px);
    }
    to {
      transform: translateY(0);
    }
  }

  .welcome-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.6);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
    animation: fadeIn 0.3s ease-in-out;
  }

  .welcome-message {
    background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
    padding: 3rem;
    border-radius: 16px;
    text-align: center;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
    max-width: 500px;
    animation: slideIn 0.4s ease-in-out;
  }

  h2 {
    font-size: 2rem;
    font-weight: 600;
    margin-bottom: 1rem;
    color: #333;
  }

  p {
    font-size: 1.1rem;
    line-height: 1.6;
    color: #555;
    margin-bottom: 1.5rem;
  }

  button {
    background-color: #4a90e2;
    color: white;
    border: none;
    padding: 0.8rem 1.6rem;
    font-size: 1rem;
    font-weight: 500;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.2s ease-in-out;
  }

  button:hover {
    background-color: #357abd;
  }
</style>
