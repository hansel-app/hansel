<template>
  <div class="container">
    <p class="header">Drop a gem</p>
    <!-- <div class="spacer" /> -->
    <p class="sub-header">Choose a receiver</p>
    <Search placeholder="Search a friend" />
    <CellGroup>
      <FriendCell
        v-for="user in mockFriends"
        :key="user.id"
        :friend="user"
        isClickable="true"
      />
    </CellGroup>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  setup() {},
  components: {
    CellGroup,
    FriendCell,
    Search,
  },
  data() {
    return {
      mockFriends,
    };
  },
});
</script>
