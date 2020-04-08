<template>
  <div id="app">
    
    <router-view ></router-view>
  </div>
</template>

<script>


export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding-top: 60px;
  position: relative;
  background-image: url('./assets/b.jpg');
  background-repeat: repeat;
  background-size: cover;
  width: 100%;
  height: 1000px;
 

  
}
</style>
