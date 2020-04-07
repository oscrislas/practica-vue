<template>
    <div class="row">
        <div class="col-2"></div>
        <br>
        <div class="col-4">
            <p>Hola {{User.Nombre}} {{User.Apellidos}} preciona el buton para checar Asistencia</p>
               <b-button class="b" v-on:click="Entrada"> 
                   <p v-if="check==false">Checar Entrada</p> 
                   <p v-else>Checar salida</p>
                </b-button>
        </div>

     
        <div class="col-4">
            <b-button v-on:click="Salir">Salir</b-button>
            <br>
            <br>
            <table class="table-dark">
                <tr>
                    <th>Hora</th>
                    <th>Lunes</th>
                    <th>Marte</th>
                    <th>Miercoles</th>
                    <th>Jueves</th>
                    <th>Viernes</th>
                </tr>
                 <tr >
                     <th>Entrada</th>
                     
                 </tr>
                 <tr>
                     <th>
                         Salida
                     </th>
                 </tr>
            </table>
        </div>
        <div class="col-2"></div>
    </div>
</template>

<script>
export default {
    name: 'empleado',
    data(){
        return{
            User: {
                Nombre: "",
                Apellidos: "",
                Correo: ""
            },
            check: true
        }
    },
beforeMount(){

},
    mounted(){
        if (localStorage.getItem('token')){
             this.User = JSON.parse(localStorage.getItem('token'));
             //console.log(JSON.parse(localStorage.getItem('token')))
        }
        else{
            this.$router.push('/')
        }
            this.$axios.post('http://localhost:3000/check',this.User)
            .then(res=>{
                console.log(res.data)
                if(res.Body==true){
                    this.check=true
                }

            });
       
    },
    methods:{
        Entrada(){
            this.$axios.post('http://localhost:3000/registroEntrada',this.User)
            .then(res=>{
                console.log(res.Body)
            });
            console.log("checado")
        },
        Salir(){
            console.log("salio")
            localStorage.removeItem('token')
            this.$router.push('/')
        }

    }
}
</script>

<style scoped>
.b{
    text-align: center;
    padding: 10%;
    
}
.table,th,td{
  padding: 10px;
  text-align: left;
  border: solid black;

}
</style>