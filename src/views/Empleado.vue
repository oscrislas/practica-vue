<template>
    <div class="row">
        <div class="col-2"></div>
        <br>
        <div class="col-4">
            <p>Hola {{User.Nombre}} {{User.Apellidos}} preciona el buton para checar Asistencia</p>
               <b-button class="b" v-on:click="Entrada">checar Entrada</b-button>
        </div>

     
        <div class="col-4">
            <table>
                <tr>
                    <th>Hora</th>
                    <th>Lunes</th>
                    <th>Marte</th>
                    <th>Miercoles</th>
                    <th>Jueves</th>
                    <th>Viernes</th>
                </tr>
                 <tr>
                     
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
            }
        }
    },
    mounted(){
        if (localStorage.getItem('token')){
             this.User = JSON.parse(localStorage.getItem('token'));
             //console.log(JSON.parse(localStorage.getItem('token')))
        }
       
    },
    methods:{
        Entrada(){
            this.$axios.post('http://localhost:3000/registroEntrada',this.User)
            .then(res=>{
                console.log(res.Body)
            });
            console.log("checado")
        }
    }
}
</script>

<style scoped>
.b{
    text-align: center;
    padding: 10%;
    
}
</style>