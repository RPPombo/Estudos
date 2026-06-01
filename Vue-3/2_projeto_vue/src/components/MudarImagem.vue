<template>
    <div>
        <button @click="$emit('mudarImagem')">Trocar imagem</button>
    </div>
</template>

<script>
    export default {
        name: 'MudarImagem',
        emits: ['mudarImagem']      
    }

</script>