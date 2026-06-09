import React from "react";
import ListaCarros from "./components/ListaCarros";
import ListaCarrosFetch from "./components/ListaCarrosFetch";

function App() {
  return (
    <div>
      <ListaCarros />
      <hr/>
      <ListaCarrosFetch />
    </div>
  );
}

export default App;
