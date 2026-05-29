<template>
    <div>
        <h1>Meu nome é: {{ nome }}</h1>
    </div>
</template>

<script>
    export default {
        name: 'LifeCycle',
        data() {
            return {
                nome: "Ainda não sei"
            }
        },
        created() {
            setTimeout(() => {
                this.nome = "Antonio"
            }, 2000)
        },
        mounted() {
            this.nome = "Carlos"
        }
    }
</script>
