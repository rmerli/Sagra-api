/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		'src/view/**/*.templ',
		'src/view/**/*.go',
	],
	darkMode: 'class',
	theme: {
		container: {
			center: true,
		},
	},
	plugins: [],
	corePlugins: {
		preflight: true,
	}
}

