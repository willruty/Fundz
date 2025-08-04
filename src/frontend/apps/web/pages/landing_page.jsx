import "../assets/styles/global.css";
import "../assets/landing/navbar.css";
import "../assets/landing/banner.css";
import "../assets/landing/benefits.css";
import "../assets/landing/security.css";
import "../assets/landing/stats.css";
import "../assets/landing/features.css";
import "../assets/landing/cta.css";
import "../assets/landing/final_call.css";
import "../assets/landing/footer.css";
import BenefitsSection from "../components/benefits_cards";

export default function LandingPage() {
    console.log("Landing Page");
    return (
        <main style={{ background: "#FEFDF9", overflowX: "hidden" }}>
            <div id="banner-bg"></div>

            <header id="navbar">
                <img src="/fundz/y_full_logo.png" alt="logo" />
                <div className="navbar-buttons">
                    <a href="/fundz/auth?mode=login">
                        <button className="default-button" id="login">
                            Login
                        </button>
                    </a>
                    <a href="/fundz/auth">
                        <button className="default-button" id="signup">
                            Cadastre-se
                        </button>
                    </a>
                </div>
            </header>

            <div id="container">

                <section id="banner">
                    <div id="banner-text">

                        <h1>Os melhores anos da sua vida merecem equilíbrio.</h1>
                        <p>
                            O Fundz te ajuda curtir o rolê, controlar sua grana e não perder a cabeça com a faculdade, tudo isso em um só lugar...
                        </p>
                        <a href="/fundz/auth">
                            <button className="sign-up default-button">
                                Acesso antecipado
                            </button>
                        </a>
                    </div>

                    <div id="banner-img">
                        <img src="/fundz/filled_phone_mockup.png" alt="phone-version" id="phone" />
                        <div id="desktop-img">
                            <img src="/fundz/finance_homepage.png" alt="desktop-version" id="desktop" />
                        </div>
                    </div>
                </section>

                <BenefitsSection />

                <section id="security">

                    <div id="security-bg"></div>

                    <div id="security-content">

                        <div id="security-info">
                            <h1>Confiança de cofre,<br /> usabilidade de boteco.</h1>
                            <div id="security-cards">
                                <div className="security-card">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="35" height="35" fill="#FFD100" className="bi bi-emoji-angry" viewBox="0 0 16 16">
                                        <path d="M7.293 1.5a1 1 0 0 1 1.414 0L11 3.793V2.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v3.293l2.354 2.353a.5.5 0 0 1-.708.708L8 2.207l-5 5V13.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 1 0 1h-4A1.5 1.5 0 0 1 2 13.5V8.207l-.646.647a.5.5 0 1 1-.708-.708z" />
                                        <path d="M10 13a1 1 0 0 1 1-1v-1a2 2 0 0 1 4 0v1a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zm3-3a1 1 0 0 0-1 1v1h2v-1a1 1 0 0 0-1-1" />
                                    </svg>
                                    <h3 className="security-title">Seus dados, <br />sua segurança</h3>
                                    <p className="security-text">Tudo criptografado de ponta a ponta, com acesso apenas seu. Segurança real pra você usar sem preocupação.</p>
                                </div>
                                <div className="security-card">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="35" height="35" fill="#FFD100" className="bi bi-crosshair" viewBox="0 0 16 16">
                                        <path d="M8.5.5a.5.5 0 0 0-1 0v.518A7 7 0 0 0 1.018 7.5H.5a.5.5 0 0 0 0 1h.518A7 7 0 0 0 7.5 14.982v.518a.5.5 0 0 0 1 0v-.518A7 7 0 0 0 14.982 8.5h.518a.5.5 0 0 0 0-1h-.518A7 7 0 0 0 8.5 1.018zm-6.48 7A6 6 0 0 1 7.5 2.02v.48a.5.5 0 0 0 1 0v-.48a6 6 0 0 1 5.48 5.48h-.48a.5.5 0 0 0 0 1h.48a6 6 0 0 1-5.48 5.48v-.48a.5.5 0 0 0-1 0v.48A6 6 0 0 1 2.02 8.5h.48a.5.5 0 0 0 0-1zM8 10a2 2 0 1 0 0-4 2 2 0 0 0 0 4" />
                                    </svg>
                                    <h3 className="security-title">Simples até de ressaca, <br /> porque menos é mais.</h3>
                                    <p className="security-text">Interface leve, fluida e fácil de entender até nos dias mais puxados. Rápido, direto e sem stress.</p>
                                </div>
                            </div>
                        </div>

                        <img src="/fundz/filled_phone_mockup.png" alt="phone-version" id="phone" />
                    </div>
                </section>

                <section id="stats">
                    <h1>Por que o Fundz faz sentido?</h1>

                    <div id="stats-cards">
                        <div className="stats-card">
                            <p className="stats-number">74%</p>
                            <p className="stats-text">dizem que têm, dificuldade em controlar gastos mensais.</p>
                            <p className="stats-font">"DataFolha/USP 2022"</p>
                        </div>
                        <div className="stats-card">
                            <p className="stats-number">6 de 10</p>
                            <p className="stats-text">universitários esquecem prazos de provas ou tarefas pelo menos 1 vez por semestre.</p>
                            <p className="stats-font">"Educa Insights/Universia Brasil 2023"</p>
                        </div>
                        <div className="stats-card">
                            <p className="stats-number">82%</p>
                            <p className="stats-text">participam de eventos sociais semanais e perdem dinheiro por falta de planejamento.</p>
                            <p className="stats-font">"UNE - Comportamento universitário 2023"</p>
                        </div>
                        <div className="stats-card">
                            <p className="stats-number">10 min</p>
                            <p className="stats-text">é o que você precisa pra se organizar com o fundz.</p>
                            <p className="stats-font">"confia, por que eu ia mentir?"</p>
                        </div>
                    </div>

                </section>

                <section id="features">
                    <div className="feature">
                        <img src="/fundz/finance_homepage.png" alt="" className="graficos features-image" />
                        <p className="graficos features-text">
                            Veja seus gastos virando gráficos tão claros que até
                            aquele seu amigo de humanas entenderia. Saiba onde foi parar seu dinheiro:
                            festas, comida, xerox ou aquele Uber que nem precisava.
                        </p>
                    </div>
                    <div className="feature">
                        <img src="/fundz/finance_homepage.png" alt="" className="semestre features-image" />
                        <p className="semestre features-text">
                            Crie alertas pra provas, entregas e trabalhos em grupo (mesmo que você vá fazer tudo sozinho de novo).
                            Seu semestre, visualizado de um jeito que até o professor aprovaria.
                        </p>
                    </div>
                    <div className="feature">
                        <img src="/fundz/finance_homepage.png" alt="" className="roles features-image" />
                        <p className="roles features-text">
                            Marque os rolês, conte as partidas de truco e divida os gastos com precisão cirúrgica.
                            Fundz ajuda você a manter a amizade mesmo depois da bebedeira.
                        </p>
                    </div>
                    <div className="feature">
                        <img src="/fundz/finance_homepage.png" alt="" className="portabilidade features-image" />
                        <p className="portabilidade features-text">
                            Use o Fundz no celular no meio da aula, no PC
                            enquanto finge que estuda ou no bar entre uma rodada e outra. Sincronizado em tudo, porque universitário vive em movimento.
                        </p>
                    </div>

                </section>

                <section id="final-call">
                    <img src="/fundz/y_full_logo.png" alt="" />
                    <div>
                        <p>Você vive a rotina. A gente deixa ela mais fácil. Vem pro Fundz.</p>
                        <a href="/fundz/auth"><button className="default-button">Bora organizar isso!</button></a>
                    </div>
                </section>

                <section id="cta">

                    <div id="simples" className="cta-card">
                        <div id="simples-left-div">
                            <div id="simples-text">
                                <h1>Simples.</h1>
                                <p>Sem complicação: <br />e-mail, senha e pronto.</p>
                            </div>
                            <a href="/fundz/auth"><button className="default-button">Organizar minha vida</button></a>
                        </div>
                        <img src="/fundz/filled_phone_mockup.png" alt="" />
                    </div>

                    <div id="rapido" className="cta-card">
                        <img src="/fundz/filled_phone_mockup.png" alt="" />
                        <div id="rapido-bottom-div">
                            <h1>Rápido.</h1>
                            <p>Escolha o que controlar: <br />grana, provas e festas...</p>
                        </div>
                    </div>

                    <div id="pratico" className="cta-card">
                        <div id="pratico-top-div">
                            <h1>Prático.</h1>
                            <p>Visualize tudo em gráficos, calendários e notificações, direto ao ponto.</p>
                        </div>
                        <img src="/fundz/filled_phone_mockup.png" alt="" />
                    </div>
                </section>

                <footer>
                    <img src="/fundz/w_full_logo.png" alt="" />
                    <h3>"Feito por quem também tá devendo nota."</h3>
                    <p>© 2025 Fundz — Feito por William R. Macedo</p>
                </footer>
            </div>
        </main>
    )
}