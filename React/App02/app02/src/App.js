import React, {useState} from "react";
import ClasseGenerica from "./components/BaseClasse";
import Carro from "./components/Carro";

function App() {
  const [montado, setCarro] = useState(true)
  const [classe, setClasse] = useState(true)

  return (
    <div>
      <h1>Componentes de classe</h1>
      {classe? <ClasseGenerica info3="banana"/> : <p>Sem componente genérico</p>}
      <button onClick={()=>setClasse(!classe)}>{classe? "Desmontar" : "Montar"} classe</button>
      {montado? <Carro fator={5} />: <p>Sem carro</p>}
      <button onClick={()=>setCarro(!montado)}>{montado? "Desmontar" : "Montar"} carro</button>
    </div>
  );
}

export default App;
