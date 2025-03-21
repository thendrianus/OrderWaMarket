import Home from './Home.svelte';
import Store from './Store.svelte';
import Admin from './Admin.svelte';
import Dashboard from './Dashboard.svelte';
import NotFound from './NotFound.svelte';

const routes = {
  '/': Home,
  '/store/:id': Store,
  '/admin': Admin,
  '/admin/dashboard': Dashboard,
  '*': NotFound
};

export default routes;