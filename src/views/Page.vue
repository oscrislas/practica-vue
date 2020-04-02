<template>
<div class="row">
  <div class="col-3">

  </div>
  <div class="col-6">
    <p>
Lorem ipsum dolor, sit amet consectetur adipisicing elit. Voluptatibus, ullam recusandae? Ipsa quidem praesentium dolorum neque magni quas est natus, in libero officia ab nulla blanditiis maiores similique deserunt nostrum!
</p>
  </div>
  <div class="col-3">

  </div>

</div>

 
 
</template>

<script>
export default {
    name: 'page',
    data(){
      return {

      }
    },
    methods:{

    
    },
    beforeCreate(){
     
        this.$emit('menuAb')
        console.log("ejecutado")
      }

}
</script>