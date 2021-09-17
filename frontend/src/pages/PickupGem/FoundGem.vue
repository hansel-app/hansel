<template>
  <div class="bg-container">
    <video autoplay muted poster="@/assets/images/gem-purple-found.png" id="gem-video">
      <source src="@/assets/videos/gem-purple-found.mp4" type="video/mp4">
    </video>
    <div id="found-gem">
      <h1>You found a gem!</h1>
    </div>
    <div id="view-button">
      <van-button>View contents</van-button>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";
import { Button } from "vant";

export default defineComponent({
  props: {
    gemColor: {
      // TODO: change bg src according to colour of gem
      type: String,
      enum: ['purple', 'pink', 'blue', 'black', 'yellow', 'green']
    }
  },
  components: { 
    "van-button": Button,
  },
});
</script>
<style scoped>
.bg-container {
  font-size: 100%;
  text-align: center;
  height:100vh;
  position: relative;
}

#found-gem {
  box-sizing: border-box;
  padding: 2em;
  display: inline-block;
  left: 0;
  position:absolute;
  height: 80%;
  width: 100%;
}

#gem-video {
  object-fit: cover;
  width: 100vw;
  height: 100vh;
  position: fixed;
  top: 0;
  left: 0;
}

.van-button {
  padding: 1em 2em;
  background-color: black;
  border: none;
  border-radius: 2em;
  color: white;
}

#view-button {
  position: absolute; 
  bottom: 20%; 
  left: 50%; 
  -webkit-transform: translateX(-50%);
  transform: translateX(-50%);
}

h1 {
  font-family: CircularStd-Black;
  font-size: 3em;
  color: rgb(87, 40, 80);
}
</style>