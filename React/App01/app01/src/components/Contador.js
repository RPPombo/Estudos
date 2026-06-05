import React, {useEffect, useState} from "react";

function Contador() {
    const [contar, setContar] = useState(0);
    
    useEffect(
        () => {
            console.log("Página carregada");
            setContar(contar + 1);
        }
    )

    return (
        <>
            <p>Contador: {contar}</p>
        </>
    )
}

export default Contador;