import { useState } from 'react'

import "../../../assets/styles/App.css";
import "../../../assets/styles/landing/global.css";
import "../../../assets/styles/landing/banner.css";
import "../../../assets/styles/landing/about.css";
import "../../../assets/styles/landing/benefits.css";
import "../../../assets/styles/landing/functions.css";

import notebookBanner from "../../../assets/img/notebook1.png";
import tela1 from "../../../assets/img/tela1.png";
import tela2 from "../../../assets/img/tela2.png";
import tela3 from "../../../assets/img/tela3.png";
import tela4 from "../../../assets/img/tela4.png";

import Navbar from "../src/components/landing/Navbar.jsx";

function LandingPage() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Navbar />
      <div className="container">

        <div id="banner">
          <h1>Assuma o controle da sua <br />vida com inteligência.</h1>
          <a href="#about"><p><i class="bi bi-arrow-down"></i>Saber mais</p></a>
          <img src={notebookBanner} alt="Banner" />
        </div>

        <div id="about">
          <h2>O que é o Fundz?</h2>
          <p>
            O Fundz é uma plataforma feita por e para universitários. Organize suas finanças, acompanhe a faculdade e
            ainda arrume tempo (e grana) pro rolê — tudo no mesmo lugar, porque a vida já é difícil o bastante.
          </p>
          <ul>
            <li className='about-card'><h3><i class="bi bi-cash" style={{ color: 'var(--detail-finance)' }}></i> Finanças sem complicação</h3>
              <p>
                Chega de planilhas chatas! No Fundz, você organiza seus gastos, define metas e acompanha sua grana de um jeito simples e visual.
                A vida já é corrida — suas finanças não precisam ser.</p>
            </li>
            <li className='about-card'><h3><i class="bi bi-book" style={{ color: 'var(--detail-academic)' }}></i> Vida acadêmica em dia</h3>
              <p>
                Fique no controle das suas matérias, prazos e projetos com uma visão clara da sua rotina universitária.
                O Fundz te ajuda a estudar sem surtar (ou quase isso).
              </p>
            </li>
            <li className='about-card'><h3><i class="bi bi-cup-straw" style={{ color: 'var(--detail-fun)' }}></i> Lazer com responsabilidade</h3>
              <p>
                Estudar e economizar não significa abrir mão da diversão. Aqui você também planeja rolês, viagens e tudo que
                faz a vida valer a pena — com o pé no chão e a cabeça no mundo.
              </p>
            </li>
          </ul>
        </div>

        <section id="benefits">

          <h1 className='benefits-title' style={{ gridColumn: "1 / span 2", gridRow: 1 }}>BENEFÍCIOS<hr /></h1>

          <div className="benefit-card-text" style={{ gridColumn: 1, gridRow: 2 }}>

            <p className='benefits-topic' style={{ color: 'var(--detail-finance)' }}>Finanças</p>

            <h3 className='benefit-card-title'>Pela primeira vez, você vai saber pra onde seu dinheiro foi.</h3>
            <p className='benefit-card-description'>
              Chega de adivinhar onde a grana sumiu. Acompanhe seus gastos com clareza, defina metas e veja seu saldo subir sem dor de cabeça.
            </p>

          </div>

          <img src={tela1} alt="tela1" className='benefit-card-image' style={{ gridColumn: 2, gridRow: 2 }} />

          <div className="benefit-card-text" style={{ gridColumn: 2, gridRow: 3 }}>

            <p className='benefits-topic' style={{ color: 'var(--detail-academic)' }}>Estudos</p>

            <h3 className='benefit-card-title'>Chega de perder prazos ou esquecer matéria.</h3>
            <p className='benefit-card-description'>
              Visualize todas suas aulas, entregas e provas em um lugar só. A vida acadêmica fica mais leve quando você enxerga o todo.
            </p>

          </div>

          <img src={tela4} alt="tela4" className='benefit-card-image' style={{ gridColumn: 1, gridRow: 3 }} />

          <div className="benefit-card-text" style={{ gridColumn: 1, gridRow: 4 }}>

            <p className='benefits-topic' style={{ color: 'var(--detail-fun)' }}>Lazer</p>

            <h3 className='benefit-card-title'>Planeje o dia da cervejada para não chegar de ressaca na prova.</h3>
            <p className='benefit-card-description'>
              Planeje viagens, festas e encontros sem culpa. O Fundz ajuda você a curtir sem ferrar seu orçamento.
            </p>

          </div>

          <img src={tela3} alt="tela3" className='benefit-card-image' style={{ gridColumn: 2, gridRow: 4 }} />


        </section>

        <section id="functions">
          <h1>FUNCIONALIDADES</h1>

        </section>

        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
        <br /><br /><br /><br /><br />
      </div>
    </>
  )
}

export default LandingPage
