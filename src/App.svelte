<script>
  import Router from 'svelte-spa-router';
  import { routes } from './routes';
</script>

<main>
  <Router {routes} />
</main>

<style>
  /* Global container styles */
  main {
    @apply min-h-screen;
  }
</style>