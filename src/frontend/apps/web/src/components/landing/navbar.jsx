import NavbarItem from "./navbar_item";
import "../../../../../assets/styles/landing/navbar.css";

export default function Navbar() {
    const menuItems = ["Início", "Benefícios", "Funcionalidades", "Identidade", "Contato"];

    return (
        <header className="navbar">

            <div id="navbar-logo">
                <img src="../../../../../assets/styles/img/logo.png" alt="Logo" />
                <span>Fundz</span>
            </div>

            <nav>
                {menuItems.map((item, index) => (
                    <NavbarItem key={index} label={item} href={`#${item.toLowerCase()}`} />
                ))}
            </nav>

            <div id="navbar-cta">
                <a href="#login" className="btn-shine"><i class="bi bi-box-arrow-in-right"></i> Acesso antecipado</a>
            </div>
            
        </header>
    );
}
