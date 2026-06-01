<template>
    <div>
        <ul>
            <li><a href="#">Algum lugar</a></li>
            <li><a href="#">Outro lugar</a></li>
            <li><a href="#">Mais um lugar</a></li>
        </ul>
    </div>
</template>

<script>
    export default {
        name: 'Header',

    }
</script>

<style scoped>
    ul {
        list-style: none;
        display:flex;
    }

    li {
        margin-right: 10px;
    }

    a {
        color: #DDDDDD;
        text-decoration: none;
        transition: .5s;
    }

    a:hover {
        color: #567986
    }
</style>