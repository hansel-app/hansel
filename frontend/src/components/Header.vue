<template>
    <van-row class="header">
        <img id="icon" :src="arrowLeft">
        <div id="title" class="block">
            <p>{{ title }}</p>
        </div>
    </van-row>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import arrowLeft from '@/assets/icons/arrow-left.svg';
import { Row } from "vant";

export default defineComponent({
    setup() {
        return { arrowLeft }
    },
    props: {
        title: {
            type: String,
            required: true,
        }
    }, 
    components: {
        "van-row": Row,
    },
})
</script>
<style scoped>
#icon {
    fill: #585969;
    padding: 0.5em;
}
</style>