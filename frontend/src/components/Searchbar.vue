<template>
    <div class="searchbar">
        <van-search
            :placeholder="placeholder" 
            background="none"
            shape="round"
        />
    </div>
</template>
<script>
import { defineComponent } from 'vue'
import { Search } from "vant";

export default defineComponent({
    components: {
        "van-search": Search,
    },
    props: {
        placeholder: {
            type: String,
            required: true,
        }
    }
})
</script>
<style scoped lang="less">
.van-search {
    margin: 1em;
}
</style>