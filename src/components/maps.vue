<template>
  <GoogleMapLoader
    :mapConfig="mapConfig"
    apiKey="AIzaSyB9oLb6_-F8PvVs0Kvi0Zb1ggtWobNF514"
  />
</template>

<script>
import GoogleMapLoader from './GoogleMapLoader'
import { mapSettings } from '@/constants/configMapa'

export default {
  components: {
    GoogleMapLoader
  },
  computed: {
    mapConfig () {
      return {
        ...mapSettings,
        center: { lat: -30, lng: 0 }
      }
    },
  },
}
</script>