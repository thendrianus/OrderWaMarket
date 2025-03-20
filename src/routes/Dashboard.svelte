<script>
  import { onMount } from 'svelte';
  import { push, location } from 'svelte-spa-router';
  
  // Auth check
  let isAuthenticated = false;
  let isLoading = true;
  
  // Store data
  let store = null;
  
  // Products management
  let products = [];
  let showAddProductModal = false;
  let editingProduct = null;
  
  // Form fields
  let productName = '';
  let productDescription = '';
  let productPrice = '';
  let productCategory = '';
  let productImage = '';
  
  // Error state
  let error = '';
  
  // Format price to IDR
  const formatPrice = (price) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
    }).format(price);
  };
  
  // Initialize modal for adding product
  const openAddProductModal = () => {
    editingProduct = null;
    productName = '';
    productDescription = '';
    productPrice = '';
    productCategory = '';
    productImage = 'https://via.placeholder.com/300';
    showAddProductModal = true;
  };
  
  // Initialize modal for editing product
  const openEditProductModal = (product) => {
    editingProduct = product;
    productName = product.name;
    productDescription = product.description;
    productPrice = product.price.toString();
    productCategory = product.category;
    productImage = product.image;
    showAddProductModal = true;
  };
  
  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();
    
    if (!productName || !productDescription || !productPrice || !productCategory) {
      error = 'Please fill in all fields';
      return;
    }
    
    const productData = {
      name: productName,
      description: productDescription,
      price: parseFloat(productPrice),
      category: productCategory,
      image: productImage || 'https://via.placeholder.com/300'
    };
    
    if (editingProduct) {
      // Update existing product
      const index = products.findIndex(p => p.id === editingProduct.id);
      
      if (index !== -1) {
        products[index] = {
          ...products[index],
          ...productData
        };
        
        products = [...products]; // Trigger reactivity
      }
    } else {
      // Add new product
      const newProduct = {
        id: Date.now(), // Generate unique ID
        ...productData
      };
      
      products = [...products, newProduct];
    }
    
    // Close modal
    showAddProductModal = false;
  };
  
  // Delete product
  const deleteProduct = (productId) => {
    if (confirm('Are you sure you want to delete this product?')) {
      products = products.filter(p => p.id !== productId);
    }
  };
  
  // Handle logout
  const handleLogout = () => {
    localStorage.removeItem('auth_token');
    push('/admin');
  };
  
  // Copy store URL to clipboard
  const copyStoreUrl = () => {
    const storeUrl = `${window.location.origin}/store/${store.id}`;
    navigator.clipboard.writeText(storeUrl)
      .then(() => {
        alert('Store URL copied to clipboard!');
      })
      .catch(err => {
        console.error('Could not copy URL: ', err);
      });
  };
  
  // Check authentication
  onMount(() => {
    const token = localStorage.getItem('auth_token');
    
    if (!token) {
      push('/admin');
      return;
    }
    
    isAuthenticated = true;
    
    // In a real app, we'd validate the token with the backend
    // and load the store data for this admin
    
    // Demo data
    setTimeout(() => {
      store = {
        id: 'demo',
        name: 'Toko Sembako Jaya',
        description: 'Your friendly neighborhood grocery store with the best prices',
        whatsappNumber: '6281234567890',
        logo: 'https://via.placeholder.com/150'
      };
      
      products = [
        {
          id: 1,
          name: 'Indomie Goreng',
          description: 'Instant fried noodles',
          price: 3500,
          image: 'https://via.placeholder.com/300',
          category: 'Noodles'
        },
        {
          id: 2,
          name: 'Beras Cap Bunga 5kg',
          description: 'Premium quality rice',
          price: 68000,
          image: 'https://via.placeholder.com/300',
          category: 'Rice'
        },
        {
          id: 3,
          name: 'Minyak Goreng Filma 1L',
          description: 'Cooking oil',
          price: 24000,
          image: 'https://via.placeholder.com/300',
          category: 'Oil'
        },
        {
          id: 4,
          name: 'Telur Ayam 1kg',
          description: 'Fresh chicken eggs',
          price: 30000,
          image: 'https://via.placeholder.com/300',
          category: 'Eggs'
        },
        {
          id: 5,
          name: 'Gula Pasir 1kg',
          description: 'White sugar',
          price: 16000,
          image: 'https://via.placeholder.com/300',
          category: 'Sugar'
        },
        {
          id: 6,
          name: 'Teh Pucuk Harum 350ml',
          description: 'Jasmine tea',
          price: 5000,
          image: 'https://via.placeholder.com/300',
          category: 'Beverages'
        }
      ];
      
      isLoading = false;
    }, 1000);
  });
</script>

