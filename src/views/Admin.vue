<template>
<div>

        <div class="main"><h1>Hola {{Admin.nombre}} {{Admin.apellidos}}</h1>
            <div class="b">
                <b-button v-b-modal.modal-1  >Agregar Empleado</b-button>
                <b-button v-on:click="Salir" >Salir</b-button>
            </div>

        </div>
    <div class="row">

        <div class="col-1">

        </div>
        <div class="col-10">

    
            
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
                    <td>{{e.nombre}}</td>
                    <td>{{e.apellidos}} </td>
                    <td>{{e.correo}}</td>
                    <td>{{e.telefono | tel}}</td>
                    <td>{{e.fechaInicio | hora}}</td>
                    <td>{{e.fechaFin | hora}}</td>
                    <td> <b-button v-b-modal.modal-2 v-on:click="precionado(index)" class="b">Editar</b-button></td>
                    <td><b-button v-on:click="borrar(e.id)" class="b">Borrar</b-button></td>
                  
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
                id: '',
                nombre: '',
                apellidos: '',
                telefono: '',
                correo: '',
                contrasena: '',
                admin: ''
            },
            row: 0,
            id:0,
            U:0,
            Admin:{
                Nombre:"",
                Apellidos:"",
                Admin:""
            },
            token: null
        }
    },
    filters:{
        hora: function(value){
            if (value!= undefined){
                if(value=="Sin Checar"){
                    return "Sin checar"
                }
                var ret=value.split(" ");
                var res =ret[1].split(":");

            return ((parseInt(res[0], 10)>12)?(parseInt(res[0], 10)-12):(parseInt(res[0], 10)))+":"+res[1]+" "+ (parseInt(res[0], 10)/12>1?"pm":"am")
            }
            return "Sin checar"
        },
        tel: function(value){
            if (value!= undefined){
                if(value.length==10){
                value='('+value
                var uno=value.slice(0, -7);
                var dos=value.slice(-7,-4);
                var tres= value.slice(-4)
                return  uno+') '+dos+" "+tres
                }
            }

            return "Incorrecto"

        }
    },
    mounted(){
this.Admin  = VueJwtDecode.decode(localStorage.getItem('token')).Username;

        if (this.Admin.admin==="Admin"){
            
            this.$axios.get('http://localhost:3000/Empleados',{ headers: { 
                'Authorization': "Bearer "+localStorage.getItem('token'),
                'Content-Type': 'application/json' } })
                .then(res=>{
                    console.log(res.data)
                    this.User=res.data

                    });

        }
        else{
            this.$router.push('/')
        }
    },
    methods:{
        precionado(index){
            if(index!=null)
           this.U= this.User[index]
           else
           this.U=this.Usuario
        },
        borrar(id){
            console.log("el id es: "+ id)
            this.$axios.post('http://localhost:3000/borrarEmpleado',id,{ headers: { 
                'Authorization': "Bearer "+localStorage.getItem('token'),
                'Content-Type': 'application/json' } })
            .then(res=>{
               // this.User = JSON.parse(res.data);
            this.User=res.data
            });
            console.log(this.User)
        },
        Salir(){
            console.log("salio")
            localStorage.removeItem('token')
            this.$router.push('/')
        },
        closemodal(){
            this.$axios.get('http://localhost:3000/Empleados',{ headers: { 
                'Authorization': "Bearer "+localStorage.getItem('token'),
                'Content-Type': 'application/json' } })
                .then(res=>{this.User=res.data});
        }
    }

    
}
</script>

<style scoped>
table{
border-collapse: separate !important;
      opacity: .8 !important;
    font-size: 15px;
  text-align: left;
  color: aliceblue;
  border-radius: 15px;
  border: solid black;
background-color: #001F49;
  position: absolute;
  width: 100%;
  margin:  auto;

}
th{
  padding: 10px;
  text-align: left;
  
  background-color: black;
}
td{
    padding-left: 10px;
    text-align: left;

  border: solid black;

}

tr:hover {
    background-color: #01295F;
    border-radius: 100% solid black;
}

h1{
  position: absolute;
  left: 10px;
}

.b{
position: absolute;
  right: 10px;
  width: 300px;
  height: auto;
  display: inline;
}



</style>

