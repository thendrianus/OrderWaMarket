<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  
  let user = null;
  let loading = {
    user: true,
    submit: false
  };
  let error = {
    user: null,
    submit: null
  };
  
  // Store form data
  let formData = {
    name: '',
    description: '',
    location: '',
    whatsappNumber: '',
    businessHours: '',
    logo: '',
    tags: ''
  };
  
  // Form validation state
  let formErrors = {
    name: '',
    whatsappNumber: ''
  };
  
  // Get authentication token
  function getAuthHeaders() {
    const token = localStorage.getItem('token');
    return {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    };
  }
  
  // Handle logout
  function handleLogout() {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    sessionStorage.removeItem('user');
    push('/admin/login');
  }
  
  // Fetch user data
  async function fetchUser() {
    loading.user = true;
    error.user = null;
    
    try {
      // Get user from storage
      const userJson = localStorage.getItem('user') || sessionStorage.getItem('user');
      if (userJson) {
        user = JSON.parse(userJson);
      } else {
        throw new Error('User data not found, please login again');
      }
    } catch (err) {
      console.error('Failed to load user:', err);
      error.user = err.message;
      // If there's an error loading user data, redirect to login
      setTimeout(() => handleLogout(), 2000);
    } finally {
      loading.user = false;
    }
  }
  
  // Validate form data
  function validateForm() {
    let isValid = true;
    formErrors = {
      name: '',
      whatsappNumber: ''
    };
    
    // Validate store name
    if (!formData.name.trim()) {
      formErrors.name = 'Store name is required';
      isValid = false;
    }
    
    // Validate WhatsApp number
    if (!formData.whatsappNumber.trim()) {
      formErrors.whatsappNumber = 'WhatsApp number is required';
      isValid = false;
    } else {
      // Basic phone number validation
      const phoneRegex = /^[0-9+\-\s()]*$/;
      if (!phoneRegex.test(formData.whatsappNumber)) {
        formErrors.whatsappNumber = 'Please enter a valid phone number';
        isValid = false;
      }
    }
    
    return isValid;
  }
  
  // Handle form submission
  async function handleSubmit(event) {
    event.preventDefault();
    
    if (!validateForm()) {
      return;
    }
    
    loading.submit = true;
    error.submit = null;
    
    try {
      // Convert tags from string to array if provided
      const tagsArray = formData.tags
        ? formData.tags.split(',').map(tag => tag.trim()).filter(tag => tag)
        : [];
      
      // Prepare request body
      const storeData = {
        name: formData.name,
        description: formData.description,
        location: formData.location,
        whatsappNumber: formData.whatsappNumber,
        businessHours: formData.businessHours,
        logo: formData.logo,
        tags: tagsArray
      };
      
      // Send request to create store
      const response = await fetch('/api/stores', {
        method: 'POST',
        headers: getAuthHeaders(),
        body: JSON.stringify(storeData)
      });
      
      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.message || `Error ${response.status}: ${response.statusText}`);
      }
      
      // Navigate to dashboard
      push('/admin/dashboard');
    } catch (err) {
      console.error('Failed to create store:', err);
      error.submit = err.message || 'Failed to create store. Please try again.';
      window.scrollTo(0, 0);
    } finally {
      loading.submit = false;
    }
  }
  
  // Initialize component
  onMount(() => {
    fetchUser();
  });
</script>

<svelte:head>
  <title>Store Setup | WhatsApp Catalogue</title>
</svelte:head>

