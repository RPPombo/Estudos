<template>
    <div id="burger-table">
        <Message :msg="msg" v-show="msg"/>
        <div>
            <div id="burger-table-heading">
                <div class="order-id">#:</div>
                <div>cliente:</div>
                <div>Pão:</div>
                <div>Carne:</div>
                <div>Opicionais:</div>
                <div>Ações:</div>
            </div>
        </div>
        <div id="burger-table-rows">
            <div class="burger-table-row" v-for="item in burgers" :key="item.id">
                <div class="order-number">{{ item.id }}</div>
                <div>{{ item.nome }}</div>
                <div>{{ item.pao }}</div>
                <div>{{ item.carne }}</div>
                <div>
                    <ul>
                        <li v-for="(opcional, index) in item.opcionais" :key="index">{{ opcional }}</li>
                    </ul>
                </div>
                <div>
                    <select name="status" class="status" @change="updateBurger($event, item.id)">
                        <option value="">Selecione</option>
                        <option :value="s.tipo" v-for="s in status" :key="s.id" :selected="item.status == s.tipo">{{ s.tipo }}</option>
                    </select>
                    <button class="delete-btn" @click="deleteBurger(item.id)">Cancelar</button> 
                </div>
            </div>

        </div>
    </div>
</template>

<script>
    import Message from './Message.vue';

    export default {
        name: 'Dashboard',
        components: {
            Message
        },
        data() {
            return {
                burgers: null,
                burger_id: null,
                status: [],
                msg: null
            }
        },
        methods: {
            async getPedidos() {
                const req = await fetch('http://localhost:3000/burgers');

                const data = await req.json();

                this.burgers = data;
            },
            async getStatus() {
                const req = await fetch('http://localhost:3000/status');

                const data = await req.json();

                this.status = data;
            },
            async deleteBurger(id) {
                const req = await fetch(`http://localhost:3000/burgers/${id}`,
                    {
                        method: "DELETE"
                    });

                    const res = await req.json();

                    this.getPedidos();

                    this.msg = `Pedido de nº${id} foi cancelado!`;

                    setTimeout(() => this.msg = null, 3000);
            },
            async updateBurger(event, id) {
                const option = event.target.value;

                const dataJSON = JSON.stringify({ status: option})

                const req = await fetch(`http://localhost:3000/burgers/${id}`,
                    {
                        method: "PATCH",
                        headers: {"Content-Type": "application/json"},
                        body: dataJSON
                    });

                const res = await req.json()

                this.msg = `Pedido de nº${res.id} foi atualizado!`

                setTimeout(() => this.msg = null, 3000)
            } 
        },
        mounted() {
            this.getPedidos();
            this.getStatus()
        }
    }
</script>

<style scoped>
    #burger-table {
        max-width: 1200px;
        margin: 0 auto;
    }

    #burger-table-heading, #burger-table-rows, .burger-table-row {
        display: flex;
        flex-wrap:wrap;
    }
    #burger-table-heading {
        font-weight: bold;
        padding: 12px;
        border-bottom: 3px solid #333;
    }

    #burger-table-heading div, .burger-table-row div{
        width: 19%;
    }

    .burger-table-row {
        width: 100%;
        padding: 12px;
        border-bottom: 1px solid #CCC
    }

    #burger-table-heading .order-id, .burger-table-row .order-number {
        width: 5%;
    }

    select {
        padding: 12px 6px;
        margin-right: 12px;
    }

    .delete-btn {
        background-color: #222;
        color: #FCBA03;
        font-weight: bold;
        border: 2px solid #222;
        padding: 10px;
        margin: 0 auto;
        cursor: pointer;
        transition: .5s;
    }

    .delete-btn:hover {
        background-color: transparent;
        color:#222;
    }
</style>