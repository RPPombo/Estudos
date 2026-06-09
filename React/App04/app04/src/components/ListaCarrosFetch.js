import React from "react";

class ListaCarrosFetch extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            carros: []
        }
    }

    componentDidMount() {
        fetch('http://localhost:3000').then((resultado) => {
            resultado.json().then(
                (res) => this.setState({carros: res})
            )
        })
    }

    render() {
        return (
            <div>
                {this.state.carros.map((carro)=>
                    <div key={carro.id}>{carro.id} - {carro.marca} - {carro.modelo}</div>
                )}
            </div>
        )
    }
}

export default ListaCarrosFetch;