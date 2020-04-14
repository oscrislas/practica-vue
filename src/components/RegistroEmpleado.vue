<template>
<div class="container">

    <div class="row">
            <div class="col-3"></div>
            <div class="col-6"></div>
            <div class="col-3"></div>
    </div>
    <div class="row">
            <div class="col-3"></div>
            <div class="col-6">
                <form @submit.prevent="registra">
                    <md-field>
                    <label>Nombre</label>
                    <md-input v-model="Usuario2.nombre"></md-input>
                    <span class="md-helper-text">Nombre</span>
                    </md-field>    <md-field>
                    <label>Apellidos</label>
                    <md-input v-model="Usuario2.apellidos"></md-input>
                    <span class="md-helper-text">Apellidos</span>
                    </md-field>
                    <md-field>
                    <label>Telefono</label>
                    <md-input  type="number" v-model="Usuario2.telefono"></md-input>
                    <span class="md-helper-text">Telefono</span>
                    </md-field>    
                    <md-field>
                    <label>Correo</label>
                    <md-input v-model="Usuario2.correo"></md-input>
                    <span class="md-helper-text">Correo</span>
                    </md-field>
                    <md-field>
                    <label>Contraseña</label>
                    <md-input type="password" v-model="Usuario2.contrasena"></md-input>
                    <span class="md-helper-text">Contraseña</span>
                    </md-field>
                    <md-button type="submit" :disabled="ssubmit"  v-if="Usuario2.id==''" class="md-primary">Registrar</md-button>
                    <md-button type="submit" :disabled="submit" v-if="Usuario2.id!=''" class="md-primary">Actuliza</md-button>
                </form>

            </div>
            <div class="col-3"></div>
    </div>
<notifications group="foo" />
</div>
</template>

<script>
export default {
    name: 'registros',
    props: ['Usuario'],
    data(){
        return {
            registro: true,
            Usuario2:{
                id: "",
                nombre: "",
                apellidos: "",
                telefono: "",
                correo: "",
                contrasena: "",
                admin: ""
            },
            Usuario3:{
                id: "",
                nombre: "",
                apellidos: "",
                telefono: "",
                correo: "",
                contrasena: "",
                admin: ""
            }
        }
                
    },
    methods:{
        registra(){
            this.$axios.post('http://localhost:3000/Empleado',this.Usuario2,{ headers: { 
                'Authorization': "Bearer "+localStorage.getItem('token'),
                'Content-Type': 'application/json' } })
            .then(res=>{
                if(res.data===true){
                    this.$notify({
                    group: 'foo',
                    title: 'Exito',
                    text: 'Operacion Exitosa'
                    });
                    this.Usuario2=this.Usuario3
                }else{
                    this.$notify({
                    group: 'foo',
                    title: 'Fallo',
                    text: 'Ubo un error intente nuevamente'
                    });
                }
            });
        //this.Usuario2="";
        }

    },
    computed:{
        ssubmit: function(){
            if(this.Usuario2.correo!=""&&this.Usuario2.contrasena!=""){
                console.log("conputed false")
                return false
            }
            console.log("computed true")
            return true
        },
        submit: function(){
            console.log("2 entro")

            if(JSON.stringify(this.Usuario)==JSON.stringify(this.Usuario2)){
                console.log(JSON.stringify(this.Usuario)+" "+JSON.stringify(this.Usuario2))
                return true
            }
            return false
        }
    },
    mounted(){
           this.Usuario2= Object.assign({},this.Usuario)
           
 
        
console.log("montado")
    }

}

</script>

<style >
.red{
    background-color: red;
}
</style>