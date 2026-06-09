import React from "react";
import axios from "axios";

class ListaCarros extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            carros: []
        }
    }

    componentDidMount() {
        axios.get("http://localhost:3000").then((res) => {
            const dadosCarros = res.data
            this.setState({carros: dadosCarros})
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

export default ListaCarros;