import React, {useState} from "react";

function SalvarNome(props) {
    const [nome, setNome] = useState();
    const [chave, setChave] = useState()

    return (
        <>
        <label>Digite uma chave: </label>
        <input type="text" onChange={(e)=>setChave(e.target.value)}/><br/>
        <label>Digite um nome: </label>
        <input type="text" onChange={(e)=>setNome(e.target.value)}/><br/>
        <button onClick={()=>props.armazenar(chave, nome)}>Guardar</button>
        <button onClick={()=>props.remover(chave)}>Remover</button>
        <button onClick={()=>props.consultar(chave)}>Consultar</button>
        </>
    )
}

export default SalvarNome;