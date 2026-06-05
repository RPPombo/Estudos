import React, {useState} from "react";

function Formulario() {
    const [form, setForm] = useState({"nome": "", "curso": "", "ano": ""})
    const handleChangeForm = (e) => {
        const atributo = e.target.getAttribute('name');

        if (atributo == 'fnome') {
            setForm({"nome": e.target.value, "curso": form.curso, "ano": form.ano})
        } else if (atributo == 'fcurso') {
            setForm({"nome": form.nome, "curso": e.target.value, "ano": form.ano})
        } else if (atributo == 'fano') {
            setForm({"nome": form.nome, "curso": form.curso, "ano": e.target.value})
        }
    }

    const [carro, setCarro] = useState('HRV')
    
    return (
        <form>
            <label>Digite seu nome: </label>
            <input type="text" name="fnome" value={form.nome} onChange={(e)=>handleChangeForm(e)}/>
            <label>Digite seu curso: </label>
            <input type="text" name="fcurso" value={form.curso} onChange={(e)=>handleChangeForm(e)}/>
            <label>Digite seu ano: </label>
            <input type="text" name="fano" value={form.ano} onChange={(e)=>handleChangeForm(e)}/>
            <p>Nome digitado: {form.nome}</p>
            <p>Curso digitado: {form.curso}</p>
            <p>Ano digitado: {form.ano}</p>

            <label>Selecione o carro: </label>
            <select value={carro} onChange={(e)=>setCarro(e.target.value)}>
                <option value="HRV">HRV</option>
                <option value="Argos">Argos</option>
                <option value="Spin">Spin</option>
            </select>
            <p>Carro selecionado: {carro}</p>
        </form>
    )
}

export default Formulario;