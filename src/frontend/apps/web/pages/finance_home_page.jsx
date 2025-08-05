import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import Navbar from "../components/navbar"
import Service_container from "../components/service_container"
import "../assets/styles/global.css"
import "../assets/styles/finance_page.css"

import {
    BalanceSummaryCard,
    CategoryExpenseCard,
    PercentageCard,
    LineGraphCard,
    TransactionsTable,
    MainGoalCard,
    GoalCard,
    CategoryBarChart,
    AccountCard,
    AddAccountButton,
    DailyAverageCard
}
    from "../components/finance_cards"

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

    const transactions = [
        {
            transaction_id: "uuid1",
            date: "2025-08-01T10:23:00",
            description: "Mensalidade da faculdade",
            amount: 850.00,
            type: "saida",
            category: "Educação",
            receiver: "Universidade ABC"
        },
        {
            transaction_id: "uuid2",
            date: "2025-08-01T15:12:00",
            description: "Bolsa de estudos",
            amount: 500.00,
            type: "entrada",
            category: "Bolsa",
            receiver: "Governo Federal"
        },
        {
            transaction_id: "uuid3",
            date: "2025-08-02T09:30:00",
            description: "Almoço RU",
            amount: 5.50,
            type: "saida",
            category: "Alimentação",
            receiver: "Restaurante Universitário"
        },
        {
            transaction_id: "uuid4",
            date: "2025-08-02T20:45:00",
            description: "Cinema com amigos",
            amount: 25.00,
            type: "saida",
            category: "Lazer",
            receiver: "CineX"
        },
        {
            transaction_id: "uuid5",
            date: "2025-08-03T18:00:00",
            description: "PIX da mãe",
            amount: 200.00,
            type: "entrada",
            category: "Família",
            receiver: "Mãe"
        },
        {
            transaction_id: "uuid6",
            date: "2025-08-03T12:45:00",
            description: "Compra de caderno",
            amount: 24.90,
            type: "saida",
            category: "Educação",
            receiver: "Papelaria Estudantil"
        },
        {
            transaction_id: "uuid7",
            date: "2025-08-04T07:50:00",
            description: "Passagem de ônibus",
            amount: 4.40,
            type: "saida",
            category: "Transporte",
            receiver: "Empresa de ônibus"
        },
        {
            transaction_id: "uuid8",
            date: "2025-08-04T20:00:00",
            description: "Freelance design",
            amount: 300.00,
            type: "entrada",
            category: "Freelance",
            receiver: "Cliente João"
        },
        {
            transaction_id: "uuid9",
            date: "2025-08-05T13:25:00",
            description: "Almoço RU",
            amount: 5.50,
            type: "saida",
            category: "Alimentação",
            receiver: "Restaurante Universitário"
        },
        {
            transaction_id: "uuid10",
            date: "2025-08-06T18:40:00",
            description: "Recarga celular",
            amount: 20.00,
            type: "saida",
            category: "Serviços",
            receiver: "Operadora"
        },
        {
            transaction_id: "uuid11",
            date: "2025-08-07T08:00:00",
            description: "Café na padaria",
            amount: 7.50,
            type: "saida",
            category: "Alimentação",
            receiver: "Padaria do Campus"
        },
        {
            transaction_id: "uuid12",
            date: "2025-08-07T21:15:00",
            description: "Bar com colegas",
            amount: 42.00,
            type: "saida",
            category: "Lazer",
            receiver: "Bar Universitário"
        },
        {
            transaction_id: "uuid13",
            date: "2025-08-08T10:00:00",
            description: "PIX do pai",
            amount: 150.00,
            type: "entrada",
            category: "Família",
            receiver: "Pai"
        },
        {
            transaction_id: "uuid14",
            date: "2025-08-09T15:20:00",
            description: "Livro de cálculo",
            amount: 90.00,
            type: "saida",
            category: "Educação",
            receiver: "Livraria Técnica"
        },
        {
            transaction_id: "uuid15",
            date: "2025-08-10T12:30:00",
            description: "Marmita congelada",
            amount: 12.00,
            type: "saida",
            category: "Alimentação",
            receiver: "Mercado"
        },
        {
            transaction_id: "uuid16",
            date: "2025-08-11T22:00:00",
            description: "Assinatura Spotify",
            amount: 9.90,
            type: "saida",
            category: "Assinaturas",
            receiver: "Spotify"
        },
        {
            transaction_id: "uuid17",
            date: "2025-08-12T17:30:00",
            description: "Venda de livro usado",
            amount: 30.00,
            type: "entrada",
            category: "Extra",
            receiver: "Colega Ana"
        },
        {
            transaction_id: "uuid18",
            date: "2025-08-13T09:00:00",
            description: "Lanche na cantina",
            amount: 6.50,
            type: "saida",
            category: "Alimentação",
            receiver: "Cantina Campus"
        },
        {
            transaction_id: "uuid19",
            date: "2025-08-14T19:30:00",
            description: "Corte de cabelo",
            amount: 25.00,
            type: "saida",
            category: "Pessoal",
            receiver: "Salão da Esquina"
        },
        {
            transaction_id: "uuid20",
            date: "2025-08-15T10:00:00",
            description: "PIX da vó",
            amount: 100.00,
            type: "entrada",
            category: "Família",
            receiver: "Vó"
        },
        {
            transaction_id: "uuid21",
            date: "2025-08-16T08:45:00",
            description: "Ônibus intermunicipal",
            amount: 15.00,
            type: "saida",
            category: "Transporte",
            receiver: "Rodoviária"
        },
        {
            transaction_id: "uuid22",
            date: "2025-08-17T20:00:00",
            description: "Freela edição de vídeo",
            amount: 250.00,
            type: "entrada",
            category: "Freelance",
            receiver: "Agência Criativa"
        },
        {
            transaction_id: "uuid23",
            date: "2025-08-18T11:00:00",
            description: "Material de aula",
            amount: 40.00,
            type: "saida",
            category: "Educação",
            receiver: "Papelaria X"
        },
        {
            transaction_id: "uuid24",
            date: "2025-08-19T18:00:00",
            description: "Pipoca e refrigerante",
            amount: 18.00,
            type: "saida",
            category: "Lazer",
            receiver: "Cinema"
        },
        {
            transaction_id: "uuid25",
            date: "2025-08-20T14:30:00",
            description: "Doação para evento",
            amount: 10.00,
            type: "saida",
            category: "Solidariedade",
            receiver: "Centro Acadêmico"
        },
        {
            transaction_id: "uuid26",
            date: "2025-08-21T12:00:00",
            description: "Venda de lanche",
            amount: 15.00,
            type: "entrada",
            category: "Extra",
            receiver: "Colega Pedro"
        },
        {
            transaction_id: "uuid27",
            date: "2025-08-22T22:00:00",
            description: "Festa da república",
            amount: 30.00,
            type: "saida",
            category: "Lazer",
            receiver: "República Alpha"
        },
        {
            transaction_id: "uuid28",
            date: "2025-08-23T10:20:00",
            description: "PIX do tio",
            amount: 80.00,
            type: "entrada",
            category: "Família",
            receiver: "Tio"
        },
        {
            transaction_id: "uuid29",
            date: "2025-08-24T16:00:00",
            description: "Compra de lápis e borracha",
            amount: 8.00,
            type: "saida",
            category: "Educação",
            receiver: "Papelaria Y"
        },
        {
            transaction_id: "uuid30",
            date: "2025-08-25T13:00:00",
            description: "Almoço RU",
            amount: 5.50,
            type: "saida",
            category: "Alimentação",
            receiver: "Restaurante Universitário"
        }
    ];

    return (
        <main>
            <Navbar />
            <Service_container titulo="Eae, qual a boa de hoje?" frase="Bem vindo ao Finance Home Page">

                <div id="content">
                    <div className="accounts-menu">
                        <h2 id="accounts-title" style={{ marginBottom: "10px" }}>Contas</h2>
                        <div className="accounts">
                            <AccountCard title={"Conta Corrente"} />
                            <AccountCard title={"Poupança"} />
                            <AccountCard title={"Cartão de crédito"} />
                            <AddAccountButton />
                        </div>
                    </div>

                    <div className="balance">
                        <BalanceSummaryCard iconSrc="64x.ico" description="Saldo atual" amount="1.234.567,89" />
                    </div>

                    <div className="category">
                        <CategoryExpenseCard iconSrc="64x.ico" description="Categoria mais frequente" amount="1.234.567,89" categoryLabel="Alimentação" />
                    </div>

                    <div className="comparation-stats">
                        <div className="percentage-stats">
                            <PercentageCard topic="Em comparação com o mês passado" percentage="12" />
                            <PercentageCard topic="Em comparação com a semana passada" percentage="125" />
                        </div>

                        <div className="daily-avg">
                            <DailyAverageCard dailyAverage={"1.234.567,89"} transactions={transactions}/>
                        </div>

                        <div className="category-chart">
                            <CategoryBarChart title={"Categorias principais"} width={900} height={450} />
                        </div>
                    </div>

                    <div className="line-graphic">
                        <LineGraphCard width={1560} height={350} title={"Seu saldo ao longo dos ultimos 30 dias"} />
                    </div>

                    <div className="transaction-table">
                        <TransactionsTable transactions={transactions} />
                    </div>

                    <h1 id="goals-title">Metas</h1>
                    <div className="goal-section">
                        <MainGoalCard title="Meta principal" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                        <div className="goal-cards">
                            <GoalCard title="Meta secundária" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                            <GoalCard title="Meta secundária" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                            <GoalCard title="Meta secundária" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                            <GoalCard title="Meta secundária" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                            <GoalCard title="Meta secundária" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                            <GoalCard title="Meta secundária" goalAmount="1.234.567,89" currentAmount="534.567,89" date="2025-08-25" />
                        </div>
                    </div>

                </div>
            </Service_container>
        </main>
    )
}