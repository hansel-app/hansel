<template>
  <h1>Drop a gem</h1>
  <h3>Choose a receiver</h3>
  <Search placeholder="Search a friend" />
  <CellGroup>
    <FriendCell
      v-for="user in mockFriends"
      :key="user.id"
      :friend="user"
      :isClickable="true"
      @click="this.$emit('SetReceiverEvent', user.id)"
    />
  </CellGroup>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
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
