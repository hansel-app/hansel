<template>
  <div class="container">
    <p class="header">Drop a gem</p>
    <p class="sub-header">Choose a receiver</p>
    <Search v-model="searchQuery" placeholder="Search a friend" />
    <CellGroup>
      <FriendCell
        v-for="friend in filteredFriends"
        :key="friend.id"
        :friend="friend"
        :isClickable="true"
        :shouldDisplayUsername="true"
        :onClick="() => $emit('nextStage')"
      />
    </CellGroup>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import { User } from "@/interfaces";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  components: {
    CellGroup,
    FriendCell,
    Search,
  },
  data() {
    return {
      allFriends: mockFriends,
      searchQuery: "",
    };
  },
  computed: {
    filteredFriends(): User[] {
      if (this.searchQuery.length === 0) {
        return this.allFriends;
      }
      return this.allFriends.filter((friend) => {
        const matchUsername = friend.username
          .toLowerCase()
          .includes(this.searchQuery.toLowerCase());
        const matchDisplayname = friend.displayName
          .toLowerCase()
          .includes(this.searchQuery.toLowerCase());
        return matchUsername || matchDisplayname;
      });
    },
  },
});
</script>
