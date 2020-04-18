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
                    <th class="priority-5" >Correo</th>
                    <th class="priority-4" >Telefono</th>
                    <th>Entrada</th>
                    <th>Salida</th>
                    <th class="priority-3">Editar</th>
                    <th class="priority-3">Borrar</th>
                </tr>
                <tr v-for="(e,index) of  User" :key="e.id">
                    <td>{{e.nombre}}</td>
                    <td >{{e.apellidos}} </td>
                    <td class="priority-5">{{e.correo}}</td>
                    <td class="priority-4">{{e.telefono | tel}}</td>
                    <td>{{e.fechaInicio | hora}}</td>
                    <td>{{e.fechaFin | hora}}</td>
                    <td class="priority-3"> <b-button v-b-modal.modal-2 v-on:click="precionado(index)" >Editar</b-button></td>
                    <td class="priority-3"><b-button v-on:click="borrar(e.id)" >Borrar</b-button></td>
                  
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

.b{
position: absolute;
  right: 10px;
  width: 300px;
  height: auto;
  display: inline;
}

@media screen and (max-width: 1225px) and (min-width: 1045px) {
                .priority-5{
                        display:none;
                }
        }
        
        @media screen and (max-width: 1045px) and (min-width: 835px) {
                .priority-5{
                        display:none;
                }
                .priority-4{
                        display:none;
                }
        }
        
        @media screen and (max-width: 835px) and (min-width: 300px) {
                .priority-5{
                        display:none;
                }
                .priority-4{
                        display:none;
                }
                .priority-3{
                        display:none;
                }
        }
        
        @media screen and (max-width: 300px) {
                .priority-5{
                        display:none;
                }
                .priority-4{
                        display:none;
                }
                .priority-3{
                        display:none;
                }
                .priority-2{
                        display:none;
                }
        
        }



</style>

