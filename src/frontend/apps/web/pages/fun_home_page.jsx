import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import Navbar from "../components/navbar"

export default function FunHomePage() {

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
            <h1>Fun Home Page</h1>
        </main>
    )
}