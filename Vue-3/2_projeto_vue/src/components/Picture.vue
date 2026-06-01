<template>
    <div>
        <img v-bind:src="imagem" :alt="descricao">
        <MudarImagem @mudar-imagem="trocar"/>
    </div>
</template>

<script>
    import MudarImagem from './MudarImagem.vue';
    export default {
        name: 'Picture',
        components: {
            MudarImagem
        },
        data() {
            return {
                imagem: "/img/Imagem-1.png",
                descricao: "Jorginho"
            }
        },
        methods: {
            trocar() {
                this.imagem = "/img/Imagem-2.png"
            }
        }
    }
</script>
