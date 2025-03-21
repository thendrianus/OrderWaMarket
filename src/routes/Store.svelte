<script>
  import { onMount } from 'svelte';
  import { location } from 'svelte-spa-router';
  
  // Get store ID from the route parameters
  $: storeId = $location.split('/').pop();
  
  let store = null;
  let products = [];
  let selectedProducts = [];
  let cart = [];
  let loading = {
    store: true,
    products: true
  };
  let error = {
    store: null,
    products: null
  };
  
  // Format price to Indonesian Rupiah
  function formatPrice(price) {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
    }).format(price);
  }
  
  // Add product to cart
  function addToCart(product) {
    const existingItem = cart.find(item => item.id === product.id);
    
    if (existingItem) {
      cart = cart.map(item => 
        item.id === product.id 
          ? { ...item, quantity: item.quantity + 1 } 
          : item
      );
    } else {
      cart = [...cart, { ...product, quantity: 1 }];
    }
    
    // Update selected products for WhatsApp order
    selectedProducts = [...selectedProducts, product.id];
  }
  
  // Remove product from cart
  function removeFromCart(productId) {
    const existingItem = cart.find(item => item.id === productId);
    
    if (existingItem && existingItem.quantity > 1) {
      cart = cart.map(item => 
        item.id === productId 
          ? { ...item, quantity: item.quantity - 1 } 
          : item
      );
    } else {
      cart = cart.filter(item => item.id !== productId);
    }
    
    // Update selected products for WhatsApp order
    if (selectedProducts.includes(productId)) {
      selectedProducts = selectedProducts.filter(id => id !== productId);
    }
  }
  
  // Generate WhatsApp order message
  function generateOrderMessage() {
    if (!store || cart.length === 0) return '';
    
    const storeInfo = `*Order from ${store.name}*\n\n`;
    
    const itemsList = cart.map(item => 
      `*${item.name}*\n` +
      `${formatPrice(item.price)} x ${item.quantity} = ${formatPrice(item.price * item.quantity)}`
    ).join('\n\n');
    
    const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
    
    return encodeURIComponent(
      storeInfo + 
      itemsList + 
      '\n\n*Total: ' + formatPrice(total) + '*' +
      '\n\nThank you!'
    );
  }
  
  // Generate WhatsApp order URL
  function getWhatsAppOrderUrl() {
    if (!store || !store.whatsappNumber) return '#';
    
    const message = generateOrderMessage();
    return `https://wa.me/${store.whatsappNumber}?text=${message}`;
  }
  
  // Fetch store data
  async function fetchStore() {
    loading.store = true;
    error.store = null;
    
    try {
      const response = await fetch(`/api/stores/${storeId}`);
      
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
    loading.products = true;
    error.products = null;
    
    try {
      const response = await fetch(`/api/stores/${storeId}/products`);
      
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
  onMount(() => {
    fetchStore();
    fetchProducts();
  });
</script>

<svelte:head>
  {#if store}
    <title>{store.name} | WhatsApp Catalogue</title>
    <meta name="description" content={store.description || `Browse products from ${store.name} and order via WhatsApp`} />
  {:else}
    <title>Store | WhatsApp Catalogue</title>
  {/if}
</svelte:head>

<div class="bg-gray-50 min-h-screen pb-20">
  <!-- Store Header -->
  {#if loading.store}
    <div class="bg-gradient-to-r from-[#128C7E] to-[#25D366] py-16">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-center">
          <div class="animate-pulse flex space-x-4 items-center">
            <div class="rounded-full bg-white/20 h-24 w-24"></div>
            <div class="space-y-2">
              <div class="h-8 bg-white/20 rounded w-64"></div>
              <div class="h-4 bg-white/20 rounded w-40"></div>
              <div class="h-4 bg-white/20 rounded w-52"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  {:else if error.store}
    <div class="bg-red-600 py-16 text-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-2xl font-bold mb-4">Error Loading Store</h1>
        <p class="mb-6">{error.store}</p>
        <button 
          class="px-4 py-2 bg-white text-red-600 rounded-md font-medium hover:bg-gray-100"
          on:click={() => fetchStore()}
        >
          Try Again
        </button>
      </div>
    </div>
  {:else if store}
    <div class={store.coverImage ? 'relative' : 'bg-gradient-to-r from-[#128C7E] to-[#25D366]'}>
      {#if store.coverImage}
        <div 
          class="absolute inset-0 bg-center bg-cover"
          style={`background-image: url(${store.coverImage})`}
        >
          <div class="absolute inset-0 bg-black/40"></div>
        </div>
      {/if}
      
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 relative">
        <div class="flex flex-col md:flex-row items-center md:items-start text-center md:text-left gap-6">
          <div class="w-28 h-28 bg-white rounded-full overflow-hidden shadow-lg flex-shrink-0">
            {#if store.logo}
              <img src={store.logo} alt={store.name} class="w-full h-full object-cover" />
            {:else}
              <div class="w-full h-full bg-[#25D366] flex items-center justify-center">
                <span class="text-white font-bold text-3xl">
                  {store.name ? store.name.charAt(0).toUpperCase() : 'S'}
                </span>
              </div>
            {/if}
          </div>
          
          <div class="text-white">
            <h1 class="text-3xl font-bold mb-2">{store.name}</h1>
            {#if store.description}
              <p class="text-white/90 mb-4 max-w-2xl">{store.description}</p>
            {/if}
            
            <div class="flex flex-wrap justify-center md:justify-start gap-4 text-sm">
              {#if store.location}
                <div class="flex items-center">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd" />
                  </svg>
                  <span>{store.location}</span>
                </div>
              {/if}
              
              {#if store.businessHours}
                <div class="flex items-center">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
                  </svg>
                  <span>{store.businessHours}</span>
                </div>
              {/if}
            </div>
          </div>
        </div>
      </div>
    </div>
  {/if}
  
  <!-- Products Section -->
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    {#if loading.products}
      <div class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-[#25d366]"></div>
      </div>
    {:else if error.products}
      <div class="bg-red-50 p-6 rounded-lg text-center text-red-700 max-w-lg mx-auto">
        <p class="text-xl font-semibold mb-2">Failed to load products</p>
        <p>{error.products}</p>
        <button 
          class="mt-4 px-4 py-2 bg-red-100 text-red-800 rounded hover:bg-red-200 transition-colors"
          on:click={() => fetchProducts()}
        >
          Try Again
        </button>
      </div>
    {:else if products.length === 0}
      <div class="bg-gray-100 p-10 rounded-lg text-center max-w-lg mx-auto">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <h3 class="text-xl font-semibold text-gray-700 mb-2">No Products Available</h3>
        <p class="text-gray-600">This store hasn't added any products yet.</p>
      </div>
    {:else}
      <div class="flex flex-col-reverse lg:flex-row">
        <!-- Products Grid -->
        <div class="w-full lg:w-2/3 pr-0 lg:pr-8">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">Products</h2>
          
          <!-- Featured Products -->
          {#if products.some(p => p.featured)}
            <div class="mb-8">
              <h3 class="text-lg font-medium text-gray-900 mb-4">Featured</h3>
              <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
                {#each products.filter(p => p.featured && p.active) as product (product.id)}
                  <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
                    {#if product.image}
                      <div class="aspect-w-16 aspect-h-9">
                        <img src={product.image} alt={product.name} class="w-full h-48 object-cover" />
                      </div>
                    {:else}
                      <div class="w-full h-48 bg-gray-200 flex items-center justify-center">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                      </div>
                    {/if}
                    
                    <div class="p-4">
                      <div class="flex justify-between items-start">
                        <div>
                          <h3 class="text-lg font-medium text-gray-900">{product.name}</h3>
                          <p class="text-[#25D366] font-semibold">{formatPrice(product.price)}</p>
                        </div>
                        
                        <span class="bg-yellow-100 text-yellow-800 text-xs font-semibold px-2 py-1 rounded">Featured</span>
                      </div>
                      
                      {#if product.description}
                        <p class="text-gray-600 my-2 text-sm line-clamp-2">{product.description}</p>
                      {/if}
                      
                      <div class="mt-3 flex">
                        <button 
                          on:click={() => addToCart(product)}
                          class="flex-1 bg-[#25D366] text-white py-2 rounded-md hover:bg-[#1da051] transition-colors"
                        >
                          Add to Cart
                        </button>
                      </div>
                    </div>
                  </div>
                {/each}
              </div>
            </div>
          {/if}
          
          <!-- All Products -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">All Products</h3>
            <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
              {#each products.filter(p => p.active) as product (product.id)}
                <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
                  {#if product.image}
                    <div class="aspect-w-16 aspect-h-9">
                      <img src={product.image} alt={product.name} class="w-full h-48 object-cover" />
                    </div>
                  {:else}
                    <div class="w-full h-48 bg-gray-200 flex items-center justify-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  {/if}
                  
                  <div class="p-4">
                    <div class="flex justify-between items-start">
                      <div>
                        <h3 class="text-lg font-medium text-gray-900">{product.name}</h3>
                        <p class="text-[#25D366] font-semibold">{formatPrice(product.price)}</p>
                      </div>
                      
                      {#if product.category}
                        <span class="bg-gray-100 text-gray-800 text-xs font-semibold px-2 py-1 rounded">{product.category}</span>
                      {/if}
                    </div>
                    
                    {#if product.description}
                      <p class="text-gray-600 my-2 text-sm line-clamp-2">{product.description}</p>
                    {/if}
                    
                    <div class="mt-3 flex">
                      <button 
                        on:click={() => addToCart(product)}
                        class="flex-1 bg-[#25D366] text-white py-2 rounded-md hover:bg-[#1da051] transition-colors"
                      >
                        Add to Cart
                      </button>
                    </div>
                  </div>
                </div>
              {/each}
            </div>
          </div>
        </div>
        
        <!-- Cart Sidebar -->
        <div class="w-full lg:w-1/3 mb-6 lg:mb-0">
          <div class="bg-white rounded-lg shadow-md p-6 sticky top-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">Your Order</h2>
            
            {#if cart.length === 0}
              <div class="py-8 text-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400 mx-auto mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
                <p class="text-gray-500">Your cart is empty</p>
                <p class="text-gray-500 text-sm mt-1">Add some products to order</p>
              </div>
            {:else}
              <div class="mb-4 divide-y divide-gray-200">
                {#each cart as item (item.id)}
                  <div class="py-3 flex justify-between items-center">
                    <div>
                      <p class="text-gray-900 font-medium">{item.name}</p>
                      <p class="text-gray-600 text-sm">{formatPrice(item.price)} x {item.quantity}</p>
                    </div>
                    <div class="flex items-center">
                      <p class="font-semibold text-gray-900 mr-3">{formatPrice(item.price * item.quantity)}</p>
                      <div class="flex border border-gray-300 rounded">
                        <button 
                          on:click={() => removeFromCart(item.id)}
                          class="px-2 py-1 text-gray-600 hover:bg-gray-100"
                        >
                          -
                        </button>
                        <span class="px-2 py-1 border-l border-r border-gray-300 bg-gray-50">{item.quantity}</span>
                        <button 
                          on:click={() => addToCart(item)}
                          class="px-2 py-1 text-gray-600 hover:bg-gray-100"
                        >
                          +
                        </button>
                      </div>
                    </div>
                  </div>
                {/each}
              </div>
              
              <div class="py-3 border-t border-gray-200">
                <div class="flex justify-between items-center font-bold text-lg text-gray-900">
                  <span>Total</span>
                  <span>{formatPrice(cart.reduce((sum, item) => sum + (item.price * item.quantity), 0))}</span>
                </div>
              </div>
              
              <div class="mt-4">
                <a 
                  href={getWhatsAppOrderUrl()}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="block w-full text-center bg-[#25D366] text-white py-3 rounded-md font-medium hover:bg-[#1da051] transition-colors"
                >
                  Order via WhatsApp
                </a>
                <p class="text-xs text-gray-500 mt-2 text-center">
                  Your order details will be sent to the store's WhatsApp.
                </p>
              </div>
            {/if}
          </div>
        </div>
      </div>
    {/if}
  </div>
  
  <!-- Footer -->
  <footer class="bg-white py-8 mt-16 border-t border-gray-200">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
      <p class="text-gray-500 text-sm">
        © {new Date().getFullYear()} WhatsApp Catalogue. All rights reserved.
      </p>
      <p class="text-gray-500 text-sm mt-2">
        WhatsApp is a trademark of Meta Platforms, Inc.
      </p>
    </div>
  </footer>
</div>