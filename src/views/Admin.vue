<template>
    <div class="row">
        <div class="col-1">

        </div>
        <div class="col-10">
            <b-button v-b-modal.modal-1 >Agregar Empleado</b-button>
            
            <br>
            <br>
            <table >
                <tr>
                    <th>Nombre</th>
                    <th>Apellido</th>
                    <th>Correo</th>
                    <th>Telefono</th>
                    <th>Entrada</th>
                    <th>Salida</th>
                    <th>Editar</th>
                    <th>Borrar</th>
                </tr>
                <tr v-for="(e,index) of  User" :key="e.id">
                    <td>{{e.Nombre}}</td>
                    <td>{{e.Apellidos}} </td>
                    <td>{{e.Correo}}</td>
                    <td>{{e.Telefono}}</td>
                    <td>{{e.FechaInicio | hora}}</td>
                    <td>{{e.FechaFin | hora}}</td>
                    <th><b-button v-b-modal.modal-2 v-on:click="precionado(index)">Editar</b-button></th>
                    <th><b-button v-on:click="borrar(e.id)">Borrar</b-button></th>
                  
                </tr>
            </table>
        </div>

        <div class="col-1"></div>
        <div>
  


</div>

<b-modal id="modal-1" hide-footer title="Registro Empleado">
  <my-registro></my-registro>
</b-modal>
<b-modal id="modal-2" hide-footer title="Actuliza Empleado">
  <my-registro  v-bind:id="id" > </my-registro>
</b-modal>

    </div>


    
</template>

<script>
import registroEmpleado from '@/components/RegistroEmpleado'
export default {
    name: 'registroEmpleado',
    components: {
    'my-registro':registroEmpleado
    },
    data(){
        return {
            User:[],
            row: 0,
            id:0
        }
    },
    filters:{
        hora: function(value){
            if(value==""){
                return "Sin checar"
            }
            var res =value.split(":");

         return parseInt(res[1], 10)%12+":"+res[2]+" "+ (parseInt(res[1], 10)%12==0?"am":"pm")
        }
    },
    beforeCreate(){

            this.$axios.get('http://localhost:3000/Empleados')
            .then(res=>{
               // this.User = JSON.parse(res.data);
            this.User=res.data
            });
        //console.log(JSON.parse(this.User))
        //console.log(JSON.stringify(user))
    },
    methods:{
        precionado(index){
            if(index!=null)
           this.id= this.User[index].id
        },
        borrar(id){
            console.log("el id es: "+ id)
            this.$axios.post('http://localhost:3000/borrarEmpleado',id)
            .then(res=>{
               // this.User = JSON.parse(res.data);
            this.User=res.data
            });
        }
    }

    
}
</script>

<style scoped>
.table{
    margin: 50px;
}
.center{
    position: fixed;
}
</style>

