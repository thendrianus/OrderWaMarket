<script>
  import { push } from 'svelte-spa-router';
</script>

<div class="min-h-screen flex flex-col">
  <!-- Hero Section -->
  <section class="bg-primary text-white py-16">
    <div class="container-custom">
      <div class="max-w-3xl mx-auto text-center">
        <h1 class="text-4xl md:text-5xl font-bold mb-6">WhatsApp Catalogue System</h1>
        <p class="text-xl mb-8">Create a beautiful online product catalogue for your business and receive orders directly through WhatsApp</p>
        <button 
          class="btn bg-white text-primary font-semibold px-6 py-3 hover:bg-gray-100"
          on:click={() => push('/admin')}
        >
          Create Your Store
        </button>
      </div>
    </div>
  </section>

  <!-- Features Section -->
  <section class="py-16 bg-white">
    <div class="container-custom">
      <h2 class="text-3xl font-bold text-center mb-12">Why Choose Our Platform?</h2>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <!-- Feature 1 -->
        <div class="bg-gray-50 p-6 rounded-lg text-center">
          <div class="bg-primary/10 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-primary" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="2" y="3" width="20" height="14" rx="2" ry="2"></rect>
              <path d="M8 21h8"></path>
              <path d="M12 17v4"></path>
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-2">Mobile-Friendly</h3>
          <p class="text-gray-600">Perfect responsive design for all devices with fast loading times</p>
        </div>
        
        <!-- Feature 2 -->
        <div class="bg-gray-50 p-6 rounded-lg text-center">
          <div class="bg-primary/10 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-primary" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"></path>
              <path d="m15 5 4 4"></path>
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-2">Easy Management</h3>
          <p class="text-gray-600">Simple dashboard to add, edit, and organize your products</p>
        </div>
        
        <!-- Feature 3 -->
        <div class="bg-gray-50 p-6 rounded-lg text-center">
          <div class="bg-primary/10 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-primary" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M3 18v-6a9 9 0 0 1 18 0v6"></path>
              <path d="M21 19a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3zM3 19a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2v-3a2 2 0 0 0-2-2H3z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-2">WhatsApp Integration</h3>
          <p class="text-gray-600">Receive orders directly to your WhatsApp without any third-party apps</p>
        </div>
      </div>
    </div>
  </section>

  <!-- How It Works Section -->
  <section class="py-16 bg-gray-50">
    <div class="container-custom">
      <h2 class="text-3xl font-bold text-center mb-12">How It Works</h2>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
        <!-- Step 1 -->
        <div class="relative">
          <div class="bg-white p-6 rounded-lg shadow-sm">
            <div class="absolute -top-5 -left-5 w-10 h-10 rounded-full bg-primary text-white flex items-center justify-center font-bold text-xl">
              1
            </div>
            <h3 class="text-xl font-semibold mb-3 mt-2">Register Your Store</h3>
            <p class="text-gray-600">Create an account and set up your store with basic information and your WhatsApp number</p>
          </div>
        </div>
        
        <!-- Step 2 -->
        <div class="relative">
          <div class="bg-white p-6 rounded-lg shadow-sm">
            <div class="absolute -top-5 -left-5 w-10 h-10 rounded-full bg-primary text-white flex items-center justify-center font-bold text-xl">
              2
            </div>
            <h3 class="text-xl font-semibold mb-3 mt-2">Add Your Products</h3>
            <p class="text-gray-600">Upload your products with names, descriptions, prices and categorize them as needed</p>
          </div>
        </div>
        
        <!-- Step 3 -->
        <div class="relative">
          <div class="bg-white p-6 rounded-lg shadow-sm">
            <div class="absolute -top-5 -left-5 w-10 h-10 rounded-full bg-primary text-white flex items-center justify-center font-bold text-xl">
              3
            </div>
            <h3 class="text-xl font-semibold mb-3 mt-2">Share Your Catalogue</h3>
            <p class="text-gray-600">Share your unique store link with customers to browse your products and order via WhatsApp</p>
          </div>
        </div>
      </div>
      
      <div class="mt-12 text-center">
        <button 
          class="btn btn-primary px-6 py-3"
          on:click={() => push('/admin')}
        >
          Get Started Today
        </button>
      </div>
    </div>
  </section>

  <!-- Sample Stores Section -->
  <section class="py-16 bg-white">
    <div class="container-custom">
      <h2 class="text-3xl font-bold text-center mb-12">Stores Using Our Platform</h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <!-- Sample Store 1 -->
        <div class="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
          <div class="h-48 bg-gray-100 flex items-center justify-center">
            <img src="https://via.placeholder.com/300x200" alt="Toko Sembako Jaya" class="object-cover w-full h-full" />
          </div>
          <div class="p-4">
            <h3 class="text-lg font-semibold mb-2">Toko Sembako Jaya</h3>
            <p class="text-gray-600 mb-4">Grocery store with all your daily necessities at affordable prices</p>
            <button 
              class="text-primary font-medium hover:underline"
              on:click={() => push('/store/demo')}
            >
              Visit Store →
            </button>
          </div>
        </div>
        
        <!-- Sample Store 2 -->
        <div class="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
          <div class="h-48 bg-gray-100 flex items-center justify-center">
            <img src="https://via.placeholder.com/300x200" alt="Warung Bu Tini" class="object-cover w-full h-full" />
          </div>
          <div class="p-4">
            <h3 class="text-lg font-semibold mb-2">Warung Bu Tini</h3>
            <p class="text-gray-600 mb-4">Delicious homemade meals and snacks for pickup or delivery</p>
            <button 
              class="text-primary font-medium hover:underline"
              on:click={() => push('/store/demo')}
            >
              Visit Store →
            </button>
          </div>
        </div>
        
        <!-- Sample Store 3 -->
        <div class="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
          <div class="h-48 bg-gray-100 flex items-center justify-center">
            <img src="https://via.placeholder.com/300x200" alt="Kosmetik Cantik" class="object-cover w-full h-full" />
          </div>
          <div class="p-4">
            <h3 class="text-lg font-semibold mb-2">Kosmetik Cantik</h3>
            <p class="text-gray-600 mb-4">Quality beauty products and cosmetics for all skin types</p>
            <button 
              class="text-primary font-medium hover:underline"
              on:click={() => push('/store/demo')}
            >
              Visit Store →
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- CTA Section -->
  <section class="py-16 bg-primary text-white">
    <div class="container-custom text-center">
      <h2 class="text-3xl font-bold mb-6">Ready to Grow Your Business?</h2>
      <p class="text-xl mb-8 max-w-2xl mx-auto">Join hundreds of business owners who are expanding their reach with our WhatsApp catalogue system</p>
      <button 
        class="btn bg-white text-primary font-semibold px-6 py-3 hover:bg-gray-100"
        on:click={() => push('/admin')}
      >
        Create Your Store For Free
      </button>
    </div>
  </section>

  <!-- Footer -->
  <footer class="bg-gray-800 text-white py-8">
    <div class="container-custom">
      <div class="flex flex-col md:flex-row justify-between items-center">
        <div class="mb-4 md:mb-0">
          <h2 class="text-xl font-bold">WhatsApp Catalogue</h2>
          <p class="text-gray-400 mt-1">© 2025 WhatsApp Catalogue. All rights reserved.</p>
        </div>
        
        <div>
          <ul class="flex space-x-6">
            <li><a href="#" class="hover:text-primary-light">About</a></li>
            <li><a href="#" class="hover:text-primary-light">Features</a></li>
            <li><a href="#" class="hover:text-primary-light">Pricing</a></li>
            <li><a href="#" class="hover:text-primary-light">Contact</a></li>
          </ul>
        </div>
      </div>
    </div>
  </footer>
</div>