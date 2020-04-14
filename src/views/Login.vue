<template >
<div class="container ">

    <div class="row">
            <div class="col-3"></div>
            <div class="col-6"  >Login</div>
            <div class="col-3"></div>
    </div>
    <div class="row">
            <div class="col-3"></div>
            <div class="col-6">
                <form @submit.prevent="validateUser">
                    <div class="form-grup">
                        <md-field>
                        <label >Usuario</label>
                        <md-input  v-model="usuario" class="form-control"  ></md-input>
                        <span class="md-helper-text" >Usuario</span>
                        <p v-bind:class="false?'valido ' : 'invalido'"></p>
                        </md-field>    
                        <md-field>
                        <label>Contraseña</label>
                        <md-input v-model="contrasena" class="form-control" type="password"></md-input>
                        <span class="md-helper-text">Contraseña</span>
                        </md-field>
                        <md-button type="submit" :disabled="ssubmitStatus" class="md-primary">Iniciar sesion</md-button>
                    </div>

                </form>

            </div>
            <div class="col-3"></div>
            <notifications group="foo" />
    </div>
            
</div>
</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
export default {
    name: "Login",
    data(){
        return {
            usuario: "",
            contrasena: "",
            submitStatus: true,
            token: null
        }
    },
    methods:{
        validateUser(){
            this.$axios.post('http://localhost:3000/login',{usuario: this.usuario,contrasena: this.contrasena})
            .then(res=>{
                this.token=VueJwtDecode.decode(res.data.token).Username
                if(this.token.id!=""){
                    console.log(res.data)
                    if(this.token.admin==="Admin"){
                        this.$router.push('admin')
                        localStorage.setItem('token', res.data.token)
                    }else{
                        this.$router.push('user')
                        localStorage.setItem('token', res.data.token)
                    }

                    
                }else{
                    this.$notify({
                    group: 'foo',
                    title: 'Error de Usuario o contraseña',
                    text: 'Porfavor ingrese una contraseña o usuario correcto'
                    });
                }
            }).catch(()=>{
                    this.$notify({
                    group: 'foo',
                    title: 'Error al Iniciar sesión',
                    text: "Credenciales inválidas de autenticación revise su información"
                    });
            });
          
        }
    },
    computed: {
        ssubmitStatus: function () {
            if(this.usuario!=""&&this.contrasena!=""){
                console.log("true")
                return false
            }else{
            console.log("false")
                    return true
            }
            
        }
    }
}
</script>

<style>
.valido{
    color: green;
}
.invalido{

    color: red;
}

</style>