<div class="min-h-screen bg-gray-100">
  <!-- Header -->
  <header class="bg-white shadow">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex items-center justify-between">
      <div class="flex items-center">
        <div class="bg-[#25d366] p-2 rounded-full mr-3">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path>
          </svg>
        </div>
        <h1 class="text-xl font-semibold text-gray-800">Store Setup</h1>
      </div>
      
      <!-- User menu -->
      <div class="flex items-center">
        {#if user}
          <span class="mr-4 text-gray-600">{user.username}</span>
        {/if}
        <button 
          on:click={handleLogout}
          class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded-md text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
        >
          Logout
        </button>
      </div>
    </div>
  </header>
  
  <!-- Main content -->
  <main class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    {#if loading.user}
      <div class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-[#25d366]"></div>
      </div>
    {:else if error.user}
      <div class="bg-red-50 p-6 rounded-lg text-center text-red-700 my-8">
        <p class="text-xl font-semibold mb-2">Authentication Error</p>
        <p>{error.user}</p>
        <p class="text-sm mt-2">Redirecting to login...</p>
      </div>
    {:else}
      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="px-4 py-5 sm:p-6">
          <h2 class="text-2xl font-bold text-gray-900 mb-1">Create Your Store</h2>
          <p class="text-gray-600 mb-6">Set up your WhatsApp catalogue to start selling your products</p>
          
          {#if error.submit}
            <div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm mb-6">
              {error.submit}
            </div>
          {/if}
          
          <form on:submit={handleSubmit} class="space-y-6">
            <!-- Store Name -->
            <div>
              <label for="name" class="block text-sm font-medium text-gray-700">
                Store Name <span class="text-red-600">*</span>
              </label>
              <div class="mt-1">
                <input
                  type="text"
                  id="name"
                  bind:value={formData.name}
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="My Awesome Store"
                  required
                />
              </div>
              {#if formErrors.name}
                <p class="mt-1 text-sm text-red-600">{formErrors.name}</p>
              {/if}
            </div>
            
            <!-- Store Description -->
            <div>
              <label for="description" class="block text-sm font-medium text-gray-700">
                Description
              </label>
              <div class="mt-1">
                <textarea
                  id="description"
                  bind:value={formData.description}
                  rows="3"
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="Tell customers about your store..."
                ></textarea>
              </div>
            </div>
            
            <!-- Store Location -->
            <div>
              <label for="location" class="block text-sm font-medium text-gray-700">
                Location
              </label>
              <div class="mt-1">
                <input
                  type="text"
                  id="location"
                  bind:value={formData.location}
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="Jakarta, Indonesia"
                />
              </div>
            </div>
            
            <!-- WhatsApp Number -->
            <div>
              <label for="whatsappNumber" class="block text-sm font-medium text-gray-700">
                WhatsApp Number <span class="text-red-600">*</span>
              </label>
              <div class="mt-1">
                <input
                  type="text"
                  id="whatsappNumber"
                  bind:value={formData.whatsappNumber}
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="628123456789"
                  required
                />
              </div>
              {#if formErrors.whatsappNumber}
                <p class="mt-1 text-sm text-red-600">{formErrors.whatsappNumber}</p>
              {:else}
                <p class="mt-1 text-xs text-gray-500">Enter your WhatsApp number including country code (e.g., 628123456789)</p>
              {/if}
            </div>
            
            <!-- Business Hours -->
            <div>
              <label for="businessHours" class="block text-sm font-medium text-gray-700">
                Business Hours
              </label>
              <div class="mt-1">
                <input
                  type="text"
                  id="businessHours"
                  bind:value={formData.businessHours}
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="Mon-Fri: 9AM-5PM, Sat: 10AM-3PM"
                />
              </div>
            </div>
            
            <!-- Logo URL -->
            <div>
              <label for="logo" class="block text-sm font-medium text-gray-700">
                Logo URL
              </label>
              <div class="mt-1">
                <input
                  type="url"
                  id="logo"
                  bind:value={formData.logo}
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="https://example.com/logo.png"
                />
              </div>
              <p class="mt-1 text-xs text-gray-500">Enter a URL to your store logo image</p>
            </div>
            
            <!-- Tags -->
            <div>
              <label for="tags" class="block text-sm font-medium text-gray-700">
                Tags
              </label>
              <div class="mt-1">
                <input
                  type="text"
                  id="tags"
                  bind:value={formData.tags}
                  class="shadow-sm focus:ring-[#25d366] focus:border-[#25d366] block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="food, grocery, organic"
                />
              </div>
              <p class="mt-1 text-xs text-gray-500">Separate tags with commas</p>
            </div>
            
            <!-- Form Actions -->
            <div class="flex items-center justify-end space-x-4 pt-4">
              <button 
                type="button"
                on:click={() => push('/')}
                class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
              >
                Cancel
              </button>
              <button 
                type="submit"
                disabled={loading.submit}
                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366] disabled:opacity-50"
              >
                {#if loading.submit}
                  <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  Creating Store...
                {:else}
                  Create Store
                {/if}
              </button>
            </div>
          </form>
        </div>
      </div>
    {/if}
  </main>
  
  <!-- Footer -->
  <footer class="mt-16 pt-8 pb-8 border-t border-gray-200 text-center text-gray-500 text-sm">
    <p>© {new Date().getFullYear()} WhatsApp Catalogue. All rights reserved.</p>
    <p class="mt-2">WhatsApp is a trademark of Meta Platforms, Inc.</p>
  </footer>
</div>