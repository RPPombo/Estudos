import React from "react";

class FormPeso extends React.Component {
    constructor(props){
        super(props)

        this.state = {
            info: " "
        }
    }

    render() {
        return (
            <div>
                <label>Peso: </label>
                <input type="text" value={this.state.info} onChange={(e)=>{
                    this.setState({info: e.target.value})
                    this.props.atribuirPeso(e.target.value)}}/>
            </div>
        )
    }
}

export default FormPeso;