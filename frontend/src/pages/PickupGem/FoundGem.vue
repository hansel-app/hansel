<template>
  <div class="bg-container">
    <div class="found-gem">
      <h1>You found a gem!</h1>
    </div>
    <div id="hi">
      <input type="button" class="button" value="View contents">
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
  font-size: 100%;
  text-align: center;
  height:100vh;
  position: relative;
  background-image: url("../../assets/purple-gem.png"); 
  background-size: cover;
  background-position: 50% 50%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}

.found-gem {
  box-sizing: border-box;
  padding: 2em;
  display: inline-block;
  left: 0;
  position:absolute;
  height: 80%;
  width: 100%;
}

.large-button {
  position: absolute;
  bottom: 0;
  margin: 0 auto;
}

#hi {
  position: absolute; 
  bottom: 20%; 
  left: 50%; 
  -webkit-transform: translateX(-50%);
  transform: translateX(-50%);
}

.button {
  padding: 1em 2em;
  background-color: black;
  border: none;
  border-radius: 2em;
  color: white;
}

h1 {
  font-family: CircularStd-Black;
  font-size: 3em;
  color: rgb(87, 40, 80);
}
</style>