import React from "react";

function Relogio() {
    return (
        <p>
            {new Date().toLocaleTimeString()}
        </p>
    )
}

export default Relogio;