<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  
  // Login form data
  let username = '';
  let password = '';
  let isLoggingIn = true;
  let isLoading = false;
  let error = '';
  
  // Registration form data
  let name = '';
  let whatsappNumber = '';
  let confirmPassword = '';
  
  // Handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault();
    error = '';
    isLoading = true;
    
    try {
      if (isLoggingIn) {
        // Login
        if (!username || !password) {
          throw new Error('Please enter both username and password');
        }
        
        // In a real app, we would make an API call to the backend
        // For now, we'll mock a successful login
        localStorage.setItem('auth_token', 'dummy_token');
        push('/admin/dashboard');
      } else {
        // Registration
        if (!username || !password || !name || !whatsappNumber || !confirmPassword) {
          throw new Error('Please fill in all fields');
        }
        
        if (password !== confirmPassword) {
          throw new Error('Passwords do not match');
        }
        
        if (whatsappNumber && !whatsappNumber.startsWith('+')) {
          throw new Error('WhatsApp number should start with country code (e.g. +62)');
        }
        
        // In a real app, we would make an API call to the backend
        // For now, we'll mock a successful registration
        localStorage.setItem('auth_token', 'dummy_token');
        push('/admin/dashboard');
      }
    } catch (err) {
      error = err.message;
    } finally {
      isLoading = false;
    }
  };
  
  // Check if already logged in
  onMount(() => {
    const token = localStorage.getItem('auth_token');
    if (token) {
      push('/admin/dashboard');
    }
  });
</script>

<div class="min-h-screen bg-gray-50 flex flex-col">
  <header class="bg-primary text-white py-4">
    <div class="container-custom">
      <div class="flex justify-between items-center">
        <h1 class="text-xl font-bold">WhatsApp Catalogue</h1>
        <button 
          class="text-white hover:underline"
          on:click={() => push('/')}
        >
          Back to Home
        </button>
      </div>
    </div>
  </header>
  
  <main class="flex-grow flex items-center justify-center py-12 px-4">
    <div class="bg-white rounded-lg shadow-md p-8 max-w-md w-full">
      <!-- Form Header -->
      <div class="text-center mb-8">
        <h2 class="text-2xl font-bold text-gray-800">
          {isLoggingIn ? 'Sign In to Your Account' : 'Create a New Account'}
        </h2>
        <p class="text-gray-600 mt-2">
          {isLoggingIn ? 'Manage your store and products' : 'Start selling your products with WhatsApp'}
        </p>
      </div>
      
      <!-- Error Message -->
      {#if error}
        <div class="bg-red-100 text-red-800 px-4 py-3 rounded-md mb-4">
          {error}
        </div>
      {/if}
      
      <!-- Form -->
      <form on:submit={handleSubmit}>
        {#if !isLoggingIn}
          <!-- Store Name (Registration only) -->
          <div class="mb-4">
            <label for="name" class="block text-gray-700 mb-1 font-medium">Store Name</label>
            <input
              type="text"
              id="name"
              bind:value={name}
              class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
              placeholder="e.g. Toko Sembako Jaya"
            />
          </div>
          
          <!-- WhatsApp Number (Registration only) -->
          <div class="mb-4">
            <label for="whatsappNumber" class="block text-gray-700 mb-1 font-medium">WhatsApp Number</label>
            <input
              type="text"
              id="whatsappNumber"
              bind:value={whatsappNumber}
              class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
              placeholder="e.g. +628123456789"
            />
            <p class="text-xs text-gray-500 mt-1">Include your country code (e.g. +62 for Indonesia)</p>
          </div>
        {/if}
        
        <!-- Username field -->
        <div class="mb-4">
          <label for="username" class="block text-gray-700 mb-1 font-medium">Username</label>
          <input
            type="text"
            id="username"
            bind:value={username}
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
            placeholder="Enter your username"
          />
        </div>
        
        <!-- Password field -->
        <div class="mb-4">
          <label for="password" class="block text-gray-700 mb-1 font-medium">Password</label>
          <input
            type="password"
            id="password"
            bind:value={password}
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
            placeholder="Enter your password"
          />
        </div>
        
        {#if !isLoggingIn}
          <!-- Confirm Password field (Registration only) -->
          <div class="mb-4">
            <label for="confirmPassword" class="block text-gray-700 mb-1 font-medium">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              bind:value={confirmPassword}
              class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
              placeholder="Confirm your password"
            />
          </div>
        {/if}
        
        <!-- Submit button -->
        <button
          type="submit"
          class="btn btn-primary w-full py-2.5 mt-2"
          disabled={isLoading}
        >
          {#if isLoading}
            <span class="inline-block h-4 w-4 border-2 border-white border-t-transparent rounded-full animate-spin mr-2"></span>
            {isLoggingIn ? 'Signing In...' : 'Creating Account...'}
          {:else}
            {isLoggingIn ? 'Sign In' : 'Create Account'}
          {/if}
        </button>
      </form>
      
      <!-- Form footer (switch between login/register) -->
      <div class="text-center mt-6">
        <p class="text-gray-600">
          {#if isLoggingIn}
            Don't have an account?
            <button 
              class="text-primary font-medium hover:underline focus:outline-none"
              on:click={() => {
                isLoggingIn = false;
                error = '';
              }}
            >
              Register
            </button>
          {:else}
            Already have an account?
            <button 
              class="text-primary font-medium hover:underline focus:outline-none"
              on:click={() => {
                isLoggingIn = true;
                error = '';
              }}
            >
              Sign In
            </button>
          {/if}
        </p>
      </div>
    </div>
  </main>
  
  <footer class="bg-gray-800 text-white py-4">
    <div class="container-custom text-center">
      <p>© 2025 WhatsApp Catalogue. All rights reserved.</p>
    </div>
  </footer>
</div>