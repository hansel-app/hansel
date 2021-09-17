<template>
  <NavBar left-arrow @click-left="goBack" />
  <h2>Add friends</h2>
  <Search placeholder="Search by username" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar, Search } from "vant";

export default defineComponent({
  components: {
    NavBar,
    Search,
  },
  methods: {
    goBack() {
      this.$router.back();
    },
  },
});
</script>
