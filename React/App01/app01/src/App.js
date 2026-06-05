import React, {useState} from "react";
import Header from "./components/Header";
import Corpo from "./components/Corpo";
import Relogio from "./components/Relogio";
import Numero from "./components/Numero";
import Carros from "./components/Carros";
import Formulario from "./components/Formulario";
import Contador from "./components/Contador";
import SalvarNome from "./components/Salvar_nomes"
import "./App.css"

function App() {

  const [num, setNum] = useState(10)

  const armazenar = (chave, valor)=>{
        localStorage.setItem(chave, valor);
    }

    const remover = (chave)=>{
        localStorage.removeItem(chave)
    }

    const consultar = (chave)=>{
        alert(localStorage.getItem(chave))
    }

  return (
    <>
      <section className="container">
        <Header />
        <Numero num={num} setNum={setNum}/>
        <Relogio />
        <SalvarNome armazenar={armazenar} consultar={consultar} remover={remover}/>
        <Corpo />
        <Carros />
        <Formulario />
      </section>
    </>
  );
}

export default App;
