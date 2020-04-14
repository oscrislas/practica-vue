<template>
    <div class="row">
        <div class="col-2"></div>
        <br>
        <div class="col-4">
            <p v-if="check==false">Hola {{User.nombre}} {{User.apellidos}} preciona el buton para checar tu hora de entrada</p>
            <p  v-else>Hola {{User.nombre}} {{User.apellidos}} preciona el buton para checar tu hora de salida</p>
               <b-button class="b" v-on:click="Entrada" v-if="check==false"> Checar Entrada
                </b-button>
                <b-button class="b" v-on:click="Entrada" v-else> Checar Salida
                </b-button>
        </div>

     
        <div class="col-4">
            <b-button v-on:click="Salir">Salir</b-button>
            <br>
            <br>
            
        </div>
        <div class="col-2"></div>
        <notifications group="foo" />
    </div>
</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
export default {
    name: 'empleado',
    data(){
        return{
            User: {
                nombre: "",
                apellidos: "",
                correo: ""
            },
            check: false
        }
    },
beforeMount(){

},
    mounted(){
        if (localStorage.getItem('token')){
             this.User = VueJwtDecode.decode(localStorage.getItem('token')).Username;
             //console.log(JSON.parse(localStorage.getItem('token')))
        }
        else{
            this.$router.push('/')
        }
            this.$axios.post('http://localhost:3000/check',this.User,{ headers: { 
                'Authorization': "Bearer "+localStorage.getItem('token'),
                'Content-Type': 'application/json' } })
            .then(res=>{
                if(res.data==true){
                    this.check=true
                }else{
                     this.check=false
                }

            });
       
    },
    methods:{
        Entrada(){
            this.$axios.post('http://localhost:3000/registroEntrada',this.User,{ headers: { 
                'Authorization': "Bearer "+localStorage.getItem('token'),
                'Content-Type': 'application/json' } })
            .then(res=>{
                console.log(res.data)
                var d = new Date();
              
                this.$notify({
                    group: 'foo',
                    title: 'Check',
                    text: "Checado a la hora "+d.getHours()+":"+d.getMinutes()
                });
                this.check=true
            });
         
            
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