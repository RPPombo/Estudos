import React, {useState} from "react";
import Tabela from "./components/Tabela";
import FormPeso from "./components/FormPeso";
import FormAltura from "./components/FormAltura";
import CalcularIMC from "./components/CalcularIMC";

function App() {
  const [peso, setPeso] = useState(0)
  const [altura, setAltura] = useState(0)
  const [resultado , setResultado] = useState(0)

  return (
    <>
      <h1>Calculadora de IMC</h1>
      <FormPeso atribuirPeso={setPeso}/>
      <FormAltura atribuirAltura={setAltura} />
      <br/>
      <CalcularIMC peso={peso} altura={altura} setRes={setResultado}/>
      <p>Resultado: {resultado.toFixed(2)}</p>
      <br/>
      <Tabela />
    </>
  )
}

export default App;