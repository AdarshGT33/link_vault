import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  content: [
    "./src/**/*.{js,ts,jsx,tsx}", // adjust based on your file structure
  ],
  theme: {
    extend: {
      colors: {
        matcha: "#C5E1B4",
        blush: "#F4C2C2",
        cream: "#FFF5E1",
        mocha: "#A1887F",
        textPrimary: "#2E2E2E", // deep grey for readable text
      },
      fontFamily: {
        heading: ["Poppins", "sans-serif"],
        body: ["Inter", "Nunito", "sans-serif"],
      },
    },
  },
  plugins: [react(), tailwindcss(),],
})
