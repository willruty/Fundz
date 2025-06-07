import NavbarItem from "./navbar_item";
import "../../../../../assets/styles/landing/navbar.css";

export default function Navbar() {
    const menuItems = ["Início", "Benefícios", "Funcionalidades", "Identidade" , "Contato"];

    return (
        <header className="navbar">
            <nav className="navbar-menu">
                {menuItems.map((item, index) => (
                    <NavbarItem key={index} label={item} href={`#${item.toLowerCase()}`} />
                ))}
            </nav>
        </header>
    );
}
