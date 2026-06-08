import React, {useState, useEffect} from "react";
import Globais from "./components/Globais";
import Pagina1 from "./components/Pagina_1";
import Pagina2 from "./components/Pagina_2";
import Exercicio from "./components/Exercicio";

function App() {

  const [info, setInfo] = useState(Globais.variavel2)

  const gravarInfo = () => {
    Globais.variavel2 = info
  }

  const verVariavel = () => {
    alert(Globais.variavel2)
  }

  const [pagina, setPagina] = useState(0)

  useEffect(() => {
    const url = window.location.href
    const res = url.split('?')
    setPagina(res[1])
  })

  const retornarPagina = () => {
    if (pagina == 1) {
      return <Pagina1 />
    } else if (pagina == 2) {
      return <Pagina2 />
    } else {
      return <p>Nenhuma página carregada</p>
    }
  }

  const trocarPagina = (p) => {
    if (p == 1) {
      window.open("http://localhost:3000?1", "_self")
    } else if (p == 2){
      window.open("http://localhost:3000?2", "_self")
    } else {
      window.open("http://localhost:3000?0", "_self")
    }
  }

  return (
    <div>
      <h1>Variáveis globais</h1>
      <p>{'Variável 1: ' + Globais.variavel1}</p>
      <p>{'Variável 2: ' + Globais.variavel2}</p>
      <p>{'Variável 3: ' + Globais.variavel3}</p>
      <hr />
      <input type="text" value={info} onChange={(e) => setInfo(e.target.value)} />
      <button onClick={() => gravarInfo()}>Gravar variável global</button>
      <button onClick={() => verVariavel()}>Ver variável global</button>
      <hr />
      {retornarPagina()}
      <button onClick={() => trocarPagina(0)}>Sem página</button>
      <button onClick={() => trocarPagina(1)}>Página 1</button>
      <button onClick={() => trocarPagina(2)}>Página 2</button>
      <hr />
      <h1>Exercício</h1>
      <Exercicio />
    </div>
  );
}

export default App;
