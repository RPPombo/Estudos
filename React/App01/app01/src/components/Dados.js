import React from "react";

function Dados(props) {
    return (
        <>
        <p>Comida 1: {props.comida_1()}</p>
        <p>Comida 2: {props.comida_2}</p>
        <p>Comida 3: {props.comida_3}</p>
        <p>2 + 2 = {props.somar(2,2)}</p>
        </>
    )
}

export default Dados;