<template>
  <vant-row>
    <p class="big-header overlay">{{ title }}</p>
  </vant-row>
  <div class="stack-container">
    <div class="semi-circle"></div>
    <Image :src="gemImgSrc" class="overlay" />
  </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from "vue";
// import { Image } from "vant";
import { GemColor } from "@/interfaces";
import { getEnumKeyByEnumValue } from "@/utils/enum";

export default defineComponent({
  // components: { Image },
  props: {
    title: {
      type: String,
      required: true,
    },
    color: {
      type: String as PropType<GemColor>,
      required: true,
    },
  },

  setup(props) {
    const gemColorName = getEnumKeyByEnumValue(GemColor, props.color);
    return {
      gemImgSrc: require(`@/assets/images/${gemColorName.toLowerCase()}_64.png`),
    };
  },
});
</script>

<style scoped>
.stack-container {
  position: relative;
  height: 325px;
  z-index: -1;
  overflow-x: hidden;
}

.overlay {
  position: absolute;
  left: 0;
  right: 0;
  margin-left: auto;
  margin-right: auto;
}

.semi-circle {
  transform: translate(-25%, -50%);
  height: 80vh;
  width: 200vw;
  border-radius: 0 0 50% 50%;
  background-color: #e4bfbf;
}

.gem-display {
  width: 100px;
  height: 100px;
}
</style>
