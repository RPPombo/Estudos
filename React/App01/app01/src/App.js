import React from "react";
import Header from "./components/Header";
import Corpo from "./components/Corpo";
import Relogio from "./components/Relogio";
import "./App.css"

function App() {

  return (
    <>
      <section className="container">
        <Header />
        <Relogio />
        <Corpo />
      </section>
    </>
  );
}

export default App;