<div class="min-h-screen flex flex-col">
  <!-- Header -->
  <header class="bg-primary text-white py-4">
    <div class="container-custom flex justify-between items-center">
      <h1 class="text-xl font-bold">Dashboard</h1>
      
      <button 
        class="btn bg-white text-primary"
        on:click={handleLogout}
      >
        Logout
      </button>
    </div>
  </header>

  <!-- Main Content -->
  <main class="flex-grow py-8">
    <!-- Loading State -->
    {#if isLoading}
      <div class="container-custom flex justify-center items-center h-64">
        <div class="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-primary"></div>
      </div>
    {/if}
    
    <!-- Dashboard Content (when loaded) -->
    {#if !isLoading && store}
      <div class="container-custom">
        <!-- Store Information -->
        <section class="bg-white p-6 rounded-lg shadow-md mb-8">
          <div class="flex flex-col md:flex-row gap-6 items-start">
            <div class="flex-grow">
              <h2 class="text-2xl font-semibold mb-4">Store Information</h2>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <p class="text-gray-600 mb-1">Store Name</p>
                  <p class="font-medium">{store.name}</p>
                </div>
                
                <div>
                  <p class="text-gray-600 mb-1">Description</p>
                  <p class="font-medium">{store.description}</p>
                </div>
                
                <div>
                  <p class="text-gray-600 mb-1">WhatsApp Number</p>
                  <p class="font-medium">{store.whatsappNumber}</p>
                </div>
              </div>
            </div>
            
            <div class="flex flex-col items-center">
              <div class="bg-gray-100 p-2 rounded-md mb-3">
                <img src={store.logo} alt="{store.name} logo" class="w-20 h-20 object-contain" />
              </div>
              <button class="text-sm text-primary">Update Logo</button>
            </div>
          </div>
          
          <div class="mt-6 flex flex-col sm:flex-row gap-3">
            <button 
              class="btn btn-secondary"
              on:click={() => push(`/store/${store.id}`)}
            >
              View Store
            </button>
            
            <button 
              class="btn border border-gray-300 bg-white"
              on:click={copyStoreUrl}
            >
              Copy Store URL
            </button>
            
            <button 
              class="btn border border-gray-300 bg-white"
            >
              Edit Store
            </button>
          </div>
        </section>
        
        <!-- Products Management -->
        <section class="bg-white p-6 rounded-lg shadow-md">
          <div class="flex justify-between items-center mb-6">
            <h2 class="text-2xl font-semibold">Products</h2>
            <button 
              class="btn btn-primary"
              on:click={openAddProductModal}
            >
              Add Product
            </button>
          </div>
          
          {#if products.length === 0}
            <div class="py-12 text-center text-gray-500">
              <p>No products yet. Click "Add Product" to create your first product.</p>
            </div>
          {:else}
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Product</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Category</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Price</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  {#each products as product}
                    <tr>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="flex-shrink-0 h-10 w-10">
                            <img src={product.image} alt={product.name} class="h-10 w-10 rounded-md object-cover" />
                          </div>
                          <div class="ml-4">
                            <div class="text-sm font-medium text-gray-900">{product.name}</div>
                            <div class="text-sm text-gray-500">{product.description}</div>
                          </div>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                          {product.category}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {formatPrice(product.price)}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <button 
                          class="text-primary hover:text-primary-dark mr-3"
                          on:click={() => openEditProductModal(product)}
                        >
                          Edit
                        </button>
                        <button 
                          class="text-red-600 hover:text-red-800"
                          on:click={() => deleteProduct(product.id)}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>
          {/if}
        </section>
      </div>
    {/if}
  </main>

  <!-- Add/Edit Product Modal -->
  {#if showAddProductModal}
    <div class="fixed inset-0 z-50 overflow-auto bg-black bg-opacity-50 flex justify-center items-center p-4">
      <div class="bg-white rounded-lg max-w-md w-full">
        <div class="p-6">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-xl font-semibold">
              {editingProduct ? 'Edit Product' : 'Add New Product'}
            </h3>
            <button 
              class="text-gray-500 hover:text-gray-700"
              on:click={() => showAddProductModal = false}
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          {#if error}
            <div class="bg-red-100 text-red-700 p-3 rounded-md mb-4">
              {error}
            </div>
          {/if}
          
          <form on:submit={handleSubmit}>
            <div class="mb-4">
              <label for="productName" class="block text-gray-700 mb-2">Product Name</label>
              <input
                type="text"
                id="productName"
                bind:value={productName}
                class="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                placeholder="e.g. Indomie Goreng"
                required
              />
            </div>
            
            <div class="mb-4">
              <label for="productDescription" class="block text-gray-700 mb-2">Description</label>
              <textarea
                id="productDescription"
                bind:value={productDescription}
                class="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                placeholder="e.g. Instant fried noodles"
                rows="2"
                required
              ></textarea>
            </div>
            
            <div class="mb-4">
              <label for="productPrice" class="block text-gray-700 mb-2">Price (IDR)</label>
              <input
                type="number"
                id="productPrice"
                bind:value={productPrice}
                class="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                placeholder="e.g. 3500"
                min="0"
                step="100"
                required
              />
            </div>
            
            <div class="mb-4">
              <label for="productCategory" class="block text-gray-700 mb-2">Category</label>
              <input
                type="text"
                id="productCategory"
                bind:value={productCategory}
                class="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                placeholder="e.g. Noodles"
                required
              />
            </div>
            
            <div class="mb-6">
              <label for="productImage" class="block text-gray-700 mb-2">Image URL</label>
              <input
                type="url"
                id="productImage"
                bind:value={productImage}
                class="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                placeholder="https://..."
              />
            </div>
            
            <div class="flex justify-end">
              <button
                type="button"
                class="btn border border-gray-300 bg-white mr-2"
                on:click={() => showAddProductModal = false}
              >
                Cancel
              </button>
              <button
                type="submit"
                class="btn btn-primary"
              >
                {editingProduct ? 'Save Changes' : 'Add Product'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  {/if}

  <footer class="bg-gray-100 py-4 mt-auto border-t">
    <div class="container-custom text-center text-gray-600 text-sm">
      <p>&copy; 2025 WhatsApp Catalogue. All rights reserved.</p>
    </div>
  </footer>
</div>