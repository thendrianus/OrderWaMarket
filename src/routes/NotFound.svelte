<script>
  import { push } from 'svelte-spa-router';
</script>

<svelte:head>
  <title>Page Not Found | WhatsApp Catalogue</title>
</svelte:head>

<div class="min-h-screen flex flex-col justify-center items-center bg-gray-100 px-4 py-16">
  <div class="text-center">
    <div class="bg-[#25d366] p-4 rounded-full inline-flex mb-6">
      <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
        <line x1="12" y1="9" x2="12" y2="13"></line>
        <line x1="12" y1="17" x2="12.01" y2="17"></line>
      </svg>
    </div>
    
    <h1 class="text-4xl font-bold text-gray-900 mb-4">404 - Page Not Found</h1>
    <p class="text-lg text-gray-600 mb-8">
      The page you're looking for doesn't exist or has been moved.
    </p>
    
    <div class="flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-4">
      <button 
        on:click={() => push('/')}
        class="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-[#25d366] hover:bg-[#1da051]"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
          <path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
        </svg>
        Back to Home
      </button>
      
      <button 
        on:click={() => window.history.back()}
        class="inline-flex items-center justify-center px-5 py-3 border border-gray-300 text-base font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l4.293 4.293a1 1 0 010 1.414z" clip-rule="evenodd" />
        </svg>
        Go Back
      </button>
    </div>
  </div>
  
  <div class="mt-16 text-sm text-gray-500 text-center">
    <p>&copy; {new Date().getFullYear()} WhatsApp Catalogue. All rights reserved.</p>
    <p class="mt-2">WhatsApp is a trademark of Meta Platforms, Inc.</p>
  </div>
</div>