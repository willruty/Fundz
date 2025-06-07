export default function NavbarItem({ label, href }) {
  return (
    <a
      href={href}
      className="text-gray-700 font-medium hover:text-blue-600 transition-colors duration-200"
    >
      {label}
    </a>
  );
}