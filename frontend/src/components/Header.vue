<template>
    <van-row class="page-header">
        <img id="icon" :src="arrowLeft">
        <div id="title" class="header">
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
.page-header {
    padding-top: 3em;
}
#icon {
    padding: 1em 1em 1em 2em;
}

#title {
    padding-bottom: 0em;
    margin-bottom: 0em;
}
</style>