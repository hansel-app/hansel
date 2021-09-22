<template>
    <div class="page-header container">
        <van-row>
            <img id="icon-left" :class="isBackIconDisplayed" :src="arrowLeft" @click="goBack">
            <img id="icon-right" :class="isCrossIconDisplayed" :src="cross" @click="goToHome">
        </van-row>
        <van-row>
            <div id="title" class="header">
                <p>{{ title }}</p>
            </div>
        </van-row>
    </div>
</template>
<script>
import { defineComponent } from "vue";
import { Row } from "vant";
import arrowLeft from "@/assets/icons/arrow-left.svg";
import cross from "@/assets/icons/cross.svg";

export default defineComponent({
    setup() {
        return { cross, arrowLeft };
    },
    props: {
        title: {
            type: String,
            required: true,
        },
        isCloseWindow: {
            type: Boolean,
            default: () => false,
        },
        backLink: {
            type: String,
            required: false,
        }
    }, 
    components: {
        "van-row": Row,
    },
    computed: {
        getIcon() {
            return require("@/assets/icons/cross.svg");
        },
        goBack() {
            return this.$router.back();
        },
        isCrossIconDisplayed() {
            return {
                'hide-icon': !this.isCloseWindow
            }
        },
        isBackIconDisplayed() {
            return {
                'hide-icon': this.isCloseWindow
            }
        }
    },
    methods: {
        goToHome() {
            this.$router.push({ name: 'home'});
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
    position: absolute;
    right: 1.5em;
}

.hide-icon {
    display: none;
}
</style>
