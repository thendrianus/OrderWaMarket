// vite.config.js
export default {
  plugins: [],
  server: {
    proxy: {
      // Proxy API requests to the backend
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false
      }
    }
  },
  resolve: {
    alias: {
      '@': './src'
    }
  },
  optimizeDeps: {
    exclude: ['@sveltejs/vite-plugin-svelte']
  },
  build: {
    rollupOptions: {
      external: ['@sveltejs/vite-plugin-svelte']
    }
  }
}