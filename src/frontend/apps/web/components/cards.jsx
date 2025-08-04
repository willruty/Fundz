import "../assets/styles/cards.css";
import React, { useState } from "react";
import { FiPlus, FiFilter, FiMaximize } from "react-icons/fi";
import { GrMoreVertical, GrBarChart } from "react-icons/gr";
import { TbTargetArrow } from "react-icons/tb";
import { IoWallet } from "react-icons/io5";
import { CiCalendarDate } from "react-icons/ci";
import { BiSolidCategory } from "react-icons/bi";
import { HiChartSquareBar } from "react-icons/hi";

import {
    LineChart,
    Line,
    ReferenceLine,
    BarChart,
    Bar,
    Rectangle,
    ResponsiveContainer,
    XAxis,
    YAxis,
    Legend,
    Tooltip,
} from 'recharts';

// Card de 25% de largura: mostra saldo
export function BalanceSummaryCard({ iconSrc, description, amount }) {
    return (
        <div className="balance-card">
            <IoWallet style={{ color: "var(--secondary-color)", width: "30px", height: "30px" }} />

            <div id="description">
                <h1>R$ {amount}</h1>
                <p>{description}</p>
            </div>
        </div>
    );
}

// Card de 50%: mostra valor por categoria
export function CategoryExpenseCard({ description, amount, categoryLabel }) {
    return (
        <div className="category-card">
            <BiSolidCategory style={{ color: "var(--secondary-color)", width: "30px", height: "30px" }} />
            <div id="description">
                <h1>R$ {amount} em<span style={{ color: "var(--secondary-color)" }}> {categoryLabel}</span></h1>
                <p>{description}</p>
            </div>
        </div>
    );
}

export function CategoryBarChart({ width, height, title }) {

    const categoryData = [
        { category: "Alimentação", amount: 1200 },
        { category: "Transporte", amount: 800 },
        { category: "Educação", amount: 500 },
        { category: "Lazer", amount: 300 },
        { category: "Saúde", amount: 650 },
    ];

    return (
        <div className="category-graphic-card ">
            <h1>{title}</h1>
            <ResponsiveContainer width="100%" height="100%">
                <BarChart
                    width={width}
                    height={height}
                    data={categoryData}
                    margin={{
                        top: 30,
                        right: 30,
                        left: 10,
                        bottom: 40,
                    }}
                >
                    <XAxis dataKey="category" />
                    <YAxis />
                    <Legend />
                    <Tooltip contentStyle={{ backgroundColor: "var(--white)", border: "none", borderRadius: "15px" }} />
                    <Bar dataKey="amount" name="quantia" fill="var(--primary-color)" activeBar={<Rectangle fill="var(--secondary-color)" />} />
                </BarChart>
            </ResponsiveContainer>
        </div>
    );
}

// Card de 10%: mostra uma porcentagem comparada com algum tópico
export function PercentageCard({ topic, percentage }) {
    return (
        <div className="percentage-card">
            <p>{topic}</p>

            <div id="percentage">
                <h1>{percentage}%</h1>
                {/* arrow up or down */}
            </div>
        </div>
    );
}

const data = [
    { day: "01/08", saldo: 150 },
    { day: "02/08", saldo: 100 },
    { day: "03/08", saldo: 60 },
    { day: "04/08", saldo: 20 },
    { day: "05/08", saldo: 1320 },
    { day: "06/08", saldo: 1130 },
    { day: "07/08", saldo: 1030 },
    { day: "08/08", saldo: 960 },
    { day: "09/08", saldo: 880 },
    { day: "10/08", saldo: 850 },
    { day: "11/08", saldo: 760 },
    { day: "12/08", saldo: 700 },
    { day: "13/08", saldo: 690 },
    { day: "14/08", saldo: 620 },
    { day: "15/08", saldo: 590 },
    { day: "16/08", saldo: 540 },
    { day: "17/08", saldo: 740 },
    { day: "18/08", saldo: 690 },
    { day: "19/08", saldo: 650 },
    { day: "20/08", saldo: 620 },
    { day: "21/08", saldo: 720 },
    { day: "22/08", saldo: 640 },
    { day: "23/08", saldo: 580 },
    { day: "24/08", saldo: 500 },
    { day: "25/08", saldo: 460 },
    { day: "26/08", saldo: 410 },
    { day: "27/08", saldo: 710 },
    { day: "28/08", saldo: 640 },
    { day: "29/08", saldo: 600 },
    { day: "30/08", saldo: 550 },
    { day: "31/08", saldo: 500 },
];

