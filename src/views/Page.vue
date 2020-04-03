<template>
<div>
<my-menu></my-menu>
<div class="row">
  
  <div class="col-3">

  </div>
  <div class="col-6">
<h1>plomero</h1>
    <div>
        <md-field>
          <label for="Servicio">Servicio</label>
          <md-select name="movies" md-dense multiple>
            <md-option value="fight-club">Labado costo: 234</md-option>
            <md-option value="godfather">Drenaje costo: 234 </md-option>
            <md-option value="godfather-ii">baño costo: 234</md-option>
            <md-option value="godfather-iii">cosina costo: 234</md-option>
            <md-option value="godfellas">otro</md-option>

          </md-select>
        </md-field>

       
        <div class="total">
          <label for="costo" >Total del costo: $2434</label>
        </div>
         <my-mapas class="mapa"></my-mapas>
        

    </div>

  </div>
  <div class="col-3">

  </div>

</div>

 
</div>
</template>

<script>
import menu from '@/components/menu'
import mapas from '@/components/maps'
export default {
    name: 'page',
    data(){
      return {

      }
    },
    components: {
    'my-menu':menu,
    'my-mapas':mapas
    },
    methods:{

    
    },
    beforeCreate(){
     
    
        console.log("ejecutado")
      }

}
</script>

<style >
.total{
  position: absolute;;
  /*bottom:  -200%;*/
  right:  0;
}
.mapa{
  width: 100%;
  height: 200px;
  background-color: black;
}
</style>