<template>
    <div>
        <h2>Esta é a descrição da pessoa: {{ nome }}</h2> 
        <p>Estou trabalhando no momento.</p>
        <p>Sei as seguintes linguagens de programação:</p>
        <ul>
        <li>Python</li>
            <li>C</li>
            <li>Go</li>
            <li>SQL</li>
        </ul>
    </div>
</template>

<script>
    export default {
        name: 'Pessoa',
        data() {
            return {
                nome: 'Rodrigo'
            }
        }
    }
</script>
