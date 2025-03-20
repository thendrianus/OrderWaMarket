// Import route components
import Home from './Home.svelte';
import NotFound from './NotFound.svelte';
import Store from './Store.svelte';
import Admin from './Admin.svelte';
import Dashboard from './Dashboard.svelte';

// Define routes
export const routes = {
  // Home page
  '/': Home,
  
  // Store page (public-facing catalog)
  '/store/:storeId': Store,
  
  // Admin pages
  '/admin': Admin,
  '/admin/dashboard': Dashboard,
  
  // 404 page
  '*': NotFound
};