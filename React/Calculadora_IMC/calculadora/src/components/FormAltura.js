import React from "react";

class FormAltura extends React.Component {
    constructor(props){
        super(props)

        this.state = {
            info: " "
        }
    }

    render() {
        return (
            <div>
                <label>Altura: </label>
                <input type="text" value={this.state.info} onChange={(e)=>{
                    this.setState({info: e.target.value})
                    this.props.atribuirAltura(e.target.value)}}/>
            </div>
        )
    }
}

export default FormAltura;