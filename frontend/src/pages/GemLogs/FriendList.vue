<template>
    <Header title="Gem logs" />
    <Searchbar placeholder="Gem logs" />
</template>
<script>
import { defineComponent } from 'vue';
import Header from '@/components/Header';
import Searchbar from '@/components/Searchbar';

export default defineComponent({
    // setup() {
        
    // },
    components: {
        Header,
        Searchbar,
    },
})
</script>
