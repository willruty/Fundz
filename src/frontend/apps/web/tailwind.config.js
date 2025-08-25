/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./app.jsx",
        "./main.jsx",
        "./src/**/*.{js,ts,jsx,tsx}",
        "./frontend/apps/web/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                'primary-color': 'var(--primary-color)',
                'secondary-color': 'var(--secondary-color)',
                'transparent-blue': 'var(--transparent-blue)',
                'transparent-blue-hover': 'var(--transparent-blue-hover)',
                'white': 'var(--white)',
            },
            borderRadius: {
                'default': 'var(--border-radius)',
            },
            fontSize: {
                'regular': 'var(--regular-font-size)',
                'l': 'var(--l-font-size)',
                'xl': 'var(--xl-font-size)',
                'xxl': 'var(--xxl-font-size)',
            },
        },
    },
    plugins: [],
}