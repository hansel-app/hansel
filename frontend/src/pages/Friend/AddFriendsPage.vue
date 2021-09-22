<template>
  <div class="background-gradient">
    <Header title="Add friends" />
    <div class="container">
      <Search placeholder="Search by username" />
      <CellGroup>
        <FriendCell
          v-for="user in filteredUsers"
          :key="user.id"
          :friend="user"
          :shouldDisplayUsername="true"
        >
          <!-- TODO: find correct add user icon -->
          <van-icon name="user-o" @click="addFriend" />
        </FriendCell>
      </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import { User } from "@/interfaces";
import { useStore } from "vuex";
import FriendCell from "@/components/FriendCell.vue";
import Header from "@/components/Header.vue";

export default defineComponent({
  components: {
    CellGroup,
    FriendCell,
    Header,
    Search,
  },
  data() {
    return {
      mockFriends,
      searchQuery: "",
      store: useStore(),
    };
  },
  computed: {
    filteredUsers(): User[] {
      if (!this.searchQuery) {
        // Don't display any users when search bar is empty.
        return [];
      }
      return mockFriends.filter((user) => {
        const matchUsername = user.username
          .toLowerCase()
          .includes(this.searchQuery.toLowerCase());
        const matchDisplayname = user.displayName
          .toLowerCase()
          .includes(this.searchQuery.toLowerCase());
        return matchUsername || matchDisplayname;
      });
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    addFriend(receiverId: number) {
      this.store.dispatch("sendFriendRequest", receiverId);
    },
  },
});
</script>
