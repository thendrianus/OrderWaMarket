/** @type {import('@sveltejs/vite-plugin-svelte').Config} */
const config = {
  // Consult https://svelte.dev/docs#compile-time-svelte-preprocess
  // for more information about preprocessors
  preprocess: [
    {
      style: ({ content }) => {
        return { code: content };
      }
    }
  ]
};

export default config;