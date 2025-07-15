import { useState, useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import '../assets/styles/auth.css';
import '../assets/styles/global.css';

export default function AuthPage() {

    const navigate = useNavigate();
    const location = useLocation();
    const [form, setForm] = useState({ name: '', email: '', password: '' })
    const [isRegister, setIsRegister] = useState(true)
    const [error, setErr] = useState('')

    const handleChange = (e) => {
        const { name, value } = e.target;
        setForm((prev) => ({ ...prev, [name]: value }))
    }

    useEffect(() => {
        const params = new URLSearchParams(location.search);
        const mode = params.get('mode');
        if (mode === 'login') {
            setIsRegister(false);
        } else {
            setIsRegister(true);
        }
    }, [location.search]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const endpoint = isRegister ? 'user/auth/register' : 'user/auth/login'
        const payload = isRegister ? form : { email: form.email, password: form.password };

        try {
            const response = await fetch(endpoint, {
                method: "POST",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload),
            })

            const raw = await response.text();
            let data;

            try {
                data = JSON.parse(raw);
            } catch {
                data = { message: raw }
            }

            if (!response.ok) throw new Error(data.erro || data.message || raw)

            localStorage.setItem("token", data.token);
            setErr("");
            navigate("/home");

        } catch (err) {
            setErr(err.message)
        }
    }

    return (
        <div className="auth-container">

            <a href="/fundz/" className="go-back">
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="currentColor" className="bi bi-arrow-left" viewBox="0 0 16 16">
                    <path fillRule="evenodd" d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8" />
                </svg>
            </a>

            <div className="auth-box">
                <h1 className="auth-title">"Gastar, estudar, beber, repetir."</h1>
                <p>Faça seu login e comece a entender como organizar tudo isso!</p>

                <form onSubmit={handleSubmit} className="auth-input-group">

                    {isRegister && (
                        <input
                            type="text"
                            name="name"
                            placeholder="Nome"
                            value={form.name}
                            onChange={handleChange}
                            className="auth-input"
                        />
                    )}

                    <input
                        type="email"
                        name="email"
                        placeholder="Email"
                        value={form.email}
                        onChange={handleChange}
                        className="auth-input"
                    />

                    <input
                        type="password"
                        name="password"
                        placeholder="Senha"
                        value={form.password}
                        onChange={handleChange}
                        className="auth-input"
                    />

                    <button type="submit" className="auth-button">
                        {isRegister ? 'Cadastrar' : 'Entrar'}
                    </button>

                    {error && <p className="auth-error" style={{ margin: '15px', color: '#ff4444' }}>{error}</p>}

                </form>

                <div className="auth-footer">
                    <button
                        type="button"
                        onClick={() => setIsRegister(prev => !prev)}
                        className="auth-link"
                    >
                        {isRegister ? 'Já tenho conta' : 'Ainda não tenho conta'}
                    </button>

                    {!isRegister && (
                        <a href="/fundz/auth/forgot-password" className="auth-link">
                            Esqueci minha senha
                        </a>
                    )}
                </div>
            </div>
        </div>
    );
}