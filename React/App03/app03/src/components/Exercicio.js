import React from "react";

class Exercicio extends React.Component {
    constructor(props) {
        super(props)
        this.carros = [
            {categoria: "Esporte", preco: "300000", modelo: "Cayman"},
            {categoria: "Esporte", preco: "100000", modelo: "Camaro"},
            {categoria: "Esporte", preco: "500000", modelo: "911-Coupé"},
            {categoria: "SUV", preco: "100000", modelo: "HRV"},
            {categoria: "SUV", preco: "130000", modelo: "Santafé"},
            {categoria: "SUV", preco: "1500000", modelo: "L35"},
            {categoria: "Utilitário", preco: "300000", modelo: "S10"},
            {categoria: "Utilitário", preco: "350000", modelo: "Hilux"},
            {categoria: "Utilitário", preco: "450000", modelo: "RAM"}
        ]

        this.state = {
            categoria: ""
        }
    }

    linhas(cat) {
        const li = []

        this.carros.forEach((carro) => {
            if (carro.categoria.toUpperCase() === cat.toUpperCase() || cat === ''){
                li.push(
                    <tr>
                        <td>{carro.categoria}</td><td>{carro.preco}</td><td>{carro.modelo}</td>
                    </tr>
                )
            }
        });

        return li
    }

    tabelaCarros(cat) {
        return (
            <table border='1' style={{borderCollapse: "collapse"}}>
                <thead>
                    <tr>
                        <th>Categoria</th><th>Preço</th><th>Modelo</th>
                    </tr>
                </thead>
                <tbody>
                    {this.linhas(cat)}
                </tbody>
            </table>
        )
    }

   pesquisaCategoria(cat, scat) {
    return (
        <div>
            <label>Digite a categoria desejada: </label>
            <input
                type="text"
                value={cat}
                onChange={(e) => scat(e.target.value)}
            />
        </div>
    )
    }

    render() {
    return (
        <>
            {this.pesquisaCategoria(this.state.categoria, (valor) => this.setState({ categoria: valor }))}
            <br/>
            {this.tabelaCarros(this.state.categoria)}
        </>
    )
    }
}

export default Exercicio;