<template>
  <div>
    <h4>Hello Friendos</h4>
    Latitude: {{ currPosition.lat.toFixed(2) }},
    Longitude: {{ currPosition.lng.toFixed(2) }}
  </div>
  <div ref="mapDiv" style="width:100%; height:80vh" />
</template>

<script>
/* eslint-disable no-undef */
import { computed, ref, onMounted } from 'vue'
import { Loader } from '@googlemaps/js-api-loader'
import { useGeolocation } from './useGeolocation'

const GOOGLE_API_KEY=process.env.VUE_APP_GOOGLE_API_KEY

export default {
  name: 'App',
  setup() {
    const { coords } = useGeolocation()
    const currPosition = computed(() => ({
      lat: coords.value.latitude,
      lng: coords.value.longitude
    }))

    const loader = new Loader({ apiKey: GOOGLE_API_KEY })
    const mapDiv = ref(null)
    onMounted(async () => {
      await loader.load()
      new google.maps.Map(mapDiv.value, {
        center: currPosition.value,
        zoom: 7
      })
    })
    return { currPosition, mapDiv }
  },
  components: {
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
