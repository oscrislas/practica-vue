<template>
<div class="container">

    <div class="row">
            <div class="col-3">1</div>
            <div class="col-6">Login</div>
            <div class="col-3">3</div>
    </div>
    <div class="row">
            <div class="col-3"></div>
            <div class="col-6">
                <form @submit.prevent="validateUser">
                    <md-field>
                    <label>Usuario</label>
                    <md-input v-model="usuario"></md-input>
                    <span class="md-helper-text">Usuario</span>
                    </md-field>    <md-field>
                    <label>Contraseña</label>
                    <md-input v-model="contrasena"></md-input>
                    <span class="md-helper-text">Contraseña</span>
                    </md-field>
                    <md-button type="submit" class="md-primary">Iniciar sesion</md-button>
                </form>

            </div>
            <div class="col-3"></div>
    </div>
            
</div>
</template>

<script>
export default {
    name: "Login",
    data(){
        return {
            usuario: "",
            contrasena: ""
        }
    },
    methods:{
        validateUser(){
            this.$axios.post('http://localhost:3000/login',{user: "oscar",pass: "1234"})
            .then(res=>{
                console.log(res)
            });
            console.log(this.usuario +" "+ this.contrasena)
        }
    }
}
</script>

<style lang="stylus" scoped>

</style>