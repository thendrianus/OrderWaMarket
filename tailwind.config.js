/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{svelte,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'whatsapp': {
          DEFAULT: '#25D366',
          light: '#DCF8C6',
          dark: '#1da051',
        }
      },
    },
  },
  plugins: [],
}