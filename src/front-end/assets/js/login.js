const formContainer = document.getElementById("formContainer");
const form = document.getElementById("userForm");
const btnLogin = document.getElementById("showLogin");
const btnSignup = document.getElementById("showSignup");
const submitBtn = document.getElementById("submitBtn");

// Modo padrão
let modo = "login";

// Alternar modo com clique nos botões
btnLogin.addEventListener("click", () => {
    formContainer.classList.remove("signup");
    submitBtn.textContent = "Entrar";
    modo = "login";
    btnLogin.style.display = "none";
    btnSignup.style.display = "inline";
});

btnSignup.addEventListener("click", () => {
    formContainer.classList.add("signup");
    submitBtn.textContent = "Cadastrar-se";
    modo = "signup";
    btnSignup.style.display = "none";
    btnLogin.style.display = "inline";
});

// Verifica se a URL contém ?modo=signup
const urlParams = new URLSearchParams(window.location.search);
const urlModo = urlParams.get("modo");

if (urlModo === "signup") {
    formContainer.classList.add("signup");
    submitBtn.textContent = "Cadastrar-se";
    modo = "signup";
    btnSignup.style.display = "none";
    btnLogin.style.display = "inline";
} else {
    btnLogin.style.display = "none";
    btnSignup.style.display = "inline";
}

// Função para mostrar mensagens na tela
function mostrarMensagem(texto, sucesso = true) {
    const msgDiv = document.getElementById("mensagem");
    msgDiv.textContent = texto;
    msgDiv.style.display = "block";
    msgDiv.style.backgroundColor = sucesso ? "#D4EDDA" : "#F8D7DA";
    msgDiv.style.color = sucesso ? "#155724" : "#721C24";
    msgDiv.style.border = sucesso ? "1px solid #C3E6CB" : "1px solid #F5C6CB";

    setTimeout(() => {
        msgDiv.style.display = "none";
    }, 4000);
}

// Envio do formulário
form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const email = document.getElementById("email").value;
    const senha = document.getElementById("senha").value;
    const nome = document.getElementById("nome")?.value;

    const payload = modo === "signup"
        ? { user_name: nome, user_email: email, user_password: senha }
        : { user_email: email, user_password: senha };

    const url = modo === "signup"
        ? "/Fundz/register"
        : "/Fundz/login";

    try {
        const res = await fetch(url, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload),
        });

        const data = await res.json(); // <-- movido para dentro do escopo de try

        if (res.ok) {
            if (modo === "signup") {
                mostrarMensagem("Cadastro feito com sucesso!");
                form.reset();
            } else {
                mostrarMensagem("Login realizado com sucesso");
                localStorage.setItem("token", data.token);
                console.log(localStorage.getItem("token"));


                // Redireciona após 2 segundos
                setTimeout(() => {
                    window.location.href = "/Fundz/dashboard";
                }, 2000);

                if (data.token) {
                    localStorage.setItem("token", data.token);
                    window.location.href = "/Fundz/dashboard";
                }
            }
        } else {
            mostrarMensagem(data.erro || "Email ou senha inválidos.", false);
        }
    } catch (err) {
        console.error("Erro:", err);
        mostrarMensagem("Erro ao conectar com servidor", false);
    }
});
