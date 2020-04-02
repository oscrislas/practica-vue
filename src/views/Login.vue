<template>
<div class="container">

    <div class="row">
            <div class="col-3">1</div>
            <div class="col-6"  :class="{'tex':usuario.length<4}">Login</div>
            <div class="col-3">3</div>
    </div>
    <div class="row">
            <div class="col-3"></div>
            <div class="col-6">
                <form @submit.prevent="validateUser">
                    <div class="form-grup">
                        <md-field>
                        <label>Usuario</label>
                        <md-input  v-model="usuario" class="form-control"></md-input>
                        <span class="md-helper-text">Usuario</span>
                        </md-field>    <md-field>
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

export default {
    name: "Login",
    data(){
        return {
            usuario: "",
            contrasena: "",
            submitStatus: true
        }
    },
    methods:{
        validateUser(){
            this.$axios.post('http://localhost:3000/login',{usuario: this.usuario,contrasena: this.contrasena})
            .then(res=>{
                if(res.data==true){
                    this.$router.push('home')
                    
                }else{
                    this.$notify({
                    group: 'foo',
                    title: 'Error de Usuario o contraseña',
                    text: 'Porfavor ingrese una contraseña o usuario correcto'
                    });
                }
            }).catch(error=>{
                    this.$notify({
                    group: 'foo',
                    title: 'Error al Iniciar sesión',
                    text: error
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
.tex {
  background-color: red;
}
</style>