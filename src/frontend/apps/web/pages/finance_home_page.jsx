import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import Navbar from "../components/navbar"
import Service_container from "../components/service_container"
import "../assets/styles/global.css"
import "../assets/styles/finance_page.css"

import { BalanceSummaryCard, CategoryExpenseCard, LineGraphCard, PercentageCard }
    from "../components/cards"

export default function FinanceHomePage() {

    const navigate = useNavigate()

    useEffect(() => {
        const token = localStorage.getItem('token')
        if (!token) {
            console.log("Token nao encontrado na home page, voltando para o login")
            navigate("/login")
            return
        }

        fetch("http://localhost:8080/fundz/user/auth/validate", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
            .then(async (res) => {
                if (!res.ok) {
                    throw new Error("Token inválido ou expirado")
                }

                const contentType = res.headers.get("content-type")
                if (contentType && contentType.includes("application/json")) {
                    const data = await res.json()
                    console.log("Token válido:", data)
                } else {
                    throw new Error("Resposta não é JSON")
                }
            })
            .catch((err) => {
                console.error("Erro na validação do token:", err.message)
                navigate("/auth")
            })
    }, [])

    return (
        <main>
            <Navbar />
            <Service_container titulo="Eae, qual a boa de hoje?" frase="Bem vindo ao Finance Home Page">

                <div id="content">
                    <BalanceSummaryCard iconSrc="64x.ico" description="Saldo atual" amount="1.234.567,89" />
                    <CategoryExpenseCard iconSrc="64x.ico" description="Categoria mais frequente" amount="1.234.567,89" categoryLabel="Alimentação" />
                    <PercentageCard topic="Em comparação com o mês passado" percentage="12" />
                    <PercentageCard topic="Em comparação com a semana p[assada" percentage="125" />
                    <LineGraphCard />
                </div>
                
            </Service_container>
        </main>
    )
}