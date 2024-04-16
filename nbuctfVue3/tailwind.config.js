/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{vue,js,ts,jsx,tsx}', './index.html'],
  important: true, // 使tailwind的css样式优先级最高
  theme: {
    extend: {}
  },
  plugins: []
}
