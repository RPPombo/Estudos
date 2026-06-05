import React from "react";

function Carros() {
    const lista = ['HRV', 'Golf', 'Focus', 'Cruze', 'Argos', 'Spin', 'Santafe', 'X6'];

    const carros = lista.map((c,i)=> 
        <li key={i}>{i + ': ' + c}</li>
    )

    return (
        <>
        <ul>
            {carros}
        </ul>
        </>
    );
}

export default Carros;