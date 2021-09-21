<template>
    <div class="page-header container">
        <van-row>
            <img id="icon-left" :src="arrowLeft">
            <img id="icon-right" :src="cross" @click="goToHome">
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
import cross from "@/assets/icons/cross.svg"
import { HOME_ROUTE } from "@/constants"


export default defineComponent({
    setup() {
        // return { arrowLeft, cross }
        const backMode = true;
        return { backMode, cross, arrowLeft };
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
    computed: {
        getIcon() {
            return require("@/assets/icons/cross.svg");
        },
        getIconPosition() {
            return !this.backMode? 'icon-left': 'icon-right';
        }
    },
    methods: {
        goToHome() {
            console.log('hi');
            this.$router.replace(HOME_ROUTE);
        },
    }
})
</script>
<style scoped>
.page-header {
    padding-top: 3em;
}

#icon-left {
    margin: -0.5em;
}

#icon-right {
    padding-left: 77%;
    padding-right: 0;
    margin-right: 0;
}

.hide-icon {
    display: none;
}
</style>
