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
.main{
    background-color: #001F49;
    padding: 20px;
    color: white;
    border-radius: 10px;
    padding-left: 0px;
    box-shadow: 10px 10px 5px 0px black;
    text-shadow: 0.5em 0.5em 0.10em black;
    font-family: "Homer Simpson UI";
    position: relative;
    width: auto;
    height: 90px;
   
    
   
}

</style>
