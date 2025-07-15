import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

export default function HomePage() {
    const navigate = useNavigate()

    useEffect(() => {
        const token = localStorage.getItem('token')
        if (!token) {
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
        <div>
            <h1>Home Page</h1>
            <p>Bem-vindo! Você está autenticado 🎉</p>
        </div>
    )
}
