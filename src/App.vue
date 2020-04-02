<template>
  <div id="app">
    <div v-if="menuV">
<my-menu></my-menu>
    </div>
    
    <router-view ></router-view>
  </div>
</template>

<script>
import menu from '@/components/menu'

export default {
  name: 'App',
  components: {
    'my-menu':menu
  },
  data(){
    return {
      menuV: true
      }
  },
  methods:{
    cambi(){
      console.log("cambiado")
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
