<template>
    <div>
        <h2>Meu formulário:</h2>
        <form @submit="enviarFormulario($event)">
            <InputText />
            <InputText />
            <Submit />
        </form>
    </div>
</template>

<script>
    import Submit from './form/Submit.vue';
    import InputText from './form/InputText.vue';
    export default {
        name: 'Form',
        components: {
            Submit,
            InputText
        },
        methods: {
            enviarFormulario(e) {
                e.preventDefault();

                console.log("Formulário enviado!");
            }
        }
    }
</script>