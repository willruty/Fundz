import "../assets/landing/global.css";
import "../assets/landing/navbar.css";
import "../assets/landing/banner.css";

export default function LandingPage() {
    return (
        <main>
            <div id="banner-bg"></div>

            <header id="navbar">
                <img src="/y_full_logo.png" alt="logo" />
                <a href="/fundz/auth">
                    <button className="default-button">
                        Sign Up
                    </button>
                </a>
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
                        <img src="/filled_phone_mockup.png" alt="phone-version" id="phone" />
                        <div id="desktop-img">
                            <img src="/finance_homepage.png" alt="desktop-version" id="desktop" />
                        </div>
                    </div>
                </section>

                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
                <br /><br /><br /><br />
            </div>
        </main>
    )
}