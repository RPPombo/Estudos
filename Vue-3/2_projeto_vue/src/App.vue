<template>
  <div>
    <Header />
    <MultiplosEventos />
    <PrimeiroComponente />
    <LifeCycle />
    <Pessoa />
  </div>
  

</template>

<script>
  import PrimeiroComponente from './components/PrimeiroComponente.vue'
  import LifeCycle from './components/LifeCycle.vue'
  import Pessoa from './components/Pessoa.vue';
  import Header from './components/Header.vue';
  import MultiplosEventos from './components/MultiplosEventos.vue';

  export default {
    name: 'App',
    components: {
      PrimeiroComponente,
      LifeCycle,
      Pessoa,
      Header,
      MultiplosEventos
    }
  }
</script>

<style>
  body {
    background-color: #333333;
    color: #FFFFFF; 
  }

  a {
    color:red;
  }
</style>