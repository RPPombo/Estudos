import React, {useState} from "react";
import Imagem_1 from "./img/Imagem-1.png";
import Imagem_2 from "./img/Imagem-2.png";

function Header() {
    const [ligado, setLigado] = useState(true);

    return (
        <header>
            <h1 style={{color: '#f00', fontSize: '5em'}}>Imagem Aleatória</h1>
            <img style={{width:'100px', height:'100px'}} src={ligado?Imagem_1: Imagem_2}/>
            <button onClick={()=>setLigado(!ligado)}>{ligado?"Desligar":"Ligar"}</button>
        </header>
    );
}

export default Header;