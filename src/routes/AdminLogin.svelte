<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  
  let formData = {
    username: '',
    password: ''
  };
  let rememberMe = false;
  let loading = false;
  let error = null;
  
  // Handle form submission
  async function handleSubmit(event) {
    event.preventDefault();
    
    if (!formData.username || !formData.password) {
      error = 'Please fill in all fields';
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
        body: JSON.stringify(formData)
      });
      
      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.message || `Error ${response.status}: ${response.statusText}`);
      }
      
      // Store authentication token
      const storage = rememberMe ? localStorage : sessionStorage;
      storage.setItem('token', data.token);
      storage.setItem('user', JSON.stringify(data.user));
      
      // Navigate to dashboard
      push('/admin/dashboard');
    } catch (err) {
      console.error('Login failed:', err);
      error = err.message || 'Login failed. Please try again.';
    } finally {
      loading = false;
    }
  }
  
  // Check if user is already logged in
  onMount(() => {
    const token = localStorage.getItem('token') || sessionStorage.getItem('token');
    if (token) {
      push('/admin/dashboard');
    }
  });
</script>

<svelte:head>
  <title>Admin Login | WhatsApp Catalogue</title>
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col">
  <!-- Header -->
  <header class="bg-white shadow">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex items-center justify-between">
      <div class="flex items-center cursor-pointer" on:click={() => push('/')}>
        <div class="bg-[#25d366] p-2 rounded-full mr-3">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path>
          </svg>
        </div>
        <h1 class="text-xl font-semibold text-gray-800">WhatsApp Catalogue</h1>
      </div>
    </div>
  </header>
  
  <!-- Main Content -->
  <main class="flex-1 flex items-center justify-center px-4 sm:px-6 lg:px-8 py-12">
    <div class="max-w-md w-full">
      <div class="text-center mb-8">
        <h2 class="text-3xl font-extrabold text-gray-900">Sign in to your account</h2>
        <p class="mt-2 text-sm text-gray-600">
          Or
          <a href="/admin/register" class="font-medium text-[#25d366] hover:text-[#1da051]">
            create a new account
          </a>
        </p>
      </div>
      
      {#if error}
        <div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm mb-6">
          {error}
        </div>
      {/if}
      
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" on:submit={handleSubmit}>
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">
              Username
            </label>
            <div class="mt-1">
              <input
                id="username"
                name="username"
                type="text"
                required
                bind:value={formData.username}
                class="whatsapp-input"
                placeholder="your_username"
              />
            </div>
          </div>
          
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">
              Password
            </label>
            <div class="mt-1">
              <input
                id="password"
                name="password"
                type="password"
                required
                bind:value={formData.password}
                class="whatsapp-input"
                placeholder="••••••••"
              />
            </div>
          </div>
          
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="remember-me"
                name="remember-me"
                type="checkbox"
                bind:checked={rememberMe}
                class="h-4 w-4 text-[#25d366] focus:ring-[#25d366] border-gray-300 rounded"
              />
              <label for="remember-me" class="ml-2 block text-sm text-gray-900">
                Remember me
              </label>
            </div>
            
            <div class="text-sm">
              <a href="#" class="font-medium text-[#25d366] hover:text-[#1da051]">
                Forgot your password?
              </a>
            </div>
          </div>
          
          <div>
            <button
              type="submit"
              disabled={loading}
              class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366] disabled:opacity-50"
            >
              {#if loading}
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Signing in...
              {:else}
                Sign in
              {/if}
            </button>
          </div>
        </form>
      </div>
    </div>
  </main>
  
  <!-- Footer -->
  <footer class="py-8 text-center text-gray-500 text-sm">
    <p>© {new Date().getFullYear()} WhatsApp Catalogue. All rights reserved.</p>
    <p class="mt-2">WhatsApp is a trademark of Meta Platforms, Inc.</p>
  </footer>
</div>