<script>
  import { onMount } from 'svelte';
  import { push, location } from 'svelte-spa-router';
  
  // State
  let store = null;
  let products = [];
  let categories = [];
  let isLoading = true;
  let error = null;
  
  // Filter states
  let selectedCategory = 'all';
  let searchQuery = '';
  
  // Cart state
  let cart = [];
  let showCart = false;
  
  // Format price to IDR
  const formatPrice = (price) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
    }).format(price);
  };
  
  // Filter products by category and search query
  $: filteredProducts = products.filter(product => {
    // Category filter
    const matchesCategory = selectedCategory === 'all' || product.category === selectedCategory;
    
    // Search filter
    const matchesSearch = searchQuery === '' || 
      product.name.toLowerCase().includes(searchQuery.toLowerCase()) || 
      product.description.toLowerCase().includes(searchQuery.toLowerCase());
    
    return matchesCategory && matchesSearch;
  });
  
  // Add to cart
  const addToCart = (product) => {
    const existingIndex = cart.findIndex(item => item.id === product.id);
    
    if (existingIndex !== -1) {
      // Product already in cart, increment quantity
      cart[existingIndex].quantity += 1;
      cart = [...cart]; // Trigger reactivity
    } else {
      // Add new product to cart
      cart = [...cart, { ...product, quantity: 1 }];
    }
  };
  
  // Remove from cart
  const removeFromCart = (productId) => {
    cart = cart.filter(item => item.id !== productId);
  };
  
  // Update cart item quantity
  const updateQuantity = (productId, newQuantity) => {
    if (newQuantity < 1) return;
    
    const index = cart.findIndex(item => item.id === productId);
    if (index !== -1) {
      cart[index].quantity = newQuantity;
      cart = [...cart]; // Trigger reactivity
    }
  };
  
  // Calculate total price
  $: totalPrice = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
  
  // Generate WhatsApp message from cart
  const generateWhatsAppMessage = () => {
    if (cart.length === 0) return '';
    
    let message = `*New Order from ${store.name}*\n\n`;
    
    // Add each item
    cart.forEach((item, index) => {
      message += `${index + 1}. ${item.name} (${item.quantity}x) - ${formatPrice(item.price * item.quantity)}\n`;
    });
    
    // Add total
    message += `\n*Total: ${formatPrice(totalPrice)}*`;
    
    return encodeURIComponent(message);
  };
  
  // Checkout via WhatsApp
  const checkoutViaWhatsApp = () => {
    if (!store || !store.whatsappNumber || cart.length === 0) return;
    
    const message = generateWhatsAppMessage();
    const whatsappUrl = `https://wa.me/${store.whatsappNumber.replace(/\+/g, '')}?text=${message}`;
    
    window.open(whatsappUrl, '_blank');
  };
  
  // Load store data
  onMount(async () => {
    try {
      // Extract store ID from URL
      const storeId = $location.split('/').pop();
      
      // In a real app, we would fetch the data from the backend API
      // For now, we'll use demo data
      setTimeout(() => {
        store = {
          id: 'demo',
          name: 'Toko Sembako Jaya',
          description: 'Your friendly neighborhood grocery store with the best prices',
          whatsappNumber: '628123456789',
          logo: 'https://via.placeholder.com/150'
        };
        
        products = [
          {
            id: '1',
            name: 'Indomie Goreng',
            description: 'Instant fried noodles',
            price: 3500,
            image: 'https://via.placeholder.com/300',
            category: 'Noodles'
          },
          {
            id: '2',
            name: 'Beras Cap Bunga 5kg',
            description: 'Premium quality rice',
            price: 68000,
            image: 'https://via.placeholder.com/300',
            category: 'Rice'
          },
          {
            id: '3',
            name: 'Minyak Goreng Filma 1L',
            description: 'Cooking oil',
            price: 24000,
            image: 'https://via.placeholder.com/300',
            category: 'Oil'
          },
          {
            id: '4',
            name: 'Telur Ayam 1kg',
            description: 'Fresh chicken eggs',
            price: 30000,
            image: 'https://via.placeholder.com/300',
            category: 'Eggs'
          },
          {
            id: '5',
            name: 'Gula Pasir 1kg',
            description: 'White sugar',
            price: 16000,
            image: 'https://via.placeholder.com/300',
            category: 'Sugar'
          },
          {
            id: '6',
            name: 'Teh Pucuk Harum 350ml',
            description: 'Jasmine tea',
            price: 5000,
            image: 'https://via.placeholder.com/300',
            category: 'Beverages'
          }
        ];
        
        // Extract categories from products
        categories = ['all', ...new Set(products.map(p => p.category))];
        
        isLoading = false;
      }, 1000);
    } catch (err) {
      error = 'Failed to load store data. Please try again later.';
      isLoading = false;
    }
  });
