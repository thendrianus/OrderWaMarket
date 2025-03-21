<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  
  let username = '';
  let password = '';
  let error = null;
  let loading = false;
  
  // Check if user is already logged in
  onMount(() => {
    const token = localStorage.getItem('whatsapp_catalogue_token');
    if (token) {
      // Redirect to dashboard if token exists
      push('/admin/dashboard');
    }
  });
  
  async function handleLogin(event) {
    event.preventDefault();
    
    if (!username || !password) {
      error = 'Please enter both username and password';
      return;
    }
    
    loading = true;
    error = null;
    
    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      });
      
      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.message || 'Login failed');
      }
      
      // Save token and user info to localStorage
      localStorage.setItem('whatsapp_catalogue_token', data.token);
      localStorage.setItem('whatsapp_catalogue_user', JSON.stringify(data.user));
      
      // Redirect to dashboard
      push('/admin/dashboard');
    } catch (err) {
      console.error('Login error:', err);
      error = err.message || 'Failed to login. Please try again.';
    } finally {
      loading = false;
    }
  }
  
  function goToRegister() {
    push('/admin/register');
  }
</script>

<div class="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
  <div class="max-w-md w-full space-y-8">
    <div>
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        Admin Login
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        Sign in to manage your store catalogue
      </p>
    </div>
    
    {#if error}
      <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
        <span class="block sm:inline">{error}</span>
      </div>
    {/if}
    
    <form class="mt-8 space-y-6" on:submit={handleLogin}>
      <div class="rounded-md shadow-sm -space-y-px">
        <div>
          <label for="username" class="sr-only">Username</label>
          <input 
            id="username" 
            name="username" 
            type="text" 
            bind:value={username}
            required 
            class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-[#25D366] focus:border-[#25D366] focus:z-10 sm:text-sm" 
            placeholder="Username"
          >
        </div>
        <div>
          <label for="password" class="sr-only">Password</label>
          <input 
            id="password" 
            name="password" 
            type="password" 
            bind:value={password}
            required 
            class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-[#25D366] focus:border-[#25D366] focus:z-10 sm:text-sm" 
            placeholder="Password"
          >
        </div>
      </div>

      <div class="flex items-center justify-end">
        <div class="text-sm">
          <button 
            type="button" 
            on:click={goToRegister}
            class="font-medium text-[#25D366] hover:text-[#128C7E]"
          >
            Create new account?
          </button>
        </div>
      </div>

      <div>
        <button 
          type="submit" 
          disabled={loading}
          class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-[#25D366] hover:bg-[#128C7E] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25D366] disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {#if loading}
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
              <div class="animate-spin h-5 w-5 border-2 border-white border-t-transparent rounded-full"></div>
            </span>
            Loading...
          {:else}
            Sign in
          {/if}
        </button>
      </div>
    </form>
    
    <div class="mt-6">
      <div class="relative">
        <div class="absolute inset-0 flex items-center">
          <div class="w-full border-t border-gray-300"></div>
        </div>
        <div class="relative flex justify-center text-sm">
          <span class="px-2 bg-gray-50 text-gray-500">
            Or go to
          </span>
        </div>
      </div>
      <div class="mt-6 flex justify-center">
        <button
          on:click={() => push('/')}
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25D366]"
        >
          Home Page
        </button>
      </div>
    </div>
  </div>
</div>