import React from "react";

class Carro extends React.Component {
    constructor(props) {
        super(props);
        this.modelo = "Golf";
        this.state = {
            ligado: false,
            velocidade: 0
        }

        this.ligarBind = this.ligar.bind(this)
    }

    ligar() {
        this.setState({ligado: !this.state.ligado})
        
    }

    acelerar() {
        this.setState((state, props)=>({
            velocidade: state.velocidade + props.fator
        }))
    }

    desacelerar() {
        this.setState((state, props)=>({
            velocidade:state.velocidade - props.fator
        }))
    }

    componentDidMount() {
        console.log("O carro foi montado!")
    }

    componentDidUpdate() {
        console.log("O carro foi atualizado!")
    }

    componentWillUnmount() {
        console.log("O carro foi desmontado!")
    }

    render() {
        return(
            <>
                <h2>Meu Carro</h2>
                <p>Modelo: {this.modelo}</p>
                <h3>Informações: </h3>
                <p>{this.state.ligado? "Ligado": "Desligado"}</p>
                <p>Velocidade: {this.state.velocidade} Km/h</p>
                <button onClick={this.ligarBind}>{this.state.ligado? "Desligar":"Ligar"}</button>
                <button onClick={()=>this.acelerar()}>Acelerar</button>
                <button onClick={()=>this.desacelerar()}>Desacelerar</button>
            </>
        )
    }
}

export default Carro;