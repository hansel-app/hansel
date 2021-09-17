<template>
    <div class="drawer" :style="{ height: bottomSheetHeight }" @click="toggleBottomSheet">
        <span class="message-content">
            <p>Me gustan los camarones</p>
        </span>
        <span class="message-details" v-if="!collapsed">
            <p>hello</p>
        </span>
    </div>
</template>
<script>
import { defineComponent } from "vue";
import { collapsed, toggleBottomSheet, bottomSheetHeight } from "@/components/BottomSheet/state";

export default defineComponent({
  setup() {
    return { collapsed, toggleBottomSheet, bottomSheetHeight};
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

    transition: 0.3s ease;

    display: flex;
    flex-direction: row;
}
</style>