import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {init} from "@telegram-apps/sdk-react";

try {
    window.Telegram.WebApp.ready();
    window.Telegram.WebApp.expand();
    window.Telegram.WebApp.enableClosingConfirmation();
    window.Telegram.WebApp.requestFullscreen();
} catch (e) {
    console.log("test environment")
}

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <App/>
    </StrictMode>,
)
