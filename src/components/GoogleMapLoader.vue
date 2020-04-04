<template>
  <div>
    <div class="google-map" ref="googleMap"></div>
  </div>
</template>

<script>
import GoogleMapsApiLoader from 'google-maps-api-loader'


export default {
  props: {
    mapConfig: Object,
    apiKey: String,
  },
  data() {
  return {
    google: null,
    map: null,
    lat: 34.4,
    lng: 34.3
  }
},
async mounted() {
    const googleMapApi = await GoogleMapsApiLoader({
      apiKey: this.apiKey
    })

    this.google = await googleMapApi
    console.log("paso por aqui")
     await navigator.geolocation.getCurrentPosition(async location=> {
      console.log("obten direccion entro")

      this.lng= await location.coords.longitude
      this.lat= await location.coords.latitude
      console.log(this.lng+" "+this.lat)
      console.log("obten direccion")
      this.initializeMap()
  })
     
  },

  methods: {
    initializeMap() {
      console.log("dos")
      const mapContainer = this.$refs.googleMap
      
      var myAdreess={
          center: {
            lat: this.lat,
            lng: this.lng
          }
      }

      Object.assign(this.mapConfig, myAdreess);
      this.map = new this.google.maps.Map(mapContainer, this.mapConfig)
      
      
    }
  }

}
</script>

<style>
.google-map{
    width: 100%;
  height: 400px;
 
}
</style>