export function LineGraphCard({ width, height, title }) {
    return (
        <div className="line-graph-card">
            <h1>{title}</h1>
            <LineChart width={width} height={height} data={data}
                margin={{ top: 50, right: 30, left: 20, bottom: 5 }}>
                <XAxis dataKey="day" stroke="#fff" interval={3} />
                <YAxis stroke="#fff" />
                <ReferenceLine y={400} stroke="#f9fafb" strokeDasharray="3 3" />
                <Tooltip contentStyle={{ backgroundColor: "#1f2937", border: "none", color: "#ffffff", borderRadius: "15px" }} />
                <Legend />
                <Line type="linear" dataKey="saldo" stroke="#FFD100" strokeWidth={2} />
            </LineChart>
        </div>
    );
}

export function TransactionsTable({ transactions }) {
    const [expanded, setExpanded] = useState(false);

    // Lógica para mostrar só as últimas 10 ou expandir para os últimos 30 dias
    const visibleTransactions = expanded
        ? transactions
        : transactions.slice(0, 10);

    return (
        <div className="transactions-table">
            <div className="table-bar">
                <div className="left">
                    <HiChartSquareBar style={{ color: "var(--secondary-color)", width: "30px", height: "30px" }} />
                    <span className="title">Transações recentes</span>
                </div>
                <div className="right">
                    <FiPlus className="action-icon" />
                    <FiFilter className="action-icon" />
                    <GrBarChart className="action-icon" />
                    <FiMaximize className="action-icon" />
                    <GrMoreVertical className="action-icon" />
                </div>
            </div>

            <table>
                <thead>
                    <tr>
                        <th>Data</th>
                        <th>Hora</th>
                        <th>Valor</th>
                        <th>Categoria</th>
                        <th>Tipo</th>
                        <th>Descrição</th>
                    </tr>
                </thead>
                <tbody>
                    {visibleTransactions.map((tx) => {
                        const date = new Date(tx.date);
                        return (
                            <tr key={tx.transaction_id}>
                                <td>{date.toLocaleDateString("pt-BR")}</td>
                                <td>{date.toLocaleTimeString("pt-BR", { hour: "2-digit", minute: "2-digit" })}</td>
                                <td className={tx.type === "entrada" ? "entrada" : "saida"}>
                                    R$ {tx.amount.toFixed(2)}
                                </td>
                                <td>{tx.category || "-"}</td>
                                <td>{tx.type === "entrada" ? "Ganho" : "Gasto"}</td>
                                <td className="desc">{tx.description}</td>
                            </tr>
                        );
                    })}
                </tbody>
            </table>

            <button className="expand-button" onClick={() => setExpanded(!expanded)}>
                {expanded ? "Mostrar menos" : "Visualizar todas"}
            </button>
        </div>
    );
}

export function MainGoalCard({ description, goalAmount, currentAmount, date }) {
    const goal = parseFloat(goalAmount.replace(/\./g, '').replace(',', '.'));
    const current = parseFloat(currentAmount.replace(/\./g, '').replace(',', '.'));
    const percentage = Math.round((current / goal) * 100);

    return (
        <div className="main-goal-card"><h1>Meta Principal</h1>
            <p>{description}</p>
            <div className="stats"><TbTargetArrow style={{ color: "var(--secondary-color)", width: "25px", height: "25px" }} />
                <p>Valor desejado: R$ {goalAmount}</p>
            </div>
            <div className="stats"><IoWallet style={{ color: "var(--secondary-color)", width: "25px", height: "25px" }} />
                <p>Valor acumulado: R$ {currentAmount}</p>
            </div>
            <div className="stats"><CiCalendarDate style={{ color: "var(--secondary-color)", width: "25px", height: "25px" }} />
                <p>Data: {date}</p>
            </div>

            <div className="progress-bar">
                <p style={{ color: "var(--secondary-color)" }}>Progresso: {percentage}%</p>
                <div className="bar-container">
                    <div className="bar-fill" style={{ width: `${percentage}%` }}></div>
                </div>
            </div>
        </div>
    );
}

export function GoalCard({ description, goalAmount, date }) {

    return (
        <div className="goal-card">
            <div className="card-stats"><TbTargetArrow style={{ color: "var(--secondary-color)", width: "25px", height: "25px" }} />
            <h1>Meta</h1>
            <p>{description}</p>
                <p>Valor desejado: R$ {goalAmount}</p>
                <p>Data: {date}</p>
            </div>
        </div>
    );
}