import React from "react";
import { Switch, Route, Link } from "react-router-dom";
import Pagina1 from "./components/Pagina_1";
import Pagina2 from "./components/Pagina_2";
import Pagina3 from "./components/Pagina_3";

function App() {
  return (
    <>
    <header>
      <Link to='/'>Home</Link>
      <Link to='/pagina1'>Página 1</Link>
      <Link to='/pagina2'>Página 2</Link>
      <Link to='/pagina3'>Página 3</Link>
    </header>
    <main>
      <Switch>
        <Route path='/pagina1' component={Pagina1} />
        <Route path='/pagina2' component={Pagina2} />
        <Route path='/pagina3' component={Pagina3} />
      </Switch>
    </main>
    </>
  );
}

export default App;
