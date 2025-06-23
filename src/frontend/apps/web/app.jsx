import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import LandingPage from "./pages/landing_page.jsx"
import AuthPage from "./pages/auth"

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/auth" element={<AuthPage />} />
      </Routes>
    </Router>
  )
}
