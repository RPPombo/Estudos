import React from "react";

function Relogio() {
    const hora = new Date().toLocaleTimeString()

    const cumprimento = () => {
        const h = new Date().getHours()
        if (h >= 0 && h < 13) {
            return <p>Bom dia!</p>
        } else if (h >= 13 && h < 19) {
            return <p>Boa Tarde!</p>
        } else {
            return <p>Boa Noite!</p>
        }
    }

    return (
        <>
            <p>{hora}</p>
            {cumprimento()}
        </>
    )
}

export default Relogio;