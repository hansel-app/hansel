<template>
    <div class="drawer" @click="toggleBottomSheet">
        <div class="wrapper">
            <van-row ref="infoBox" class="message-content">
                <p>{{ gemPackage.message }}</p>
            </van-row>
            <div v-if="!collapsed">
              <van-row class="message-details">
                  <p>{{ gemPackage.dropper }}</p>
                  <p>{{ gemPackage.dropTime }}</p>
              </van-row>
              <van-row class="buttons">
                  <van-button @click="backToHome">Continue</van-button>
              </van-row>
            </div>
        </div>
    </div>
</template>
<script lang="ts">
import { GemInfo } from "@/interfaces";
import { defineComponent, PropType, ref, computed } from "vue";
import { Row, Button } from "vant";

export default defineComponent({
  setup() {
    const collapsed = ref(true)
    const toggleBottomSheet = () => (collapsed.value = !collapsed.value)
    // how to set this dynamically...
    const BOTTOM_SHEET_HEIGHT = 160
    const BOTTOM_SHEET_HEIGHT_COLLAPSED = 50
    const bottomSheetHeight = computed(
        () => `${collapsed.value ? BOTTOM_SHEET_HEIGHT_COLLAPSED : BOTTOM_SHEET_HEIGHT}px`
    )

    return { 
      collapsed, toggleBottomSheet, bottomSheetHeight,
    };
  },
  components: {
      "van-button": Button,
      "van-row": Row,
  },
  props: {
    gemPackage: {
      type: Object as PropType<GemInfo>,
      required: true,
    }
  },
  methods: {
    backToHome() {
      this.$router.replace({ name: "home" });
    },
  },
});
</script>
<style scoped>
.drawer {
    color: black;
    background-color: rgba(255, 255, 255, 0.3);
    backdrop-filter: blur(5px);
    border-radius: 2em 2em 0 0;

    float: left;
    position: fixed;
    z-index: 1;
    left: 0;
    right: 0;
    bottom: 0;
    padding: 0.5em;

    transition: max-height 0.3s;

    display: flex;
    flex-direction: row;
}

.wrapper {
    float: left;
    width: 100%;
    text-align: left;
    padding: 0 1em;
}

.Button {
    position: absolute;
    bottom: 0;
}
</style>