import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom';
import './assets/styles/global.css'
import App from './app.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter basename='/fundz'>
      <App />
    </BrowserRouter>
  </StrictMode>,
)
