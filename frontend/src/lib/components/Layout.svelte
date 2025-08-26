<script>
	import 'bootstrap/dist/css/bootstrap.min.css';
	import { HouseDoor, Box, Archive } from 'svelte-bootstrap-icons';
</script>

<div class="d-flex">
	<div class="sidebar vh-100 p-3 bg-light">
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

	<main class="p-4 w-100">
		<slot />
	</main>
</div>

<style>
	.sidebar {
		width: 280px;
	}
	.nav-link {
		display: flex;
		align-items: center;
		gap: 10px;
	}
</style>
