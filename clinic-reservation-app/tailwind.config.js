/** @type {import('tailwindcss').Config} */
export default {
	mode: 'jit',
	content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
	theme: {
		extend: {
			colors: {
				'light-blue': '#023E8A',
				'dark-blue': '#337bde',
				'plain-white': '#F9FAFF',
				'dark-white': '#96A0AD',
				'dark-gray': '#BDC4CD',
			},
		},
	},
	plugins: [],
};
