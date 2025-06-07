import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import "../../../assets/styles/landing/landing_page.css";
import "../../../assets/styles/tailwind.css";
import LandingPage from './landing_page.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <LandingPage />
  </StrictMode>,
)