import { useState } from 'react'
import "../../../assets/styles/App.css";
import "../../../assets/styles/landing/landing_page.css";
import Navbar from "../src/components/landing/Navbar.jsx";

function LandingPage() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Navbar />
      {/* Conteúdo da landing page aqui */}
      <section>
        <h1 className="text-4xl font-bold">Bem-vindo à Fundz!</h1>
        {/* outras seções */}
      </section>
    </>
  )
}

export default LandingPage
