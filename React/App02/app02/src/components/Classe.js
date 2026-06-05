import React from "react";

class ClasseGenerica extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <>
            <h2>Primeiro componente de classe!</h2>
            <p>Palavra: {this.props.palavra}</p>
            </>
        )
    }
}

export default ClasseGenerica;