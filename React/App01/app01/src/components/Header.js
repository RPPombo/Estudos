import React from "react";
import Imagem from "./img/Imagem-2.png"

function Header() {
    return (
        <header>
            <h1 style={{color: '#f00', fontSize: '5em'}}>Imagem Aleatória</h1>
            <img src={Imagem}/>
        </header>
    );
}

export default Header;