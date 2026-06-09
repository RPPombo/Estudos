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
    if (d == '+' || d == '-' || d == '*' || d == '/') {
      if (operacao) {
        setOperacao(false)
        setValorTela(resultado + d)
        return
      } else if (valorTela[valorTela.length-1] == '+' ||
        valorTela[valorTela.length-1] == '-' ||
        valorTela[valorTela.length-1] == '*' ||
        valorTela[valorTela.length-1] == '/' ) {
          return
        }
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

  const fazerOperacao = (oper) => {
    if (oper == 'bs') {
      let vtela = valorTela
      vtela = vtela.substring(0, (vtela.length-1))
      setValorTela(vtela)
      setOperacao(false)
      return
    }

    try {
      const r = eval(valorTela)
      setAcumulador(r)
      setResultado(r)
      setOperacao(true)
    } catch {
      setResultado('ERRO')
    }
  }

  const cssContainer = {
    display: 'flex',
    justifyContent: 'flex-start',
    alignItems: 'center',
    flexDirection: 'column',
    width: 300,
    border: '1px solid #000'
  }

  const cssBotao = {
    flexDirection: 'row',
    flexWrap: 'wrap'
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
    <div style={cssContainer}>
      <h3>Calculadora matemática simples</h3>
      {Tela(valorTela, resultado)}
      <div style={cssBotao}>
        {Btn('AC', limparMemoria)}
        {Btn('(', () => addDigitoTela('('))}
        {Btn(')', () => addDigitoTela(')'))}
        {Btn('/', () => addDigitoTela('/'))}
        {Btn('7', () => addDigitoTela('7'))}
        {Btn('8', () => addDigitoTela('8'))}
        {Btn('9', () => addDigitoTela('9'))}
        {Btn('*', () => addDigitoTela('*'))}
        {Btn('4', () => addDigitoTela('4'))}
        {Btn('5', () => addDigitoTela('5'))}
        {Btn('6', () => addDigitoTela('6'))}
        {Btn('-', () => addDigitoTela('-'))}
        {Btn('1', () => addDigitoTela('1'))}
        {Btn('2', () => addDigitoTela('2'))}
        {Btn('3', () => addDigitoTela('3'))}
        {Btn('+', () => addDigitoTela('+'))}
        {Btn('0', () => addDigitoTela('0'))}
        {Btn('.', () => addDigitoTela('.'))}
        {Btn('<-', () => fazerOperacao('bs'))}
        {Btn('=', () => fazerOperacao('='))}
      </div>

    </div>
  )
}

export default App;