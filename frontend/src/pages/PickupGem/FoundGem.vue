<template>
  <!-- <div  :style="{'background-image':'url('+backgroundUrl+');'}"> -->
  <div class="bg-container">
    <div class="found-gem">
      <h1>You found a gem!!</h1>
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
html, body {
    height:100%;
    margin: 0;
    padding: 0;
}
* {
    box-sizing: border-box;
}
.bg-container {
  height:100vh;
  border:5px solid red;
  background: url("../../assets/purple-gem.png") no-repeat center center fixed; 
}
</style>