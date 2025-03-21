<script>
  import Router from 'svelte-spa-router';
  import { wrap } from 'svelte-spa-router/wrap';
  
  // Import route components
  import Home from './routes/Home.svelte';
  import Store from './routes/Store.svelte';
  import AdminLogin from './routes/AdminLogin.svelte';
  import AdminRegister from './routes/AdminRegister.svelte';
  import Dashboard from './routes/Dashboard.svelte';
  import StoreSetup from './routes/StoreSetup.svelte';
  import NotFound from './routes/NotFound.svelte';
  
  // Define routes
  const routes = {
    // Public routes
    '/': Home,
    '/store/:storeId': Store,
    
    // Admin routes
    '/admin/login': AdminLogin,
    '/admin/register': AdminRegister,
    '/admin/dashboard': wrap({
      component: Dashboard,
      conditions: [
        () => {
          // Check if user is authenticated
          return !!localStorage.getItem('token') || !!sessionStorage.getItem('token');
        }
      ],
      // Redirect to login if not authenticated
      userData: { redirectTo: '/admin/login' }
    }),
    '/admin/store/setup': wrap({
      component: StoreSetup,
      conditions: [
        () => {
          // Check if user is authenticated
          return !!localStorage.getItem('token') || !!sessionStorage.getItem('token');
        }
      ],
      // Redirect to login if not authenticated
      userData: { redirectTo: '/admin/login' }
    }),
    
    // Catch-all route for 404 errors
    '*': NotFound
  };
  
  // Handle route conditional failures
  function conditionsFailed(event) {
    const { route, userData } = event.detail;
    if (userData && userData.redirectTo) {
      window.location.href = userData.redirectTo;
    }
  }
</script>

<Router {routes} on:conditionsFailed={conditionsFailed} />