import "../apps/web/assets/landing/navbar.css";

export default function LandingPage() {
    return (
        <div className="w-full flex justify-around items-center p-5">
            <img src="/y_full_logo.png" alt="logo" className="h-12 w-auto" />
            <a href="/fundz/auth">
                <button className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
                    Sign Up
                </button>
            </a>
        </div>
    )
}