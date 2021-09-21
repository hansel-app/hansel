<template>
    <div class="background-gradient">
        <Header title="Gem logs" />
        <Searchbar placeholder="Search a friend" />
    </div>
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