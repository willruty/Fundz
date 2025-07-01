import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import LandingPage from "./pages/landing_page.jsx"
import AuthPage from "./pages/auth"
import FunHomePage from "./pages/fun_home.jsx"
import AcademicHomePage from "./pages/academic_home.jsx"
import FinanceHomePage from "./pages/finance_home.jsx"
import HomePage from "./pages/home.jsx"

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/fundz" element={<LandingPage />} />
        <Route path="/fundz/auth" element={<AuthPage />} />
        <Route path="/fundz/home" element={<HomePage />} />
        <Route path="/fundz/fun" element={<FunHomePage />} />
        <Route path="/fundz/academic" element={<AcademicHomePage />} />
        <Route path="/fundz/finance" element={<FinanceHomePage />} />
      </Routes>
    </Router>
  )
}
