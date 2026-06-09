import React, {useState} from "react";

function App() {

  const jogoInicial = [['','',''], ['','',''], ['','','']]
  const [jogo, setJogo] = useState(jogoInicial)
  const [simboloAtual, setSimboloAtual] = useState('X')
  const [jogando, setJogando] = useState(true)
  const [jogadas, setJogadas] = useState(0)

  const tabuleiro = (j) => {
    return (
      <div style={tabu}>
        <div style={tabuLinha}>
          <div style={casa} data-pos='00' onClick={(e) => jogar(e)}>{j[0][0]}</div>
          <div style={casa} data-pos='01' onClick={(e) => jogar(e)}>{j[0][1]}</div>
          <div style={casa} data-pos='02' onClick={(e) => jogar(e)}>{j[0][2]}</div>
        </div>
        <div style={tabuLinha}>
          <div style={casa} data-pos='10' onClick={(e) => jogar(e)}>{j[1][0]}</div>
          <div style={casa} data-pos='11' onClick={(e) => jogar(e)}>{j[1][1]}</div>
          <div style={casa} data-pos='12' onClick={(e) => jogar(e)}>{j[1][2]}</div>
        </div>
        <div style={tabuLinha}>
          <div style={casa} data-pos='20' onClick={(e) => jogar(e)}>{j[2][0]}</div>
          <div style={casa} data-pos='21' onClick={(e) => jogar(e)}>{j[2][1]}</div>
          <div style={casa} data-pos='22' onClick={(e) => jogar(e)}>{j[2][2]}</div>
        </div>
      </div>
    )
  }

  const verificaVitoria = () => {
    let pontos = 0
    let vitoria = false
    for (let l=0; l<3; l++) {
      pontos = 0;
      for (let c=0; c<3; c++) {
        if (jogo[l][c] == simboloAtual) {
          pontos++
        }
      }
      if (pontos == 3) {
        vitoria = true
        break
      }
    }

    for (let c=0; c<3; c++) {
      pontos = 0
      for (let l=0;l<3;l++) {
        if (jogo[l][c] == simboloAtual) {
          pontos++
        }
      }
      if (pontos == 3) {
        vitoria = true
        break
      }
    }

    pontos = 0
    for (let d=0;d<3;d++) {
      if (jogo[d][d] == simboloAtual ) {
        pontos++
      }
       
    }
    if (pontos == 3) {
        vitoria = true
    }

    pontos = 0
    let l = 0
    for (let c=2;c>=0;c--) {
      if (jogo[l][c] == simboloAtual) {
          pontos++
        }
    }
    if (pontos == 3) {
      vitoria = true
    }

    return vitoria
  }

  const trocaJogador = () => {
    simboloAtual == 'X'? setSimboloAtual('O') : setSimboloAtual('X')
  }

  const retornaPosicao = (e) => {
    const p = e.target.getAttribute('data-pos')
    const pos = [parseInt(p.substring(0,1)), parseInt(p.substring(1,2))]
    return pos
  }

  const verificaEspacoVazio = (e) => {
    if (jogo[retornaPosicao(e)[0]][retornaPosicao(e)[1]] == '') {
      return true
    } else {
      return true
    }
  }

  const jogar = (e) => {
    if (jogando) {
      if (verificaEspacoVazio(e)) {
        jogo[retornaPosicao(e)[0]][retornaPosicao(e)[1]] = simboloAtual
        setJogadas(jogadas+ 1)
        if(verificaVitoria()) {
          alert('Jogador do ' + simboloAtual + ' ganhou!')
          setJogando(false)
        } else {
          trocaJogador()
        }
      } else {
        alert('Esse espaço já está preenchido!')
      }
    }
  }

  const reiniciar = () => {
    setJogando(true)
    setJogo(jogoInicial)
    setSimboloAtual('X')
    setJogadas(0)
  }

  const btnJogarNovamente = () => {
    if (!jogando || jogadas >= 9) {
      return (
        <button onClick={() => reiniciar()}>Jogar Novamente</button>
      )
  }
  }

  const tabu = {
    display: 'flex',
    flexDirection: 'column'
  }

  const tabuLinha = {
    display: 'flex',
    flexDirection: 'row'
  }

  const casa = {
    width: 100,
    height: 100,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    flexDirection: 'row',
    cursor: 'pointer',
    fontSize: 60,
    border: '1px solid #000'
  }

  return (
    <div>
      <div>
        <p>Quem joga: {simboloAtual}</p>
      </div>
      <div>
        {tabuleiro(jogo)}
      </div>
      <div>
        {btnJogarNovamente()}
      </div>
    </div>
  );
}

export default App;
