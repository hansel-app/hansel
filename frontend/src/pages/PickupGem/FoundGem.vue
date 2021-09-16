<template>
  <!-- <div  :style="{'background-image':'url('+backgroundUrl+');'}"> -->
  <div class="bg-container">
    <div class="found-gem">
      <h1>You found a gem!</h1>
      <button>View contents</button>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      // backgroundUrl: require('../../assets/purple-gem.png')
      //image: { backgroundImage: "../../assets/purple-gem.png" }
    };
  }
};
</script>
<style scoped>
.bg-container {
  height:100vh;
  position: relative;
  background-image: url("../../assets/purple-gem.png"); 
  background-size: cover;
  background-repeat: no-repeat;
  background-position: 50% 50%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}
.found-gem {
  left: 0;
  position:absolute;
  text-align:center;
  top: 3em;
  width: 100%;
}
h1 {
  font-family: CircularStd-Black;
  font-size: 3em;
  color: black;
}
</style>