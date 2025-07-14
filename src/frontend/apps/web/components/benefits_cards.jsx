import { useState } from 'react';
import '../../web/assets/landing/benefits.css'; 

const benefitsData = [
    {
        title: 'Controle financeiro inteligente',
        description: 'Visualize, planeje e acompanhe suas finanças com gráficos e estatísticas simples.',
        image: '/fundz/money.jpg'
    },
    {
        title: 'Facilidade na vida universitária',
        description: 'Organize horários, trabalhos e provas para não se perder na bagunça acadêmica.',
        image: '/fundz/study.jpg'
    },
    {
        title: 'Eventos e festas organizadas',
        description: 'Participe e gerencie eventos com facilidade: de rolezinhos a festas da rep.',
        image: '/fundz/redcups2.jpg'
    }
];

export default function BenefitsSection() {
    const [activeIndex, setActiveIndex] = useState(0);

    return (
        <section id="benefits">
            
            <div
                id="benefits-img"
                style={{ backgroundImage: `url(${benefitsData[activeIndex].image})` }}
            ></div>

            <div id="benefits-info">
                <h1 id="benefits-main-title">Benefícios</h1>
                <div id="benefits-cards">
                    {benefitsData.map((card, index) => (
                        <div
                            key={index}
                            className={`benefits-card ${activeIndex === index ? 'active' : 'disabled'}`}
                            onClick={() => setActiveIndex(index)}
                        >
                            <h1 className="benefits-title">{card.title}</h1>
                            <p className="benefits-text">{card.description}</p>
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
}
