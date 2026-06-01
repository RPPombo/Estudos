<template>
    <div>
        <button @click="primeiro($event), segundo($event)">Ativar múltiplos eventos</button>
        <p>{{ informacao }}</p>
    </div>
</template>

<script>
    export default {
        name: 'MultiplosEventos',
        data() {
            return {
                informacao: ""
            }
        },
        methods: {
            primeiro(e) {
                e.preventDefault()
                this.informacao = "Primeiro evento ativado!"
            },
            segundo(e) {
                e.preventDefault()
                setTimeout(() => this.informacao = "Segundo evento ativado!", 2000)
            }
        }
    }
</script>       