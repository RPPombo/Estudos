import React from "react";

class BaseClasse extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            info1: 'Primeira informação',
            info2: 'Segunda informação',
            info3: this.props.info3
        }

        this.tI = this.trocarInfo1.bind(this)
    }

    trocarInfo1() {
        this.setState(state=>({
            info1: "Nova informação"
        }))
    }

    render() {
        return (
            <>
                <h2>Componente de classe genérico</h2>
                <p>{this.state.info1}</p>
                <p>{this.state.info2}</p>
                <p>{this.state.info3}</p>

                <button onClick={this.tI}>Trocar informação 1</button>
            </>
        )
    }

    componentDidMount() {
        console.log("componente genérico montado!")
    }

    componentDidUpdate() {
        console.log("componente genérico atualizado!")
    }

    componentWillUnmount() {
        console.log("componente genérico desmontado!")
    }
}

export default BaseClasse;