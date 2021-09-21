<template>
    <div class="background-gradient">
        <Header title="Gem logs"/>
        <div class="container">
            <Searchbar placeholder="Search a friend" />
        </div>
    </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import Header from '@/components/Header';
import Searchbar from '@/components/Searchbar';

export default defineComponent({
    components: {
        Header,
        Searchbar,
    },
    data() {
    return {
      isCloseWindow: true,
    };
  },
})
</script>