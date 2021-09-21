<template>
    <div class="searchbar">
        <Search :placeholder="placeholder" />
    </div>
</template>
<script>
import { defineComponent } from 'vue'
import { Search } from "vant";

export default defineComponent({
    components: {
        Search,
    },
    props: {
        placeholder: {
            type: String,
            required: true,
        }
    }
})
</script>
