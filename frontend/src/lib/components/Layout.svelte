<script>
	import 'bootstrap/dist/css/bootstrap.min.css';
	import { HouseDoor, Box, Archive } from 'svelte-bootstrap-icons';
</script>

<div class="layout">
	<div class="sidebar">
		<h4 class="mb-4">Home Inventory</h4>
		<ul class="nav flex-column">
			<li class="nav-item">
				<a class="nav-link" href="/">
					<HouseDoor />
					Home
				</a>
			</li>
			<li class="nav-item">
				<a class="nav-link" href="/storages">
					<Archive />
					Storages
				</a>
			</li>
		</ul>
	</div>

	<main class="main-content">
		<slot />
	</main>
</div>

<style>
	.layout {
		display: flex;
		min-height: 100vh;
		background-color: #f0f2f5;
	}

	.sidebar {
		width: 250px;
		background-color: #ffffff;
		padding: 1.5rem;
		box-shadow: 2px 0 5px rgba(0, 0, 0, 0.05);
	}

	.sidebar h4 {
		font-weight: 600;
		color: #333;
	}

	.nav-link {
		display: flex;
		align-items: center;
		gap: 12px;
		color: #555;
		padding: 0.75rem 1rem;
		border-radius: 8px;
		transition: background-color 0.2s ease, color 0.2s ease;
	}

	.nav-link:hover {
		background-color: #e9ecef;
		color: #000;
	}

	.main-content {
		flex-grow: 1;
		padding: 2rem;
	}
</style>
