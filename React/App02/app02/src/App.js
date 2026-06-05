import React from "react";
import ClasseGenerica from "./components/Classe";
import Carro from "./components/Carro";

function App() {
  return (
    <div>
      <h1>Componentes de classe</h1>
      <ClasseGenerica palavra="banana"/>
      <Carro fator={1} />
    </div>
  );
}

export default App;
