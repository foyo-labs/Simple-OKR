<script>
	import '../app.postcss';
	import Header from './Header.svelte';
</script>

<div class="min-h-full">
	<Header />
	<main>
		<div class="max-w-7xl py-6 sm:px-6 lg:px-8">
			<slot />
		</div>
	</main>
</div>
