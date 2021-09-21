<template>
  <div>
    <Header
      :isCloseWindow="true" 
      title="Drop a gem" 
    />
    <div class="container">
      <p class="sub-header">Choose a receiver</p>
      <Searchbar
        placeholder="Search a friend" 
      />
    <CellGroup>
      <FriendCell
        v-for="friend in filteredFriends"
        :key="friend.id"
        :friend="friend"
        :isClickable="true"
        :shouldDisplayUsername="true"
        @click="
          () => {
            $emit('set-receiver-event', friend.id);
            $emit('next-stage');
          }
        "
      />
    </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import { User } from "@/interfaces";
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
<style scoped>
.van-cell-group {
  background: none;
}
</style>