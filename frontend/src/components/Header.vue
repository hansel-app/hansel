<template>
    <div class="page-header container">
        <van-row>
            <img id="icon-left" :src="arrowLeft">
        </van-row>
        <van-row>
            <div id="title" class="header">
                <p>{{ title }}</p>
            </div>
        </van-row>
    </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { Row } from "vant";
import arrowLeft from "@/assets/icons/arrow-left.svg"

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
.page-header {
    padding-top: 3em;
}
#icon-left {
    /* padding: 1em 1em 1em 2em; */
    fill: black;
}

/* #title {
    padding-bottom: 0em;
    margin-bottom: 0em;
} */
</style>