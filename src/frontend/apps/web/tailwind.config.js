import { colors } from '../../assets/colors/main_colors.js'

/** @type {import('tailwindcss').Config} */
export default {
    content: ["./src/**/*.{js,ts,jsx,tsx}"],
    theme: {
        extend: {
            colors: {
                primary: colors.primary,
                secondary: colors.secondary,
                background: colors.background,
                p_dark: colors.paragraph_dark,
                p_medium: colors.paragraph_medium,
                p_light: colors.p_light,
                detail: colors.detail
            },
            fontFamily: {
                catchland: ['catchland', 'sans-serif'],
            },
        }
    },
    plugins: []
}
