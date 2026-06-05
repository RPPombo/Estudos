import React from "react";
import Dados from "./Dados";

function Corpo() {
    const somar = (v1, v2) => {
        return v1+v2;
    }

    const texto_destaque = {
        color: "#00f",
        fontSize: "3em"
    }

    return (
        <section>
            <h2 style={texto_destaque}>Curso de React</h2>
            <p className="texto">Melhor que a aula do Hete</p>
            <Dados 
            comida_1={() => {return "lasanha"}} 
            comida_2="Carne" 
            comida_3="Frango Assado"
            somar = {somar}
            />
        </section>
    )
}

export default Corpo;