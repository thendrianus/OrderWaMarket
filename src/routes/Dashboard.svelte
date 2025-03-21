<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  
  let user = null;
  let store = null;
  let products = [];
  let loading = {
    user: true,
    store: true,
    products: true
  };
  let error = {
    user: null,
    store: null,
    products: null
  };
  let activeTab = 'overview';
  
  // Format price to Indonesian Rupiah
  function formatPrice(price) {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
    }).format(price);
  }
  
  // Handle logout
  function handleLogout() {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    sessionStorage.removeItem('user');
    push('/admin/login');
  }
  
  // Get authentication token
  function getAuthHeaders() {
    const token = localStorage.getItem('token');
    return {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    };
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
  
  // Fetch store data
  async function fetchStore() {
    loading.store = true;
    error.store = null;
    
    try {
      const response = await fetch('/api/my-store', {
        headers: getAuthHeaders()
      });
      
      if (response.status === 404) {
        // User doesn't have a store yet
        store = null;
        return;
      }
      
      if (!response.ok) {
        throw new Error(`Error ${response.status}: ${response.statusText}`);
      }
      
      store = await response.json();
    } catch (err) {
      console.error('Failed to load store:', err);
      error.store = err.message;
    } finally {
      loading.store = false;
    }
  }
  
  // Fetch products for a store
  async function fetchProducts() {
    if (!store) return;
    
    loading.products = true;
    error.products = null;
    
    try {
      const response = await fetch(`/api/stores/${store.id}/products`, {
        headers: getAuthHeaders()
      });
      
      if (!response.ok) {
        throw new Error(`Error ${response.status}: ${response.statusText}`);
      }
      
      products = await response.json();
    } catch (err) {
      console.error('Failed to load products:', err);
      error.products = err.message;
    } finally {
      loading.products = false;
    }
  }
  
  // Initialize component
  onMount(async () => {
    await fetchUser();
    await fetchStore();
    if (store) {
      await fetchProducts();
    }
  });
  
  // Navigate to store setup if no store exists
  $: if (!loading.store && !error.store && !store) {
    push('/admin/store/setup');
  }
</script>

<svelte:head>
  <title>Admin Dashboard | WhatsApp Catalogue</title>
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
        <h1 class="text-xl font-semibold text-gray-800">Admin Dashboard</h1>
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
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M3 3a1 1 0 00-1 1v12a1 1 0 001 1h12a1 1 0 001-1V7.414l-5-5H3zm7 5a1 1 0 10-2 0v6a1 1 0 102 0V8zm5 2a1 1 0 10-2 0v4a1 1 0 102 0v-4z" clip-rule="evenodd" />
          </svg>
          Logout
        </button>
      </div>
    </div>
  </header>
  
  <!-- Main content -->
  <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    {#if loading.user || loading.store}
      <div class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-[#25d366]"></div>
      </div>
    {:else if error.user}
      <div class="bg-red-50 p-6 rounded-lg text-center text-red-700 my-8">
        <p class="text-xl font-semibold mb-2">Authentication Error</p>
        <p>{error.user}</p>
        <p class="text-sm mt-2">Redirecting to login...</p>
      </div>
    {:else if error.store}
      <div class="bg-red-50 p-6 rounded-lg text-center text-red-700 my-8">
        <p class="text-xl font-semibold mb-2">Failed to load store data</p>
        <p>{error.store}</p>
        <button 
          class="mt-4 px-4 py-2 bg-red-100 text-red-800 rounded hover:bg-red-200 transition-colors"
          on:click={() => window.location.reload()}
        >
          Try Again
        </button>
      </div>
    {:else if store}
      <!-- Store exists, show dashboard -->
      <div class="bg-white shadow rounded-lg overflow-hidden">
        <!-- Store header -->
        <div class="p-6 bg-gray-50 border-b border-gray-200">
          <div class="flex flex-col md:flex-row md:items-center md:justify-between">
            <div>
              <h2 class="text-2xl font-bold text-gray-800">{store.name}</h2>
              <p class="text-gray-600 mt-1">{store.description || 'No description available'}</p>
            </div>
            
            <div class="mt-4 md:mt-0 flex items-center">
              <span class="{store.active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'} px-3 py-1 rounded-full text-xs font-medium mr-3">
                {store.active ? 'Active' : 'Inactive'}
              </span>              
              
              <button 
                class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1.5" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                </svg>
                Edit Store
              </button>
            </div>
          </div>
        </div>
        
        <!-- Tabs -->
        <div class="border-b border-gray-200">
          <nav class="flex -mb-px">
            <button 
            class="{activeTab === 'overview' ? 'border-[#25d366] text-[#25d366]' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'} whitespace-nowrap py-4 px-6 border-b-2 font-medium text-sm"
            on:click={() => activeTab = 'overview'}
          >
            Overview
          </button>
          
          <button 
            class="{activeTab === 'products' ? 'border-[#25d366] text-[#25d366]' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'} whitespace-nowrap py-4 px-6 border-b-2 font-medium text-sm"
            on:click={() => activeTab = 'products'}
          >
            Products
          </button>
          
          <button 
            class="{activeTab === 'settings' ? 'border-[#25d366] text-[#25d366]' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'} whitespace-nowrap py-4 px-6 border-b-2 font-medium text-sm"
            on:click={() => activeTab = 'settings'}
          >
            Settings
          </button>          
          </nav>
        </div>
        
        <!-- Tab content -->
        <div class="p-6">
          {#if activeTab === 'overview'}
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <!-- Store Info Card -->
              <div class="bg-white overflow-hidden shadow rounded-lg">
                <div class="px-4 py-5 sm:p-6">
                  <h3 class="text-lg font-medium text-gray-900 mb-3">Store Information</h3>
                  
                  <div class="space-y-3">
                    {#if store.location}
                      <div>
                        <dt class="text-sm font-medium text-gray-500">Location</dt>
                        <dd class="mt-1 text-sm text-gray-900">{store.location}</dd>
                      </div>
                    {/if}
                    
                    {#if store.whatsappNumber}
                      <div>
                        <dt class="text-sm font-medium text-gray-500">WhatsApp Number</dt>
                        <dd class="mt-1 text-sm text-gray-900">{store.whatsappNumber}</dd>
                      </div>
                    {/if}
                    
                    {#if store.businessHours}
                      <div>
                        <dt class="text-sm font-medium text-gray-500">Business Hours</dt>
                        <dd class="mt-1 text-sm text-gray-900">{store.businessHours}</dd>
                      </div>
                    {/if}
                    
                    <div>
                      <dt class="text-sm font-medium text-gray-500">Created</dt>
                      <dd class="mt-1 text-sm text-gray-900">{new Date(store.createdAt).toLocaleDateString()}</dd>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Product Stats Card -->
              <div class="bg-white overflow-hidden shadow rounded-lg">
                <div class="px-4 py-5 sm:p-6">
                  <h3 class="text-lg font-medium text-gray-900 mb-3">Product Stats</h3>
                  
                  {#if loading.products}
                    <div class="flex justify-center py-4">
                      <div class="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-[#25d366]"></div>
                    </div>
                  {:else if error.products}
                    <p class="text-red-600 text-sm">{error.products}</p>
                  {:else}
                    <div class="space-y-4">
                      <div>
                        <p class="text-sm font-medium text-gray-500">Total Products</p>
                        <p class="text-3xl font-semibold text-gray-900">{products.length}</p>
                      </div>
                      
                      <div>
                        <p class="text-sm font-medium text-gray-500">Active Products</p>
                        <p class="text-3xl font-semibold text-gray-900">{products.filter(p => p.active).length}</p>
                      </div>
                      
                      <div>
                        <p class="text-sm font-medium text-gray-500">Featured Products</p>
                        <p class="text-3xl font-semibold text-gray-900">{products.filter(p => p.featured).length}</p>
                      </div>
                    </div>
                  {/if}
                </div>
              </div>
              
              <!-- Store Preview Card -->
              <div class="bg-white overflow-hidden shadow rounded-lg">
                <div class="px-4 py-5 sm:p-6">
                  <h3 class="text-lg font-medium text-gray-900 mb-3">Store Preview</h3>
                  
                  <div class="flex flex-col items-center">
                    {#if store.logo}
                      <img src={store.logo} alt={store.name} class="w-32 h-32 object-cover rounded-lg mb-4" />
                    {:else}
                      <div class="w-32 h-32 bg-gray-200 rounded-lg flex items-center justify-center mb-4">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                          <path fill-rule="evenodd" d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z" clip-rule="evenodd" />
                        </svg>
                      </div>
                    {/if}
                    
                    <button 
                      on:click={() => window.open(`/store/${store.id}`, '_blank')}
                      class="mt-2 inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
                    >
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1.5" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                        <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd" />
                      </svg>
                      View Store
                    </button>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Recent products -->
            {#if !loading.products && !error.products && products.length > 0}
              <div class="mt-8">
                <h3 class="text-lg font-medium text-gray-900 mb-4">Recent Products</h3>
                
                <div class="overflow-x-auto">
                  <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                      <tr>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Product
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Price
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Category
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Status
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Stock
                        </th>
                      </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                      {#each products.slice(0, 5) as product}
                        <tr>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="flex items-center">
                              <div class="flex-shrink-0 h-10 w-10 bg-gray-200 rounded-md overflow-hidden">
                                {#if product.image}
                                  <img src={product.image} alt={product.name} class="h-10 w-10 object-cover" />
                                {:else}
                                  <div class="h-10 w-10 flex items-center justify-center text-gray-400">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                                    </svg>
                                  </div>
                                {/if}
                              </div>
                              <div class="ml-4">
                                <div class="text-sm font-medium text-gray-900">{product.name}</div>
                                <div class="text-xs text-gray-500 truncate max-w-xs">{product.description || 'No description'}</div>
                              </div>
                            </div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="text-sm text-gray-900">{formatPrice(product.price)}</div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="text-sm text-gray-900">{product.category || 'Uncategorized'}</div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <span class="{product.active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'} px-2 inline-flex text-xs leading-5 font-semibold rounded-full py-1">
                              {product.active ? 'Active' : 'Inactive'}
                            </span>
                            
                            {#if product.featured}
                              <span class="ml-1 bg-yellow-100 text-yellow-800 px-2 inline-flex text-xs leading-5 font-semibold rounded-full py-1">
                                Featured
                              </span>
                            {/if}
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {product.stock !== undefined ? product.stock : 'N/A'}
                          </td>
                        </tr>
                      {/each}
                    </tbody>
                  </table>
                </div>
                
                {#if products.length > 5}
                  <div class="mt-4 text-center">
                    <button 
                      on:click={() => activeTab = 'products'}
                      class="text-sm text-[#25d366] hover:text-[#1da051] hover:underline"
                    >
                      View all products
                    </button>
                  </div>
                {/if}
              </div>
            {/if}
            
          {:else if activeTab === 'products'}
            <div class="flex justify-between items-center mb-6">
              <h3 class="text-lg font-medium text-gray-900">Products</h3>
              
              <button 
                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1.5" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
                </svg>
                Add Product
              </button>
            </div>
            
            {#if loading.products}
              <div class="flex justify-center py-12">
                <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-[#25d366]"></div>
              </div>
            {:else if error.products}
              <div class="bg-red-50 p-4 rounded-lg text-center text-red-700">
                <p>{error.products}</p>
                <button 
                  class="mt-2 px-4 py-2 bg-red-100 text-red-800 rounded hover:bg-red-200 transition-colors"
                  on:click={() => fetchProducts()}
                >
                  Try Again
                </button>
              </div>
            {:else if products.length === 0}
              <div class="bg-gray-50 p-8 rounded-lg text-center text-gray-600">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
                </svg>
                <p class="text-xl">No products available yet</p>
                <p class="mt-2 text-gray-500">Add your first product to start selling</p>
                <button 
                  class="mt-4 px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
                >
                  Add First Product
                </button>
              </div>
            {:else}
              <div class="bg-white shadow overflow-hidden rounded-lg">
                <div class="overflow-x-auto">
                  <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                      <tr>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Product
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Price
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Category
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Status
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Stock
                        </th>
                        <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Actions
                        </th>
                      </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                      {#each products as product}
                        <tr>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="flex items-center">
                              <div class="flex-shrink-0 h-10 w-10 bg-gray-200 rounded-md overflow-hidden">
                                {#if product.image}
                                  <img src={product.image} alt={product.name} class="h-10 w-10 object-cover" />
                                {:else}
                                  <div class="h-10 w-10 flex items-center justify-center text-gray-400">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                                    </svg>
                                  </div>
                                {/if}
                              </div>
                              <div class="ml-4">
                                <div class="text-sm font-medium text-gray-900">{product.name}</div>
                                <div class="text-xs text-gray-500 truncate max-w-xs">{product.description || 'No description'}</div>
                              </div>
                            </div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="text-sm text-gray-900">{formatPrice(product.price)}</div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="text-sm text-gray-900">{product.category || 'Uncategorized'}</div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <span class="{product.active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'} px-2 inline-flex text-xs leading-5 font-semibold rounded-full py-1">
                              {product.active ? 'Active' : 'Inactive'}
                            </span>
                            
                            {#if product.featured}
                              <span class="ml-1 bg-yellow-100 text-yellow-800 px-2 inline-flex text-xs leading-5 font-semibold rounded-full py-1">
                                Featured
                              </span>
                            {/if}
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {product.stock !== undefined ? product.stock : 'N/A'}
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                            <button class="text-[#25d366] hover:text-[#1da051] mr-3">Edit</button>
                            <button class="text-red-600 hover:text-red-900">Delete</button>
                          </td>
                        </tr>
                      {/each}
                    </tbody>
                  </table>
                </div>
              </div>
            {/if}
            
          {:else if activeTab === 'settings'}
            <div class="bg-white shadow rounded-lg overflow-hidden">
              <div class="px-4 py-5 sm:p-6">
                <h3 class="text-lg font-medium text-gray-900 mb-4">Store Settings</h3>
                
                <div class="max-w-xl">
                  <div class="divide-y divide-gray-200">
                    <div class="py-4">
                      <h4 class="text-sm font-medium text-gray-900">Store Visibility</h4>
                      <p class="mt-1 text-sm text-gray-500">Control whether your store is visible to customers</p>
                      <div class="mt-3 flex items-center">
                        <button 
                          class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]"
                          class:bg-[#25d366]={store.active}
                          class:bg-gray-200={!store.active}
                          aria-pressed={store.active}
                          aria-labelledby="store-active-label"
                        >
                          <span 
                            class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200"
                            class:translate-x-5={store.active}
                            class:translate-x-0={!store.active}
                          ></span>
                        </button>
                        <span class="ml-3" id="store-active-label">
                          <span class="text-sm font-medium text-gray-900">{store.active ? 'Active' : 'Inactive'}</span>
                        </span>
                      </div>
                    </div>
                    
                    <div class="py-4">
                      <h4 class="text-sm font-medium text-gray-900">Store Information</h4>
                      <p class="mt-1 text-sm text-gray-500">Update your store details</p>
                      <div class="mt-3">
                        <button class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]">
                          Edit Store Information
                        </button>
                      </div>
                    </div>
                    
                    <div class="py-4">
                      <h4 class="text-sm font-medium text-gray-900">Change WhatsApp Number</h4>
                      <p class="mt-1 text-sm text-gray-500">Update the WhatsApp number customers use to contact you</p>
                      <div class="mt-3">
                        <button class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366]">
                          Change Number
                        </button>
                      </div>
                    </div>
                    
                    <div class="py-4">
                      <h4 class="text-sm font-medium text-red-600">Delete Store</h4>
                      <p class="mt-1 text-sm text-gray-500">Permanently remove your store and all its products</p>
                      <div class="mt-3">
                        <button class="inline-flex items-center px-3 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
                          Delete Store
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          {/if}
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