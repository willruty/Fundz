import { Routes, Route, Navigate, Outlet } from "react-router-dom";

import LandingPage from "./pages/landing_page.jsx";
import AuthPage from "./pages/auth_page.jsx";
import FunHomePage from "./pages/fun_home_page.jsx";
import AcademicHomePage from "./pages/academic_home_page.jsx";
import FinanceHomePage from "./pages/finance_home_page.jsx";
import ProfilePage from "./pages/profile_page.jsx";
import HomePage from "./pages/home_page.jsx";

function PrivateRoute() {
  const token = localStorage.getItem("token");
  return token ? <Outlet /> : <Navigate to="/fundz/auth?mode=login" />; 
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/auth" element={<AuthPage />} />

      <Route element={<PrivateRoute />}>
        <Route path="/home" element={<HomePage />} />
        <Route path="/fun" element={<FunHomePage />} />
        <Route path="/academic" element={<AcademicHomePage />} />
        <Route path="/finance" element={<FinanceHomePage />} />
        <Route path="/profile" element={<ProfilePage />} />
      </Route>

    </Routes>
  );
}