</script>

<div class="min-h-screen flex flex-col bg-gray-50">
  <!-- Header with store info -->
  <header class="bg-primary text-white py-4 sticky top-0 z-10 shadow-md">
    <div class="container-custom">
      <div class="flex justify-between items-center">
        <!-- Store name/logo -->
        <div class="flex items-center">
          {#if store?.logo}
            <img src={store.logo} alt="{store?.name} logo" class="w-10 h-10 rounded-full mr-3 bg-white p-1" />
          {/if}
          <h1 class="text-xl font-bold">{store?.name || 'Store'}</h1>
        </div>
        
        <!-- Cart button -->
        <button 
          class="relative"
          on:click={() => showCart = !showCart}
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
          
          {#if cart.length > 0}
            <span class="absolute -top-2 -right-2 bg-secondary text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center">
              {cart.reduce((sum, item) => sum + item.quantity, 0)}
            </span>
          {/if}
        </button>
      </div>
    </div>
  </header>

  <!-- Loading state -->
  {#if isLoading}
    <div class="flex-grow flex items-center justify-center">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
    </div>
  
  <!-- Error state -->
  {:else if error}
    <div class="flex-grow flex items-center justify-center p-4">
      <div class="bg-red-100 text-red-800 p-4 rounded-md max-w-md text-center">
        <p>{error}</p>
        <button 
          class="mt-4 btn btn-primary"
          on:click={() => window.location.reload()}
        >
          Try Again
        </button>
      </div>
    </div>
  
  <!-- Content loaded -->
  {:else if store}
    <main class="flex-grow py-6">
      <div class="container-custom">
        <!-- Store description -->
        <div class="mb-6 bg-white p-4 rounded-lg shadow-sm">
          <p class="text-gray-700">{store.description}</p>
        </div>
        
        <!-- Search and filter bar -->
        <div class="flex flex-col md:flex-row gap-4 mb-6">
          <!-- Search input -->
          <div class="flex-grow">
            <input
              type="text"
              bind:value={searchQuery}
              placeholder="Search products..."
              class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
            />
          </div>
          
          <!-- Category filter -->
          <div class="md:w-1/3">
            <select 
              bind:value={selectedCategory}
              class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
            >
              {#each categories as category}
                <option value={category}>{category === 'all' ? 'All Categories' : category}</option>
              {/each}
            </select>
          </div>
        </div>
        
        <!-- Products grid -->
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {#each filteredProducts as product (product.id)}
            <div class="bg-white rounded-lg overflow-hidden shadow-sm hover:shadow-md transition-shadow">
              <div class="h-48 bg-gray-100 flex items-center justify-center">
                <img 
                  src={product.image} 
                  alt={product.name} 
                  class="h-full w-full object-cover"
                />
              </div>
              
              <div class="p-4">
                <h3 class="font-semibold text-lg mb-1">{product.name}</h3>
                <p class="text-gray-600 text-sm mb-2">{product.description}</p>
                <p class="text-primary font-bold">{formatPrice(product.price)}</p>
                
                <button 
                  class="mt-3 w-full btn btn-primary py-1.5"
                  on:click={() => addToCart(product)}
                >
                  Add to Cart
                </button>
              </div>
            </div>
          {/each}
          
          {#if filteredProducts.length === 0}
            <div class="col-span-full py-12 text-center text-gray-500">
              <p>No products found. Try a different search or category.</p>
            </div>
          {/if}
        </div>
      </div>
    </main>
    
    <!-- Footer -->
    <footer class="bg-gray-800 text-white py-4 mt-auto">
      <div class="container-custom text-center text-sm">
        <p>© 2025 {store.name}. Powered by WhatsApp Catalogue.</p>
      </div>
    </footer>
    
    <!-- Shopping Cart Sidebar -->
    {#if showCart}
      <div class="fixed inset-0 bg-black bg-opacity-50 z-20" on:click={() => showCart = false}></div>
      
      <div class="fixed top-0 right-0 h-full w-full sm:w-96 bg-white shadow-lg z-30 transform transition-transform duration-300">
        <div class="flex flex-col h-full">
          <!-- Cart header -->
          <div class="bg-primary text-white p-4 flex justify-between items-center">
            <h2 class="text-xl font-bold">Your Cart</h2>
            <button on:click={() => showCart = false}>
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          <!-- Cart content -->
          <div class="flex-grow overflow-y-auto p-4">
            {#if cart.length === 0}
              <div class="h-full flex flex-col items-center justify-center text-gray-500">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
                <p>Your cart is empty</p>
                <button 
                  class="mt-4 text-primary font-medium hover:underline"
                  on:click={() => showCart = false}
                >
                  Continue Shopping
                </button>
              </div>
            {:else}
              <ul class="space-y-4">
                {#each cart as item (item.id)}
                  <li class="flex border-b pb-4">
                    <div class="w-20 h-20 flex-shrink-0 bg-gray-100 rounded-md overflow-hidden">
                      <img src={item.image} alt={item.name} class="w-full h-full object-cover" />
                    </div>
                    
                    <div class="ml-4 flex-grow">
                      <h4 class="font-medium">{item.name}</h4>
                      <p class="text-gray-600 text-sm">{formatPrice(item.price)}</p>
                      
                      <div class="mt-2 flex items-center">
                        <button 
                          class="w-8 h-8 flex items-center justify-center border border-gray-300 rounded-l-md hover:bg-gray-100"
                          on:click={() => updateQuantity(item.id, item.quantity - 1)}
                        >
                          -
                        </button>
                        <input 
                          type="number" 
                          min="1" 
                          bind:value={item.quantity}
                          on:input={(e) => updateQuantity(item.id, parseInt(e.target.value) || 1)}
                          class="w-12 h-8 text-center border-y border-gray-300"
                        />
                        <button 
                          class="w-8 h-8 flex items-center justify-center border border-gray-300 rounded-r-md hover:bg-gray-100"
                          on:click={() => updateQuantity(item.id, item.quantity + 1)}
                        >
                          +
                        </button>
                        
                        <button 
                          class="ml-auto text-red-500 hover:text-red-700"
                          on:click={() => removeFromCart(item.id)}
                        >
                          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                          </svg>
                        </button>
                      </div>
                    </div>
                  </li>
                {/each}
              </ul>
            {/if}
          </div>
          
          <!-- Cart footer -->
          {#if cart.length > 0}
            <div class="border-t p-4">
              <div class="flex justify-between items-center mb-4">
                <span class="font-medium">Total:</span>
                <span class="font-bold text-xl">{formatPrice(totalPrice)}</span>
              </div>
              
              <button 
                class="btn btn-primary w-full py-3 flex items-center justify-center"
                on:click={checkoutViaWhatsApp}
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413Z" />
                </svg>
                Order via WhatsApp
              </button>
            </div>
          {/if}
        </div>
      </div>
    {/if}
  {/if}
</div>