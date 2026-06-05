import React from "react";

function Numero(props) {
    return (
        <>
            <p>Valor do state num: {props.num}</p>
            <button onClick={()=>props.setNum(props.num+10)}>num+10</button>
        </>
    );
}

export default Numero;