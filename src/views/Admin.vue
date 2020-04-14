<template>
    <div class="row">
        <div class="col-1">

        </div>
        <div class="col-10">
            <b-button v-b-modal.modal-1 class="b">Agregar Empleado</b-button>
            <b-button v-on:click="Salir">Salir</b-button>
            <p>Hola {{Admin.Nombre}} {{Admin.Apellidos}}</p>
            
            <br>
            <br>
            <table class="table-dark tabla" >
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
                    <td>{{e.Telefono | tel}}</td>
                    <td>{{e.FechaInicio | hora}}</td>
                    <td>{{e.FechaFin | hora}}</td>
                    <th><b-button v-b-modal.modal-2 v-on:click="precionado(index)" class="b">Editar</b-button></th>
                    <th><b-button v-on:click="borrar(e.id)" class="b">Borrar</b-button></th>
                  
                </tr>
            </table>
        </div>

        <div class="col-1"></div>
        <div>
  


</div>

<b-modal id="modal-1" hide-footer title="Registro Empleado" @hidden="closemodal" >
  <my-registro v-bind:Usuario="Usuario"></my-registro>
</b-modal>
<b-modal id="modal-2" hide-footer title="Actuliza Empleado" @hidden="closemodal">
  <my-registro  v-bind:Usuario="U" > </my-registro>
</b-modal>

    </div>


    
</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
import registroEmpleado from '@/components/RegistroEmpleado'
export default {
    name: 'registroEmpleado',
    components: {
    'my-registro':registroEmpleado
    },
    data(){
        return {
            User:[],
            Usuario:{
                Id: '',
                Nombre: '',
                Apellidos: '',
                Telefono: '',
                Correo: '',
                Contrasena: '',
                Admin: ''
            },
            row: 0,
            id:0,
            U:0,
            Admin:{
                Nombre:"",
                Apellidos:"",
                Admin:""
            }
        }
    },
    filters:{
        hora: function(value){
            if(value=="Sin Checar"){
                return "Sin checar"
            }
            var ret=value.split(" ");
            var res =ret[1].split(":");

         return ((parseInt(res[0], 10)>12)?(parseInt(res[0], 10)-12):(parseInt(res[0], 10)))+":"+res[1]+" "+ (parseInt(res[0], 10)/12>1?"pm":"am")
        },
        tel: function(value){
            if(value.length==10){
            value='('+value
            var uno=value.slice(0, -7);
            var dos=value.slice(-7,-4);
            var tres= value.slice(-4)
            return  uno+') '+dos+" "+tres
            }
            return "Incorrecto"

        }
    },
    mounted(){


        if (localStorage.getItem('token')){
            this.Admin  = VueJwtDecode.decode(localStorage.getItem('token')).Username;
            this.$axios.get('http://localhost:3000/Empleados')
                .then(res=>{this.User=res.data});

        }
        else{
            this.$router.push('/')
        }
    },
    methods:{
        precionado(index){
            if(index!=null)
           this.U= this.User[index]
        },
        borrar(id){
            console.log("el id es: "+ id)
            this.$axios.post('http://localhost:3000/borrarEmpleado',id)
            .then(res=>{
               // this.User = JSON.parse(res.data);
            this.User=res.data
            });
            console.log(this.User)
        },
        Salir(){
            console.log("salio")
            localStorage.removeItem('admin')
            this.$router.push('/')
        },
        closemodal(){
            this.$axios.get('http://localhost:3000/Empleados')
                .then(res=>{this.User=res.data});
        }
    }

    
}
</script>

<style scoped>
.table,th,td{
  padding: 10px;
  text-align: left;
  border: solid black;
 

}

tr:hover {
    background-color: black;
    border-radius: 100% solid black;
}
.th{
    width: 10px;
    background-color: black;
    
}
.center{
    position: fixed;
}
.tabla{
    opacity: .9;
    font-size: 15px;
}
.b{
    opacity: .9;
    font-size: 14px;
}
</style>

