<template>
    <div>
        <Picture />
        <p v-if="esta_trabalhando">Estou trabalhando no momento.</p>
        <p v-else>Não estou trabalhando atualmente.</p>
        <p>Sei as seguintes linguagens de programação:</p>
        <ul>
        <li v-for="(item, index) in lista" :key="index">{{ item }}</li>
        </ul>
        <div>
            <button @click="showEmail">{{ texto_botao }}</button>
        </div>
        <p v-show="mostrar_email">Meu email: {{ email }}</p>
        <p v-show="mostrar_email">Email auxiliar: {{ compEmail }}</p>
        <p>Para acessar meu portifólio <a v-bind:href="meu_link" target="_blank">clicar aqui</a></p>
    </div>
</template>

<script>
    import Picture from './Picture.vue';
    export default {
        name: 'Info',
        components: {
            Picture
        },
        props: {
            compEmail: String
        },
        data() {
            return {
                esta_trabalhando: false,
                mostrar_email: true,
                email: "alguma_coisa@email.com",
                meu_link: "https://google.com",
                texto_botao: "Esconder e-mail",
                lista: ["Python", "C", "Go"]
            }
        },
        methods: {
            showEmail() {
                this.mostrar_email = !this.mostrar_email
                if (!this.mostrar_email) {
                    this.texto_botao = "Mostrar e-mail"
                } else {
                    this.texto_botao = "Esconder e-mail"
                }
            }
        }
    }
</script>
