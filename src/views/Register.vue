<template>
<div class="container">

    <div class="row">
            <div class="col-3"></div>
            <div class="col-6">Registro</div>
            <div class="col-3"></div>
    </div>
    <div class="row">
            <div class="col-3"></div>
            <div class="col-6">
                <form @submit.prevent="registra">
                    <md-field>
                    <label>Nombre</label>
                    <md-input v-model="Usuario.nombre"></md-input>
                    <span class="md-helper-text">Nombre</span>
                    </md-field>    <md-field>
                    <label>Email</label>
                    <md-input v-model="Usuario.email"></md-input>
                    <span class="md-helper-text">Email</span>
                    </md-field>
                                        <md-field>
                    <label>Contraseña</label>
                    <md-input v-model="Usuario.password"></md-input>
                    <span class="md-helper-text">Contraseña</span>
                    </md-field>    <md-field>
                    <label>Re-Contraseña</label>
                    <md-input v-model="contrasena"></md-input>
                    <span class="md-helper-text">Re-Contraseña</span>
                    </md-field>
                    <md-button type="submit" class="md-primary">Registro</md-button>
                </form>

            </div>
            <div class="col-3"></div>
    </div>
            
</div>
</template>

<script>
export default {
    name: 'registros',
    data(){
        return {
            Usuario: {
                nombre: null,
                email: null,
                password: null

                },
            contrasena: null
        }
    },
    methods:{
        registra(){
            console.log("registra")
        }
    }
}

</script>

<style lang="stylus" scoped>

</style>