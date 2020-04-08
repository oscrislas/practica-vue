<template>
    <div class="row">
        <div class="col-2"></div>
        <br>
        <div class="col-4">
            <p v-if="check==false">Hola {{User.Nombre}} {{User.Apellidos}} preciona el buton para checar tu hora de entrada</p>
            <p  v-else>Hola {{User.Nombre}} {{User.Apellidos}} preciona el buton para checar tu hora de salida</p>
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
                if(res.data==true){
                    this.check=true
                }else{
                     this.check=false
                }

            });
       
    },
    methods:{
        Entrada(){
            this.$axios.post('http://localhost:3000/registroEntrada',this.User)
            .then(res=>{
                console.log(res.data)
                var d = new Date();
              
                this.$notify({
                    group: 'foo',
                    title: 'Check',
                    text: "Checado a la hora "+d.getHours()+":"+d.getMinutes()
                });
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