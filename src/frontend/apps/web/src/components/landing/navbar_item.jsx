export default function NavbarItem({ label, href }) {
  return (
    <a
      href={href}
      className="nav-item"
    >
      {label}
    </a>
  );
}