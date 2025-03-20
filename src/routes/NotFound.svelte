<script>
  import { push } from 'svelte-spa-router';
</script>

<div class="min-h-screen flex flex-col justify-center items-center bg-gray-50 px-4">
  <div class="text-center">
    <h1 class="text-6xl font-bold text-primary mb-4">404</h1>
    <h2 class="text-2xl font-semibold mb-4">Page Not Found</h2>
    <p class="text-gray-600 mb-8">The page you're looking for doesn't exist or has been moved.</p>
    <button 
      class="btn btn-primary"
      on:click={() => push('/')}
    >
      Go Home
    </button>
  </div>
</div>