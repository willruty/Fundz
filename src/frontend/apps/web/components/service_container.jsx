import "../assets/styles/service_container.css"

export default function Service_container({ titulo, frase, children }) {

    return (
        <div id="service_container">
            <div id="title">
                <h1>{titulo}</h1>

                <div id="service_text">
                    <p>{frase}</p>
                </div>
            </div>

            <div id="service_info">
                {children}
            </div>
        </div>
    );
}