<template>
  <div class="background-gradient">
    <Header title="Drop a gem" />
    <div class="container">
      <p class="sub-header">Choose a receiver</p>
      <Searchbar placeholder="Search a friend" />
      <CellGroup>
        <FriendCell
          v-for="user in mockFriends"
          :key="user.id"
          :friend="user"
          isClickable="true"
        />
      </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import FriendCell from "@/components/FriendCell.vue";
import Searchbar from "@/components/Searchbar.vue";
import Header from "@/components/Header.vue";

export default defineComponent({
  components: {
    CellGroup,
    FriendCell,
    Header,
    Searchbar,
  },
  data() {
    return {
      mockFriends,
    };
  },
});
</script>
<style scoped>
.van-cell-group {
  background: none;
}
</style>