import { getByDisplayValue } from "@testing-library/dom";
import React, {useState} from "react";

function App() {
  const [valorTela, setValorTela] = useState("")
  const [resultado, setResultado] = useState(0)
  const [acumulador, setAcumulador] = useState(0)
  const [operacao, setOperacao] = useState(false)

  const Btn = (label, onClick) => {
    return <button style={cssBtn} onClick={onClick}>{label}</button>
  }

  const Tela = (valor, resultado) => {
    return (
      <div style={cssTela}>
        <span style={cssTelaOp}>{valor}</span>
        <span style={cssTelaResult}>{resultado}</span>
      </div>
    )
  }

  const addDigitoTela = (d) => {
    if ((d == '+' || d == '-' || d == '*' || d == '/') && operacao) {
      setOperacao(false)
      setValorTela(resultado + d)
      return
    } else if (operacao) {
      setValorTela(d)
      setOperacao(false)
      return
    }

    const valorDigitadoTela = valorTela + d
    setValorTela(valorDigitadoTela)
  }

  const limparMemoria = () => {
    setOperacao(false)
    setValorTela('')
    setResultado(0)
    setAcumulador(0)

    return
  }

  

  const cssTela = {
    display: 'flex',
    paddingLeft: 20,
    paddingRight: 20,
    justifyContent: 'center',
    alignItems: 'flex-start',
    backgroundColor: '#444',
    flexDirection: 'column',
    width: 260
  }

  const cssTelaOp = {
    fontSize: 25,
    color: '#fff',
    height: 20, 
  }

  const cssTelaResult = {
    fontSize: 50,
    color: '#fff'
  }

  const cssBtn = {
    fontSize: 30,
    height: 75,
    width: 75,
    padding: 20,
    backgroundColor: '#000',
    color: '#fff',
    colorBorder: '#fff',
    textAlign: 'center',
    outline: 'none'
  }

  return (
    <>
    </>
  )
}

export default App;