<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  
  let featuredStores = [];
  let loading = true;
  let error = null;
  
  // Fetch featured stores
  async function fetchFeaturedStores() {
    loading = true;
    error = null;
    
    try {
      const response = await fetch('/api/stores/featured');
      
      if (!response.ok) {
        throw new Error(`Error ${response.status}: ${response.statusText}`);
      }
      
      featuredStores = await response.json();
    } catch (err) {
      console.error('Failed to load featured stores:', err);
      error = err.message;
    } finally {
      loading = false;
    }
  }
  
  onMount(() => {
    fetchFeaturedStores();
  });
</script>

<svelte:head>
  <title>WhatsApp Catalogue | Browse Products & Order via WhatsApp</title>
  <meta name="description" content="Browse products from local businesses and order directly via WhatsApp" />
</svelte:head>

<div class="bg-gray-50 min-h-screen">
  <!-- Hero Section -->
  <section class="bg-gradient-to-r from-[#128C7E] to-[#25D366] text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24 text-center md:text-left">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
        <div>
          <h1 class="text-4xl md:text-5xl font-bold mb-6">
            Simple Product Catalogues for <span class="whitespace-nowrap">WhatsApp Shops</span>
          </h1>
          <p class="text-xl mb-8 text-gray-100">
            Browse products from local businesses and order directly via WhatsApp. No complicated checkout process!
          </p>
          <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 justify-center md:justify-start">
            <button 
              on:click={() => push('/admin/register')}
              class="px-8 py-3 bg-white text-[#128C7E] font-semibold rounded-md hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-75 shadow-md"
            >
              Create Your Catalogue
            </button>
            <button 
              on:click={() => document.getElementById('featured').scrollIntoView({ behavior: 'smooth' })}
              class="px-8 py-3 bg-transparent border-2 border-white text-white font-semibold rounded-md hover:bg-white/10 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-75"
            >
              Browse Stores
            </button>
          </div>
        </div>
        
        <div class="hidden md:block">
          <div class="relative">
            <div class="absolute -top-6 -left-6 w-64 h-64 bg-[#34B7F1] opacity-20 rounded-full filter blur-3xl"></div>
            <div class="absolute -bottom-8 -right-8 w-64 h-64 bg-[#25D366] opacity-20 rounded-full filter blur-3xl"></div>
            
            <div class="relative bg-white p-6 rounded-xl shadow-2xl transform rotate-3">
              <div class="flex items-center mb-4">
                <div class="w-12 h-12 bg-[#25D366] rounded-full flex items-center justify-center">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                  </svg>
                </div>
                <div class="ml-4">
                  <h3 class="text-gray-900 font-semibold">Toko Sembako Abadi</h3>
                  <p class="text-gray-500 text-sm">Open 8am - 9pm</p>
                </div>
              </div>
              
              <div class="grid grid-cols-2 gap-3 mb-4">
                <div class="bg-gray-100 p-2 rounded-lg">
                  <div class="aspect-w-1 aspect-h-1 mb-2">
                    <div class="w-full h-24 bg-gray-200 rounded-md flex items-center justify-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  </div>
                  <p class="text-gray-800 font-medium text-sm">Rice 5kg</p>
                  <p class="text-[#25D366] font-semibold">Rp 65.000</p>
                </div>
                
                <div class="bg-gray-100 p-2 rounded-lg">
                  <div class="aspect-w-1 aspect-h-1 mb-2">
                    <div class="w-full h-24 bg-gray-200 rounded-md flex items-center justify-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  </div>
                  <p class="text-gray-800 font-medium text-sm">Cooking Oil 1L</p>
                  <p class="text-[#25D366] font-semibold">Rp 28.000</p>
                </div>
              </div>
              
              <button class="w-full bg-[#25D366] text-white py-3 rounded-lg font-medium">
                Order via WhatsApp
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  
  <!-- How It Works -->
  <section class="py-20 bg-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-16">How It Works</h2>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
        <!-- For Customers -->
        <div class="text-center">
          <div class="bg-[#DCF8C6] w-16 h-16 mx-auto rounded-full flex items-center justify-center mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-[#128C7E]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-3 text-gray-900">Browse Products</h3>
          <p class="text-gray-600">
            Find products from local stores. View prices, descriptions, and availability without leaving the page.
          </p>
        </div>
        
        <!-- Order -->
        <div class="text-center">
          <div class="bg-[#DCF8C6] w-16 h-16 mx-auto rounded-full flex items-center justify-center mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-[#128C7E]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-3 text-gray-900">Select and Order</h3>
          <p class="text-gray-600">
            Choose your products and click "Order via WhatsApp" to connect directly with the store owner.
          </p>
        </div>
        
        <!-- Receive -->
        <div class="text-center">
          <div class="bg-[#DCF8C6] w-16 h-16 mx-auto rounded-full flex items-center justify-center mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-[#128C7E]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-3 text-gray-900">Confirm and Receive</h3>
          <p class="text-gray-600">
            Confirm your order via WhatsApp and arrange delivery or pickup with the store owner.
          </p>
        </div>
      </div>
    </div>
  </section>
  
  <!-- Featured Stores -->
  <section id="featured" class="py-20 bg-gray-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-16">Featured Stores</h2>
      
      {#if loading}
        <div class="flex justify-center py-12">
          <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-[#25d366]"></div>
        </div>
      {:else if error}
        <div class="bg-red-50 p-6 rounded-lg text-center text-red-700 max-w-lg mx-auto">
          <p class="text-xl font-semibold mb-2">Oops! Something went wrong</p>
          <p>{error}</p>
          <button 
            class="mt-4 px-4 py-2 bg-red-100 text-red-800 rounded hover:bg-red-200 transition-colors"
            on:click={() => fetchFeaturedStores()}
          >
            Try Again
          </button>
        </div>
      {:else if featuredStores.length === 0}
        <div class="bg-gray-100 p-10 rounded-lg text-center max-w-lg mx-auto">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
          </svg>
          <h3 class="text-xl font-semibold text-gray-700 mb-2">No Stores Yet</h3>
          <p class="text-gray-600 mb-6">Be the first to create a store on our platform!</p>
          <button 
            on:click={() => push('/admin/register')}
            class="px-6 py-2 bg-[#25D366] text-white font-medium rounded-md hover:bg-[#1da051] transition-colors"
          >
            Create Your Store
          </button>
        </div>
      {:else}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          {#each featuredStores as store}
            <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
              <div 
                class="h-40 bg-gray-200 bg-center bg-cover"
                style={store.coverImage ? `background-image: url(${store.coverImage})` : ''}
              >
                {#if !store.coverImage}
                  <div class="h-full flex items-center justify-center bg-gradient-to-r from-[#128C7E] to-[#25D366]">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-white opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                    </svg>
                  </div>
                {/if}
              </div>
              
              <div class="p-6">
                <div class="flex items-center mb-4">
                  <div class="w-12 h-12 bg-gray-200 rounded-full overflow-hidden flex-shrink-0">
                    {#if store.logo}
                      <img src={store.logo} alt={store.name} class="w-full h-full object-cover" />
                    {:else}
                      <div class="w-full h-full bg-[#25D366] flex items-center justify-center">
                        <span class="text-white font-bold text-lg">
                          {store.name ? store.name.charAt(0).toUpperCase() : 'S'}
                        </span>
                      </div>
                    {/if}
                  </div>
                  
                  <div class="ml-4">
                    <h3 class="font-semibold text-lg text-gray-900">{store.name}</h3>
                    {#if store.location}
                      <p class="text-gray-600 text-sm">{store.location}</p>
                    {/if}
                  </div>
                </div>
                
                {#if store.description}
                  <p class="text-gray-600 mb-6 line-clamp-2">{store.description}</p>
                {:else}
                  <p class="text-gray-500 italic mb-6">No description available</p>
                {/if}
                
                <div class="flex justify-between items-center">
                  {#if store.tags && store.tags.length > 0}
                    <div class="flex flex-wrap gap-2">
                      {#each store.tags.slice(0, 2) as tag}
                        <span class="inline-block bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded">
                          {tag}
                        </span>
                      {/each}
                      {#if store.tags.length > 2}
                        <span class="inline-block bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded">
                          +{store.tags.length - 2}
                        </span>
                      {/if}
                    </div>
                  {:else}
                    <div></div>
                  {/if}
                  
                  <button 
                    on:click={() => push(`/store/${store.id}`)}
                    class="text-[#25D366] font-medium hover:text-[#1da051] transition-colors flex items-center"
                  >
                    Visit Store
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-1" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M10.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L12.586 11H5a1 1 0 110-2h7.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </section>
  
  <!-- For Businesses -->
  <section class="py-20 bg-[#128C7E] text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center max-w-3xl mx-auto">
        <h2 class="text-3xl font-bold mb-6">For Business Owners</h2>
        <p class="text-lg mb-8 text-gray-100">
          Create your own WhatsApp catalogue for free. Showcase your products and receive orders directly via WhatsApp.
        </p>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
          <div class="bg-white/10 backdrop-blur-sm rounded-lg p-6">
            <div class="w-12 h-12 bg-white/20 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold mb-2">Easy Setup</h3>
            <p class="text-white/80">Create your store and add products in minutes, no technical skills required.</p>
          </div>
          
          <div class="bg-white/10 backdrop-blur-sm rounded-lg p-6">
            <div class="w-12 h-12 bg-white/20 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold mb-2">Mobile Friendly</h3>
            <p class="text-white/80">Manage your store and respond to orders from your smartphone.</p>
          </div>
          
          <div class="bg-white/10 backdrop-blur-sm rounded-lg p-6">
            <div class="w-12 h-12 bg-white/20 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold mb-2">No Commission Fees</h3>
            <p class="text-white/80">Keep all your profits - we don't take any commission on your sales.</p>
          </div>
        </div>
        
        <button 
          on:click={() => push('/admin/register')}
          class="px-8 py-3 bg-white text-[#128C7E] font-semibold rounded-md hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-75 shadow-md"
        >
          Get Started - It's Free
        </button>
      </div>
    </div>
  </section>
  
  <!-- Footer -->
  <footer class="bg-gray-800 text-white py-12">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <div>
          <div class="flex items-center mb-4">
            <div class="bg-[#25d366] p-2 rounded-full mr-3">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path>
              </svg>
            </div>
            <span class="text-lg font-semibold">WhatsApp Catalogue</span>
          </div>
          <p class="text-gray-400 mb-4">
            Simple product catalogues for WhatsApp shops. Browse products and order directly.
          </p>
          <p class="text-gray-500 text-sm">
            © {new Date().getFullYear()} WhatsApp Catalogue<br>
            WhatsApp is a trademark of Meta Platforms, Inc.
          </p>
        </div>
        
        <div>
          <h3 class="text-lg font-medium mb-4">For Businesses</h3>
          <ul class="space-y-2">
            <li>
              <a href="/admin/register" class="text-gray-400 hover:text-white transition-colors">Create a Store</a>
            </li>
            <li>
              <a href="/admin/login" class="text-gray-400 hover:text-white transition-colors">Login</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Pricing</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Features</a>
            </li>
          </ul>
        </div>
        
        <div>
          <h3 class="text-lg font-medium mb-4">For Customers</h3>
          <ul class="space-y-2">
            <li>
              <a href="#featured" class="text-gray-400 hover:text-white transition-colors">Featured Stores</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Categories</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">How It Works</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">FAQ</a>
            </li>
          </ul>
        </div>
        
        <div>
          <h3 class="text-lg font-medium mb-4">Legal</h3>
          <ul class="space-y-2">
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Privacy Policy</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Terms of Service</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Cookie Policy</a>
            </li>
            <li>
              <a href="#" class="text-gray-400 hover:text-white transition-colors">Contact Us</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </footer>
</div>