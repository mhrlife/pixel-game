/** @type {import('tailwindcss').Config} */
export default {
    darkMode: ["class"],
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                tg: {
                    bg: 'var(--tg-theme-bg-color)',
                    text: 'var(--tg-theme-text-color)',
                    hint: 'var(--tg-theme-hint-color)',
                    link: 'var(--tg-theme-link-color)',
                    button: 'var(--tg-theme-button-color)',
                    buttonText: 'var(--tg-theme-button-text-color)',
                    secondaryBg: 'var(--tg-theme-secondary-bg-color)',
                    headerBg: 'var(--tg-theme-header-bg-color)',
                    bottomBarBg: 'var(--tg-theme-bottom-bar-bg-color)',
                    accentText: 'var(--tg-theme-accent-text-color)',
                    sectionBg: 'var(--tg-theme-section-bg-color)',
                    sectionHeaderText: 'var(--tg-theme-section-header-text-color)',
                    sectionSeparator: 'var(--tg-theme-section-separator-color)',
                    subtitleText: 'var(--tg-theme-subtitle-text-color)',
                    destructiveText: 'var(--tg-theme-destructive-text-color)',
                },
                // Updated to pull from Telegram's CSS variables:
                background: "var(--tg-theme-bg-color)",
                foreground: "var(--tg-theme-text-color)",
                card: {
                    DEFAULT: "var(--tg-theme-section-bg-color)",
                    foreground: "var(--tg-theme-section-header-text-color)",
                },
                popover: {
                    DEFAULT: "var(--tg-theme-secondary-bg-color)",
                    foreground: "var(--tg-theme-text-color)",
                },
                primary: {
                    DEFAULT: "var(--tg-theme-button-color)",
                    foreground: "var(--tg-theme-button-text-color)",
                },
                secondary: {
                    DEFAULT: "var(--tg-theme-link-color)",
                    foreground: "var(--tg-theme-text-color)",
                },
                muted: {
                    DEFAULT: "var(--tg-theme-hint-color)",
                    foreground: "var(--tg-theme-text-color)",
                },
                accent: {
                    DEFAULT: "var(--tg-theme-accent-text-color)",
                    foreground: "var(--tg-theme-bg-color)", // or any suitable contrast
                },
                destructive: {
                    DEFAULT: "var(--tg-theme-destructive-text-color)",
                    foreground: "var(--tg-theme-button-text-color)", // or "#fff"
                },
                border: "var(--tg-theme-section-separator-color)",
                input: "var(--tg-theme-bg-color)", // or another variable you prefer
                ring: "var(--tg-theme-link-color)",

                // Example: Chart colors could also reference Telegram colors, but
                // you may not have direct matches. Adjust to taste:
                chart: {
                    "1": "var(--tg-theme-button-color)",
                    "2": "var(--tg-theme-link-color)",
                    "3": "var(--tg-theme-accent-text-color)",
                    "4": "var(--tg-theme-destructive-text-color)",
                    "5": "var(--tg-theme-hint-color)",
                },

                sidebar: {
                    DEFAULT: "var(--tg-theme-header-bg-color)",
                    foreground: "var(--tg-theme-text-color)",
                    primary: "var(--tg-theme-button-color)",
                    "primary-foreground": "var(--tg-theme-button-text-color)",
                    accent: "var(--tg-theme-link-color)",
                    "accent-foreground": "var(--tg-theme-text-color)",
                    border: "var(--tg-theme-section-separator-color)",
                    ring: "var(--tg-theme-link-color)",
                },
            },
            borderRadius: {
                lg: 'var(--radius)',
                md: 'calc(var(--radius) - 2px)',
                sm: 'calc(var(--radius) - 4px)'
            },
            keyframes: {
                'accordion-down': {
                    from: {
                        height: '0'
                    },
                    to: {
                        height: 'var(--radix-accordion-content-height)'
                    }
                },
                'accordion-up': {
                    from: {
                        height: 'var(--radix-accordion-content-height)'
                    },
                    to: {
                        height: '0'
                    }
                }
            },
            animation: {
                'accordion-down': 'accordion-down 0.2s ease-out',
                'accordion-up': 'accordion-up 0.2s ease-out'
            }
        }
    },
    plugins: [require("tailwindcss-animate")],
}

