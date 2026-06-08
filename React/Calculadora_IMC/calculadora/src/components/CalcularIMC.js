import React from "react";

class CalcularIMC extends React.Component {
    constructor(props) {
        super(props)
    }

    calc(peso, altura) {
        const resultado = peso/(altura*altura)
        this.props.setRes(resultado)
    }
    render() {
        return (
                <button onClick={()=>{this.calc(this.props.peso, this.props.altura)}}>Calcular IMC</button>
        )
    }
}

export default CalcularIMC;