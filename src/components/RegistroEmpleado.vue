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
                    <md-input v-model="Usuario.Nombre"></md-input>
                    <span class="md-helper-text">Nombre</span>
                    </md-field>    <md-field>
                    <label>Apellidos</label>
                    <md-input v-model="Usuario.Apellidos"></md-input>
                    <span class="md-helper-text">Apellidos</span>
                    </md-field>
                    <md-field>
                    <label>Telefono</label>
                    <md-input  type="number" v-model="Usuario.Telefono"></md-input>
                    <span class="md-helper-text">Telefono</span>
                    </md-field>    
                    <md-field>
                    <label>Correo</label>
                    <md-input v-model="Usuario.Correo"></md-input>
                    <span class="md-helper-text">Correo</span>
                    </md-field>
                    <md-field>
                    <label>Contraseña</label>
                    <md-input type="password" v-model="Usuario.Contrasena"></md-input>
                    <span class="md-helper-text">Contraseña</span>
                    </md-field>
                    <md-button type="submit" v-if="id==null" class="md-primary">Registrar</md-button>
                    <md-button type="submit" v-if="id!=null" class="md-primary">Actuliza</md-button>
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
    props: ['id'],
    data(){
        return {
            registro: true,
            Usuario: {
                Id: null,
                Nombre: null,
                Apellidos: null,
                Telefono: null,
                Correo: null,
                Contrasena: null,
                Admin: null
                }
        }
    },
    methods:{
        registra(){
            this.$axios.post('http://localhost:3000/Empleado',this.Usuario)
            .then(res=>{
                if(res.data===true){
                    this.$notify({
                    group: 'foo',
                    title: 'Exito',
                    text: 'Operacion Exitosa'
                    });
                }else{
                    this.$notify({
                    group: 'foo',
                    title: 'Fallo',
                    text: 'Ubo un error intente nuevamente'
                    });
                }
            });

        }

    },
    mounted(){
        console.log("el id es "+this.id)
        if(this.id!=null){
                    this.$axios.post('http://localhost:3000/getEmpleado', this.id)
            .then(res=>{
               // this.User = JSON.parse(res.data);
            this.Usuario=res.data
            console.log(this.Usuario)
            });
        }
        
console.log("montado")
    }

}

</script>

<style >
.red{
    background-color: red;
}
</style>