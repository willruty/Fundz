import '../assets/styles/auth.css';
import '../assets/styles/global.css';

export default function AuthPage() {
    return (
        <div className="auth-container">

            <a href="/fundz" className='go-back'>
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="currentColor" class="bi bi-arrow-left" viewBox="0 0 16 16">
                    <path fill-rule="evenodd" d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8" />
                </svg>
            </a>

            <div className="auth-box">
                <h1 className="auth-title">"Gastar, estudar, beber, repetir."</h1>
                <p>Faça seu login e comece a entender como organizar tudo isso!</p>

                <div className="auth-input-group">
                    <input type="text" placeholder="Nome" className="auth-input" />
                    <input type="email" placeholder="Email" className="auth-input" />
                    <input type="password" placeholder="Senha" className="auth-input" />
                </div>

                <button className="auth-button">Entrar</button>

                <div className="auth-footer">
                    <a href="/fundz/auth/register" className="auth-link">Ainda não tenho conta</a>
                    <a href="/fundz/auth/forgot-password" className="auth-link">Esqueci minha senha</a>
                </div>
            </div>
        </div>
    );
}
