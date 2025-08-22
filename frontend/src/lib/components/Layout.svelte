<nav>
  <a href="/">Items</a>
  <a href="/storages">Storages</a>
</nav>

<div class="container">
  <slot />
</div>

<style>
  .container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }

  nav {
    display: flex;
    justify-content: center;
    padding: 10px;
    border-bottom: 1px solid #ccc;
  }

  nav a {
    margin: 0 10px;
    text-decoration: none;
    color: #333;
  }
</style>