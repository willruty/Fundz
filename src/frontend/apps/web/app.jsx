import { Routes, Route } from "react-router-dom";

import LandingPage from "./pages/landing_page.jsx";
import AuthPage from "./pages/auth_page.jsx";
import FunHomePage from "./pages/fun_home.jsx";
import AcademicHomePage from "./pages/academic_home.jsx";
import FinanceHomePage from "./pages/finance_home.jsx";
import HomePage from "./pages/home.jsx";

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/auth" element={<AuthPage />} />
      <Route path="/home" element={<HomePage />} />
      <Route path="/fun" element={<FunHomePage />} />
      <Route path="/academic" element={<AcademicHomePage />} />
      <Route path="/finance" element={<FinanceHomePage />} />
    </Routes>
  );
}
