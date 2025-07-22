import "../assets/styles/cards.css";

// Card de 25% de largura: mostra saldo
export function BalanceSummaryCard({ iconSrc, description, amount }) {
    return (
        <div className="balance-card">
            <img src={iconSrc} id="card-img" alt="Ícone do card" />
            {/* menu de três pontos, se precisar adicionar depois */}

            <div id="description">
                <h1>R$ {amount}</h1>
                <p>{description}</p>
            </div>
        </div>
    );
}

// Card de 50%: mostra valor por categoria
export function CategoryExpenseCard({ iconSrc, description, amount, categoryLabel }) {
    return (
        <div className="category-card">
            <img src={iconSrc} id="card-img" alt="Ícone da categoria" />
            {/* menu de três pontos, se quiser adicionar */}

            <div id="description">
                <h1>R$ {amount} em<span style={{ color: "var(--secondary-color)" }}> {categoryLabel}</span></h1>
                <p>{description}</p>
            </div>
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

// Card de 25%: mostra a meta principal
export function GoalCard({ description, goalAmount, currentAmount, date }) {

    const percentage = Math.round((currentAmount / goalAmount) * 100);

    return (
        <div className="goal-card">
            <h1>Meta Principal</h1>
            <p>{description}</p>

            <div className="stats">
                <img src="" alt="" />
                <p>Valor desejado: R$ {goalAmount}</p>
            </div>

            <div className="stats">
                <img src="" alt="" />
                <p>Valor acumulado: R$ {currentAmount}</p>
            </div>

            <div className="stats">
                <img src="" alt="" />
                <p>Data: {date}</p>
            </div>

            <div className="progress-bar">
                <p>Progresso: {percentage}%</p>
            </div>
        </div>
    );
}

