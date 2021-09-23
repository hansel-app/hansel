<template>
    <div class="notification center-text">
        <!-- {{ notificationCount }} -->
        12
    </div>
</template>
<script>
import { defineComponent } from "vue"

export default defineComponent({
})
</script>
<style scoped lang="less">
.notification {
    background-color: white;
    // border: 1px red solid;
    color: black;
    height: 1.5em;
    width: 1.5em;
    margin-left: 1em;
    border-radius: 0.3em;
    vertical-align: middle;
}

.center-text {
    text-align: middle;
    line-height: 1.5em;
}
</style>