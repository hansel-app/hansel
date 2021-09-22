<template>
    <div class="backswipe-area" @click="goBack" v-touch:swipe.right="goBack">
    </div>
</template>
<script>
import { defineComponent } from "vue";
import { HOME_ROUTE } from "@/constants";

export default defineComponent({
    methods: {
        goBack() {
            console.log('backing');
            return this.$router.push(this.path);
        }
    },
    props: {
        path: {
            type: String,
            default: () => HOME_ROUTE,
        }
    }
})
</script>

<style scoped>
.backswipe-area {
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    height: 100vh;
    width: 6%;
    z-index: 1;
}
</style>