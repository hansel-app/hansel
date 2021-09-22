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
